// Copyright 2025 武晓晨 <wuxc.eng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/iWuxc/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package apiserver

import (
	"context"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	handler "github.com/iWuxc/miniblog/internal/apiserver/handler/grpc"
	"github.com/iWuxc/miniblog/internal/pkg/server"
	"net/http"
)

// ginServer 定义一个使用 Gin 框架开发的 HTTP 服务器.
type ginServer struct {
	srv server.Server
}

// 确保 *ginServer 实现了 server.Server 接口.
var _ server.Server = (*ginServer)(nil)

// NewGinServer 初始化一个新的 Gin 服务器实例.
func (c *ServerConfig) NewGinServer() server.Server {
	//创建Gin引擎
	engine := gin.New()

	//注册REST API路由
	c.InstallRESTAPI(engine)

	httpsrv := server.NewHttpServer(c.cfg.HTTPOptions, engine)

	return &ginServer{srv: httpsrv}
}

// 注册 API 路由。路由的路径和 HTTP 方法，严格遵循 REST 规范.
func (c *ServerConfig) InstallRESTAPI(engine *gin.Engine) {
	//注册业务五官的 API 接口
	InstallGenericAPI(engine)

	//创建核心业务处理器
	handler.NewHandler()

	//注册健康检查接口
	//engine.GET("/healthz", handler.Healthz)
}

// InstallGenericAPI 注册业务无关的路由，例如 pprof、404 处理等.
func InstallGenericAPI(engine *gin.Engine) {
	// 注册 pprof 路由
	pprof.Register(engine)

	// 注册 404 路由处理
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "Page not found.")
	})
}

// 注册 API 路由。路由的路径和 HTTP 方法，严格遵循 REST 规范.// RunOrDie 启动 Gin 服务器，出错则程序崩溃退出.
func (s *ginServer) RunOrDie() {
	select {}
}

// GracefulStop 优雅停止服务器.
func (s *ginServer) GracefulStop(ctx context.Context) {
	s.srv.GracefulStop(ctx)
}
