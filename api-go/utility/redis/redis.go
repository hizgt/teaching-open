package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

// Redis Redis工具类
type Redis struct {
	client *gredis.Redis
}

// New 创建Redis实例
func New() *Redis {
	return &Redis{
		client: g.Redis(),
	}
}

// Set 设置键值
func (r *Redis) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	_, err := r.client.Set(ctx, key, value)
	if err != nil {
		return err
	}
	if ttl > 0 {
		_, err = r.client.Expire(ctx, key, int64(ttl.Seconds()))
	}
	return err
}

// Get 获取值
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

// Del 删除键
func (r *Redis) Del(ctx context.Context, keys ...string) error {
	_, err := r.client.Del(ctx, keys...)
	return err
}

// Exists 检查键是否存在
func (r *Redis) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key)
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// Expire 设置过期时间
func (r *Redis) Expire(ctx context.Context, key string, ttl time.Duration) error {
	_, err := r.client.Expire(ctx, key, int64(ttl.Seconds()))
	return err
}

// TTL 获取剩余过期时间
func (r *Redis) TTL(ctx context.Context, key string) (time.Duration, error) {
	result, err := r.client.TTL(ctx, key)
	if err != nil {
		return 0, err
	}
	return time.Duration(result) * time.Second, nil
}

// Token相关操作

// TokenKey 生成Token缓存键
func TokenKey(token string) string {
	return fmt.Sprintf("token:%s", token)
}

// UserTokenKey 生成用户Token缓存键
func UserTokenKey(userId string) string {
	return fmt.Sprintf("user_token:%s", userId)
}

// SetToken 缓存Token
func (r *Redis) SetToken(ctx context.Context, token, userId string, ttl time.Duration) error {
	// 存储token -> userId映射
	err := r.Set(ctx, TokenKey(token), userId, ttl)
	if err != nil {
		return err
	}
	// 存储userId -> token映射 (用于单设备登录)
	return r.Set(ctx, UserTokenKey(userId), token, ttl)
}

// GetTokenUserId 获取Token对应的用户ID
func (r *Redis) GetTokenUserId(ctx context.Context, token string) (string, error) {
	return r.Get(ctx, TokenKey(token))
}

// ValidateToken 验证Token是否有效
func (r *Redis) ValidateToken(ctx context.Context, token string) (bool, error) {
	return r.Exists(ctx, TokenKey(token))
}

// InvalidateToken 使Token失效
func (r *Redis) InvalidateToken(ctx context.Context, token string) error {
	// 获取用户ID
	userId, err := r.GetTokenUserId(ctx, token)
	if err == nil && userId != "" {
		// 删除用户Token映射
		r.Del(ctx, UserTokenKey(userId))
	}
	// 删除Token
	return r.Del(ctx, TokenKey(token))
}

// InvalidateUserTokens 使用户所有Token失效
func (r *Redis) InvalidateUserTokens(ctx context.Context, userId string) error {
	// 获取用户当前Token
	token, err := r.Get(ctx, UserTokenKey(userId))
	if err == nil && token != "" {
		r.Del(ctx, TokenKey(token))
	}
	return r.Del(ctx, UserTokenKey(userId))
}

// RefreshToken 刷新Token过期时间
func (r *Redis) RefreshToken(ctx context.Context, token string, ttl time.Duration) error {
	err := r.Expire(ctx, TokenKey(token), ttl)
	if err != nil {
		return err
	}
	// 同时刷新用户Token映射
	userId, err := r.GetTokenUserId(ctx, token)
	if err == nil && userId != "" {
		r.Expire(ctx, UserTokenKey(userId), ttl)
	}
	return nil
}

// 权限缓存相关操作

// PermissionKey 生成权限缓存键
func PermissionKey(userId string) string {
	return fmt.Sprintf("permission:%s", userId)
}

// SetUserPermissions 缓存用户权限
func (r *Redis) SetUserPermissions(ctx context.Context, userId string, permissions []string, ttl time.Duration) error {
	return r.Set(ctx, PermissionKey(userId), permissions, ttl)
}

// GetUserPermissions 获取用户权限
func (r *Redis) GetUserPermissions(ctx context.Context, userId string) ([]string, error) {
	result, err := r.client.Get(ctx, PermissionKey(userId))
	if err != nil {
		return nil, err
	}
	return result.Strings(), nil
}

// InvalidateUserPermissions 使用户权限缓存失效
func (r *Redis) InvalidateUserPermissions(ctx context.Context, userId string) error {
	return r.Del(ctx, PermissionKey(userId))
}

// 字典缓存相关操作

// DictKey 生成字典缓存键
func DictKey(dictCode string) string {
	return fmt.Sprintf("dict:%s", dictCode)
}

// SetDictItems 缓存字典项
func (r *Redis) SetDictItems(ctx context.Context, dictCode string, items interface{}, ttl time.Duration) error {
	return r.Set(ctx, DictKey(dictCode), items, ttl)
}

// GetDictItems 获取字典项
func (r *Redis) GetDictItems(ctx context.Context, dictCode string) (string, error) {
	return r.Get(ctx, DictKey(dictCode))
}

// InvalidateDictCache 使字典缓存失效
func (r *Redis) InvalidateDictCache(ctx context.Context, dictCode string) error {
	return r.Del(ctx, DictKey(dictCode))
}

// 验证码相关操作

// CaptchaKey 生成验证码缓存键
func CaptchaKey(key string) string {
	return fmt.Sprintf("captcha:%s", key)
}

// SetCaptcha 缓存验证码
func (r *Redis) SetCaptcha(ctx context.Context, key, code string, ttl time.Duration) error {
	return r.Set(ctx, CaptchaKey(key), code, ttl)
}

// GetCaptcha 获取验证码
func (r *Redis) GetCaptcha(ctx context.Context, key string) (string, error) {
	return r.Get(ctx, CaptchaKey(key))
}

// ValidateCaptcha 验证验证码
func (r *Redis) ValidateCaptcha(ctx context.Context, key, code string) (bool, error) {
	storedCode, err := r.GetCaptcha(ctx, key)
	if err != nil {
		return false, err
	}
	return storedCode == code, nil
}

// DeleteCaptcha 删除验证码
func (r *Redis) DeleteCaptcha(ctx context.Context, key string) error {
	return r.Del(ctx, CaptchaKey(key))
}
