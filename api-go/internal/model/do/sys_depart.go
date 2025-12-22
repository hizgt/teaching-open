// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDepart is the golang structure of table sys_depart for DAO operations like Where/Data.
type SysDepart struct {
	g.Meta         `orm:"table:sys_depart, do:true"`
	Id             interface{}
	ParentId       interface{}
	DepartName     interface{}
	DepartNameEn   interface{}
	DepartNameAbbr interface{}
	DepartOrder    interface{}
	Description    interface{}
	OrgCategory    interface{}
	OrgType        interface{}
	OrgCode        interface{}
	Mobile         interface{}
	Fax            interface{}
	Address        interface{}
	Memo           interface{}
	Status         interface{}
	DelFlag        interface{}
	CreateBy       interface{}
	CreateTime     *gtime.Time
	UpdateBy       interface{}
	UpdateTime     *gtime.Time
}
