// Package sys provides API definitions for teaching management.
package sys

import (
	"github.com/gogf/gf/v2/frame/g"

	"teaching-open/internal/model/entity"
)

// =============================================================================
// Course API - 课程管理接口
// =============================================================================

// CourseListReq 课程列表请求
type CourseListReq struct {
	g.Meta         `path:"/teaching/teachingCourse/list" method:"get" tags:"课程管理" summary:"课程列表"`
	CourseName     string `json:"courseName"     dc:"课程名称"`
	CourseType     string `json:"courseType"     dc:"课程类型"`
	CourseCategory string `json:"courseCategory" dc:"课程分类"`
	IsShared       int    `json:"isShared"       dc:"是否共享: -1全部 0否 1是" d:"-1"`
	ShowHome       int    `json:"showHome"       dc:"首页展示: -1全部 0否 1是" d:"-1"`
	PageNo         int    `json:"pageNo"         dc:"页码" d:"1"`
	PageSize       int    `json:"pageSize"       dc:"每页数量" d:"10"`
}

// CourseListRes 课程列表响应
type CourseListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*entity.TeachingCourse `json:"records"`
	Total    int                      `json:"total"`
	PageNo   int                      `json:"current"`
	PageSize int                      `json:"size"`
}

// CourseHomeCourseReq 首页课程请求
type CourseHomeCourseReq struct {
	g.Meta `path:"/teaching/teachingCourse/getHomeCourse" method:"get" tags:"课程管理" summary:"首页课程列表"`
}

// CourseHomeCourseRes 首页课程响应
type CourseHomeCourseRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.TeachingCourse `json:"list"`
}

// CourseAddReq 添加课程请求
type CourseAddReq struct {
	g.Meta         `path:"/teaching/teachingCourse/add" method:"post" tags:"课程管理" summary:"添加课程"`
	CourseName     string `json:"courseName"     v:"required#课程名称不能为空" dc:"课程名称"`
	CourseDesc     string `json:"courseDesc"     dc:"课程描述"`
	CourseIcon     string `json:"courseIcon"     dc:"课程图标"`
	CourseCover    string `json:"courseCover"    dc:"课程封面"`
	ShowType       int    `json:"showType"       dc:"展示类型" d:"1"`
	CourseMap      string `json:"courseMap"      dc:"课程地图"`
	IsShared       int    `json:"isShared"       dc:"是否共享" d:"0"`
	ShowHome       int    `json:"showHome"       dc:"首页展示" d:"0"`
	OrderNum       int    `json:"orderNum"       dc:"排序" d:"1"`
	DepartIds      string `json:"departIds"      dc:"授权部门ID"`
	CourseType     string `json:"courseType"     dc:"课程类型"`
	CourseCategory string `json:"courseCategory" dc:"课程分类"`
}

// CourseAddRes 添加课程响应
type CourseAddRes struct {
	g.Meta `mime:"application/json"`
}

// CourseEditReq 编辑课程请求
type CourseEditReq struct {
	g.Meta         `path:"/teaching/teachingCourse/edit" method:"put" tags:"课程管理" summary:"编辑课程"`
	Id             string `json:"id"             v:"required#课程ID不能为空" dc:"课程ID"`
	CourseName     string `json:"courseName"     v:"required#课程名称不能为空" dc:"课程名称"`
	CourseDesc     string `json:"courseDesc"     dc:"课程描述"`
	CourseIcon     string `json:"courseIcon"     dc:"课程图标"`
	CourseCover    string `json:"courseCover"    dc:"课程封面"`
	ShowType       int    `json:"showType"       dc:"展示类型"`
	CourseMap      string `json:"courseMap"      dc:"课程地图"`
	IsShared       int    `json:"isShared"       dc:"是否共享"`
	ShowHome       int    `json:"showHome"       dc:"首页展示"`
	OrderNum       int    `json:"orderNum"       dc:"排序"`
	DepartIds      string `json:"departIds"      dc:"授权部门ID"`
	CourseType     string `json:"courseType"     dc:"课程类型"`
	CourseCategory string `json:"courseCategory" dc:"课程分类"`
}

// CourseEditRes 编辑课程响应
type CourseEditRes struct {
	g.Meta `mime:"application/json"`
}

// CourseDeleteReq 删除课程请求
type CourseDeleteReq struct {
	g.Meta `path:"/teaching/teachingCourse/delete" method:"delete" tags:"课程管理" summary:"删除课程"`
	Id     string `json:"id" v:"required#课程ID不能为空" dc:"课程ID"`
}

// CourseDeleteRes 删除课程响应
type CourseDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CourseDeleteBatchReq 批量删除课程请求
type CourseDeleteBatchReq struct {
	g.Meta `path:"/teaching/teachingCourse/deleteBatch" method:"delete" tags:"课程管理" summary:"批量删除课程"`
	Ids    string `json:"ids" v:"required#课程ID不能为空" dc:"课程ID，多个用逗号分隔"`
}

// CourseDeleteBatchRes 批量删除课程响应
type CourseDeleteBatchRes struct {
	g.Meta `mime:"application/json"`
}

// CourseGetByIdReq 获取课程详情请求
type CourseGetByIdReq struct {
	g.Meta `path:"/teaching/teachingCourse/queryById" method:"get" tags:"课程管理" summary:"课程详情"`
	Id     string `json:"id" v:"required#课程ID不能为空" dc:"课程ID"`
}

// CourseGetByIdRes 获取课程详情响应
type CourseGetByIdRes struct {
	g.Meta `mime:"application/json"`
	*entity.TeachingCourse
}

// CoursePublishReq 发布/下架课程请求
type CoursePublishReq struct {
	g.Meta   `path:"/teaching/teachingCourse/publish" method:"put" tags:"课程管理" summary:"发布/下架课程"`
	Id       string `json:"id"       v:"required#课程ID不能为空" dc:"课程ID"`
	ShowHome int    `json:"showHome" dc:"首页展示: 0下架 1发布"`
}

// CoursePublishRes 发布/下架课程响应
type CoursePublishRes struct {
	g.Meta `mime:"application/json"`
}

// CourseSetSharedReq 设置共享请求
type CourseSetSharedReq struct {
	g.Meta   `path:"/teaching/teachingCourse/setShared" method:"put" tags:"课程管理" summary:"设置课程共享"`
	Id       string `json:"id"       v:"required#课程ID不能为空" dc:"课程ID"`
	IsShared int    `json:"isShared" dc:"是否共享: 0否 1是"`
}

// CourseSetSharedRes 设置共享响应
type CourseSetSharedRes struct {
	g.Meta `mime:"application/json"`
}

// CourseAuthorizeDeptReq 授权部门请求
type CourseAuthorizeDeptReq struct {
	g.Meta    `path:"/teaching/teachingCourse/authorizeDept" method:"post" tags:"课程管理" summary:"授权课程给部门"`
	Id        string `json:"id"        v:"required#课程ID不能为空" dc:"课程ID"`
	DepartIds string `json:"departIds" dc:"部门ID，多个用逗号分隔"`
}

// CourseAuthorizeDeptRes 授权部门响应
type CourseAuthorizeDeptRes struct {
	g.Meta `mime:"application/json"`
}
