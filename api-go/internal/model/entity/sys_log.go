package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog 系统日志表
type SysLog struct {
	Id           string      `json:"id"           orm:"id,primary"     description:"日志ID"`
	LogType      int         `json:"logType"      orm:"log_type"       description:"日志类型(1登录2操作)"`
	LogContent   string      `json:"logContent"   orm:"log_content"    description:"日志内容"`
	OperateType  int         `json:"operateType"  orm:"operate_type"   description:"操作类型"`
	Userid       string      `json:"userid"       orm:"userid"         description:"用户ID"`
	Username     string      `json:"username"     orm:"username"       description:"用户名"`
	Ip           string      `json:"ip"           orm:"ip"             description:"IP地址"`
	Method       string      `json:"method"       orm:"method"         description:"请求方法"`
	RequestUrl   string      `json:"requestUrl"   orm:"request_url"    description:"请求URL"`
	RequestParam string      `json:"requestParam" orm:"request_param"  description:"请求参数"`
	RequestType  string      `json:"requestType"  orm:"request_type"   description:"请求类型"`
	CostTime     int64       `json:"costTime"     orm:"cost_time"      description:"耗时(ms)"`
	CreateBy     string      `json:"createBy"     orm:"create_by"      description:"创建人"`
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"    description:"创建时间"`
}

// TableName 表名
func (e *SysLog) TableName() string {
	return "sys_log"
}
