// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package internal

import (
	"crypto/sha1"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

type serviceAccountTokenProvider struct {
	oauthTokenProvider

	ctx    context.Context
	config *jwt.Config
}

// NewServiceAccountTokenProvider returns TokenProvider that supports service accounts.
func NewServiceAccountTokenProvider(ctx context.Context, jsonKey []byte, scopes []string) (TokenProvider, error) {
	config, err := google.JWTConfigFromJSON(jsonKey, scopes...)
	if err != nil {
		return nil, err
	}
	return &serviceAccountTokenProvider{
		oauthTokenProvider: oauthTokenProvider{
			interactive: false,
			tokenFlavor: "service_account",
		},
		ctx:    ctx,
		config: config,
	}, nil
}

func (p *serviceAccountTokenProvider) CacheSeed() []byte {
	seed := sha1.Sum(p.config.PrivateKey)
	return seed[:]
}

func (p *serviceAccountTokenProvider) MintToken() (Token, error) {
	src := p.config.TokenSource(p.ctx)
	tok, err := src.Token()
	if err != nil {
		return nil, err
	}
	return makeToken(tok), nil
}

func (p *serviceAccountTokenProvider) RefreshToken(Token) (Token, error) {
	// JWT tokens are self sufficient, there's no need for refresh_token. Minting
	// a token and "refreshing" it is a same thing.
	return p.MintToken()
}
