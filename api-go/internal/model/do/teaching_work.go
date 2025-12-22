// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingWork is the golang structure of table teaching_work for DAO operations.
type TeachingWork struct {
	g.Meta       `orm:"table:teaching_work, do:true"`
	Id           interface{} // 主键
	CreateBy     interface{} // 创建人
	CreateTime   *gtime.Time // 创建日期
	UpdateBy     interface{} // 更新人
	UpdateTime   *gtime.Time // 更新日期
	SysOrgCode   interface{} // 所属部门
	UserId       interface{} // 用户ID
	DepartId     interface{} // 班级ID
	CourseId     interface{} // 课程ID
	WorkName     interface{} // 作业名
	WorkType     interface{} // 作业类型
	WorkFile     interface{} // 作业文件
	WorkCover    interface{} // 作业封面
	WorkStatus   interface{} // 作业状态
	StarNum      interface{} // 点赞次数
	CollectNum   interface{} // 收藏次数
	DelFlag      interface{} // 删除标识
	ViewNum      interface{} // 查看次数
	AdditionalId interface{} // 附加作业ID
	WorkScene    interface{} // 来源场景
	HasCloudData interface{} // 是否包含云变量
}