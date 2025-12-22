package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 新闻公告管理 ====================

// NewsListReq 新闻列表请求
type NewsListReq struct {
	g.Meta     `path:"/teaching/teachingNews/list" method:"get" tags:"新闻公告" summary:"新闻列表"`
	NewsTitle  string `json:"newsTitle" dc:"标题关键词"`
	NewsStatus int    `json:"newsStatus" dc:"状态 -1全部 0草稿 1发布" d:"-1"`
	PageNo     int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize   int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// NewsListRes 新闻列表响应
type NewsListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}

// NewsAddReq 添加新闻请求
type NewsAddReq struct {
	g.Meta      `path:"/teaching/teachingNews/add" method:"post" tags:"新闻公告" summary:"添加新闻"`
	NewsTitle   string `json:"newsTitle" v:"required#标题不能为空" dc:"标题"`
	NewsContent string `json:"newsContent" dc:"内容"`
	NewsStatus  int    `json:"newsStatus" dc:"状态 0草稿 1发布" d:"0"`
}

// NewsAddRes 添加新闻响应
type NewsAddRes struct {
	g.Meta `mime:"application/json"`
	Id     string `json:"id"`
}

// NewsEditReq 编辑新闻请求
type NewsEditReq struct {
	g.Meta      `path:"/teaching/teachingNews/edit" method:"put" tags:"新闻公告" summary:"编辑新闻"`
	Id          string `json:"id" v:"required#ID不能为空" dc:"新闻ID"`
	NewsTitle   string `json:"newsTitle" dc:"标题"`
	NewsContent string `json:"newsContent" dc:"内容"`
	NewsStatus  int    `json:"newsStatus" dc:"状态" d:"-1"`
}

// NewsEditRes 编辑新闻响应
type NewsEditRes struct {
	g.Meta `mime:"application/json"`
}

// NewsDeleteReq 删除新闻请求
type NewsDeleteReq struct {
	g.Meta `path:"/teaching/teachingNews/delete" method:"delete" tags:"新闻公告" summary:"删除新闻"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"新闻ID"`
}

// NewsDeleteRes 删除新闻响应
type NewsDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// NewsDeleteBatchReq 批量删除新闻请求
type NewsDeleteBatchReq struct {
	g.Meta `path:"/teaching/teachingNews/deleteBatch" method:"delete" tags:"新闻公告" summary:"批量删除新闻"`
	Ids    string `json:"ids" v:"required#IDs不能为空" dc:"新闻ID列表，逗号分隔"`
}

// NewsDeleteBatchRes 批量删除新闻响应
type NewsDeleteBatchRes struct {
	g.Meta `mime:"application/json"`
}

// NewsGetByIdReq 新闻详情请求
type NewsGetByIdReq struct {
	g.Meta `path:"/teaching/teachingNews/queryById" method:"get" tags:"新闻公告" summary:"新闻详情"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"新闻ID"`
}

// NewsGetByIdRes 新闻详情响应
type NewsGetByIdRes struct {
	g.Meta `mime:"application/json"`
	*NewsInfo
}

// NewsInfo 新闻信息
type NewsInfo struct {
	Id          string      `json:"id"`
	NewsTitle   string      `json:"newsTitle"`
	NewsContent string      `json:"newsContent"`
	NewsStatus  int         `json:"newsStatus"`
	CreateBy    string      `json:"createBy"`
	CreateTime  interface{} `json:"createTime"`
	UpdateBy    string      `json:"updateBy"`
	UpdateTime  interface{} `json:"updateTime"`
}

// NewsPublishReq 发布新闻请求
type NewsPublishReq struct {
	g.Meta `path:"/teaching/teachingNews/publish" method:"put" tags:"新闻公告" summary:"发布新闻"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"新闻ID"`
}

// NewsPublishRes 发布新闻响应
type NewsPublishRes struct {
	g.Meta `mime:"application/json"`
}

// NewsOfflineReq 下架新闻请求
type NewsOfflineReq struct {
	g.Meta `path:"/teaching/teachingNews/offline" method:"put" tags:"新闻公告" summary:"下架新闻"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"新闻ID"`
}

// NewsOfflineRes 下架新闻响应
type NewsOfflineRes struct {
	g.Meta `mime:"application/json"`
}

// NewsPublicListReq 公开新闻列表请求（无需登录）
type NewsPublicListReq struct {
	g.Meta   `path:"/teaching/teachingNews/publicList" method:"get" tags:"新闻公告" summary:"公开新闻列表"`
	PageNo   int `json:"pageNo" dc:"页码" d:"1"`
	PageSize int `json:"pageSize" dc:"每页数量" d:"10"`
}

// NewsPublicListRes 公开新闻列表响应
type NewsPublicListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}
