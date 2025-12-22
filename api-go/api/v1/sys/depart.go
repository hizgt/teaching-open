// =================================================================================
// API definitions for sys depart module
// =================================================================================

package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DepartInfo 部门信息
type DepartInfo struct {
	Id             string        `json:"id"`
	ParentId       string        `json:"parentId"`
	DepartName     string        `json:"departName"`
	DepartNameEn   string        `json:"departNameEn"`
	DepartNameAbbr string        `json:"departNameAbbr"`
	DepartOrder    int           `json:"departOrder"`
	Description    string        `json:"description"`
	OrgCategory    string        `json:"orgCategory"`
	OrgType        string        `json:"orgType"`
	OrgCode        string        `json:"orgCode"`
	Mobile         string        `json:"mobile"`
	Fax            string        `json:"fax"`
	Address        string        `json:"address"`
	Memo           string        `json:"memo"`
	Status         string        `json:"status"`
	CreateTime     string        `json:"createTime"`
	Children       []*DepartInfo `json:"children,omitempty"`
}

// DepartTreeReq 部门树请求
type DepartTreeReq struct {
	g.Meta `path:"/sys/sysDepart/queryTreeList" method:"get" tags:"部门管理" summary:"获取部门树"`
}

// DepartTreeRes 部门树响应
type DepartTreeRes struct {
	List []*DepartInfo `json:"list"`
}

// DepartAddReq 添加部门请求
type DepartAddReq struct {
	g.Meta         `path:"/sys/sysDepart/add" method:"post" tags:"部门管理" summary:"添加部门"`
	ParentId       string `json:"parentId"`
	DepartName     string `json:"departName" v:"required#部门名称不能为空"`
	DepartNameEn   string `json:"departNameEn"`
	DepartNameAbbr string `json:"departNameAbbr"`
	DepartOrder    int    `json:"departOrder"`
	Description    string `json:"description"`
	OrgCategory    string `json:"orgCategory" d:"1"`
	OrgType        string `json:"orgType"`
	OrgCode        string `json:"orgCode"`
	Mobile         string `json:"mobile"`
	Fax            string `json:"fax"`
	Address        string `json:"address"`
	Memo           string `json:"memo"`
	Status         string `json:"status" d:"1"`
}

// DepartAddRes 添加部门响应
type DepartAddRes struct{}

// DepartEditReq 编辑部门请求
type DepartEditReq struct {
	g.Meta         `path:"/sys/sysDepart/edit" method:"put" tags:"部门管理" summary:"编辑部门"`
	Id             string `json:"id" v:"required#部门ID不能为空"`
	ParentId       string `json:"parentId"`
	DepartName     string `json:"departName" v:"required#部门名称不能为空"`
	DepartNameEn   string `json:"departNameEn"`
	DepartNameAbbr string `json:"departNameAbbr"`
	DepartOrder    int    `json:"departOrder"`
	Description    string `json:"description"`
	OrgCategory    string `json:"orgCategory"`
	OrgType        string `json:"orgType"`
	OrgCode        string `json:"orgCode"`
	Mobile         string `json:"mobile"`
	Fax            string `json:"fax"`
	Address        string `json:"address"`
	Memo           string `json:"memo"`
	Status         string `json:"status"`
}

// DepartEditRes 编辑部门响应
type DepartEditRes struct{}

// DepartDeleteReq 删除部门请求
type DepartDeleteReq struct {
	g.Meta `path:"/sys/sysDepart/delete" method:"delete" tags:"部门管理" summary:"删除部门"`
	Id     string `json:"id" v:"required#部门ID不能为空"`
}

// DepartDeleteRes 删除部门响应
type DepartDeleteRes struct{}

// DepartGetReq 获取部门详情请求
type DepartGetReq struct {
	g.Meta `path:"/sys/sysDepart/getById" method:"get" tags:"部门管理" summary:"获取部门详情"`
	Id     string `json:"id" v:"required#部门ID不能为空"`
}

// DepartGetRes 获取部门详情响应
type DepartGetRes struct {
	*DepartInfo
}

// DepartIdTreeReq 部门ID树请求
type DepartIdTreeReq struct {
	g.Meta `path:"/sys/sysDepart/queryIdTree" method:"get" tags:"部门管理" summary:"获取部门ID树"`
}

// DepartIdTreeRes 部门ID树响应
type DepartIdTreeRes struct {
	List []*DepartIdNode `json:"list"`
}

// DepartIdNode 部门ID节点
type DepartIdNode struct {
	Key      string          `json:"key"`
	Value    string          `json:"value"`
	Title    string          `json:"title"`
	Children []*DepartIdNode `json:"children,omitempty"`
}

// DepartSearchReq 部门搜索请求
type DepartSearchReq struct {
	g.Meta  `path:"/sys/sysDepart/searchBy" method:"get" tags:"部门管理" summary:"部门关键字搜索"`
	Keyword string `json:"keyword"`
}

// DepartSearchRes 部门搜索响应
type DepartSearchRes struct {
	List []*DepartInfo `json:"list"`
}

// UserDepartReq 获取用户部门请求
type UserDepartReq struct {
	g.Meta `path:"/sys/sysDepart/getUserDeparts" method:"get" tags:"部门管理" summary:"获取用户部门"`
	UserId string `json:"userId" v:"required#用户ID不能为空"`
}

// UserDepartRes 获取用户部门响应
type UserDepartRes struct {
	DepartIds []string      `json:"departIds"`
	Departs   []*DepartInfo `json:"departs"`
}

// SaveUserDepartReq 保存用户部门请求
type SaveUserDepartReq struct {
	g.Meta    `path:"/sys/sysDepart/saveUserDepart" method:"post" tags:"部门管理" summary:"保存用户部门"`
	UserId    string   `json:"userId" v:"required#用户ID不能为空"`
	DepartIds []string `json:"departIds"`
}

// SaveUserDepartRes 保存用户部门响应
type SaveUserDepartRes struct{}
