# Copyright 2019 The LUCI Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Swarming related supporting structs and functions."""

load('@stdlib//internal/luci/lib/validate.star', 'validate')
load('@stdlib//internal/time.star', 'time')


# See https://chromium.googlesource.com/infra/luci/luci-py/+/master/appengine/swarming/server/config.py
_DIMENSION_VALUE_LENGTH = 256
_DIMENSION_KEY_RE = r'^[a-zA-Z\-\_\.]{1,64}$'


# See https://chromium.googlesource.com/infra/infra/+/master/appengine/cr-buildbucket/swarming/swarmingcfg.py
#
# Except we drop {1,4096} because Go regexp engine refuses to use such high
# repetition numbers (and 4K for a name is effectively unbounded anyway). Let
# the server deal with such crazy values itself.
_CACHE_NAME_RE = r'^[a-z0-9_]+$'


# A struct returned by swarming.cache(...).
#
# See swarming.cache(...) function for all details.
#
# Fields:
#   path: string, where to mount the cache.
#   name: string, name of the cache to mount.
#   wait_for_warm_cache: duration or None, how long to wait for a warm cache.
_cache_ctor = genstruct('swarming.cache')


# A struct returned by swarming.dimension(...).
#
# See swarming.dimension(...) function for all details.
#
# Fields:
#   value: string, value of the dimension.
#   expiration: duration or None, when the dimension expires.
_dimension_ctor = genstruct('swarming.dimension')


def _cache(path, name=None, wait_for_warm_cache=None):
  """A request for the bot to mount a named cache to a path.

  Each bot has a LRU of named caches: think of them as local named directories
  in some protected place that survive between builds.

  A build can request one or more such caches to be mounted (in read/write mode)
  at the requested path relative to some known root. In recipes-based builds,
  the path is relative to api.paths['cache'] dir.

  If it's the first time a cache is mounted on this particular bot, it will
  appear as an empty directory. Otherwise it will contain whatever was left
  there by the previous build that mounted exact same named cache on this bot,
  even if that build is completely irrelevant to the current build and just
  happened to use the same named cache (sometimes this is useful to share state
  between different builders).

  At the end of the build the cache directory is unmounted. If at that time the
  bot is running out of space, caches (in their entirety, the named cache
  directory and all files inside) are evicted in LRU manner until there's enough
  free disk space left. Renaming a cache is equivalent to clearing it from the
  builder perspective. The files will still be there, but eventually will be
  purged by GC.

  Additionally, Buildbucket always implicitly requests to mount a special
  builder cache to 'builder' path:

      swarming.cache('builder', name=some_hash('<project>/<bucket>/<builder>'))

  This means that any LUCI builder has a "personal disk space" on the bot.
  Builder cache is often a good start before customizing caching. In recipes, it
  is available at api.path['cache'].join('builder').

  In order to share the builder cache directory among multiple builders, some
  explicitly named cache can be mounted to 'builder' path on these builders.
  Buildbucket will not try to override it with its auto-generated builder cache.

  For example, if builders 'a' and 'b' both declare they use named cache
  swarming.cache('builder', name='my_shared_cache'), and an 'a' build ran on
  a bot and left some files in the builder cache, then when a 'b' build runs on
  the same bot, the same files will be available in its builder cache.

  If the pool of swarming bots is shared among multiple LUCI projects and
  projects mount same named cache, the cache will be shared across projects.
  To avoid affecting and being affected by other projects, prefix the cache
  name with something project-specific, e.g. "v8-".

  Args:
    path: path where the cache should be mounted to, relative to some known
        root (in recipes this root is api.path['cache']). Must use POSIX
        format (forward slashes). In most cases, it does not need slashes at
        all. Must be unique in the given builder definition (cannot mount
        multiple caches to the same path).
    name: identifier of the cache to mount to the path. Default is same
        value as 'path' itself. Must be unique in the given builder definition
        (cannot mount the same cache to multiple paths).
    wait_for_warm_cache: how long to wait (with minutes precision) for a bot
        that has this named cache already to become available and pick up the
        build, before giving up and starting looking for any matching bot
        (regardless whether it has the cache or not). If there are no bots with
        this cache at all, the build will skip waiting and will immediately
        fallback to any matching bot. By default (if unset or zero), there'll
        be no attempt to find a bot with this cache already warm: the build
        may or may not end up on a warm bot, there's no guarantee one way or
        another.

  Returns:
    swarming.cache struct.
  """
  path = validate.string('path', path)
  name = validate.string('name', name, default=path, regexp=_CACHE_NAME_RE, required=False)

  # See also _validate_relative_path in appengine/cr-buildbucket/swarming/swarmingcfg.py.
  if not path:
    fail('bad cache path: must not be empty')
  if '\\' in path:
    fail('bad cache path %r: must not contain "\\", for Windows sake' % path)
  if '..' in path.split('/'):
    fail('bad cache path %r: must not contain ".."' % path)
  if path.startswith('/'):
    fail('bad cache path %r: must not start with "/"' % path)

  return _cache_ctor(
      path = path,
      name = name,
      wait_for_warm_cache = validate.duration(
          'wait_for_warm_cache',
          wait_for_warm_cache,
          precision=time.minute,
          min=time.minute,
          required=False,
      ),
  )


def _dimension(value, expiration=None):
  """A value of some Swarming dimension, annotated with its expiration time.

  Intended to be used as a value in 'dimensions' dict when using dimensions
  that expire:

      dimensions = {
          ...
          'device': swarming.dimension('preferred', expiration=5*time.minute),
          ...
      }

  Args:
    value: string value of the dimension.
    expiration: how long to wait (with minutes precision) for a bot with this
        dimension to become available and pick up the build, or None to wait
        until the overall build expiration timeout.

  Returns:
    swarming.dimension struct.
  """
  # See also 'validate_dimension_value' in appengine/swarming/server/config.py.
  val = validate.string('value', value)
  if not val:
    fail('bad dimension value: must not be empty')
  if len(val) > _DIMENSION_VALUE_LENGTH:
    fail('bad dimension value %r: must be at most %d chars' % (val, _DIMENSION_VALUE_LENGTH))
  if val != val.strip():
    fail('bad dimension value %r: must not have leading or trailing whitespace' % (val,))
  return _dimension_ctor(
      value = val,
      expiration = validate.duration(
          'expiration',
          expiration,
          precision=time.minute,
          min=time.minute,
          required=False,
      ),
  )


def _validate_caches(attr, caches):
  """Validates a list of caches.

  Ensures each entry is swarming.cache struct, and no two entries use same
  name or path.

  Args:
    attr: field name with caches, for error messages.
    caches: a list of swarming.cache(...) entries to validate.

  Returns:
    Validate list of caches (may be an empty list, never None).
  """
  per_path = {}
  per_name = {}
  caches = validate.list(attr, caches)
  for c in caches:
    validate.struct(attr, c, _cache_ctor)
    if c.path in per_path:
      fail('bad "caches": caches %s and %s use same path' % (c, per_path[c.path]))
    if c.name in per_name:
      fail('bad "caches": caches %s and %s use same name' % (c, per_name[c.name]))
    per_path[c.path] = c
    per_name[c.name] = c
  return caches


def _validate_dimensions(attr, dimensions):
  """Validates and normalizes a dict with dimensions.

  The dict should have string keys and values are swarming.dimension, a string
  or a list of thereof (for repeated dimensions).

  Args:
    attr: field name with dimensions, for error messages.
    dimensions: a dict {string: string|swarming.dimension}.

  Returns:
    Validated and normalized dict in form {string: [swarming.dimension]}.
  """
  out = {}
  for k, v in validate.str_dict(attr, dimensions).items():
    # See also 'validate_dimension_key' in appengine/swarming/server/config.py.
    validate.string(attr, k, regexp=_DIMENSION_KEY_RE)
    if type(v) == 'list':
      out[k] = [_as_dim(k, x) for x in v]
    else:
      out[k] = [_as_dim(k, v)]
  return out


def _as_dim(key, val):
  """string|swarming.dimension -> swarming.dimension."""
  if type(val) == 'string':
    return _dimension(val)
  return validate.struct(key, val, _dimension_ctor)


def _validate_tags(attr, tags):
  """Validates a list of "k:v" pairs with Swarming tags.

  Args:
    attr: field name with tags, for error messages.
    tags: a list of tags to validate.

  Returns:
    Validated list of tags in same order, with duplicates removed.
  """
  out = set()  # note: in starlark sets/dicts remember the order
  for t in validate.list(attr, tags):
    validate.string(attr, t, regexp=r'.+\:.+')
    out = out.union([t])
  return list(out)


# Public API exposed to end-users and to other LUCI modules.
swarming = struct(
    cache = _cache,
    dimension = _dimension,

    # Validators are useful for macros that modify caches, dimensions, etc.
    validate_caches = _validate_caches,
    validate_dimensions = _validate_dimensions,
    validate_tags = _validate_tags,
)