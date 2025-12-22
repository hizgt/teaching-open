package response

import (
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