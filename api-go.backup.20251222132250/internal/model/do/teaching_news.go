// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingNews is the golang structure of table teaching_news for DAO operations.
type TeachingNews struct {
	g.Meta      `orm:"table:teaching_news, do:true"`
	Id          interface{} // 主键
	NewsTitle   interface{} // 标题
	NewsContent interface{} // 内容
	NewsStatus  interface{} // 状态 0草稿 1发布
	CreateBy    interface{} // 创建人
	CreateTime  *gtime.Time // 创建日期
	UpdateBy    interface{} // 更新人
	UpdateTime  *gtime.Time // 更新日期
}
