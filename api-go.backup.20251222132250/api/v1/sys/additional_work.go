package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 附加作业管理 ====================

// AdditionalWorkListReq 附加作业列表请求
type AdditionalWorkListReq struct {
	g.Meta   `path:"/teaching/teachingAdditionalWork/list" method:"get" tags:"附加作业" summary:"附加作业列表"`
	WorkName string `json:"workName" dc:"作业名称关键词"`
	CodeType string `json:"codeType" dc:"代码类型"`
	Status   *int   `json:"status" dc:"状态 0未发布 1发布"`
	WorkDept string `json:"workDept" dc:"班级"`
	PageNo   int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// AdditionalWorkListRes 附加作业列表响应
type AdditionalWorkListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}

// AdditionalWorkAddReq 添加附加作业请求
type AdditionalWorkAddReq struct {
	g.Meta          `path:"/teaching/teachingAdditionalWork/add" method:"post" tags:"附加作业" summary:"添加附加作业"`
	CodeType        string `json:"codeType" dc:"代码类型"`
	WorkName        string `json:"workName" v:"required#作业名称不能为空" dc:"作业名"`
	WorkDesc        string `json:"workDesc" dc:"作业描述"`
	WorkCover       string `json:"workCover" dc:"作业封面"`
	WorkUrl         string `json:"workUrl" dc:"作业文件"`
	WorkDept        string `json:"workDept" v:"required#分配班级不能为空" dc:"分配班级逗号分割"`
	WorkDocumentUrl string `json:"workDocumentUrl" dc:"作业资料"`
}

// AdditionalWorkAddRes 添加附加作业响应
type AdditionalWorkAddRes struct {
	g.Meta `mime:"application/json"`
	Id     string `json:"id"`
}

// AdditionalWorkEditReq 编辑附加作业请求
type AdditionalWorkEditReq struct {
	g.Meta          `path:"/teaching/teachingAdditionalWork/edit" method:"put" tags:"附加作业" summary:"编辑附加作业"`
	Id              string `json:"id" v:"required#ID不能为空" dc:"附加作业ID"`
	CodeType        string `json:"codeType" dc:"代码类型"`
	WorkName        string `json:"workName" dc:"作业名"`
	WorkDesc        string `json:"workDesc" dc:"作业描述"`
	WorkCover       string `json:"workCover" dc:"作业封面"`
	WorkUrl         string `json:"workUrl" dc:"作业文件"`
	WorkDept        string `json:"workDept" dc:"分配班级逗号分割"`
	WorkDocumentUrl string `json:"workDocumentUrl" dc:"作业资料"`
	Status          *int   `json:"status" dc:"状态"`
}

// AdditionalWorkEditRes 编辑附加作业响应
type AdditionalWorkEditRes struct {
	g.Meta `mime:"application/json"`
}

// AdditionalWorkDeleteReq 删除附加作业请求
type AdditionalWorkDeleteReq struct {
	g.Meta `path:"/teaching/teachingAdditionalWork/delete" method:"delete" tags:"附加作业" summary:"删除附加作业"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"附加作业ID"`
}

// AdditionalWorkDeleteRes 删除附加作业响应
type AdditionalWorkDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// AdditionalWorkDeleteBatchReq 批量删除附加作业请求
type AdditionalWorkDeleteBatchReq struct {
	g.Meta `path:"/teaching/teachingAdditionalWork/deleteBatch" method:"delete" tags:"附加作业" summary:"批量删除附加作业"`
	Ids    string `json:"ids" v:"required#IDs不能为空" dc:"附加作业ID列表，逗号分隔"`
}

// AdditionalWorkDeleteBatchRes 批量删除附加作业响应
type AdditionalWorkDeleteBatchRes struct {
	g.Meta `mime:"application/json"`
}

// AdditionalWorkGetByIdReq 附加作业详情请求
type AdditionalWorkGetByIdReq struct {
	g.Meta `path:"/teaching/teachingAdditionalWork/queryById" method:"get" tags:"附加作业" summary:"附加作业详情"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"附加作业ID"`
}

// AdditionalWorkGetByIdRes 附加作业详情响应
type AdditionalWorkGetByIdRes struct {
	g.Meta `mime:"application/json"`
	Info   *AdditionalWorkInfo `json:"info"`
}

// AdditionalWorkInfo 附加作业信息
type AdditionalWorkInfo struct {
	Id              string      `json:"id"`
	CodeType        string      `json:"codeType"`
	WorkName        string      `json:"workName"`
	WorkDesc        string      `json:"workDesc"`
	WorkCover       string      `json:"workCover"`
	WorkUrl         string      `json:"workUrl"`
	WorkDept        string      `json:"workDept"`
	WorkDocumentUrl string      `json:"workDocumentUrl"`
	Status          int         `json:"status"`
	SysOrgCode      string      `json:"sysOrgCode"`
	CreateBy        string      `json:"createBy"`
	CreateTime      interface{} `json:"createTime"`
	UpdateBy        string      `json:"updateBy"`
	UpdateTime      interface{} `json:"updateTime"`
}

// AdditionalWorkPublishReq 发布附加作业请求
type AdditionalWorkPublishReq struct {
	g.Meta `path:"/teaching/teachingAdditionalWork/publish" method:"put" tags:"附加作业" summary:"发布附加作业"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"附加作业ID"`
}

// AdditionalWorkPublishRes 发布附加作业响应
type AdditionalWorkPublishRes struct {
	g.Meta `mime:"application/json"`
}

// AdditionalWorkOfflineReq 下架附加作业请求
type AdditionalWorkOfflineReq struct {
	g.Meta `path:"/teaching/teachingAdditionalWork/offline" method:"put" tags:"附加作业" summary:"下架附加作业"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"附加作业ID"`
}

// AdditionalWorkOfflineRes 下架附加作业响应
type AdditionalWorkOfflineRes struct {
	g.Meta `mime:"application/json"`
}

// AdditionalWorkByDeptReq 按班级获取附加作业请求
type AdditionalWorkByDeptReq struct {
	g.Meta   `path:"/teaching/teachingAdditionalWork/listByDept" method:"get" tags:"附加作业" summary:"按班级获取附加作业"`
	WorkDept string `json:"workDept" dc:"班级"`
	CodeType string `json:"codeType" dc:"代码类型"`
	PageNo   int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// AdditionalWorkByDeptRes 按班级获取附加作业响应
type AdditionalWorkByDeptRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}
