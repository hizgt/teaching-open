// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingCourseDeptDao is the data access object for table teaching_course_dept.
type TeachingCourseDeptDao struct {
	table   string
	group   string
	columns TeachingCourseDeptColumns
}

// TeachingCourseDeptColumns defines and stores column names for table teaching_course_dept.
type TeachingCourseDeptColumns struct {
	Id         string
	CreateBy   string
	CreateTime string
	UpdateBy   string
	UpdateTime string
	SysOrgCode string
	DeptId     string
	CourseId   string
	OpenTime   string
}

var teachingCourseDeptColumns = TeachingCourseDeptColumns{
	Id:         "id",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
	SysOrgCode: "sys_org_code",
	DeptId:     "dept_id",
	CourseId:   "course_id",
	OpenTime:   "open_time",
}

func NewTeachingCourseDeptDao() *TeachingCourseDeptDao {
	return &TeachingCourseDeptDao{
		group:   "default",
		table:   "teaching_course_dept",
		columns: teachingCourseDeptColumns,
	}
}

func (dao *TeachingCourseDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

func (dao *TeachingCourseDeptDao) Table() string {
	return dao.table
}

func (dao *TeachingCourseDeptDao) Columns() TeachingCourseDeptColumns {
	return dao.columns
}

func (dao *TeachingCourseDeptDao) Group() string {
	return dao.group
}

func (dao *TeachingCourseDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

func (dao *TeachingCourseDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
