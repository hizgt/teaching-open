package password

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateSalt 生成随机盐值
func GenerateSalt() string {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		// 如果随机生成失败，使用时间戳作为备选
		return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", time.Now().UnixNano()))))[:16]
	}
	return hex.EncodeToString(bytes)
}

// Encrypt 使用MD5+Salt加密密码
// 加密方式与Java版本保持一致: MD5(username + password + salt)
func Encrypt(username, password, salt string) string {
	data := username + password + salt
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// EncryptSimple 简单MD5加密 (不含用户名)
// 加密方式: MD5(password + salt)
func EncryptSimple(password, salt string) string {
	data := password + salt
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Verify 验证密码
func Verify(username, password, salt, encryptedPassword string) bool {
	return Encrypt(username, password, salt) == encryptedPassword
}

// VerifySimple 简单验证密码 (不含用户名)
func VerifySimple(password, salt, encryptedPassword string) bool {
	return EncryptSimple(password, salt) == encryptedPassword
}

// MD5 计算字符串的MD5值
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
