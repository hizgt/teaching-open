package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"teaching-open/api/middleware"
)

func main() {
	s := g.Server()

	// 全局中间件
	s.Use(
		middleware.CORS,           // 跨域
		middleware.SecurityHeaders, // 安全头
		middleware.Logging,        // 日志
	)

	// 健康检查
	s.BindHandler("/health", func(r *ghttp.Request) {
		r.Response.WriteJson(g.Map{
			"status": "ok",
		})
	})

	// 启动服务器
	s.Run()
}