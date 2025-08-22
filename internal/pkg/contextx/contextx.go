// Copyright 2025 武晓晨 <wuxc.eng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/iWuxc/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package contextx

import "context"

// 定义用于上下文的键
type (
	// userIDKey 定义用户 ID 的上下文键.
	userIDKey struct{}

	// requestIDKey 定义请求 ID 的上下文键.
	requestIDKey struct{}
)

// WithUserID 将用户 ID 存放到上下文中.
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// UserID 从上下文中提取用户 ID.
func UserID(ctx context.Context) string {
	return ctx.Value(userIDKey{}).(string)
}

// WithRequestID 将请求 ID 存放到上下文中.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// RequestID 从上下文中提取请求 ID.
func RequestID(ctx context.Context) string {
	return ctx.Value(requestIDKey{}).(string)
}
