package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"teaching-open/internal/model/entity"
)

type CourseUnitListReq struct {
	g.Meta   `path:"/teaching/teachingCourseUnit/list" method:"get" tags:"课程单元管理" summary:"课程单元列表"`
	CourseId string `json:"courseId" v:"required#课程ID不能为空" dc:"课程ID"`
	UnitName string `json:"unitName" dc:"单元名称"`
	PageNo   int    `json:"pageNo" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
}

type CourseUnitListRes struct {
	g.Meta   `mime:"application/json"`
	List     []*entity.TeachingCourseUnit `json:"records"`
	Total    int                          `json:"total"`
	PageNo   int                          `json:"current"`
	PageSize int                          `json:"size"`
}

type CourseUnitAllReq struct {
	g.Meta   `path:"/teaching/teachingCourseUnit/queryByCourseId" method:"get" tags:"课程单元管理" summary:"获取课程所有单元"`
	CourseId string `json:"courseId" v:"required#课程ID不能为空" dc:"课程ID"`
}

type CourseUnitAllRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.TeachingCourseUnit `json:"list"`
}

type CourseUnitAddReq struct {
	g.Meta            `path:"/teaching/teachingCourseUnit/add" method:"post" tags:"课程单元管理" summary:"添加课程单元"`
	UnitName          string `json:"unitName" v:"required#单元名称不能为空"`
	UnitIntro         string `json:"unitIntro"`
	UnitCover         string `json:"unitCover"`
	CourseId          string `json:"courseId" v:"required#课程ID不能为空"`
	CourseVideo       string `json:"courseVideo"`
	CourseVideoSource int    `json:"courseVideoSource" d:"1"`
	ShowCourseVideo   int    `json:"showCourseVideo" d:"1"`
	CourseCase        string `json:"courseCase"`
	ShowCourseCase    int    `json:"showCourseCase" d:"1"`
	CoursePpt         string `json:"coursePpt"`
	ShowCoursePpt     int    `json:"showCoursePpt" d:"0"`
	CourseWorkType    int    `json:"courseWorkType"`
	CourseWork        string `json:"courseWork"`
	CourseWorkAnswer  string `json:"courseWorkAnswer"`
	CoursePlan        string `json:"coursePlan"`
	ShowCoursePlan    int    `json:"showCoursePlan" d:"0"`
	MapX              int    `json:"mapX"`
	MapY              int    `json:"mapY"`
	MediaContent      string `json:"mediaContent"`
	OrderNum          int    `json:"orderNum" d:"1"`
}

type CourseUnitAddRes struct {
	g.Meta `mime:"application/json"`
}

type CourseUnitEditReq struct {
	g.Meta            `path:"/teaching/teachingCourseUnit/edit" method:"put" tags:"课程单元管理" summary:"编辑课程单元"`
	Id                string `json:"id" v:"required#单元ID不能为空"`
	UnitName          string `json:"unitName" v:"required#单元名称不能为空"`
	UnitIntro         string `json:"unitIntro"`
	UnitCover         string `json:"unitCover"`
	CourseId          string `json:"courseId"`
	CourseVideo       string `json:"courseVideo"`
	CourseVideoSource int    `json:"courseVideoSource"`
	ShowCourseVideo   int    `json:"showCourseVideo"`
	CourseCase        string `json:"courseCase"`
	ShowCourseCase    int    `json:"showCourseCase"`
	CoursePpt         string `json:"coursePpt"`
	ShowCoursePpt     int    `json:"showCoursePpt"`
	CourseWorkType    int    `json:"courseWorkType"`
	CourseWork        string `json:"courseWork"`
	CourseWorkAnswer  string `json:"courseWorkAnswer"`
	CoursePlan        string `json:"coursePlan"`
	ShowCoursePlan    int    `json:"showCoursePlan"`
	MapX              int    `json:"mapX"`
	MapY              int    `json:"mapY"`
	MediaContent      string `json:"mediaContent"`
	OrderNum          int    `json:"orderNum"`
}

type CourseUnitEditRes struct {
	g.Meta `mime:"application/json"`
}

type CourseUnitDeleteReq struct {
	g.Meta `path:"/teaching/teachingCourseUnit/delete" method:"delete" tags:"课程单元管理" summary:"删除课程单元"`
	Id     string `json:"id" v:"required#单元ID不能为空"`
}

type CourseUnitDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type CourseUnitDeleteBatchReq struct {
	g.Meta `path:"/teaching/teachingCourseUnit/deleteBatch" method:"delete" tags:"课程单元管理" summary:"批量删除课程单元"`
	Ids    string `json:"ids" v:"required#单元ID不能为空"`
}

type CourseUnitDeleteBatchRes struct {
	g.Meta `mime:"application/json"`
}

type CourseUnitGetByIdReq struct {
	g.Meta `path:"/teaching/teachingCourseUnit/queryById" method:"get" tags:"课程单元管理" summary:"课程单元详情"`
	Id     string `json:"id" v:"required#单元ID不能为空"`
}

type CourseUnitGetByIdRes struct {
	g.Meta `mime:"application/json"`
	*entity.TeachingCourseUnit
}

type CourseUnitSortReq struct {
	g.Meta  `path:"/teaching/teachingCourseUnit/sort" method:"put" tags:"课程单元管理" summary:"课程单元排序"`
	UnitIds string `json:"unitIds" v:"required#单元ID不能为空"`
}

type CourseUnitSortRes struct {
	g.Meta `mime:"application/json"`
}