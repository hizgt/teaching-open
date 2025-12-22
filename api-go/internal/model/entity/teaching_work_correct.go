// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingWorkCorrect is the golang structure for table teaching_work_correct.
type TeachingWorkCorrect struct {
	Id         string      `json:"id"         description:"主键"`
	CreateBy   string      `json:"createBy"   description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" description:"创建日期"`
	UpdateBy   string      `json:"updateBy"   description:"更新人"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新日期"`
	SysOrgCode string      `json:"sysOrgCode" description:"所属部门"`
	WorkId     string      `json:"workId"     description:"作业ID"`
	Score      int         `json:"score"      description:"评分"`
	Comment    string      `json:"comment"    description:"评语"`
}