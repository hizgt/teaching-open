package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingWork 作业表
type TeachingWork struct {
	Id           string      `json:"id"           orm:"id,primary"      description:"作业ID"`
	UserId       string      `json:"userId"       orm:"user_id"         description:"用户ID"`
	DepartId     string      `json:"departId"     orm:"depart_id"       description:"部门ID"`
	CourseId     string      `json:"courseId"     orm:"course_id"       description:"课程ID"`
	AdditionalId string      `json:"additionalId" orm:"additional_id"   description:"附加作业ID"`
	WorkName     string      `json:"workName"     orm:"work_name"       description:"作业名称"`
	WorkType     string      `json:"workType"     orm:"work_type"       description:"作业类型"`
	WorkStatus   int         `json:"workStatus"   orm:"work_status"     description:"作业状态"`
	WorkFile     string      `json:"workFile"     orm:"work_file"       description:"作业文件"`
	WorkCover    string      `json:"workCover"    orm:"work_cover"      description:"作业封面"`
	ViewNum      int         `json:"viewNum"      orm:"view_num"        description:"浏览数"`
	StarNum      int         `json:"starNum"      orm:"star_num"        description:"点赞数"`
	CollectNum   int         `json:"collectNum"   orm:"collect_num"     description:"收藏数"`
	DelFlag      int         `json:"delFlag"      orm:"del_flag"        description:"删除标记"`
	WorkScene    string      `json:"workScene"    orm:"work_scene"      description:"作业场景"`
	HasCloudData bool        `json:"hasCloudData" orm:"has_cloud_data"  description:"是否有云端数据"`
	SysOrgCode   string      `json:"sysOrgCode"   orm:"sys_org_code"    description:"组织编码"`
	CreateBy     string      `json:"createBy"     orm:"create_by"       description:"创建人"`
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"     description:"创建时间"`
	UpdateBy     string      `json:"updateBy"     orm:"update_by"       description:"更新人"`
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"update_time"     description:"更新时间"`
}

// TableName 表名
func (e *TeachingWork) TableName() string {
	return "teaching_work"
}

// TeachingWorkComment 作业评论表
type TeachingWorkComment struct {
	Id         string      `json:"id"         orm:"id,primary"    description:"评论ID"`
	WorkId     string      `json:"workId"     orm:"work_id"       description:"作业ID"`
	UserId     string      `json:"userId"     orm:"user_id"       description:"用户ID"`
	Content    string      `json:"content"    orm:"content"       description:"评论内容"`
	CreateBy   string      `json:"createBy"   orm:"create_by"     description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time"   description:"创建时间"`
}

// TableName 表名
func (e *TeachingWorkComment) TableName() string {
	return "teaching_work_comment"
}

// TeachingWorkCorrect 作业批改表
type TeachingWorkCorrect struct {
	Id         string      `json:"id"         orm:"id,primary"    description:"批改ID"`
	WorkId     string      `json:"workId"     orm:"work_id"       description:"作业ID"`
	UserId     string      `json:"userId"     orm:"user_id"       description:"批改人ID"`
	Score      int         `json:"score"      orm:"score"         description:"分数"`
	Comment    string      `json:"comment"    orm:"comment"       description:"批改评语"`
	CreateBy   string      `json:"createBy"   orm:"create_by"     description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time"   description:"创建时间"`
}

// TableName 表名
func (e *TeachingWorkCorrect) TableName() string {
	return "teaching_work_correct"
}

// TeachingAdditionalWork 附加作业表
type TeachingAdditionalWork struct {
	Id         string      `json:"id"         orm:"id,primary"    description:"附加作业ID"`
	WorkName   string      `json:"workName"   orm:"work_name"     description:"作业名称"`
	WorkDesc   string      `json:"workDesc"   orm:"work_desc"     description:"作业描述"`
	WorkFile   string      `json:"workFile"   orm:"work_file"     description:"作业文件"`
	DepartId   string      `json:"departId"   orm:"depart_id"     description:"部门ID"`
	CreateBy   string      `json:"createBy"   orm:"create_by"     description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time"   description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   orm:"update_by"     description:"更新人"`
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time"   description:"更新时间"`
}

// TableName 表名
func (e *TeachingAdditionalWork) TableName() string {
	return "teaching_additional_work"
}
