package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// Response 统一响应结构体
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Result    interface{} `json:"result,omitempty"`
	Success   bool        `json:"success"`
	Timestamp int64       `json:"timestamp"`
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

// SuccessExit 返回成功响应并退出
func SuccessExit(r *ghttp.Request, data ...interface{}) {
	Success(r, data...)
	r.Exit()
}

// Error 返回错误响应
func Error(r *ghttp.Request, message string) {
	Json(r, 500, message)
}

// ErrorExit 返回错误响应并退出
func ErrorExit(r *ghttp.Request, message string) {
	Error(r, message)
	r.Exit()
}

// Unauthorized 返回未授权响应
func Unauthorized(r *ghttp.Request, message string) {
	if message == "" {
		message = "未登录或token已过期"
	}
	Json(r, 401, message)
}

// UnauthorizedExit 返回未授权响应并退出
func UnauthorizedExit(r *ghttp.Request, message string) {
	Unauthorized(r, message)
	r.Exit()
}

// Forbidden 返回禁止访问响应
func Forbidden(r *ghttp.Request, message string) {
	if message == "" {
		message = "无权限访问"
	}
	Json(r, 403, message)
}

// ForbiddenExit 返回禁止访问响应并退出
func ForbiddenExit(r *ghttp.Request, message string) {
	Forbidden(r, message)
	r.Exit()
}

// BadRequest 返回错误请求响应
func BadRequest(r *ghttp.Request, message string) {
	Json(r, 400, message)
}

// BadRequestExit 返回错误请求响应并退出
func BadRequestExit(r *ghttp.Request, message string) {
	BadRequest(r, message)
	r.Exit()
}

// NotFound 返回未找到响应
func NotFound(r *ghttp.Request, message string) {
	if message == "" {
		message = "资源不存在"
	}
	Json(r, 404, message)
}

// NotFoundExit 返回未找到响应并退出
func NotFoundExit(r *ghttp.Request, message string) {
	NotFound(r, message)
	r.Exit()
}

// TooManyRequests 返回请求过于频繁响应
func TooManyRequests(r *ghttp.Request, message string) {
	if message == "" {
		message = "请求过于频繁，请稍后再试"
	}
	Json(r, 429, message)
}

// TooManyRequestsExit 返回请求过于频繁响应并退出
func TooManyRequestsExit(r *ghttp.Request, message string) {
	TooManyRequests(r, message)
	r.Exit()
}