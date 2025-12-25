package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"teaching-open/internal/model/vo"
	"teaching-open/utility/jwt"
	"teaching-open/utility/password"
	"teaching-open/utility/redis"
)

// LoginService 登录服务接口
type LoginService interface {
	Login(ctx context.Context, req *vo.LoginReq) (*vo.LoginRes, error)
	PhoneLogin(ctx context.Context, req *vo.PhoneLoginReq) (*vo.LoginRes, error)
	Logout(ctx context.Context, token string) error
	GetCaptcha(ctx context.Context, key string) (*vo.CaptchaRes, error)
	ValidateCaptcha(ctx context.Context, key, code string) (bool, error)
	SendSms(ctx context.Context, req *vo.SendSmsReq) error
}

// loginServiceImpl 登录服务实现
type loginServiceImpl struct {
	userService SysUserService
	dictService SysDictService
	redisUtil   *redis.Redis
	jwtUtil     *jwt.JWT
}

// NewLoginService 创建登录服务实例
func NewLoginService() LoginService {
	return &loginServiceImpl{
		userService: NewSysUserService(),
		dictService: NewSysDictService(),
		redisUtil:   redis.New(),
		jwtUtil:     jwt.New(),
	}
}

// Login 用户名密码登录
func (s *loginServiceImpl) Login(ctx context.Context, req *vo.LoginReq) (*vo.LoginRes, error) {
	// 验证验证码
	valid, err := s.ValidateCaptcha(ctx, req.CheckKey, req.Captcha)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, gerror.New("验证码错误")
	}

	// 删除已使用的验证码
	s.redisUtil.DeleteCaptcha(ctx, req.CheckKey)

	// 获取用户
	user, err := s.userService.GetUserByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户名或密码错误")
	}

	// 检查用户有效性
	err = s.userService.CheckUserIsEffective(ctx, user)
	if err != nil {
		return nil, err
	}

	// 验证密码
	if !password.Verify(user.Username, req.Password, user.Salt, user.Password) {
		return nil, gerror.New("用户名或密码错误")
	}

	// 获取用户角色
	roleCodes, err := s.userService.GetUserRolesSet(ctx, user.Username)
	if err != nil {
		g.Log().Warning(ctx, "获取用户角色失败", err)
	}

	// 生成Token
	tokenPair, err := s.jwtUtil.GenerateTokenPair(user.Id, user.Username, roleCodes, "", "web")
	if err != nil {
		return nil, gerror.New("生成Token失败")
	}

	// 缓存Token
	expireDuration := g.Cfg().MustGet(ctx, "jwt.expire").Duration()
	if expireDuration == 0 {
		expireDuration = 2 * time.Hour
	}
	err = s.redisUtil.SetToken(ctx, tokenPair.AccessToken, user.Id, expireDuration)
	if err != nil {
		g.Log().Warning(ctx, "缓存Token失败", err)
	}

	// 获取所有字典项
	dictItems, err := s.dictService.QueryAllDictItems(ctx)
	if err != nil {
		g.Log().Warning(ctx, "获取字典项失败", err)
	}

	// 构建响应
	return &vo.LoginRes{
		Token: tokenPair.AccessToken,
		UserInfo: vo.UserInfo{
			Id:           user.Id,
			Username:     user.Username,
			Realname:     user.Realname,
			Avatar:       user.Avatar,
			Sex:          user.Sex,
			Email:        user.Email,
			Phone:        user.Phone,
			OrgCode:      user.OrgCode,
			Status:       user.Status,
			UserIdentity: user.UserIdentity,
			RoleCodes:    roleCodes,
		},
		SysAllDictItems: dictItems,
	}, nil
}

// PhoneLogin 手机号登录
func (s *loginServiceImpl) PhoneLogin(ctx context.Context, req *vo.PhoneLoginReq) (*vo.LoginRes, error) {
	// 验证短信验证码
	valid, err := s.redisUtil.ValidateCaptcha(ctx, "sms:"+req.Phone, req.Captcha)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, gerror.New("验证码错误或已过期")
	}

	// 删除已使用的验证码
	s.redisUtil.DeleteCaptcha(ctx, "sms:"+req.Phone)

	// 获取用户
	user, err := s.userService.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("手机号未注册")
	}

	// 检查用户有效性
	err = s.userService.CheckUserIsEffective(ctx, user)
	if err != nil {
		return nil, err
	}

	// 获取用户角色
	roleCodes, err := s.userService.GetUserRolesSet(ctx, user.Username)
	if err != nil {
		g.Log().Warning(ctx, "获取用户角色失败", err)
	}

	// 生成Token
	tokenPair, err := s.jwtUtil.GenerateTokenPair(user.Id, user.Username, roleCodes, "", "web")
	if err != nil {
		return nil, gerror.New("生成Token失败")
	}

	// 缓存Token
	expireDuration := g.Cfg().MustGet(ctx, "jwt.expire").Duration()
	if expireDuration == 0 {
		expireDuration = 2 * time.Hour
	}
	err = s.redisUtil.SetToken(ctx, tokenPair.AccessToken, user.Id, expireDuration)
	if err != nil {
		g.Log().Warning(ctx, "缓存Token失败", err)
	}

	// 获取所有字典项
	dictItems, err := s.dictService.QueryAllDictItems(ctx)
	if err != nil {
		g.Log().Warning(ctx, "获取字典项失败", err)
	}

	return &vo.LoginRes{
		Token: tokenPair.AccessToken,
		UserInfo: vo.UserInfo{
			Id:           user.Id,
			Username:     user.Username,
			Realname:     user.Realname,
			Avatar:       user.Avatar,
			Sex:          user.Sex,
			Email:        user.Email,
			Phone:        user.Phone,
			OrgCode:      user.OrgCode,
			Status:       user.Status,
			UserIdentity: user.UserIdentity,
			RoleCodes:    roleCodes,
		},
		SysAllDictItems: dictItems,
	}, nil
}

// Logout 退出登录
func (s *loginServiceImpl) Logout(ctx context.Context, token string) error {
	return s.redisUtil.InvalidateToken(ctx, token)
}

// GetCaptcha 获取验证码
func (s *loginServiceImpl) GetCaptcha(ctx context.Context, key string) (*vo.CaptchaRes, error) {
	// 生成随机验证码 (简化实现，实际应使用图形验证码库)
	code := generateRandomCode(4)

	// 缓存验证码，5分钟有效
	err := s.redisUtil.SetCaptcha(ctx, key, code, 5*time.Minute)
	if err != nil {
		return nil, err
	}

	// 返回验证码图片 (简化实现，返回Base64编码的图片)
	// 实际应使用验证码生成库如 github.com/mojocn/base64Captcha
	return &vo.CaptchaRes{
		Key:   key,
		Image: "data:image/png;base64,..." + code, // 简化，实际应生成图片
	}, nil
}

// ValidateCaptcha 验证验证码
func (s *loginServiceImpl) ValidateCaptcha(ctx context.Context, key, code string) (bool, error) {
	return s.redisUtil.ValidateCaptcha(ctx, key, code)
}

// SendSms 发送短信验证码
func (s *loginServiceImpl) SendSms(ctx context.Context, req *vo.SendSmsReq) error {
	// 生成6位数字验证码
	code := generateRandomCode(6)

	// 缓存验证码，5分钟有效
	err := s.redisUtil.SetCaptcha(ctx, "sms:"+req.Phone, code, 5*time.Minute)
	if err != nil {
		return err
	}

	// 发送短信 (需要集成短信服务商SDK)
	// 这里只是示例，实际需要调用阿里云短信服务等
	g.Log().Info(ctx, "发送短信验证码", g.Map{
		"phone": req.Phone,
		"code":  code,
		"mode":  req.SmsMode,
	})

	return nil
}

// generateRandomCode 生成随机验证码
func generateRandomCode(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[time.Now().UnixNano()%int64(len(charset))]
		time.Sleep(time.Nanosecond)
	}
	return string(code)
}
