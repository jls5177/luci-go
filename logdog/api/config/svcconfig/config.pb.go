// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/config/svcconfig/config.proto

package svcconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Config is the overall instance configuration.
type Config struct {
	// Configuration for the Butler's log transport.
	Transport *Transport `protobuf:"bytes,10,opt,name=transport" json:"transport,omitempty"`
	// Configuration for intermediate Storage.
	Storage *Storage `protobuf:"bytes,11,opt,name=storage" json:"storage,omitempty"`
	// Coordinator is the coordinator service configuration.
	Coordinator *Coordinator `protobuf:"bytes,20,opt,name=coordinator" json:"coordinator,omitempty"`
	// Collector is the collector fleet configuration.
	Collector *Collector `protobuf:"bytes,21,opt,name=collector" json:"collector,omitempty"`
	// Archivist microservice configuration.
	Archivist *Archivist `protobuf:"bytes,22,opt,name=archivist" json:"archivist,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Config) GetTransport() *Transport {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (m *Config) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *Config) GetCoordinator() *Coordinator {
	if m != nil {
		return m.Coordinator
	}
	return nil
}

func (m *Config) GetCollector() *Collector {
	if m != nil {
		return m.Collector
	}
	return nil
}

func (m *Config) GetArchivist() *Archivist {
	if m != nil {
		return m.Archivist
	}
	return nil
}

// Coordinator is the Coordinator service configuration.
type Coordinator struct {
	// The name of the authentication group for administrators.
	AdminAuthGroup string `protobuf:"bytes,10,opt,name=admin_auth_group,json=adminAuthGroup" json:"admin_auth_group,omitempty"`
	// The name of the authentication group for backend services.
	ServiceAuthGroup string `protobuf:"bytes,11,opt,name=service_auth_group,json=serviceAuthGroup" json:"service_auth_group,omitempty"`
	// A list of origin URLs that are allowed to perform CORS RPC calls.
	RpcAllowOrigins []string `protobuf:"bytes,20,rep,name=rpc_allow_origins,json=rpcAllowOrigins" json:"rpc_allow_origins,omitempty"`
	// The maximum amount of time after a prefix has been registered when log
	// streams may also be registered under that prefix.
	//
	// After the expiration period has passed, new log stream registration will
	// fail.
	//
	// Project configurations or stream prefix regitrations may override this by
	// providing >= 0 values for prefix expiration. The smallest configured
	// expiration will be applied.
	PrefixExpiration *google_protobuf.Duration `protobuf:"bytes,21,opt,name=prefix_expiration,json=prefixExpiration" json:"prefix_expiration,omitempty"`
	// The full path of the archival Pub/Sub topic.
	//
	// The Coordinator must have permission to publish to this topic.
	ArchiveTopic string `protobuf:"bytes,30,opt,name=archive_topic,json=archiveTopic" json:"archive_topic,omitempty"`
	// The amount of time after an archive request has been dispatched before it
	// should be executed.
	//
	// Since terminal messages can arrive out of order, the archival request may
	// be kicked off before all of the log stream data has been loaded into
	// intermediate storage. If this happens, the Archivist will retry archival
	// later autometically.
	//
	// This parameter is an optimization to stop the archivist from wasting its
	// time until the log stream has a reasonable expectation of being available.
	ArchiveSettleDelay *google_protobuf.Duration `protobuf:"bytes,31,opt,name=archive_settle_delay,json=archiveSettleDelay" json:"archive_settle_delay,omitempty"`
	// The amount of time before a log stream is candidate for archival regardless
	// of whether or not it's been terminated or complete.
	//
	// This is a failsafe designed to ensure that log streams with missing records
	// or no terminal record (e.g., Butler crashed) are eventually archived.
	//
	// This should be fairly large (days) to avoid prematurely archiving
	// long-running streams, but should be considerably smaller than the
	// intermediate storage data retention period.
	//
	// If a project's "max_stream_age" is smaller than this value, it will be used
	// on that project's streams.
	ArchiveDelayMax *google_protobuf.Duration `protobuf:"bytes,32,opt,name=archive_delay_max,json=archiveDelayMax" json:"archive_delay_max,omitempty"`
}

func (m *Coordinator) Reset()                    { *m = Coordinator{} }
func (m *Coordinator) String() string            { return proto.CompactTextString(m) }
func (*Coordinator) ProtoMessage()               {}
func (*Coordinator) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Coordinator) GetAdminAuthGroup() string {
	if m != nil {
		return m.AdminAuthGroup
	}
	return ""
}

func (m *Coordinator) GetServiceAuthGroup() string {
	if m != nil {
		return m.ServiceAuthGroup
	}
	return ""
}

func (m *Coordinator) GetRpcAllowOrigins() []string {
	if m != nil {
		return m.RpcAllowOrigins
	}
	return nil
}

func (m *Coordinator) GetPrefixExpiration() *google_protobuf.Duration {
	if m != nil {
		return m.PrefixExpiration
	}
	return nil
}

func (m *Coordinator) GetArchiveTopic() string {
	if m != nil {
		return m.ArchiveTopic
	}
	return ""
}

func (m *Coordinator) GetArchiveSettleDelay() *google_protobuf.Duration {
	if m != nil {
		return m.ArchiveSettleDelay
	}
	return nil
}

func (m *Coordinator) GetArchiveDelayMax() *google_protobuf.Duration {
	if m != nil {
		return m.ArchiveDelayMax
	}
	return nil
}

// Collector is the set of configuration parameters for Collector instances.
type Collector struct {
	// The maximum number of concurrent transport messages to process. If <= 0,
	// a default will be chosen based on the transport.
	MaxConcurrentMessages int32 `protobuf:"varint,1,opt,name=max_concurrent_messages,json=maxConcurrentMessages" json:"max_concurrent_messages,omitempty"`
	// The maximum number of concurrent workers to process each ingested message.
	// If <= 0, collector.DefaultMaxMessageWorkers will be used.
	MaxMessageWorkers int32 `protobuf:"varint,2,opt,name=max_message_workers,json=maxMessageWorkers" json:"max_message_workers,omitempty"`
	// The maximum number of log stream states to cache locally. If <= 0, a
	// default will be used.
	StateCacheSize int32 `protobuf:"varint,3,opt,name=state_cache_size,json=stateCacheSize" json:"state_cache_size,omitempty"`
	// The maximum amount of time that cached stream state is valid. If <= 0, a
	// default will be used.
	StateCacheExpiration *google_protobuf.Duration `protobuf:"bytes,4,opt,name=state_cache_expiration,json=stateCacheExpiration" json:"state_cache_expiration,omitempty"`
}

func (m *Collector) Reset()                    { *m = Collector{} }
func (m *Collector) String() string            { return proto.CompactTextString(m) }
func (*Collector) ProtoMessage()               {}
func (*Collector) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Collector) GetMaxConcurrentMessages() int32 {
	if m != nil {
		return m.MaxConcurrentMessages
	}
	return 0
}

func (m *Collector) GetMaxMessageWorkers() int32 {
	if m != nil {
		return m.MaxMessageWorkers
	}
	return 0
}

func (m *Collector) GetStateCacheSize() int32 {
	if m != nil {
		return m.StateCacheSize
	}
	return 0
}

func (m *Collector) GetStateCacheExpiration() *google_protobuf.Duration {
	if m != nil {
		return m.StateCacheExpiration
	}
	return nil
}

// Configuration for the Archivist microservice.
type Archivist struct {
	// The name of the archival Pub/Sub subscription.
	//
	// This should be connected to "archive_topic", and the Archivist must have
	// permission to consume from this subscription.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// The number of tasks to run at a time. If blank, the archivist will choose a
	// default value.
	Tasks int32 `protobuf:"varint,2,opt,name=tasks" json:"tasks,omitempty"`
	// The name of the staging storage bucket. All projects will share the same
	// staging bucket. Logs for a project will be staged under:
	//
	// gs://<gs_staging_bucket>/<app-id>/<project-name>/...
	GsStagingBucket string `protobuf:"bytes,3,opt,name=gs_staging_bucket,json=gsStagingBucket" json:"gs_staging_bucket,omitempty"`
	// Service-wide index configuration. This is used if per-project configuration
	// is not specified.
	ArchiveIndexConfig *ArchiveIndexConfig `protobuf:"bytes,10,opt,name=archive_index_config,json=archiveIndexConfig" json:"archive_index_config,omitempty"`
	// If true, always render the log entries as a binary file during archival,
	// regardless of whether a specific stream has a binary file extension.
	//
	// By default, a stream will only be rendered as a binary if its descriptor
	// includes a non-empty binary file extension field.
	//
	// The binary stream consists of each log entry's data rendered back-to-back.
	//   - For text streams, this produces a text document similar to the source
	//     text.
	//   - For binary streams and datagram streams, this reproduces the source
	//     contiguous binary file.
	//   - For datagram streams, the size-prefixed datagrams are written back-to-
	//     back.
	//
	// Enabling this option will consume roughly twice the archival space, as each
	// stream's data will be archived once as a series of log entries and once as
	// a binary file.
	//
	// Streams without an explicit binary file extension will default to ".bin" if
	// this is enabled.
	RenderAllStreams bool `protobuf:"varint,13,opt,name=render_all_streams,json=renderAllStreams" json:"render_all_streams,omitempty"`
}

func (m *Archivist) Reset()                    { *m = Archivist{} }
func (m *Archivist) String() string            { return proto.CompactTextString(m) }
func (*Archivist) ProtoMessage()               {}
func (*Archivist) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Archivist) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *Archivist) GetTasks() int32 {
	if m != nil {
		return m.Tasks
	}
	return 0
}

func (m *Archivist) GetGsStagingBucket() string {
	if m != nil {
		return m.GsStagingBucket
	}
	return ""
}

func (m *Archivist) GetArchiveIndexConfig() *ArchiveIndexConfig {
	if m != nil {
		return m.ArchiveIndexConfig
	}
	return nil
}

func (m *Archivist) GetRenderAllStreams() bool {
	if m != nil {
		return m.RenderAllStreams
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "svcconfig.Config")
	proto.RegisterType((*Coordinator)(nil), "svcconfig.Coordinator")
	proto.RegisterType((*Collector)(nil), "svcconfig.Collector")
	proto.RegisterType((*Archivist)(nil), "svcconfig.Archivist")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/logdog/api/config/svcconfig/config.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6a, 0xdb, 0x3c,
	0x1c, 0x27, 0xed, 0xd7, 0x7e, 0x8b, 0xd2, 0x43, 0xa2, 0xa5, 0x9d, 0x57, 0x58, 0x17, 0xb2, 0x9b,
	0x30, 0x8a, 0x03, 0x1d, 0x8c, 0xdd, 0xec, 0x90, 0xa5, 0xdd, 0x18, 0xa3, 0x14, 0x9c, 0xc2, 0x2e,
	0x85, 0x22, 0x2b, 0x8a, 0xa8, 0x6d, 0x19, 0x49, 0x6e, 0xbd, 0x3e, 0xc3, 0xae, 0xf6, 0xa0, 0x63,
	0x8f, 0x30, 0x74, 0xf0, 0x01, 0x7a, 0x51, 0x7a, 0x95, 0xf8, 0x77, 0xf2, 0x5f, 0xfa, 0xc9, 0x02,
	0xef, 0x99, 0x08, 0xc9, 0x5a, 0x8a, 0x94, 0x17, 0x69, 0x28, 0x24, 0x9b, 0x26, 0x05, 0xe1, 0xd3,
	0x44, 0xb0, 0x58, 0xb0, 0x29, 0xce, 0xf9, 0x94, 0x88, 0x6c, 0xc5, 0xd9, 0x54, 0xdd, 0x10, 0xff,
	0xcf, 0xfd, 0x84, 0xb9, 0x14, 0x5a, 0xc0, 0x6e, 0x8d, 0x1f, 0x7d, 0x7c, 0x74, 0x12, 0x96, 0x64,
	0xcd, 0x6f, 0x70, 0xe2, 0xb2, 0x8e, 0x3e, 0x3c, 0x3a, 0x40, 0x69, 0x21, 0x31, 0xa3, 0xde, 0xff,
	0xe9, 0xd1, 0x7e, 0x2d, 0x71, 0xa6, 0x72, 0x21, 0xb5, 0x4f, 0x38, 0x66, 0x42, 0xb0, 0x84, 0x4e,
	0xed, 0xd3, 0xb2, 0x58, 0x4d, 0xe3, 0x42, 0x62, 0xcd, 0x45, 0xe6, 0xf8, 0xf1, 0xaf, 0x0d, 0xb0,
	0x3d, 0xb7, 0x56, 0x78, 0x0a, 0xba, 0xb5, 0x3b, 0x00, 0xa3, 0xce, 0xa4, 0x77, 0x3a, 0x0c, 0xeb,
	0xe4, 0xf0, 0xaa, 0xe2, 0xa2, 0x46, 0x06, 0x4f, 0xc0, 0xff, 0x7e, 0xe2, 0xa0, 0x67, 0x1d, 0xb0,
	0xe5, 0x58, 0x38, 0x26, 0xaa, 0x24, 0xf0, 0x1d, 0xe8, 0x11, 0x21, 0x64, 0xcc, 0x33, 0xac, 0x85,
	0x0c, 0x86, 0xd6, 0x71, 0xd8, 0x72, 0xcc, 0x1b, 0x36, 0x6a, 0x4b, 0xcd, 0x6c, 0x44, 0x24, 0x09,
	0x25, 0xc6, 0x77, 0x70, 0x6f, 0xb6, 0x79, 0xc5, 0x45, 0x8d, 0xcc, 0x78, 0x5c, 0x1d, 0x5c, 0xe9,
	0xe0, 0xf0, 0x9e, 0x67, 0x56, 0x71, 0x51, 0x23, 0x1b, 0xff, 0xde, 0x04, 0xbd, 0xd6, 0x10, 0x70,
	0x02, 0xfa, 0x38, 0x4e, 0x79, 0x86, 0x70, 0xa1, 0xd7, 0x88, 0x49, 0x51, 0xe4, 0x76, 0x6b, 0xba,
	0xd1, 0x9e, 0xc5, 0x67, 0x85, 0x5e, 0x7f, 0x35, 0x28, 0x3c, 0x01, 0x50, 0x51, 0x79, 0xc3, 0x09,
	0x6d, 0x6b, 0x7b, 0x56, 0xdb, 0xf7, 0x4c, 0xa3, 0x7e, 0x0d, 0x06, 0x32, 0x27, 0x08, 0x27, 0x89,
	0xb8, 0x45, 0x42, 0x72, 0xc6, 0x33, 0x15, 0x0c, 0x47, 0x9b, 0x93, 0x6e, 0xb4, 0x2f, 0x73, 0x32,
	0x33, 0xf8, 0xa5, 0x83, 0xe1, 0x17, 0x30, 0xc8, 0x25, 0x5d, 0xf1, 0x12, 0xd1, 0x32, 0xe7, 0xae,
	0x3d, 0xbf, 0x07, 0xcf, 0x43, 0x57, 0x6f, 0x58, 0xd5, 0x1b, 0x9e, 0xf9, 0x7a, 0xa3, 0xbe, 0xf3,
	0x9c, 0xd7, 0x16, 0xf8, 0x0a, 0xec, 0xba, 0x85, 0x52, 0xa4, 0x45, 0xce, 0x49, 0x70, 0x6c, 0x87,
	0xdb, 0xf1, 0xe0, 0x95, 0xc1, 0xe0, 0x77, 0x30, 0xac, 0x44, 0x8a, 0x6a, 0x9d, 0x50, 0x14, 0xd3,
	0x04, 0xff, 0x0c, 0x5e, 0x3e, 0xf4, 0x3e, 0xe8, 0x6d, 0x0b, 0xeb, 0x3a, 0x33, 0x26, 0x78, 0x0e,
	0x06, 0x55, 0x98, 0x4d, 0x41, 0x29, 0x2e, 0x83, 0xd1, 0x43, 0x49, 0xfb, 0xde, 0x63, 0x33, 0x2e,
	0x70, 0x39, 0xfe, 0xd3, 0x01, 0xdd, 0xba, 0x61, 0xf8, 0x16, 0x3c, 0x4b, 0x71, 0x89, 0x88, 0xc8,
	0x48, 0x21, 0x25, 0xcd, 0x34, 0x4a, 0xa9, 0x52, 0x98, 0x51, 0x15, 0x74, 0x46, 0x9d, 0xc9, 0x56,
	0x74, 0x90, 0xe2, 0x72, 0x5e, 0xb3, 0x17, 0x9e, 0x84, 0x21, 0x78, 0x6a, 0x7c, 0x5e, 0x8c, 0x6e,
	0x85, 0xbc, 0xa6, 0x52, 0x05, 0x1b, 0xd6, 0x33, 0x48, 0x71, 0xe9, 0x95, 0x3f, 0x1c, 0x61, 0xaa,
	0x57, 0x1a, 0x6b, 0x8a, 0x08, 0x26, 0x6b, 0x8a, 0x14, 0xbf, 0xa3, 0xc1, 0xa6, 0x15, 0xef, 0x59,
	0x7c, 0x6e, 0xe0, 0x05, 0xbf, 0xa3, 0xf0, 0x12, 0x1c, 0xb6, 0x95, 0xad, 0x96, 0xfe, 0x7b, 0x68,
	0xad, 0xc3, 0x26, 0xaa, 0x69, 0x6a, 0xfc, 0xb7, 0x03, 0xba, 0xf5, 0xf1, 0x84, 0x63, 0xb0, 0xa3,
	0x8a, 0xa5, 0x22, 0x92, 0xe7, 0x36, 0xb4, 0xe3, 0x6a, 0x6b, 0x63, 0x70, 0x08, 0xb6, 0x34, 0x56,
	0xd7, 0xd5, 0x72, 0xdc, 0x83, 0x39, 0x65, 0x4c, 0x21, 0xa5, 0x31, 0xe3, 0x19, 0x43, 0xcb, 0x82,
	0x5c, 0x53, 0x6d, 0xd7, 0xd0, 0x8d, 0xf6, 0x99, 0x5a, 0x38, 0xfc, 0xb3, 0x85, 0xe1, 0x65, 0x53,
	0x3c, 0xcf, 0x62, 0x6a, 0x37, 0x78, 0xc5, 0x99, 0xbf, 0x08, 0x5e, 0xdc, 0xfb, 0x70, 0xe8, 0x37,
	0xa3, 0x72, 0x57, 0x47, 0x5d, 0x7e, 0x0b, 0x33, 0x1f, 0x84, 0xa4, 0x59, 0x4c, 0xa5, 0x39, 0xe5,
	0x48, 0x69, 0x49, 0x71, 0xaa, 0x82, 0xdd, 0x51, 0x67, 0xf2, 0x24, 0xea, 0x3b, 0x66, 0x96, 0x24,
	0x0b, 0x87, 0x2f, 0xb7, 0xed, 0xde, 0xbc, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0xac, 0x23, 0x61,
	0xf1, 0xbd, 0x05, 0x00, 0x00,
}
