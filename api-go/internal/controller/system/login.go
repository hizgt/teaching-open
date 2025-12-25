package system

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
)

// LoginController 登录控制器
type LoginController struct {
	loginService service.LoginService
}

// NewLoginController 创建登录控制器实例
func NewLoginController() *LoginController {
	return &LoginController{
		loginService: service.NewLoginService(),
	}
}

// Login 用户登录
// @Summary 用户登录
// @Tags 系统-登录
// @Accept json
// @Produce json
// @Param data body vo.LoginReq true "登录参数"
// @Success 200 {object} vo.LoginRes
// @Router /sys/login [post]
func (c *LoginController) Login(r *ghttp.Request) {
	var req vo.LoginReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	res, err := c.loginService.Login(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// PhoneLogin 手机号登录
// @Summary 手机号登录
// @Tags 系统-登录
// @Accept json
// @Produce json
// @Param data body vo.PhoneLoginReq true "手机号登录参数"
// @Success 200 {object} vo.LoginRes
// @Router /sys/phoneLogin [post]
func (c *LoginController) PhoneLogin(r *ghttp.Request) {
	var req vo.PhoneLoginReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	res, err := c.loginService.PhoneLogin(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// Logout 退出登录
// @Summary 退出登录
// @Tags 系统-登录
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /sys/logout [post]
func (c *LoginController) Logout(r *ghttp.Request) {
	token := r.Header.Get("X-Access-Token")
	if token == "" {
		token = r.Header.Get("Authorization")
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}
	}

	err := c.loginService.Logout(r.Context(), token)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r)
}

// GetCaptcha 获取验证码
// @Summary 获取验证码
// @Tags 系统-登录
// @Accept json
// @Produce json
// @Param key path string true "验证码key"
// @Success 200 {object} vo.CaptchaRes
// @Router /sys/randomImage/{key} [get]
func (c *LoginController) GetCaptcha(r *ghttp.Request) {
	key := r.Get("key").String()
	if key == "" {
		response.Error(r, "key不能为空")
		return
	}

	res, err := c.loginService.GetCaptcha(r.Context(), key)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// SendSms 发送短信验证码
// @Summary 发送短信验证码
// @Tags 系统-登录
// @Accept json
// @Produce json
// @Param data body vo.SendSmsReq true "发送短信参数"
// @Success 200 {object} response.Response
// @Router /sys/sms [post]
func (c *LoginController) SendSms(r *ghttp.Request) {
	var req vo.SendSmsReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	err := c.loginService.SendSms(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r)
}
