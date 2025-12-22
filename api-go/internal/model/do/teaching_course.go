// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import "github.com/gogf/gf/v2/frame/g"

// TeachingCourse is the golang structure for table teaching_course.
type TeachingCourse struct {
	g.Meta         `orm:"table:teaching_course, do:true"`
	Id             interface{} // 主键
	CreateBy       interface{} // 创建人
	CreateTime     interface{} // 创建日期
	UpdateBy       interface{} // 更新人
	UpdateTime     interface{} // 更新日期
	SysOrgCode     interface{} // 所属部门
	DelFlag        interface{} // 删除标志
	CourseName     interface{} // 科目名
	CourseDesc     interface{} // 科目介绍
	CourseIcon     interface{} // 科目图标
	CourseCover    interface{} // 科目封面
	ShowType       interface{} // 展示类型
	CourseMap      interface{} // 课程地图
	IsShared       interface{} // 是否共享课程
	ShowHome       interface{} // 首页展示
	OrderNum       interface{} // 排序
	DepartIds      interface{} // 授权部门
	CourseType     interface{} // 课程类型
	CourseCategory interface{} // 课程分类
}
