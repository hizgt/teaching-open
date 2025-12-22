// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingScratchAssetDao is the data access object for table teaching_scratch_assets.
type TeachingScratchAssetDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns TeachingScratchAssetColumns // columns contains all the column names of Table for convenient usage.
}

// TeachingScratchAssetColumns defines and stores column names for table teaching_scratch_assets.
type TeachingScratchAssetColumns struct {
	Id         string // 主键ID
	AssetType  string // 素材类型 1背景 2声音 3造型 4角色
	AssetName  string // 素材名
	AssetData  string // 素材JSON数据
	Md5Ext     string // 素材md5
	Tags       string // 标签
	CreateBy   string // 创建人
	CreateTime string // 创建时间
	UpdateBy   string // 修改人
	UdpateTime string // 修改时间
	DelFlag    string // 删除状态 0正常 1删除
}

// teachingScratchAssetColumns holds the columns for table teaching_scratch_assets.
var teachingScratchAssetColumns = TeachingScratchAssetColumns{
	Id:         "id",
	AssetType:  "asset_type",
	AssetName:  "asset_name",
	AssetData:  "asset_data",
	Md5Ext:     "md5_ext",
	Tags:       "tags",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UdpateTime: "udpate_time",
	DelFlag:    "del_flag",
}

// NewTeachingScratchAssetDao creates and returns a new DAO object for table data access.
func NewTeachingScratchAssetDao() *TeachingScratchAssetDao {
	return &TeachingScratchAssetDao{
		group:   "default",
		table:   "teaching_scratch_assets",
		columns: teachingScratchAssetColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TeachingScratchAssetDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TeachingScratchAssetDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TeachingScratchAssetDao) Columns() TeachingScratchAssetColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TeachingScratchAssetDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TeachingScratchAssetDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
func (dao *TeachingScratchAssetDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
