package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
	"teaching-open/api/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
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
					"time":   gtime.Now().String(),
				})
			})

			// 启动服务器
			s.Run()
			return nil
		},
	}
)