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
	handler "github.com/iWuxc/miniblog/internal/apiserver/handler/http"
	mw "github.com/iWuxc/miniblog/internal/pkg/middleware/gin"
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

	// 注册全局中间件，用于恢复 panic、设置 HTTP 头、添加请求 ID 等
	engine.Use(
		gin.Recovery(),
		mw.NoCache,
		mw.Cors,
		mw.Secure,
		mw.RequestIDMiddleware(),
		mw.AuthnBypasswMiddleware())

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
	handler := handler.NewHandler(c.biz, c.val)

	//注册健康检查接口
	engine.GET("/healthz", handler.Healthz)

	// 注册用户登录和令牌刷新接口。这两个接口比较简单，所以没有API版本
	engine.POST("/login", handler.Login)
	engine.PUT("refresh-token", handler.RefreshToken)

	var authMiddlewares []gin.HandlerFunc

	// 注册 v1 版本 API 路由分组
	v1 := engine.Group("v1")
	{
		// 用户相关理由
		userv1 := v1.Group("/users")
		{
			// 创建用户。这里要注意：创建用户是不用进行认证和授权的
			userv1.POST("", handler.CreateUser)
			userv1.Use(authMiddlewares...)
			userv1.PUT(":userID/change-password", handler.ChangePassword) //修改用户密码
			userv1.PUT(":userID", handler.UpdateUser)                     // 跟新用户信息
			userv1.DELETE(":userID", handler.DeleteUser)                  //删除用户
			userv1.GET(":userID", handler.GetUser)                        //获取用户信息
			userv1.GET("", handler.ListUser)                              //查询用户列表
		}
		postv1 := v1.Group("/posts", authMiddlewares...)
		{
			postv1.POST("", handler.CreatePost)          //创建文章
			postv1.PUT(":postID", handler.UpdatePost)    //修改文章
			postv1.DELETE(":postID", handler.DeletePost) //删除文章
			postv1.GET(":postID", handler.GetPost)       //获取文章
			postv1.GET("", handler.ListPost)             //查询文章列表
		}
	}
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
	s.srv.RunOrDie()
}

// GracefulStop 优雅停止服务器.
func (s *ginServer) GracefulStop(ctx context.Context) {
	s.srv.GracefulStop(ctx)
}
