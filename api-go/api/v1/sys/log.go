// =================================================================================
// API definitions for sys log module
// =================================================================================

package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LogInfo 系统日志信息
type LogInfo struct {
	Id           string `json:"id"`
	LogType      int    `json:"logType"`      // 日志类型（1登录日志，2操作日志）
	LogContent   string `json:"logContent"`   // 日志内容
	OperateType  int    `json:"operateType"`  // 操作类型
	Userid       string `json:"userid"`       // 操作用户账号
	Username     string `json:"username"`     // 操作用户名称
	Ip           string `json:"ip"`           // IP
	Method       string `json:"method"`       // 请求方法
	RequestUrl   string `json:"requestUrl"`   // 请求路径
	RequestParam string `json:"requestParam"` // 请求参数
	RequestType  string `json:"requestType"`  // 请求类型
	CostTime     int64  `json:"costTime"`     // 耗时(ms)
	CreateTime   string `json:"createTime"`   // 创建时间
}

// DataLogInfo 数据日志信息
type DataLogInfo struct {
	Id          string `json:"id"`
	CreateBy    string `json:"createBy"`    // 创建人
	CreateTime  string `json:"createTime"`  // 创建时间
	DataTable   string `json:"dataTable"`   // 表名
	DataId      string `json:"dataId"`      // 数据ID
	DataContent string `json:"dataContent"` // 数据内容
	DataVersion int    `json:"dataVersion"` // 版本号
}

// ========== 系统日志管理 ==========

// LogListReq 系统日志列表请求
type LogListReq struct {
	g.Meta      `path:"/sys/log/list" method:"get" tags:"日志管理" summary:"系统日志列表"`
	Page        int    `json:"page" d:"1" v:"min:1" dc:"页码"`
	PageSize    int    `json:"pageSize" d:"10" v:"min:1|max:100" dc:"每页数量"`
	LogType     int    `json:"logType" dc:"日志类型（1登录日志，2操作日志）"`
	Username    string `json:"username" dc:"操作用户名称"`
	Ip          string `json:"ip" dc:"IP地址"`
	StartTime   string `json:"startTime" dc:"开始时间"`
	EndTime     string `json:"endTime" dc:"结束时间"`
	LogContent  string `json:"logContent" dc:"日志内容关键字"`
	OperateType int    `json:"operateType" dc:"操作类型"`
}

// LogListRes 系统日志列表响应
type LogListRes struct {
	List     []*LogInfo `json:"list"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
}

// LogDeleteReq 删除系统日志请求
type LogDeleteReq struct {
	g.Meta `path:"/sys/log/delete" method:"delete" tags:"日志管理" summary:"删除系统日志"`
	Id     string `json:"id" v:"required#日志ID不能为空"`
}

// LogDeleteRes 删除系统日志响应
type LogDeleteRes struct{}

// LogDeleteBatchReq 批量删除系统日志请求
type LogDeleteBatchReq struct {
	g.Meta `path:"/sys/log/deleteBatch" method:"delete" tags:"日志管理" summary:"批量删除系统日志"`
	Ids    string `json:"ids" v:"required#日志ID不能为空"`
}

// LogDeleteBatchRes 批量删除系统日志响应
type LogDeleteBatchRes struct{}

// LogClearReq 清空系统日志请求
type LogClearReq struct {
	g.Meta  `path:"/sys/log/clear" method:"delete" tags:"日志管理" summary:"清空系统日志"`
	LogType int `json:"logType" dc:"日志类型（不填则清空所有日志）"`
}

// LogClearRes 清空系统日志响应
type LogClearRes struct {
	Count int64 `json:"count"` // 删除的数量
}

// ========== 数据日志管理 ==========

// DataLogListReq 数据日志列表请求
type DataLogListReq struct {
	g.Meta    `path:"/sys/dataLog/list" method:"get" tags:"日志管理" summary:"数据日志列表"`
	Page      int    `json:"page" d:"1" v:"min:1" dc:"页码"`
	PageSize  int    `json:"pageSize" d:"10" v:"min:1|max:100" dc:"每页数量"`
	DataTable string `json:"dataTable" dc:"表名"`
	DataId    string `json:"dataId" dc:"数据ID"`
	CreateBy  string `json:"createBy" dc:"创建人"`
	StartTime string `json:"startTime" dc:"开始时间"`
	EndTime   string `json:"endTime" dc:"结束时间"`
}

// DataLogListRes 数据日志列表响应
type DataLogListRes struct {
	List     []*DataLogInfo `json:"list"`
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}

// DataLogGetByIdReq 查询数据日志详情请求
type DataLogGetByIdReq struct {
	g.Meta `path:"/sys/dataLog/queryById" method:"get" tags:"日志管理" summary:"数据日志详情"`
	Id     string `json:"id" v:"required#日志ID不能为空"`
}

// DataLogGetByIdRes 查询数据日志详情响应
type DataLogGetByIdRes struct {
	*DataLogInfo
}

// DataLogHistoryReq 查询数据变更历史请求
type DataLogHistoryReq struct {
	g.Meta    `path:"/sys/dataLog/history" method:"get" tags:"日志管理" summary:"数据变更历史"`
	DataTable string `json:"dataTable" v:"required#表名不能为空"`
	DataId    string `json:"dataId" v:"required#数据ID不能为空"`
}

// DataLogHistoryRes 查询数据变更历史响应
type DataLogHistoryRes struct {
	List []*DataLogInfo `json:"list"`
}
