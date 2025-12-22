// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingWorkDao is the data access object for table teaching_work.
type TeachingWorkDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TeachingWorkColumns // columns contains all the column names of Table for convenient usage.
}

// TeachingWorkColumns defines and stores column names for table teaching_work.
type TeachingWorkColumns struct {
	Id           string // 主键
	CreateBy     string // 创建人
	CreateTime   string // 创建日期
	UpdateBy     string // 更新人
	UpdateTime   string // 更新日期
	SysOrgCode   string // 所属部门
	UserId       string // 用户ID
	DepartId     string // 班级ID
	CourseId     string // 课程ID
	WorkName     string // 作业名
	WorkType     string // 作业类型
	WorkFile     string // 作业文件
	WorkCover    string // 作业封面
	WorkStatus   string // 作业状态
	StarNum      string // 点赞次数
	CollectNum   string // 收藏次数
	DelFlag      string // 删除标识
	ViewNum      string // 查看次数
	AdditionalId string // 附加作业ID
	WorkScene    string // 来源场景
	HasCloudData string // 是否包含云变量
}

// teachingWorkColumns holds the columns for table teaching_work.
var teachingWorkColumns = TeachingWorkColumns{
	Id:           "id",
	CreateBy:     "create_by",
	CreateTime:   "create_time",
	UpdateBy:     "update_by",
	UpdateTime:   "update_time",
	SysOrgCode:   "sys_org_code",
	UserId:       "user_id",
	DepartId:     "depart_id",
	CourseId:     "course_id",
	WorkName:     "work_name",
	WorkType:     "work_type",
	WorkFile:     "work_file",
	WorkCover:    "work_cover",
	WorkStatus:   "work_status",
	StarNum:      "star_num",
	CollectNum:   "collect_num",
	DelFlag:      "del_flag",
	ViewNum:      "view_num",
	AdditionalId: "additional_id",
	WorkScene:    "work_scene",
	HasCloudData: "has_cloud_data",
}

// NewTeachingWorkDao creates and returns a new DAO object for table data access.
func NewTeachingWorkDao() *TeachingWorkDao {
	return &TeachingWorkDao{
		group:   "default",
		table:   "teaching_work",
		columns: teachingWorkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeachingWorkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current DAO.
func (dao *TeachingWorkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current DAO.
func (dao *TeachingWorkDao) Columns() TeachingWorkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current DAO.
func (dao *TeachingWorkDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for current DAO. It automatically sets the context for current operation.
func (dao *TeachingWorkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *TeachingWorkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}