package vo

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" v:"required"`
	Password string `json:"password" v:"required"`
	Captcha  string `json:"captcha"  v:"required"`
	CheckKey string `json:"checkKey" v:"required"`
}

// LoginRes 登录响应
type LoginRes struct {
	Token        string   `json:"token"`
	UserInfo     UserInfo `json:"userInfo"`
	SysAllDictItems map[string][]DictItem `json:"sysAllDictItems,omitempty"`
}

// UserInfo 用户信息
type UserInfo struct {
	Id           string   `json:"id"`
	Username     string   `json:"username"`
	Realname     string   `json:"realname"`
	Avatar       string   `json:"avatar"`
	Sex          int      `json:"sex"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	OrgCode      string   `json:"orgCode"`
	Status       int      `json:"status"`
	UserIdentity int      `json:"userIdentity"`
	RoleCodes    []string `json:"roleCodes"`
	DepartIds    []string `json:"departIds"`
}

// PhoneLoginReq 手机号登录请求
type PhoneLoginReq struct {
	Phone   string `json:"phone"   v:"required|phone"`
	Captcha string `json:"captcha" v:"required"`
}

// LogoutReq 退出登录请求
type LogoutReq struct {
	Token string `json:"token"`
}

// CaptchaRes 验证码响应
type CaptchaRes struct {
	Key   string `json:"key"`
	Image string `json:"image"`
}

// SendSmsReq 发送短信请求
type SendSmsReq struct {
	Phone   string `json:"phone"   v:"required|phone"`
	SmsMode string `json:"smsMode" v:"required"` // 0登录 1注册 2忘记密码
}

// DictItem 字典项
type DictItem struct {
	Text  string `json:"text"`
	Value string `json:"value"`
	Title string `json:"title,omitempty"`
}
