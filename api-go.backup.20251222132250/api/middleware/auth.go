package middleware

import (
	"strings"

	"teaching-open/internal/consts"
	"teaching-open/utility/jwt"
	"teaching-open/utility/response"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Auth 认证中间件
func Auth(r *ghttp.Request) {
	// 从请求头获取token
	token := r.Header.Get("X-Access-Token")
	if token == "" {
		token = r.Header.Get("Authorization")
		// 去掉Bearer前缀（如果存在）
		token = strings.TrimPrefix(token, "Bearer ")
	}

	// 检查token是否存在
	if token == "" {
		response.Unauthorized(r, "未提供访问令牌")
		return
	}

	// 验证并解析token
	claims, err := jwt.ValidateToken(token)
	if err != nil {
		if err == jwt.ErrTokenExpired {
			response.Unauthorized(r, "访问令牌已过期")
		} else {
			response.Unauthorized(r, "访问令牌无效")
		}
		return
	}

	// 将用户信息存入请求上下文
	r.SetCtxVar(consts.CtxKeyUserId, claims.UserId)
	r.SetCtxVar(consts.CtxKeyUsername, claims.Username)
	r.SetCtxVar(consts.CtxKeyRealname, claims.Realname)
	r.SetCtxVar(consts.CtxKeyToken, token)

	// 继续执行后续中间件和处理函数
	r.Middleware.Next()
}
