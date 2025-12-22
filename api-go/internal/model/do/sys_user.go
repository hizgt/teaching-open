// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta     `orm:"table:sys_user, do:true"`
	Id         interface{} `orm:"id"          description:"主键id"`
	Username   interface{} `orm:"username"    description:"登录账号"`
	Realname   interface{} `orm:"realname"    description:"真实姓名"`
	Password   interface{} `orm:"password"    description:"密码"`
	Salt       interface{} `orm:"salt"        description:"md5密码盐"`
	Avatar     interface{} `orm:"avatar"      description:"头像"`
	Birthday   *gtime.Time `orm:"birthday"    description:"生日"`
	Sex        interface{} `orm:"sex"         description:"性别(0-默认未知,1-男,2-女)"`
	Email      interface{} `orm:"email"       description:"电子邮件"`
	Phone      interface{} `orm:"phone"       description:"电话"`
	OrgCode    interface{} `orm:"org_code"    description:"机构编码"`
	Status     interface{} `orm:"status"      description:"状态(1-正常,2-冻结)"`
	DelFlag    interface{} `orm:"del_flag"    description:"删除状态(0-正常,1-已删除)"`
	WorkNo     interface{} `orm:"work_no"     description:"工号，唯一键"`
	School     interface{} `orm:"school"      description:"学校"`
	CreateBy   interface{} `orm:"create_by"   description:"创建人"`
	CreateTime *gtime.Time `orm:"create_time" description:"创建时间"`
	UpdateBy   interface{} `orm:"update_by"   description:"更新人"`
	UpdateTime *gtime.Time `orm:"update_time" description:"更新时间"`
}
