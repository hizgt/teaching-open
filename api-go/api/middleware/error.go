package middleware

import (
	"teaching-open/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Error 错误处理中间件
func Error(r *ghttp.Request) {
	r.Middleware.Next()

	// 如果有错误，统一处理
	if err := r.GetError(); err != nil {
		g.Log().Error(r.Context(), "Request error:", err)
		response.InternalError(r, err.Error())
	}
}
