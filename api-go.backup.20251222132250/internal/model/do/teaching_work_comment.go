// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingWorkComment is the golang structure of table teaching_work_comment for DAO operations.
type TeachingWorkComment struct {
	g.Meta     `orm:"table:teaching_work_comment, do:true"`
	Id         interface{} // 主键
	CreateBy   interface{} // 创建人
	CreateTime *gtime.Time // 创建日期
	UpdateBy   interface{} // 更新人
	UpdateTime *gtime.Time // 更新日期
	SysOrgCode interface{} // 所属部门
	WorkId     interface{} // 作业ID
	Comment    interface{} // 评论内容
	UserId     interface{} // 用户ID
}