// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingWork is the golang structure for table teaching_work.
type TeachingWork struct {
	Id           string      `json:"id"           description:"主键"`
	CreateBy     string      `json:"createBy"     description:"创建人"`
	CreateTime   *gtime.Time `json:"createTime"   description:"创建日期"`
	UpdateBy     string      `json:"updateBy"     description:"更新人"`
	UpdateTime   *gtime.Time `json:"updateTime"   description:"更新日期"`
	SysOrgCode   string      `json:"sysOrgCode"   description:"所属部门"`
	UserId       string      `json:"userId"       description:"用户ID"`
	DepartId     string      `json:"departId"     description:"班级ID"`
	CourseId     string      `json:"courseId"     description:"课程ID"`
	WorkName     string      `json:"workName"     description:"作业名"`
	WorkType     string      `json:"workType"     description:"作业类型"`
	WorkFile     string      `json:"workFile"     description:"作业文件"`
	WorkCover    string      `json:"workCover"    description:"作业封面"`
	WorkStatus   int         `json:"workStatus"   description:"作业状态"`
	StarNum      int         `json:"starNum"      description:"点赞次数"`
	CollectNum   int         `json:"collectNum"   description:"收藏次数"`
	DelFlag      int         `json:"delFlag"      description:"删除标识"`
	ViewNum      int         `json:"viewNum"      description:"查看次数"`
	AdditionalId string      `json:"additionalId" description:"附加作业ID"`
	WorkScene    string      `json:"workScene"    description:"来源场景"`
	HasCloudData int         `json:"hasCloudData" description:"是否包含云变量"`
}