// =================================================================================
// Code generated and target of the generated file. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDepart is the golang structure for table sys_depart.
type SysDepart struct {
	Id             string      `json:"id"             description:"ID"`
	ParentId       string      `json:"parentId"       description:"父机构ID"`
	DepartName     string      `json:"departName"     description:"机构/部门名称"`
	DepartNameEn   string      `json:"departNameEn"   description:"英文名"`
	DepartNameAbbr string      `json:"departNameAbbr" description:"缩写"`
	DepartOrder    int         `json:"departOrder"    description:"排序"`
	Description    string      `json:"description"    description:"描述"`
	OrgCategory    string      `json:"orgCategory"    description:"机构类别 1组织机构，2岗位"`
	OrgType        string      `json:"orgType"        description:"机构类型 1一级部门 2子部门"`
	OrgCode        string      `json:"orgCode"        description:"机构编码"`
	Mobile         string      `json:"mobile"         description:"手机号"`
	Fax            string      `json:"fax"            description:"传真"`
	Address        string      `json:"address"        description:"地址"`
	Memo           string      `json:"memo"           description:"备注"`
	Status         string      `json:"status"         description:"状态（1启用，0不启用）"`
	DelFlag        string      `json:"delFlag"        description:"删除状态（0，正常，1已删除）"`
	CreateBy       string      `json:"createBy"       description:"创建人"`
	CreateTime     *gtime.Time `json:"createTime"     description:"创建日期"`
	UpdateBy       string      `json:"updateBy"       description:"更新人"`
	UpdateTime     *gtime.Time `json:"updateTime"     description:"更新日期"`
}
