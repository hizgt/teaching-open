package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser 用户表
type SysUser struct {
	Id           string      `json:"id"           orm:"id,primary"      description:"用户ID"`
	Username     string      `json:"username"     orm:"username"        description:"用户名"`
	Realname     string      `json:"realname"     orm:"realname"        description:"真实姓名"`
	Password     string      `json:"-"            orm:"password"        description:"密码"`
	Salt         string      `json:"-"            orm:"salt"            description:"盐值"`
	Avatar       string      `json:"avatar"       orm:"avatar"          description:"头像"`
	Birthday     *gtime.Time `json:"birthday"     orm:"birthday"        description:"生日"`
	Sex          int         `json:"sex"          orm:"sex"             description:"性别(1男2女)"`
	Email        string      `json:"email"        orm:"email"           description:"邮箱"`
	Phone        string      `json:"phone"        orm:"phone"           description:"手机号"`
	OrgCode      string      `json:"orgCode"      orm:"org_code"        description:"组织机构编码"`
	Status       int         `json:"status"       orm:"status"          description:"状态(1正常2冻结)"`
	DelFlag      int         `json:"delFlag"      orm:"del_flag"        description:"删除标记(0正常1删除)"`
	WorkNo       string      `json:"workNo"       orm:"work_no"         description:"工号"`
	Post         string      `json:"post"         orm:"post"            description:"职务"`
	School       string      `json:"school"       orm:"school"          description:"学校"`
	Telephone    string      `json:"telephone"    orm:"telephone"       description:"座机号"`
	DepartIds    string      `json:"departIds"    orm:"depart_ids"      description:"部门ID列表"`
	ThirdId      string      `json:"thirdId"      orm:"third_id"        description:"第三方ID"`
	ThirdType    string      `json:"thirdType"    orm:"third_type"      description:"第三方类型"`
	UserIdentity int         `json:"userIdentity" orm:"user_identity"   description:"用户身份(1普通2上级)"`
	ActivitiSync int         `json:"activitiSync" orm:"activiti_sync"   description:"同步工作流"`
	CreateBy     string      `json:"createBy"     orm:"create_by"       description:"创建人"`
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"     description:"创建时间"`
	UpdateBy     string      `json:"updateBy"     orm:"update_by"       description:"更新人"`
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"update_time"     description:"更新时间"`
}

// TableName 表名
func (e *SysUser) TableName() string {
	return "sys_user"
}
