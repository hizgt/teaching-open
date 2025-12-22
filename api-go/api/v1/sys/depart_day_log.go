package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 部门日志统计 ====================

// DepartDayLogGetReportReq 统计报表请求
type DepartDayLogGetReportReq struct {
	g.Meta    `path:"/teaching/teachingDepartDayLog/getReport" method:"get" tags:"部门日志统计" summary:"统计报表"`
	DepartId  string `json:"departId" dc:"班级ID"`
	StartDate string `json:"startDate" v:"date#开始日期格式错误" dc:"开始日期 YYYY-MM-DD"`
	EndDate   string `json:"endDate" v:"date#结束日期格式错误" dc:"结束日期 YYYY-MM-DD"`
	PageNo    int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize  int    `json:"pageSize" dc:"每页数量" d:"10"`
	SortField string `json:"sortField" dc:"排序字段" d:"create_time"`
	SortOrder string `json:"sortOrder" dc:"排序方向 asc/desc" d:"desc"`
}

// DepartDayLogGetReportRes 统计报表响应
type DepartDayLogGetReportRes struct {
	g.Meta `mime:"application/json"`
	List   []*DepartDayLogInfo `json:"records"`
	Total  int64               `json:"total"`
}

// DepartDayLogGroupByDepartReq 按部门统计请求
type DepartDayLogGroupByDepartReq struct {
	g.Meta    `path:"/teaching/teachingDepartDayLog/getReportGroupByDepart" method:"get" tags:"部门日志统计" summary:"按部门统计"`
	StartDate string `json:"startDate" v:"date#开始日期格式错误" dc:"开始日期 YYYY-MM-DD"`
	EndDate   string `json:"endDate" v:"date#结束日期格式错误" dc:"结束日期 YYYY-MM-DD"`
	PageNo    int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize  int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// DepartDayLogGroupByDepartRes 按部门统计响应
type DepartDayLogGroupByDepartRes struct {
	g.Meta `mime:"application/json"`
	List   []*DepartStatInfo `json:"records"`
	Total  int64             `json:"total"`
}

// DepartDayLogGroupByMonthReq 按月份统计请求
type DepartDayLogGroupByMonthReq struct {
	g.Meta    `path:"/teaching/teachingDepartDayLog/getReportGroupByMonth" method:"get" tags:"部门日志统计" summary:"按月份统计"`
	DepartId  string `json:"departId" dc:"班级ID"`
	StartDate string `json:"startDate" v:"date#开始日期格式错误" dc:"开始日期 YYYY-MM-DD"`
	EndDate   string `json:"endDate" v:"date#结束日期格式错误" dc:"结束日期 YYYY-MM-DD"`
}

// DepartDayLogGroupByMonthRes 按月份统计响应
type DepartDayLogGroupByMonthRes struct {
	g.Meta `mime:"application/json"`
	List   []*MonthStatInfo `json:"list"`
}

// DepartDayLogUnitViewReq 单元查看日志请求
type DepartDayLogUnitViewReq struct {
	g.Meta   `path:"/teaching/teachingDepartDayLog/unitViewLog" method:"post" tags:"部门日志统计" summary:"记录单元查看日志"`
	UnitId   string `json:"unitId" v:"required#单元ID不能为空" dc:"单元ID"`
	CourseId string `json:"courseId" v:"required#课程ID不能为空" dc:"课程ID"`
	DepartId string `json:"departId" dc:"班级ID"`
}

// DepartDayLogUnitViewRes 单元查看日志响应
type DepartDayLogUnitViewRes struct {
	g.Meta `mime:"application/json"`
}

// DepartDayLogInfo 部门日志信息
type DepartDayLogInfo struct {
	Id                         string      `json:"id"`
	DepartId                   string      `json:"departId"`
	DepartName                 string      `json:"departName"`
	UnitOpenCount              int         `json:"unitOpenCount"`              // 开课次数
	CourseWorkAssignCount      int         `json:"courseWorkAssignCount"`      // 课程作业布置次数
	AdditionalWorkAssignCount  int         `json:"additionalWorkAssignCount"`  // 附加作业布置次数
	CourseWorkCorrectCount     int         `json:"courseWorkCorrectCount"`     // 课程作业批改次数
	AdditionalWorkCorrectCount int         `json:"additionalWorkCorrectCount"` // 附加作业批改次数
	CourseWorkSubmitCount      int         `json:"courseWorkSubmitCount"`      // 课程作业提交次数
	AdditionalWorkSubmitCount  int         `json:"additionalWorkSubmitCount"`  // 附加作业提交次数
	CreateTime                 interface{} `json:"createTime"`                 // 日期
}

// DepartStatInfo 部门统计信息
type DepartStatInfo struct {
	DepartId                   string `json:"departId"`
	DepartName                 string `json:"departName"`
	TotalUnitOpen              int64  `json:"totalUnitOpen"`              // 总开课次数
	TotalCourseWorkAssign      int64  `json:"totalCourseWorkAssign"`      // 总课程作业布置次数
	TotalAdditionalWorkAssign  int64  `json:"totalAdditionalWorkAssign"`  // 总附加作业布置次数
	TotalCourseWorkCorrect     int64  `json:"totalCourseWorkCorrect"`     // 总课程作业批改次数
	TotalAdditionalWorkCorrect int64  `json:"totalAdditionalWorkCorrect"` // 总附加作业批改次数
	TotalCourseWorkSubmit      int64  `json:"totalCourseWorkSubmit"`      // 总课程作业提交次数
	TotalAdditionalWorkSubmit  int64  `json:"totalAdditionalWorkSubmit"`  // 总附加作业提交次数
	DayCount                   int64  `json:"dayCount"`                   // 统计天数
}

// MonthStatInfo 月份统计信息
type MonthStatInfo struct {
	Month                      string `json:"month"`                      // 月份 YYYY-MM
	TotalUnitOpen              int64  `json:"totalUnitOpen"`              // 总开课次数
	TotalCourseWorkAssign      int64  `json:"totalCourseWorkAssign"`      // 总课程作业布置次数
	TotalAdditionalWorkAssign  int64  `json:"totalAdditionalWorkAssign"`  // 总附加作业布置次数
	TotalCourseWorkCorrect     int64  `json:"totalCourseWorkCorrect"`     // 总课程作业批改次数
	TotalAdditionalWorkCorrect int64  `json:"totalAdditionalWorkCorrect"` // 总附加作业批改次数
	TotalCourseWorkSubmit      int64  `json:"totalCourseWorkSubmit"`      // 总课程作业提交次数
	TotalAdditionalWorkSubmit  int64  `json:"totalAdditionalWorkSubmit"`  // 总附加作业提交次数
	DayCount                   int64  `json:"dayCount"`                   // 统计天数
}
