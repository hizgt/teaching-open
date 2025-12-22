package teaching

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 课程列表
type CourseListReq struct {
	g.Meta `path:"/teaching/course/list" method:"get" tags:"教学管理" summary:"课程列表"`
	Page   int    `json:"page" v:"required|min:1" dc:"页码"`
	PageSize int  `json:"pageSize" v:"required|min:1|max:100" dc:"每页数量"`
	Name    string `json:"name" dc:"课程名称"`
	Type    string `json:"type" dc:"课程类型"`
	Status  string `json:"status" dc:"状态"`
	CreateBy string `json:"createBy" dc:"创建人"`
}

type CourseListRes struct {
	Records   []CourseItem `json:"records"`
	Total     int          `json:"total"`
	Page      int          `json:"pageNo"`
	PageSize  int          `json:"pageSize"`
}

type CourseItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CoverImage  string `json:"coverImage"`
	Status      string `json:"status"`
	CreateBy    string `json:"createBy"`
	CreateTime  string `json:"createTime"`
	UnitCount   int    `json:"unitCount"`
}

// 创建课程
type CourseCreateReq struct {
	g.Meta      `path:"/teaching/course/add" method:"post" tags:"教学管理" summary:"创建课程"`
	Name        string   `json:"name" v:"required" dc:"课程名称"`
	Type        string   `json:"type" v:"required|in:scratch,python" dc:"课程类型"`
	Description string   `json:"description" dc:"课程描述"`
	CoverImage  string   `json:"coverImage" dc:"封面图"`
	DeptIds     []string `json:"deptIds" dc:"关联部门ID列表"`
}

type CourseCreateRes struct {
	Id string `json:"id"`
}

// 更新课程
type CourseUpdateReq struct {
	g.Meta      `path:"/teaching/course/edit" method:"put" tags:"教学管理" summary:"更新课程"`
	Id          string   `json:"id" v:"required" dc:"课程ID"`
	Name        string   `json:"name" v:"required" dc:"课程名称"`
	Type        string   `json:"type" v:"required|in:scratch,python" dc:"课程类型"`
	Description string   `json:"description" dc:"课程描述"`
	CoverImage  string   `json:"coverImage" dc:"封面图"`
	DeptIds     []string `json:"deptIds" dc:"关联部门ID列表"`
}

// 删除课程
type CourseDeleteReq struct {
	g.Meta `path:"/teaching/course/delete" method:"delete" tags:"教学管理" summary:"删除课程"`
	Id     string `json:"id" v:"required" dc:"课程ID"`
}

// 课程详情
type CourseDetailReq struct {
	g.Meta `path:"/teaching/course/detail" method:"get" tags:"教学管理" summary:"课程详情"`
	Id     string `json:"id" v:"required" dc:"课程ID"`
}

type CourseDetailRes struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	Description string         `json:"description"`
	CoverImage  string         `json:"coverImage"`
	Status      string         `json:"status"`
	CreateBy    string         `json:"createBy"`
	CreateTime  string         `json:"createTime"`
	UpdateBy    string         `json:"updateBy"`
	UpdateTime  string         `json:"updateTime"`
	Depts       []DeptItem     `json:"depts"`
	Units       []UnitItem     `json:"units"`
}

type DeptItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UnitItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	SortOrder   int    `json:"sortOrder"`
	ResourceUrl string `json:"resourceUrl"`
}

// 发布/下架课程
type CoursePublishReq struct {
	g.Meta `path:"/teaching/course/publish" method:"put" tags:"教学管理" summary:"发布/下架课程"`
	Id     string `json:"id" v:"required" dc:"课程ID"`
	Status string `json:"status" v:"required|in:draft,published" dc:"状态"`
}