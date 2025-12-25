package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
	"teaching-open/api/middleware"
	"teaching-open/internal/controller/system"
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

			// 注册路由
			registerRoutes(s)

			// 启动服务器
			s.Run()
			return nil
		},
	}
)

// registerRoutes 注册路由
func registerRoutes(s *ghttp.Server) {
	// 登录控制器
	loginCtrl := system.NewLoginController()
	// 用户控制器
	userCtrl := system.NewSysUserController()

	// 系统模块路由组
	s.Group("/sys", func(group *ghttp.RouterGroup) {
		// 无需认证的接口
		group.POST("/login", loginCtrl.Login)
		group.POST("/phoneLogin", loginCtrl.PhoneLogin)
		group.GET("/randomImage/:key", loginCtrl.GetCaptcha)
		group.POST("/sms", loginCtrl.SendSms)

		// 需要认证的接口
		group.Middleware(middleware.Auth)
		group.POST("/logout", loginCtrl.Logout)

		// 用户管理
		group.Group("/user", func(userGroup *ghttp.RouterGroup) {
			userGroup.GET("/list", userCtrl.List)
			userGroup.POST("/add", userCtrl.Add)
			userGroup.PUT("/edit", userCtrl.Edit)
			userGroup.DELETE("/delete", userCtrl.Delete)
			userGroup.DELETE("/deleteBatch", userCtrl.DeleteBatch)
			userGroup.PUT("/resetPassword", userCtrl.ResetPassword)
			userGroup.GET("/queryById", userCtrl.QueryById)
		})
	})
}