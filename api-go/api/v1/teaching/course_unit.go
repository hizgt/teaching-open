package teaching

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 课程单元列表
type CourseUnitListReq struct {
	g.Meta   `path:"/teaching/courseUnit/list" method:"get" tags:"教学管理" summary:"课程单元列表"`
	CourseId string `json:"courseId" v:"required" dc:"课程ID"`
}

type CourseUnitListRes struct {
	Records []UnitItem `json:"records"`
}

// 创建课程单元
type CourseUnitCreateReq struct {
	g.Meta      `path:"/teaching/courseUnit/add" method:"post" tags:"教学管理" summary:"创建课程单元"`
	CourseId    string `json:"courseId" v:"required" dc:"课程ID"`
	Name        string `json:"name" v:"required" dc:"单元名称"`
	Content     string `json:"content" dc:"单元内容"`
	SortOrder   int    `json:"sortOrder" dc:"排序"`
	ResourceUrl string `json:"resourceUrl" dc:"资源链接"`
}

type CourseUnitCreateRes struct {
	Id string `json:"id"`
}

// 更新课程单元
type CourseUnitUpdateReq struct {
	g.Meta      `path:"/teaching/courseUnit/edit" method:"put" tags:"教学管理" summary:"更新课程单元"`
	Id          string `json:"id" v:"required" dc:"单元ID"`
	Name        string `json:"name" v:"required" dc:"单元名称"`
	Content     string `json:"content" dc:"单元内容"`
	SortOrder   int    `json:"sortOrder" dc:"排序"`
	ResourceUrl string `json:"resourceUrl" dc:"资源链接"`
}

// 删除课程单元
type CourseUnitDeleteReq struct {
	g.Meta `path:"/teaching/courseUnit/delete" method:"delete" tags:"教学管理" summary:"删除课程单元"`
	Id     string `json:"id" v:"required" dc:"单元ID"`
}

// 课程单元排序
type CourseUnitSortReq struct {
	g.Meta `path:"/teaching/courseUnit/sort" method:"put" tags:"教学管理" summary:"课程单元排序"`
	CourseId string       `json:"courseId" v:"required" dc:"课程ID"`
	Units    []UnitSortItem `json:"units" v:"required" dc:"单元排序列表"`
}

type UnitSortItem struct {
	Id        string `json:"id" v:"required" dc:"单元ID"`
	SortOrder int    `json:"sortOrder" v:"required|min:0" dc:"排序"`
}