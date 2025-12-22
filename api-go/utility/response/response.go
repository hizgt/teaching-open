package response

import (
	"teaching-open/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// JsonRes 响应结构
type JsonRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
}

// PageRes 分页响应结构
type PageRes struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
	Size    int         `json:"size"`
	Current int         `json:"current"`
	Pages   int         `json:"pages"`
}

// PageResult 分页结果响应
type PageResult struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Result  PageRes `json:"result"`
	Success bool    `json:"success"`
}

// Json 返回JSON格式数据
func Json(r *ghttp.Request, code int, message string, data interface{}) {
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Result:  data,
		Success: code == consts.CodeSuccess,
	})
}

// JsonExit 返回JSON格式数据并退出当前请求
func JsonExit(r *ghttp.Request, code int, message string, data interface{}) {
	Json(r, code, message, data)
	r.Exit()
}

// Success 返回成功响应
func Success(r *ghttp.Request, data interface{}) {
	Json(r, consts.CodeSuccess, "操作成功", data)
}

// SuccessExit 返回成功响应并退出
func SuccessExit(r *ghttp.Request, data interface{}) {
	JsonExit(r, consts.CodeSuccess, "操作成功", data)
}

// Error 返回错误响应
func Error(r *ghttp.Request, message string) {
	Json(r, consts.CodeError, message, nil)
}

// ErrorExit 返回错误响应并退出
func ErrorExit(r *ghttp.Request, message string) {
	JsonExit(r, consts.CodeError, message, nil)
}

// ErrorCode 返回指定错误码的错误响应
func ErrorCode(r *ghttp.Request, code int, message string) {
	Json(r, code, message, nil)
}

// ErrorCodeExit 返回指定错误码的错误响应并退出
func ErrorCodeExit(r *ghttp.Request, code int, message string) {
	JsonExit(r, code, message, nil)
}

// Page 返回分页数据
func Page(r *ghttp.Request, records interface{}, total int64, current, size int) {
	pages := int(total) / size
	if int(total)%size != 0 {
		pages++
	}

	r.Response.WriteJson(PageResult{
		Code:    consts.CodeSuccess,
		Message: "查询成功",
		Result: PageRes{
			Records: records,
			Total:   total,
			Size:    size,
			Current: current,
			Pages:   pages,
		},
		Success: true,
	})
}

// Unauthorized 返回未授权错误
func Unauthorized(r *ghttp.Request, message string) {
	if message == "" {
		message = "未授权访问"
	}
	JsonExit(r, consts.CodeUnauthorized, message, nil)
}

// PermissionDenied 返回权限不足错误
func PermissionDenied(r *ghttp.Request, message string) {
	if message == "" {
		message = "权限不足"
	}
	JsonExit(r, consts.CodePermissionDenied, message, nil)
}

// InvalidParameter 返回参数错误
func InvalidParameter(r *ghttp.Request, message string) {
	if message == "" {
		message = "参数错误"
	}
	JsonExit(r, consts.CodeInvalidParameter, message, nil)
}

// NotFound 返回资源不存在错误
func NotFound(r *ghttp.Request, message string) {
	if message == "" {
		message = "资源不存在"
	}
	JsonExit(r, consts.CodeRecordNotFound, message, nil)
}

// InternalError 返回内部错误
func InternalError(r *ghttp.Request, message string) {
	if message == "" {
		message = "服务器内部错误"
	}
	g.Log().Error(r.Context(), message)
	JsonExit(r, consts.CodeInternalError, message, nil)
}
