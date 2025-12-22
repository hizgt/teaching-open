// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingNews is the golang structure for table teaching_news.
type TeachingNews struct {
	Id          string      `json:"id"          description:"主键"`
	NewsTitle   string      `json:"newsTitle"   description:"标题"`
	NewsContent string      `json:"newsContent" description:"内容"`
	NewsStatus  int         `json:"newsStatus"  description:"状态 0草稿 1发布"`
	CreateBy    string      `json:"createBy"    description:"创建人"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建日期"`
	UpdateBy    string      `json:"updateBy"    description:"更新人"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新日期"`
}
