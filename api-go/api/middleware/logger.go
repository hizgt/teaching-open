package middleware

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Logger 日志中间件
func Logger(r *ghttp.Request) {
	// 记录开始时间
	startTime := time.Now()

	// 继续处理请求
	r.Middleware.Next()

	// 计算处理时间
	latency := time.Since(startTime)

	// 记录日志
	g.Log().Info(r.Context(),
		"method:", r.Method,
		"path:", r.URL.Path,
		"status:", r.Response.Status,
		"latency:", latency,
		"ip:", r.GetClientIp(),
	)
}
