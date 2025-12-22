// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDepartDao is the data access object for table sys_depart.
type SysDepartDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SysDepartColumns // columns contains all the column names of Table for convenient usage.
}

// SysDepartColumns defines and stores column names for table sys_depart.
type SysDepartColumns struct {
	Id             string
	ParentId       string
	DepartName     string
	DepartNameEn   string
	DepartNameAbbr string
	DepartOrder    string
	Description    string
	OrgCategory    string
	OrgType        string
	OrgCode        string
	Mobile         string
	Fax            string
	Address        string
	Memo           string
	Status         string
	DelFlag        string
	CreateBy       string
	CreateTime     string
	UpdateBy       string
	UpdateTime     string
}

// sysDepartColumns holds the columns for table sys_depart.
var sysDepartColumns = SysDepartColumns{
	Id:             "id",
	ParentId:       "parent_id",
	DepartName:     "depart_name",
	DepartNameEn:   "depart_name_en",
	DepartNameAbbr: "depart_name_abbr",
	DepartOrder:    "depart_order",
	Description:    "description",
	OrgCategory:    "org_category",
	OrgType:        "org_type",
	OrgCode:        "org_code",
	Mobile:         "mobile",
	Fax:            "fax",
	Address:        "address",
	Memo:           "memo",
	Status:         "status",
	DelFlag:        "del_flag",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
}

// NewSysDepartDao creates and returns a new DAO object for table data access.
func NewSysDepartDao() *SysDepartDao {
	return &SysDepartDao{
		group:   "default",
		table:   "sys_depart",
		columns: sysDepartColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysDepartDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysDepartDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysDepartDao) Columns() SysDepartColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysDepartDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysDepartDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *SysDepartDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
