package response

import (
	"math"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// Response 统一响应格式
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Result    interface{} `json:"result,omitempty"`
	Success   bool        `json:"success"`
	Timestamp int64       `json:"timestamp"`
}

// PageResult 分页响应格式
type PageResult struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
	Size    int         `json:"size"`
	Current int         `json:"current"`
	Pages   int         `json:"pages"`
}

// NewPageResult 创建分页结果
func NewPageResult(records interface{}, total int64, size, current int) *PageResult {
	pages := 0
	if size > 0 {
		pages = int(math.Ceil(float64(total) / float64(size)))
	}
	return &PageResult{
		Records: records,
		Total:   total,
		Size:    size,
		Current: current,
		Pages:   pages,
	}
}

// Json 返回JSON响应
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	r.Response.WriteJson(Response{
		Code:      code,
		Message:   message,
		Result:    responseData,
		Success:   code == 200,
		Timestamp: gtime.Now().Timestamp(),
	})
}

// JsonExit 返回JSON响应并退出
func JsonExit(r *ghttp.Request, code int, message string) {
	Json(r, code, message)
	r.Exit()
}

// Success 返回成功响应
func Success(r *ghttp.Request, data ...interface{}) {
	Json(r, 200, "操作成功", data...)
}

// Error 返回错误响应
func Error(r *ghttp.Request, message string) {
	Json(r, 500, message)
}

// UnauthorizedExit 返回401未授权响应并退出
func UnauthorizedExit(r *ghttp.Request, message string) {
	JsonExit(r, 401, message)
}

// ForbiddenExit 返回403禁止访问响应并退出
func ForbiddenExit(r *ghttp.Request, message string) {
	JsonExit(r, 403, message)
}

// TooManyRequestsExit 返回429请求过于频繁响应并退出
func TooManyRequestsExit(r *ghttp.Request, message string) {
	JsonExit(r, 429, message)
}


// SuccessMsg 返回成功消息响应
func SuccessMsg(r *ghttp.Request, message string) {
	Json(r, 200, message)
}

// PageSuccess 返回分页成功响应
func PageSuccess(r *ghttp.Request, records interface{}, total int, current, size int) {
	pageResult := NewPageResult(records, int64(total), size, current)
	Success(r, pageResult)
}
