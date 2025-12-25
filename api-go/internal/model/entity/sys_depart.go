package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDepart 部门表
type SysDepart struct {
	Id             string      `json:"id"             orm:"id,primary"        description:"部门ID"`
	ParentId       string      `json:"parentId"       orm:"parent_id"         description:"父ID"`
	DepartName     string      `json:"departName"     orm:"depart_name"       description:"部门名称"`
	DepartNameEn   string      `json:"departNameEn"   orm:"depart_name_en"    description:"英文名称"`
	DepartNameAbbr string      `json:"departNameAbbr" orm:"depart_name_abbr"  description:"缩写"`
	DepartOrder    int         `json:"departOrder"    orm:"depart_order"      description:"排序"`
	Description    string      `json:"description"    orm:"description"       description:"描述"`
	OrgCategory    string      `json:"orgCategory"    orm:"org_category"      description:"机构类别"`
	OrgType        string      `json:"orgType"        orm:"org_type"          description:"机构类型"`
	OrgCode        string      `json:"orgCode"        orm:"org_code"          description:"机构编码"`
	Mobile         string      `json:"mobile"         orm:"mobile"            description:"手机号"`
	Fax            string      `json:"fax"            orm:"fax"               description:"传真"`
	Address        string      `json:"address"        orm:"address"           description:"地址"`
	Memo           string      `json:"memo"           orm:"memo"              description:"备注"`
	Status         string      `json:"status"         orm:"status"            description:"状态"`
	DelFlag        string      `json:"delFlag"        orm:"del_flag"          description:"删除标记"`
	CreateBy       string      `json:"createBy"       orm:"create_by"         description:"创建人"`
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"       description:"创建时间"`
	UpdateBy       string      `json:"updateBy"       orm:"update_by"         description:"更新人"`
	UpdateTime     *gtime.Time `json:"updateTime"     orm:"update_time"       description:"更新时间"`
}

// TableName 表名
func (e *SysDepart) TableName() string {
	return "sys_depart"
}

// SysUserDepart 用户部门关联表
type SysUserDepart struct {
	Id       string `json:"id"       orm:"id,primary" description:"ID"`
	UserId   string `json:"userId"   orm:"user_id"    description:"用户ID"`
	DepId    string `json:"depId"    orm:"dep_id"     description:"部门ID"`
}

// TableName 表名
func (e *SysUserDepart) TableName() string {
	return "sys_user_depart"
}
