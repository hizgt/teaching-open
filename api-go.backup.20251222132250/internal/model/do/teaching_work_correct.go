// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingWorkCorrect is the golang structure of table teaching_work_correct for DAO operations.
type TeachingWorkCorrect struct {
	g.Meta     `orm:"table:teaching_work_correct, do:true"`
	Id         interface{} // 主键
	CreateBy   interface{} // 创建人
	CreateTime *gtime.Time // 创建日期
	UpdateBy   interface{} // 更新人
	UpdateTime *gtime.Time // 更新日期
	SysOrgCode interface{} // 所属部门
	WorkId     interface{} // 作业ID
	Score      interface{} // 评分
	Comment    interface{} // 评语
}