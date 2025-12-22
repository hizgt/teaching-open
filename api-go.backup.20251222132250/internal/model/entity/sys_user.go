// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id         string      `json:"id"         description:"主键id"`
	Username   string      `json:"username"   description:"登录账号"`
	Realname   string      `json:"realname"   description:"真实姓名"`
	Password   string      `json:"password"   description:"密码"`
	Salt       string      `json:"salt"       description:"md5密码盐"`
	Avatar     string      `json:"avatar"     description:"头像"`
	Birthday   *gtime.Time `json:"birthday"   description:"生日"`
	Sex        int         `json:"sex"        description:"性别(0-默认未知,1-男,2-女)"`
	Email      string      `json:"email"      description:"电子邮件"`
	Phone      string      `json:"phone"      description:"电话"`
	OrgCode    string      `json:"orgCode"    description:"机构编码"`
	Status     int         `json:"status"     description:"状态(1-正常,2-冻结)"`
	DelFlag    int         `json:"delFlag"    description:"删除状态(0-正常,1-已删除)"`
	WorkNo     string      `json:"workNo"     description:"工号，唯一键"`
	School     string      `json:"school"     description:"学校"`
	CreateBy   string      `json:"createBy"   description:"创建人"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   description:"更新人"`
	UpdateTime *gtime.Time `json:"updateTime" description:"更新时间"`
}
