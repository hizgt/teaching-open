package jwt

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"teaching-open/internal/consts"
)

// Claims JWT声明
type Claims struct {
	UserId    string `json:"userId"`
	Username  string `json:"username"`
	Realname  string `json:"realname"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

var (
	ErrTokenExpired = errors.New("token已过期")
	ErrTokenInvalid = errors.New("token无效")
)

// GenerateToken 生成JWT令牌
func GenerateToken(userId, username, realname string) (string, error) {
	ctx := gctx.New()

	// 获取JWT配置
	secret := g.Cfg().MustGet(ctx, "jwt.secret").String()
	if secret == "" {
		secret = "teaching-open-secret-key"
	}

	expireTime := g.Cfg().MustGet(ctx, "jwt.expire").Int()
	if expireTime == 0 {
		expireTime = 7200 // 默认2小时
	}

	now := time.Now().Unix()
	claims := Claims{
		UserId:    userId,
		Username:  username,
		Realname:  realname,
		IssuedAt:  now,
		ExpiresAt: now + int64(expireTime),
	}

	// 序列化claims
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// Base64编码
	token := base64.StdEncoding.EncodeToString(claimsBytes)

	return token, nil
}

// ParseToken 解析JWT令牌
func ParseToken(token string) (*Claims, error) {
	if token == "" {
		return nil, ErrTokenInvalid
	}

	// Base64解码
	claimsBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	// 反序列化claims
	var claims Claims
	err = json.Unmarshal(claimsBytes, &claims)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	return &claims, nil
}

// ValidateToken 验证JWT令牌
func ValidateToken(token string) (*Claims, error) {
	claims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	// 检查是否过期
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, ErrTokenExpired
	}

	return claims, nil
}

// GetUserIdFromToken 从令牌中获取用户ID
func GetUserIdFromToken(token string) (string, error) {
	claims, err := ValidateToken(token)
	if err != nil {
		return "", err
	}

	return claims.UserId, nil
}

// GetUsername 从上下文中获取用户名
func GetUsername(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v := ctx.Value(consts.CtxKeyUsername)
	if v == nil {
		return ""
	}
	if username, ok := v.(string); ok {
		return username
	}
	return ""
}

// GetUserId 从上下文中获取用户ID
func GetUserId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v := ctx.Value(consts.CtxKeyUserId)
	if v == nil {
		return ""
	}
	if userId, ok := v.(string); ok {
		return userId
	}
	return ""
}

// GetRealname 从上下文中获取真实姓名
func GetRealname(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v := ctx.Value(consts.CtxKeyRealname)
	if v == nil {
		return ""
	}
	if realname, ok := v.(string); ok {
		return realname
	}
	return ""
}
