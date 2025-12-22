package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingCourse 课程表
type TeachingCourse struct {
	Id          string      `json:"id"          description:"课程ID"`
	Name        string      `json:"name"        description:"课程名称"`
	Type        string      `json:"type"        description:"课程类型(scratch/python)"`
	Description string      `json:"description" description:"课程描述"`
	CoverImage  string      `json:"coverImage"  description:"封面图"`
	Status      string      `json:"status"      description:"状态(draft/published)"`
	CreateBy    string      `json:"createBy"    description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateBy    string      `json:"updateBy"    description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新时间"`
}