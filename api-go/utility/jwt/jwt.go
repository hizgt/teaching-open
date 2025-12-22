package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gogf/gf/v2/frame/g"
)

// Claims JWT声明
type Claims struct {
	UserId     string   `json:"userId"`
	Username   string   `json:"username"`
	RoleIds    []string `json:"roleIds"`
	DeviceId   string   `json:"deviceId"`
	DeviceType string   `json:"deviceType"`
	jwt.RegisteredClaims
}

// TokenPair Token对
type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// JWT JWT工具类
type JWT struct{}

// New 创建JWT实例
func New() *JWT {
	return &JWT{}
}

// GenerateTokenPair 生成Token对
func (j *JWT) GenerateTokenPair(userId, username string, roleIds []string, deviceId, deviceType string) (*TokenPair, error) {
	// 生成访问Token
	accessToken, err := j.GenerateAccessToken(userId, username, roleIds, deviceId, deviceType)
	if err != nil {
		return nil, err
	}

	// 生成刷新Token
	refreshToken, err := j.GenerateRefreshToken(userId, deviceId)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// GenerateAccessToken 生成访问Token
func (j *JWT) GenerateAccessToken(userId, username string, roleIds []string, deviceId, deviceType string) (string, error) {
	expireDuration := g.Cfg().MustGet(context.Background(), "jwt.expire").Duration()
	if expireDuration == 0 {
		expireDuration = 2 * time.Hour // 默认2小时
	}

	claims := Claims{
		UserId:     userId,
		Username:   username,
		RoleIds:    roleIds,
		DeviceId:   deviceId,
		DeviceType: deviceType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "teaching-open",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := g.Cfg().MustGet(context.Background(), "jwt.secret").String()
	if secret == "" {
		secret = "teaching-open-secret-key-change-in-production"
	}

	return token.SignedString([]byte(secret))
}

// GenerateRefreshToken 生成刷新Token
func (j *JWT) GenerateRefreshToken(userId, deviceId string) (string, error) {
	refreshExpire := g.Cfg().MustGet(context.Background(), "jwt.refreshExpire").Duration()
	if refreshExpire == 0 {
		refreshExpire = 168 * time.Hour // 默认7天
	}

	claims := jwt.RegisteredClaims{
		Subject:   userId,
		ID:        deviceId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpire)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "teaching-open-refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := g.Cfg().MustGet(context.Background(), "jwt.secret").String()
	if secret == "" {
		secret = "teaching-open-secret-key-change-in-production"
	}

	return token.SignedString([]byte(secret))
}

// ValidateToken 验证Token
func (j *JWT) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		secret := g.Cfg().MustGet(context.Background(), "jwt.secret").String()
		if secret == "" {
			secret = "teaching-open-secret-key-change-in-production"
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ValidateRefreshToken 验证刷新Token
func (j *JWT) ValidateRefreshToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := g.Cfg().MustGet(context.Background(), "jwt.secret").String()
		if secret == "" {
			secret = "teaching-open-secret-key-change-in-production"
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid refresh token")
}

// ExtractTokenFromHeader 从Authorization头提取Token
func (j *JWT) ExtractTokenFromHeader(authHeader string) (string, error) {
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return "", fmt.Errorf("invalid authorization header format")
	}
	return authHeader[7:], nil
}

// RefreshAccessToken 使用刷新Token生成新的访问Token
func (j *JWT) RefreshAccessToken(refreshToken, userId, username string, roleIds []string, deviceId, deviceType string) (string, error) {
	// 验证刷新Token
	_, err := j.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	// 生成新的访问Token
	return j.GenerateAccessToken(userId, username, roleIds, deviceId, deviceType)
}