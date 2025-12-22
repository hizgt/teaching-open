package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"teaching-open/api/middleware"
	"teaching-open/internal/controller/hello"
	"teaching-open/internal/controller/sys"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "Teaching Open Backend Server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 全局中间件
			s.Use(middleware.CORS, middleware.Logger, middleware.Error)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				// 公开路由
				group.Group("/api/v1", func(group *ghttp.RouterGroup) {
					// 健康检查
					group.GET("/health", func(r *ghttp.Request) {
						r.Response.WriteJson(g.Map{
							"code":    0,
							"message": "ok",
							"result": g.Map{
								"status":  "healthy",
								"version": "3.0.0",
								"name":    "Teaching Open API",
							},
							"success": true,
						})
					})

					// 用户登录
					group.Bind(sys.SysUser.Login)

					// 原有hello示例
					group.Bind(hello.NewV1())
				})

				// 需要认证的路由
				group.Group("/api/v1", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					// 用户管理
					group.Bind(
						sys.SysUser.GetList,
						sys.SysUser.Add,
						sys.SysUser.Edit,
						sys.SysUser.Delete,
						sys.SysUser.GetById,
					)
					// 角色管理
					group.Bind(
						sys.SysRole.GetList,
						sys.SysRole.Add,
						sys.SysRole.Edit,
						sys.SysRole.Delete,
						sys.SysRole.GetById,
						sys.SysRole.GetAll,
						sys.SysRole.GetUserRoles,
						sys.SysRole.SaveUserRoles,
					)
					// 权限管理
					group.Bind(
						sys.SysPermission.GetList,
						sys.SysPermission.GetTree,
						sys.SysPermission.Add,
						sys.SysPermission.Edit,
						sys.SysPermission.Delete,
						sys.SysPermission.GetById,
						sys.SysPermission.GetRolePermissions,
						sys.SysPermission.SaveRolePermissions,
						sys.SysPermission.GetUserPermissions,
					)
					// 部门管理
					group.Bind(
						sys.SysDepart.GetTree,
						sys.SysDepart.Add,
						sys.SysDepart.Edit,
						sys.SysDepart.Delete,
						sys.SysDepart.GetById,
						sys.SysDepart.GetIdTree,
						sys.SysDepart.SearchBy,
						sys.SysDepart.GetUserDeparts,
						sys.SysDepart.SaveUserDepart,
					)
					// 字典管理
					group.Bind(
						sys.Dict.List,
						sys.Dict.Add,
						sys.Dict.Edit,
						sys.Dict.Delete,
						sys.Dict.GetById,
						sys.Dict.ItemList,
						sys.Dict.GetItemsByCode,
						sys.Dict.AddItem,
						sys.Dict.EditItem,
						sys.Dict.DeleteItem,
					)
					// 日志管理
					group.Bind(
						sys.SysLog.GetList,
						sys.SysLog.Delete,
						sys.SysLog.DeleteBatch,
						sys.SysLog.Clear,
						sys.SysLog.GetDataLogList,
						sys.SysLog.GetDataLogById,
						sys.SysLog.GetDataLogHistory,
					)
					// 文件管理
					group.Bind(
						sys.SysFile.Upload,
						sys.SysFile.UploadBatch,
						sys.SysFile.GetList,
						sys.SysFile.GetById,
						sys.SysFile.Delete,
						sys.SysFile.DeleteBatch,
						sys.SysFile.View,
						sys.SysFile.Download,
					)
					// 课程管理
					group.Bind(
						sys.TeachingCourse.GetList,
						sys.TeachingCourse.GetHomeCourse,
						sys.TeachingCourse.Add,
						sys.TeachingCourse.Edit,
						sys.TeachingCourse.Delete,
						sys.TeachingCourse.DeleteBatch,
						sys.TeachingCourse.GetById,
						sys.TeachingCourse.Publish,
						sys.TeachingCourse.SetShared,
						sys.TeachingCourse.AuthorizeDept,
					)
					// 课程单元管理
					group.Bind(
						sys.TeachingCourseUnit.GetList,
						sys.TeachingCourseUnit.GetByCourseId,
						sys.TeachingCourseUnit.Add,
						sys.TeachingCourseUnit.Edit,
						sys.TeachingCourseUnit.Delete,
						sys.TeachingCourseUnit.DeleteBatch,
						sys.TeachingCourseUnit.GetById,
						sys.TeachingCourseUnit.Sort,
					)
					// 班级课程管理
					group.Bind(
						sys.TeachingCourseDept.GetList,
						sys.TeachingCourseDept.GetByDeptId,
						sys.TeachingCourseDept.GetByCourseId,
						sys.TeachingCourseDept.AddOrUpdate,
						sys.TeachingCourseDept.Delete,
						sys.TeachingCourseDept.BatchAdd,
					)
					// 作品管理
					group.Bind(
						sys.TeachingWork.List,
						sys.TeachingWork.Mine,
						sys.TeachingWork.GreatWork,
						sys.TeachingWork.StarWork,
						sys.TeachingWork.Leaderboard,
						sys.TeachingWork.Add,
						sys.TeachingWork.Edit,
						sys.TeachingWork.Delete,
						sys.TeachingWork.DeleteBatch,
						sys.TeachingWork.GetById,
						sys.TeachingWork.Submit,
						sys.TeachingWork.StudentWorkInfo,
						sys.TeachingWork.SendWork,
						sys.TeachingWork.MineAdditional,
						sys.TeachingWork.StarToggle,
						sys.TeachingWork.CollectToggle,
						sys.TeachingWork.CorrectList,
						sys.TeachingWork.CorrectAdd,
						sys.TeachingWork.CommentList,
						sys.TeachingWork.CommentAdd,
						sys.TeachingWork.CommentDelete,
						sys.TeachingWork.TagGet,
						sys.TeachingWork.TagSet,
						sys.TeachingWork.TagDelete,
					)
					// 新闻公告管理
					group.Bind(
						sys.TeachingNews.List,
						sys.TeachingNews.PublicList,
						sys.TeachingNews.Add,
						sys.TeachingNews.Edit,
						sys.TeachingNews.Delete,
						sys.TeachingNews.DeleteBatch,
						sys.TeachingNews.QueryById,
						sys.TeachingNews.Publish,
						sys.TeachingNews.Offline,
					)
					// 附加作业管理
					group.Bind(
						sys.TeachingAdditionalWork.List,
						sys.TeachingAdditionalWork.ListByDept,
						sys.TeachingAdditionalWork.Add,
						sys.TeachingAdditionalWork.Edit,
						sys.TeachingAdditionalWork.Delete,
						sys.TeachingAdditionalWork.DeleteBatch,
						sys.TeachingAdditionalWork.QueryById,
						sys.TeachingAdditionalWork.Publish,
						sys.TeachingAdditionalWork.Offline,
					)
					// Scratch素材管理
					group.Bind(
						sys.TeachingScratchAsset.List,
						sys.TeachingScratchAsset.GetScratchAssets,
						sys.TeachingScratchAsset.Add,
						sys.TeachingScratchAsset.Edit,
						sys.TeachingScratchAsset.Delete,
						sys.TeachingScratchAsset.DeleteBatch,
						sys.TeachingScratchAsset.QueryById,
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
