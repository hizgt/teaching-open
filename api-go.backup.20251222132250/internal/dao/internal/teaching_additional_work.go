package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TeachingAdditionalWorkDao 附加作业DAO内部结构
type TeachingAdditionalWorkDao struct {
	table   string
	group   string
	columns TeachingAdditionalWorkColumns
}

// TeachingAdditionalWorkColumns 附加作业表字段
type TeachingAdditionalWorkColumns struct {
	Id              string
	CreateBy        string
	CreateTime      string
	UpdateBy        string
	UpdateTime      string
	SysOrgCode      string
	CodeType        string
	WorkName        string
	WorkDesc        string
	WorkCover       string
	WorkUrl         string
	WorkDept        string
	Status          string
	WorkDocumentUrl string
}

var teachingAdditionalWorkColumns = TeachingAdditionalWorkColumns{
	Id:              "id",
	CreateBy:        "create_by",
	CreateTime:      "create_time",
	UpdateBy:        "update_by",
	UpdateTime:      "update_time",
	SysOrgCode:      "sys_org_code",
	CodeType:        "code_type",
	WorkName:        "work_name",
	WorkDesc:        "work_desc",
	WorkCover:       "work_cover",
	WorkUrl:         "work_url",
	WorkDept:        "work_dept",
	Status:          "status",
	WorkDocumentUrl: "work_document_url",
}

func NewTeachingAdditionalWorkDao() *TeachingAdditionalWorkDao {
	return &TeachingAdditionalWorkDao{
		table:   "teaching_additional_work",
		group:   "default",
		columns: teachingAdditionalWorkColumns,
	}
}

func (d *TeachingAdditionalWorkDao) DB() gdb.DB {
	return g.DB(d.group)
}

func (d *TeachingAdditionalWorkDao) Table() string {
	return d.table
}

func (d *TeachingAdditionalWorkDao) Columns() TeachingAdditionalWorkColumns {
	return d.columns
}

func (d *TeachingAdditionalWorkDao) Group() string {
	return d.group
}

func (d *TeachingAdditionalWorkDao) Ctx(ctx context.Context) *gdb.Model {
	return d.DB().Model(d.table).Safe().Ctx(ctx)
}

func (d *TeachingAdditionalWorkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) error {
	return d.DB().Transaction(ctx, f)
}
