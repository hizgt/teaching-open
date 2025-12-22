package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TeachingAdditionalWork 附加作业数据对象
type TeachingAdditionalWork struct {
	g.Meta          `orm:"table:teaching_additional_work, do:true"`
	Id              interface{}
	CreateBy        interface{}
	CreateTime      *gtime.Time
	UpdateBy        interface{}
	UpdateTime      *gtime.Time
	SysOrgCode      interface{}
	CodeType        interface{}
	WorkName        interface{}
	WorkDesc        interface{}
	WorkCover       interface{}
	WorkUrl         interface{}
	WorkDept        interface{}
	Status          interface{}
	WorkDocumentUrl interface{}
}
