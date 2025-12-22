package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingCourseUnit 课程单元表
type TeachingCourseUnit struct {
	Id          string      `json:"id"          description:"单元ID"`
	CourseId    string      `json:"courseId"    description:"课程ID"`
	Name        string      `json:"name"        description:"单元名称"`
	Content     string      `json:"content"     description:"单元内容"`
	SortOrder   int         `json:"sortOrder"   description:"排序"`
	ResourceUrl string      `json:"resourceUrl" description:"资源链接"`
	CreateBy    string      `json:"createBy"    description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新时间"`
}