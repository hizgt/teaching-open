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
	// 角色控制器
	roleCtrl := system.NewSysRoleController()
	// 权限控制器
	permCtrl := system.NewSysPermissionController()
	// 部门控制器
	departCtrl := system.NewSysDepartController()

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

		// 角色管理
		group.Group("/role", func(roleGroup *ghttp.RouterGroup) {
			roleGroup.GET("/list", roleCtrl.List)
			roleGroup.GET("/queryAll", roleCtrl.QueryAll)
			roleGroup.GET("/queryById", roleCtrl.QueryById)
			roleGroup.POST("/add", roleCtrl.Add)
			roleGroup.PUT("/edit", roleCtrl.Edit)
			roleGroup.DELETE("/delete", roleCtrl.Delete)
			roleGroup.DELETE("/deleteBatch", roleCtrl.DeleteBatch)
			roleGroup.GET("/queryUserRole", roleCtrl.QueryUserRole)
			roleGroup.POST("/saveUserRole", roleCtrl.SaveUserRole)
			roleGroup.GET("/queryRolePermission", roleCtrl.QueryRolePermission)
			roleGroup.POST("/saveRolePermission", roleCtrl.SaveRolePermission)
		})

		// 权限管理
		group.Group("/permission", func(permGroup *ghttp.RouterGroup) {
			permGroup.GET("/list", permCtrl.List)
			permGroup.GET("/queryById", permCtrl.QueryById)
			permGroup.POST("/add", permCtrl.Add)
			permGroup.PUT("/edit", permCtrl.Edit)
			permGroup.DELETE("/delete", permCtrl.Delete)
			permGroup.GET("/getUserPermission", permCtrl.GetUserPermission)
			permGroup.GET("/getUserMenus", permCtrl.GetUserMenus)
			permGroup.GET("/getUserPermCodes", permCtrl.GetUserPermCodes)
		})

		// 部门管理
		group.Group("/sysDepart", func(departGroup *ghttp.RouterGroup) {
			departGroup.GET("/queryTreeList", departCtrl.QueryTreeList)
			departGroup.GET("/queryById", departCtrl.QueryById)
			departGroup.POST("/add", departCtrl.Add)
			departGroup.PUT("/edit", departCtrl.Edit)
			departGroup.DELETE("/delete", departCtrl.Delete)
			departGroup.GET("/queryIdTree", departCtrl.QueryIdTree)
			departGroup.GET("/searchBy", departCtrl.SearchBy)
			departGroup.GET("/queryDepartTreeSync", departCtrl.QueryDepartTreeSync)
			departGroup.GET("/queryUserDepart", departCtrl.QueryUserDepart)
			departGroup.POST("/saveUserDepart", departCtrl.SaveUserDepart)
		})
	})
}