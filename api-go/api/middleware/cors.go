package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

// CORS 跨域中间件
func CORS(r *ghttp.Request) {
	ctx := gctx.New()

	// 从配置文件读取CORS配置
	allowOrigin := g.Cfg().MustGet(ctx, "cors.allowOrigin", "*").String()
	allowMethods := g.Cfg().MustGet(ctx, "cors.allowMethods", "GET,POST,PUT,DELETE,OPTIONS").String()
	allowHeaders := g.Cfg().MustGet(ctx, "cors.allowHeaders", "Origin,Content-Type,Accept,Authorization,X-Token").String()
	allowCredentials := g.Cfg().MustGet(ctx, "cors.allowCredentials", true).Bool()
	maxAge := g.Cfg().MustGet(ctx, "cors.maxAge", 3600).Int()

	// 设置CORS头
	r.Response.CORSDefault()
	r.Response.Header().Set("Access-Control-Allow-Origin", allowOrigin)
	r.Response.Header().Set("Access-Control-Allow-Methods", allowMethods)
	r.Response.Header().Set("Access-Control-Allow-Headers", allowHeaders)

	if allowCredentials {
		r.Response.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	r.Response.Header().Set("Access-Control-Max-Age", g.NewVar(maxAge).String())

	// OPTIONS请求直接返回
	if r.Method == "OPTIONS" {
		r.Response.WriteStatus(200)
		r.ExitAll()
	}

	r.Middleware.Next()
}
