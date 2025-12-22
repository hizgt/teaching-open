package middleware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"teaching-open/utility/jwt"
	"teaching-open/utility/response"
)

// AuthConfig 认证配置
type AuthConfig struct {
	ExcludePaths       []string      // 排除路径
	MaxDevices         int           // 最大设备数
	EnableAutoRefresh  bool          // 启用自动刷新
	RefreshThreshold   time.Duration // 刷新阈值
	EnableRateLimit    bool          // 启用限流
	RateLimitDuration  time.Duration // 限流时间窗口
	RateLimitMaxRequests int         // 最大请求数
}

// DefaultAuthConfig 默认认证配置
func DefaultAuthConfig() *AuthConfig {
	return &AuthConfig{
		ExcludePaths:       []string{"/sys/login", "/sys/register", "/sys/captcha", "/health"},
		MaxDevices:         5,                // 默认最多5个设备
		EnableAutoRefresh:  true,             // 启用自动刷新
		RefreshThreshold:   30 * time.Minute, // 30分钟内自动刷新
		EnableRateLimit:    true,             // 启用限流
		RateLimitDuration:  time.Minute,      // 1分钟
		RateLimitMaxRequests: 60,             // 每分钟最多60次请求
	}
}

// Auth 认证中间件
func Auth(r *ghttp.Request) {
	config := DefaultAuthConfig()

	// 检查是否在排除路径中
	if isExcludedPath(r.URL.Path, config.ExcludePaths) {
		r.Middleware.Next()
		return
	}

	// 限流检查
	if config.EnableRateLimit {
		if !checkRateLimit(r) {
			response.TooManyRequestsExit(r, "请求过于频繁，请稍后再试")
			return
		}
	}

	// 提取Token
	token, err := extractToken(r)
	if err != nil {
		response.UnauthorizedExit(r, err.Error())
		return
	}

	// 验证Token
	claims, err := jwt.New().ValidateToken(token)
	if err != nil {
		response.UnauthorizedExit(r, "token验证失败")
		return
	}

	// 检查多设备登录限制
	if !checkDeviceLimit(claims.UserId, claims.DeviceId, config.MaxDevices) {
		response.UnauthorizedExit(r, "账号已在其他设备登录，请重新登录")
		return
	}

	// 将用户信息存入上下文
	setUserContext(r, claims)

	// 自动刷新Token
	if config.EnableAutoRefresh {
		autoRefreshToken(r, claims, config.RefreshThreshold)
	}

	r.Middleware.Next()
}

// Permission 权限验证中间件
func Permission(permission string) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		userId := r.GetCtxVar("userId").String()
		if userId == "" {
			response.UnauthorizedExit(r, "用户未登录")
			return
		}

		// 检查用户权限（这里需要业务层实现）
		// hasPermission, err := service.Permission().CheckUserPermission(r.Context(), userId, permission)
		// if err != nil || !hasPermission {
		//     response.ForbiddenExit(r, "无权限访问")
		//     return
		// }

		// 临时实现：超级管理员或特定权限
		roleIds := getUserRoleIds(r)
		if !hasPermission(roleIds, permission) {
			response.ForbiddenExit(r, "无权限访问")
			return
		}

		r.Middleware.Next()
	}
}

// Role 角色验证中间件
func Role(roles ...string) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		userRoleIds := getUserRoleIds(r)
		if !hasAnyRole(userRoleIds, roles) {
			response.ForbiddenExit(r, "无权限访问")
			return
		}

		r.Middleware.Next()
	}
}

// AdminOnly 仅管理员中间件
func AdminOnly(r *ghttp.Request) {
	userRoleIds := getUserRoleIds(r)
	if !hasRole(userRoleIds, "admin") {
		response.ForbiddenExit(r, "需要管理员权限")
		return
	}

	r.Middleware.Next()
}

// CORS 跨域中间件
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// SecurityHeaders 安全头中间件
func SecurityHeaders(r *ghttp.Request) {
	r.Response.Header().Set("X-Content-Type-Options", "nosniff")
	r.Response.Header().Set("X-Frame-Options", "DENY")
	r.Response.Header().Set("X-XSS-Protection", "1; mode=block")
	r.Response.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	r.Response.Header().Set("Content-Security-Policy", "default-src 'self'")
	r.Middleware.Next()
}

// Logging 日志中间件
func Logging(r *ghttp.Request) {
	start := time.Now()

	r.Middleware.Next()

	// 记录请求日志
	duration := time.Since(start)
	userId := r.GetCtxVar("userId").String()
	if userId == "" {
		userId = "anonymous"
	}

	g.Log().Info(r.Context(), "API Request", g.Map{
		"method":     r.Method,
		"path":       r.URL.Path,
		"userId":     userId,
		"ip":         r.GetClientIp(),
		"userAgent":  r.Header.Get("User-Agent"),
		"duration":   duration.String(),
		"statusCode": r.Response.Status,
	})
}

// 辅助函数

// isExcludedPath 检查是否为排除路径
func isExcludedPath(path string, excludePaths []string) bool {
	for _, excludePath := range excludePaths {
		if strings.HasPrefix(path, excludePath) {
			return true
		}
	}
	return false
}

// extractToken 提取Token
func extractToken(r *ghttp.Request) (string, error) {
	// 优先从Authorization头提取
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		return jwt.New().ExtractTokenFromHeader(authHeader)
	}

	// 从X-Access-Token头提取
	token := r.Header.Get("X-Access-Token")
	if token != "" {
		return token, nil
	}

	// 从查询参数提取
	token = r.Get("token").String()
	if token != "" {
		return token, nil
	}

	return "", fmt.Errorf("未找到访问令牌")
}

// checkRateLimit 限流检查
func checkRateLimit(r *ghttp.Request) bool {
	config := DefaultAuthConfig()
	ip := r.GetClientIp()
	cacheKey := fmt.Sprintf("rate_limit:%s", ip)

	count, err := gcache.Get(r.Context(), cacheKey)
	if err != nil {
		// 首次请求
		gcache.Set(r.Context(), cacheKey, 1, config.RateLimitDuration)
		return true
	}

	currentCount := count.Int() + 1
	if currentCount > config.RateLimitMaxRequests {
		return false
	}

	gcache.Set(r.Context(), cacheKey, currentCount, config.RateLimitDuration)
	return true
}

// checkDeviceLimit 检查设备登录限制
func checkDeviceLimit(userId, deviceId string, maxDevices int) bool {
	cacheKey := fmt.Sprintf("user_devices:%s", userId)

	devices, err := gcache.Get(context.Background(), cacheKey)
	if err != nil {
		// 首次登录，创建设备列表
		deviceList := []string{deviceId}
		gcache.Set(context.Background(), cacheKey, deviceList, 24*time.Hour)
		return true
	}

	deviceList := devices.Strings()
	if len(deviceList) >= maxDevices {
		// 检查当前设备是否已登录
		for _, device := range deviceList {
			if device == deviceId {
				return true
			}
		}
		// 设备数已达上限，且当前设备未登录
		return false
	}

	// 添加新设备
	deviceList = append(deviceList, deviceId)
	gcache.Set(context.Background(), cacheKey, deviceList, 24*time.Hour)
	return true
}

// setUserContext 设置用户上下文
func setUserContext(r *ghttp.Request, claims *jwt.Claims) {
	r.SetCtxVar("userId", claims.UserId)
	r.SetCtxVar("username", claims.Username)
	r.SetCtxVar("roleIds", claims.RoleIds)
	r.SetCtxVar("deviceId", claims.DeviceId)
	r.SetCtxVar("deviceType", claims.DeviceType)
}

// autoRefreshToken 自动刷新Token
func autoRefreshToken(r *ghttp.Request, claims *jwt.Claims, threshold time.Duration) {
	if claims.ExpiresAt == nil {
		return
	}

	remainingTime := time.Until(claims.ExpiresAt.Time)
	if remainingTime > threshold {
		return
	}

	// 生成新的Token对
	jwtUtil := jwt.New()
	tokenPair, err := jwtUtil.GenerateTokenPair(claims.UserId, claims.Username, claims.RoleIds, claims.DeviceId, claims.DeviceType)
	if err != nil {
		g.Log().Error(r.Context(), "自动刷新Token失败", err)
		return
	}

	// 将新Token添加到响应头
	r.Response.Header().Set("X-New-Access-Token", tokenPair.AccessToken)
	r.Response.Header().Set("X-New-Refresh-Token", tokenPair.RefreshToken)
	r.Response.Header().Set("X-Token-Refreshed", "true")
}

// getUserRoleIds 获取用户角色ID列表
func getUserRoleIds(r *ghttp.Request) []string {
	roleIds := r.GetCtxVar("roleIds")
	if roleIds.IsEmpty() {
		return []string{}
	}
	return roleIds.Strings()
}

// hasPermission 检查权限
func hasPermission(userRoleIds []string, permission string) bool {
	// 临时实现：超级管理员拥有所有权限
	if hasRole(userRoleIds, "admin") {
		return true
	}

	// 这里应该调用权限服务进行检查
	// 暂时返回true，实际项目中需要实现权限检查逻辑
	return true
}

// hasRole 检查是否拥有指定角色
func hasRole(userRoleIds []string, role string) bool {
	for _, userRoleId := range userRoleIds {
		if userRoleId == role {
			return true
		}
	}
	return false
}

// hasAnyRole 检查是否拥有任意指定角色
func hasAnyRole(userRoleIds []string, roles []string) bool {
	for _, role := range roles {
		if hasRole(userRoleIds, role) {
			return true
		}
	}
	return false
}