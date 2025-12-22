package jwt

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Claims JWT声明结构体
type Claims struct {
	UserId     string   `json:"userId"`
	Username   string   `json:"username"`
	RoleIds    []string `json:"roleIds"`
	DeviceId   string   `json:"deviceId"`   // 设备ID，用于多设备登录限制
	DeviceType string   `json:"deviceType"` // 设备类型：web, mobile, pc等
	jwt.RegisteredClaims
}

// RefreshClaims 刷新Token声明
type RefreshClaims struct {
	UserId   string `json:"userId"`
	DeviceId string `json:"deviceId"`
	jwt.RegisteredClaims
}

// TokenPair Token对
type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
	TokenType    string `json:"tokenType"`
}

// JWT工具类
type JWTUtil struct{}

// New 创建JWT工具实例
func New() *JWTUtil {
	return &JWTUtil{}
}

// getSecret 获取JWT密钥
func (j *JWTUtil) getSecret() []byte {
	secret := g.Cfg().MustGet(context.Background(), "jwt.secret").String()
	if secret == "" {
		secret = "teaching-open-default-secret-key"
	}
	return []byte(secret)
}

// getExpireTime 获取Token过期时间
func (j *JWTUtil) getExpireTime() time.Duration {
	expireStr := g.Cfg().MustGet(context.Background(), "jwt.expire").String()
	if expireStr == "" {
		return 2 * time.Hour // 默认2小时
	}
	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		return 2 * time.Hour
	}
	return expire
}

// getRefreshExpireTime 获取刷新Token过期时间
func (j *JWTUtil) getRefreshExpireTime() time.Duration {
	refreshExpireStr := g.Cfg().MustGet(context.Background(), "jwt.refreshExpire").String()
	if refreshExpireStr == "" {
		return 7 * 24 * time.Hour // 默认7天
	}
	expire, err := time.ParseDuration(refreshExpireStr)
	if err != nil {
		return 7 * 24 * time.Hour
	}
	return expire
}

// GenerateToken 生成访问Token
func (j *JWTUtil) GenerateToken(userId, username string, roleIds []string, deviceId, deviceType string) (string, error) {
	if deviceId == "" {
		deviceId = uuid.New().String()
	}

	claims := Claims{
		UserId:     userId,
		Username:   username,
		RoleIds:    roleIds,
		DeviceId:   deviceId,
		DeviceType: deviceType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.getExpireTime())),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "teaching-open",
			Subject:   userId,
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.getSecret())
}

// GenerateRefreshToken 生成刷新Token
func (j *JWTUtil) GenerateRefreshToken(userId, deviceId string) (string, error) {
	claims := RefreshClaims{
		UserId:   userId,
		DeviceId: deviceId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.getRefreshExpireTime())),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "teaching-open",
			Subject:   userId,
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.getSecret())
}

// GenerateTokenPair 生成Token对
func (j *JWTUtil) GenerateTokenPair(userId, username string, roleIds []string, deviceId, deviceType string) (*TokenPair, error) {
	if deviceId == "" {
		deviceId = uuid.New().String()
	}

	accessToken, err := j.GenerateToken(userId, username, roleIds, deviceId, deviceType)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.GenerateRefreshToken(userId, deviceId)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(j.getExpireTime().Seconds()),
		TokenType:    "Bearer",
	}, nil
}

// ParseToken 解析访问Token
func (j *JWTUtil) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.getSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ParseRefreshToken 解析刷新Token
func (j *JWTUtil) ParseRefreshToken(tokenString string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.getSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid refresh token")
}

// RefreshToken 刷新Token
func (j *JWTUtil) RefreshToken(refreshTokenString string) (*TokenPair, error) {
	_, err := j.ParseRefreshToken(refreshTokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %v", err)
	}

	// 从Redis获取用户信息（需要业务层配合）
	// 这里返回nil，需要在业务层实现
	return nil, fmt.Errorf("refresh token requires user info from cache")
}

// ValidateToken 验证Token有效性
func (j *JWTUtil) ValidateToken(tokenString string) (*Claims, error) {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	// 检查是否过期
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}

// ExtractTokenFromHeader 从请求头提取Token
func (j *JWTUtil) ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}

// IsTokenExpired 检查Token是否过期
func (j *JWTUtil) IsTokenExpired(tokenString string) bool {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return true
	}

	if claims.ExpiresAt == nil {
		return false
	}

	return claims.ExpiresAt.Before(time.Now())
}

// GetRemainingTime 获取Token剩余时间（秒）
func (j *JWTUtil) GetRemainingTime(tokenString string) (int64, error) {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return 0, err
	}

	if claims.ExpiresAt == nil {
		return 0, fmt.Errorf("token has no expiration time")
	}

	remaining := claims.ExpiresAt.Time.Sub(time.Now())
	if remaining < 0 {
		return 0, nil
	}

	return int64(remaining.Seconds()), nil
}