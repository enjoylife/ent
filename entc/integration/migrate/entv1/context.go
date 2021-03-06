// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"context"
)

type contextKey struct{}

// FromContext returns the Client stored in a context, or nil if there isn't one.
func FromContext(ctx context.Context) *Client {
	c, _ := ctx.Value(contextKey{}).(*Client)
	return c
}

// NewContext returns a new context with the given Client attached.
func NewContext(parent context.Context, c *Client) context.Context {
	return context.WithValue(parent, contextKey{}, c)
}
