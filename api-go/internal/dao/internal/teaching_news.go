// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingNewsDao is the data access object for table teaching_news.
type TeachingNewsDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TeachingNewsColumns // columns contains all the column names of Table for convenient usage.
}

// TeachingNewsColumns defines and stores column names for table teaching_news.
type TeachingNewsColumns struct {
	Id          string // 主键
	NewsTitle   string // 标题
	NewsContent string // 内容
	NewsStatus  string // 状态
	CreateBy    string // 创建人
	CreateTime  string // 创建日期
	UpdateBy    string // 更新人
	UpdateTime  string // 更新日期
}

// teachingNewsColumns holds the columns for table teaching_news.
var teachingNewsColumns = TeachingNewsColumns{
	Id:          "id",
	NewsTitle:   "news_title",
	NewsContent: "news_content",
	NewsStatus:  "news_status",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
}

// NewTeachingNewsDao creates and returns a new DAO object for table data access.
func NewTeachingNewsDao() *TeachingNewsDao {
	return &TeachingNewsDao{
		group:   "default",
		table:   "teaching_news",
		columns: teachingNewsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeachingNewsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current DAO.
func (dao *TeachingNewsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current DAO.
func (dao *TeachingNewsDao) Columns() TeachingNewsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current DAO.
func (dao *TeachingNewsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for current DAO. It automatically sets the context for current operation.
func (dao *TeachingNewsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *TeachingNewsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
