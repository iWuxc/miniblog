package grpc

import (
	"context"
	"github.com/iWuxc/miniblog/internal/pkg/contextx"
	"github.com/iWuxc/miniblog/internal/pkg/known"
	"github.com/iWuxc/miniblog/internal/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthnBypasswInterceptor 是一个 gRPC 拦截器，模拟所有请求都通过认证。
func AuthnBypasswInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 从请求头中获取用户ID
		userID := "user-00001" // 默认用户ID
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			// 获取header中指定的用户ID，假设Header名为"x-user-id"
			if values := md.Get(known.XUserID); len(values) > 0 {
				userID = values[0]
			}
		}

		log.Debugw("Simulated authentication successful", "userID", userID)

		// 将默认的用于信息存入上下文
		//noline: staticcheck
		ctx = context.WithValue(ctx, known.XUserID, userID)

		// 为 log 和 contexts 提供上下文支持
		ctx = contextx.WithUserID(ctx, userID)

		//继续处理请求
		return handler(ctx, req)
	}
}
