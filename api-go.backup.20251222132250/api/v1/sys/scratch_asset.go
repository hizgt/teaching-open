package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== Scratch素材管理 ====================

// ScratchAssetListReq 素材列表请求
type ScratchAssetListReq struct {
	g.Meta    `path:"/teaching/teachingScratchAssets/list" method:"get" tags:"Scratch素材" summary:"素材列表"`
	AssetType int    `json:"assetType" dc:"素材类型 -1全部 1背景 2声音 3造型 4角色" d:"-1"`
	AssetName string `json:"assetName" dc:"素材名关键词"`
	Tags      string `json:"tags" dc:"标签关键词"`
	PageNo    int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize  int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// ScratchAssetListRes 素材列表响应
type ScratchAssetListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}

// ScratchAssetGetReq 获取素材请求 (用于Scratch编辑器加载)
type ScratchAssetGetReq struct {
	g.Meta    `path:"/teaching/teachingScratchAssets/getScratchAssets" method:"get" tags:"Scratch素材" summary:"获取素材(Scratch编辑器)"`
	AssetType int `json:"assetType" v:"required|in:1,2,3,4#类型不能为空|类型值无效" dc:"素材类型 1背景 2声音 3造型 4角色"`
}

// ScratchAssetGetRes 获取素材响应
type ScratchAssetGetRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"list"`
}

// ScratchAssetAddReq 添加素材请求
type ScratchAssetAddReq struct {
	g.Meta    `path:"/teaching/teachingScratchAssets/add" method:"post" tags:"Scratch素材" summary:"添加素材"`
	AssetType int    `json:"assetType" v:"required|in:1,2,3,4#类型不能为空|类型值无效" dc:"素材类型 1背景 2声音 3造型 4角色"`
	AssetName string `json:"assetName" v:"required#名称不能为空" dc:"素材名"`
	AssetData string `json:"assetData" dc:"素材JSON数据"`
	Md5Ext    string `json:"md5Ext" dc:"素材md5"`
	Tags      string `json:"tags" dc:"标签，多个用逗号分隔"`
}

// ScratchAssetAddRes 添加素材响应
type ScratchAssetAddRes struct {
	g.Meta `mime:"application/json"`
	Id     string `json:"id"`
}

// ScratchAssetEditReq 编辑素材请求
type ScratchAssetEditReq struct {
	g.Meta    `path:"/teaching/teachingScratchAssets/edit" method:"put" tags:"Scratch素材" summary:"编辑素材"`
	Id        string `json:"id" v:"required#ID不能为空" dc:"素材ID"`
	AssetType int    `json:"assetType" dc:"素材类型 1背景 2声音 3造型 4角色" d:"-1"`
	AssetName string `json:"assetName" dc:"素材名"`
	AssetData string `json:"assetData" dc:"素材JSON数据"`
	Md5Ext    string `json:"md5Ext" dc:"素材md5"`
	Tags      string `json:"tags" dc:"标签"`
}

// ScratchAssetEditRes 编辑素材响应
type ScratchAssetEditRes struct {
	g.Meta `mime:"application/json"`
}

// ScratchAssetDeleteReq 删除素材请求
type ScratchAssetDeleteReq struct {
	g.Meta `path:"/teaching/teachingScratchAssets/delete" method:"delete" tags:"Scratch素材" summary:"删除素材"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"素材ID"`
}

// ScratchAssetDeleteRes 删除素材响应
type ScratchAssetDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ScratchAssetDeleteBatchReq 批量删除素材请求
type ScratchAssetDeleteBatchReq struct {
	g.Meta `path:"/teaching/teachingScratchAssets/deleteBatch" method:"delete" tags:"Scratch素材" summary:"批量删除素材"`
	Ids    string `json:"ids" v:"required#IDs不能为空" dc:"素材ID列表，逗号分隔"`
}

// ScratchAssetDeleteBatchRes 批量删除素材响应
type ScratchAssetDeleteBatchRes struct {
	g.Meta `mime:"application/json"`
}

// ScratchAssetGetByIdReq 素材详情请求
type ScratchAssetGetByIdReq struct {
	g.Meta `path:"/teaching/teachingScratchAssets/queryById" method:"get" tags:"Scratch素材" summary:"素材详情"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"素材ID"`
}

// ScratchAssetGetByIdRes 素材详情响应
type ScratchAssetGetByIdRes struct {
	g.Meta `mime:"application/json"`
	*ScratchAssetInfo
}

// ScratchAssetInfo 素材信息
type ScratchAssetInfo struct {
	Id         string      `json:"id"`
	AssetType  int         `json:"assetType"`
	AssetName  string      `json:"assetName"`
	AssetData  string      `json:"assetData"`
	Md5Ext     string      `json:"md5Ext"`
	Tags       string      `json:"tags"`
	CreateBy   string      `json:"createBy"`
	CreateTime interface{} `json:"createTime"`
	UpdateBy   string      `json:"updateBy"`
	UpdateTime interface{} `json:"updateTime"`
	DelFlag    int         `json:"delFlag"`
}
