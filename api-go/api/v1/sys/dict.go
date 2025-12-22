// =================================================================================
// API definitions for sys dict module
// =================================================================================

package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DictInfo 字典信息
type DictInfo struct {
	Id          string `json:"id"`
	DictName    string `json:"dictName"`
	DictCode    string `json:"dictCode"`
	Description string `json:"description"`
	Type        int    `json:"type"`
	CreateTime  string `json:"createTime,omitempty"`
}

// DictItemInfo 字典项信息
type DictItemInfo struct {
	Id          string `json:"id"`
	DictId      string `json:"dictId"`
	ItemText    string `json:"itemText"`
	ItemValue   string `json:"itemValue"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder"`
	Status      int    `json:"status"`
}

// ========== 字典管理 ==========

// DictListReq 字典列表请求
type DictListReq struct {
	g.Meta   `path:"/sys/dict/list" method:"get" tags:"字典管理" summary:"字典列表"`
	Page     int    `json:"page" d:"1" v:"min:1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" v:"min:1|max:100" dc:"每页数量"`
	DictName string `json:"dictName" dc:"字典名称"`
	DictCode string `json:"dictCode" dc:"字典编码"`
}

// DictListRes 字典列表响应
type DictListRes struct {
	List     []*DictInfo `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// DictAddReq 添加字典请求
type DictAddReq struct {
	g.Meta      `path:"/sys/dict/add" method:"post" tags:"字典管理" summary:"添加字典"`
	DictName    string `json:"dictName" v:"required#字典名称不能为空"`
	DictCode    string `json:"dictCode" v:"required#字典编码不能为空"`
	Description string `json:"description"`
	Type        int    `json:"type" d:"0" dc:"字典类型0为string,1为number"`
}

// DictAddRes 添加字典响应
type DictAddRes struct{}

// DictEditReq 编辑字典请求
type DictEditReq struct {
	g.Meta      `path:"/sys/dict/edit" method:"put" tags:"字典管理" summary:"编辑字典"`
	Id          string `json:"id" v:"required#字典ID不能为空"`
	DictName    string `json:"dictName"`
	DictCode    string `json:"dictCode"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}

// DictEditRes 编辑字典响应
type DictEditRes struct{}

// DictDeleteReq 删除字典请求
type DictDeleteReq struct {
	g.Meta `path:"/sys/dict/delete" method:"delete" tags:"字典管理" summary:"删除字典"`
	Id     string `json:"id" v:"required#字典ID不能为空"`
}

// DictDeleteRes 删除字典响应
type DictDeleteRes struct{}

// DictGetByIdReq 获取字典详情请求
type DictGetByIdReq struct {
	g.Meta `path:"/sys/dict/queryById" method:"get" tags:"字典管理" summary:"字典详情"`
	Id     string `json:"id" v:"required#字典ID不能为空"`
}

// DictGetByIdRes 获取字典详情响应
type DictGetByIdRes struct {
	*DictInfo
}

// ========== 字典项管理 ==========

// DictItemListReq 字典项列表请求
type DictItemListReq struct {
	g.Meta `path:"/sys/dictItem/list" method:"get" tags:"字典管理" summary:"字典项列表"`
	DictId string `json:"dictId" v:"required#字典ID不能为空"`
}

// DictItemListRes 字典项列表响应
type DictItemListRes struct {
	List []*DictItemInfo `json:"list"`
}

// DictItemsByCodeReq 根据字典编码获取字典项
type DictItemsByCodeReq struct {
	g.Meta   `path:"/sys/dict/getDictItems/:dictCode" method:"get" tags:"字典管理" summary:"获取字典项"`
	DictCode string `json:"dictCode" in:"path" v:"required#字典编码不能为空"`
}

// DictItemsByCodeRes 根据字典编码获取字典项响应
type DictItemsByCodeRes struct {
	List []*DictItemInfo `json:"list"`
}

// DictItemAddReq 添加字典项请求
type DictItemAddReq struct {
	g.Meta      `path:"/sys/dictItem/add" method:"post" tags:"字典管理" summary:"添加字典项"`
	DictId      string `json:"dictId" v:"required#字典ID不能为空"`
	ItemText    string `json:"itemText" v:"required#字典项文本不能为空"`
	ItemValue   string `json:"itemValue" v:"required#字典项值不能为空"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder" d:"1"`
	Status      int    `json:"status" d:"1" dc:"状态1启用0禁用"`
}

// DictItemAddRes 添加字典项响应
type DictItemAddRes struct{}

// DictItemEditReq 编辑字典项请求
type DictItemEditReq struct {
	g.Meta      `path:"/sys/dictItem/edit" method:"put" tags:"字典管理" summary:"编辑字典项"`
	Id          string `json:"id" v:"required#字典项ID不能为空"`
	ItemText    string `json:"itemText"`
	ItemValue   string `json:"itemValue"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder"`
	Status      int    `json:"status"`
}

// DictItemEditRes 编辑字典项响应
type DictItemEditRes struct{}

// DictItemDeleteReq 删除字典项请求
type DictItemDeleteReq struct {
	g.Meta `path:"/sys/dictItem/delete" method:"delete" tags:"字典管理" summary:"删除字典项"`
	Id     string `json:"id" v:"required#字典项ID不能为空"`
}

// DictItemDeleteRes 删除字典项响应
type DictItemDeleteRes struct{}
