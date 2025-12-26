package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

type IOAuthWechatService interface {
	GetAuthURL(ctx context.Context, state string) (string, error)
	GetAccessToken(ctx context.Context, code string) (*WechatAccessToken, error)
	GetUserInfo(ctx context.Context, accessToken, openid string) (*WechatUserInfo, error)
	BindUser(ctx context.Context, userId, openid string) error
	UnbindUser(ctx context.Context, userId string) error
	GetUserByOpenid(ctx context.Context, openid string) (string, error)
}

type WechatAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}

type WechatUserInfo struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

type sOAuthWechat struct{}

func OAuthWechat() IOAuthWechatService {
	return &sOAuthWechat{}
}

// GetAuthURL 获取微信授权URL
func (s *sOAuthWechat) GetAuthURL(ctx context.Context, state string) (string, error) {
	appId := g.Cfg().MustGet(ctx, "oauth.wechat.appId").String()
	redirectUri := g.Cfg().MustGet(ctx, "oauth.wechat.redirectUri").String()

	if appId == "" || redirectUri == "" {
		return "", fmt.Errorf("微信登录配置不完整")
	}

	authURL := fmt.Sprintf(
		"https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=%s#wechat_redirect",
		appId, redirectUri, state,
	)

	return authURL, nil
}

// GetAccessToken 通过code获取access_token
func (s *sOAuthWechat) GetAccessToken(ctx context.Context, code string) (*WechatAccessToken, error) {
	appId := g.Cfg().MustGet(ctx, "oauth.wechat.appId").String()
	appSecret := g.Cfg().MustGet(ctx, "oauth.wechat.appSecret").String()

	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appId, appSecret, code,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result WechatAccessToken
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.AccessToken == "" {
		return nil, fmt.Errorf("获取access_token失败: %s", string(body))
	}

	return &result, nil
}

// GetUserInfo 获取微信用户信息
func (s *sOAuthWechat) GetUserInfo(ctx context.Context, accessToken, openid string) (*WechatUserInfo, error) {
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
		accessToken, openid,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo WechatUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	if userInfo.Openid == "" {
		return nil, fmt.Errorf("获取用户信息失败: %s", string(body))
	}

	return &userInfo, nil
}

// BindUser 绑定用户和微信openid
func (s *sOAuthWechat) BindUser(ctx context.Context, userId, openid string) error {
	// 检查openid是否已被绑定
	existUserId, _ := s.GetUserByOpenid(ctx, openid)
	if existUserId != "" && existUserId != userId {
		return fmt.Errorf("该微信账号已被其他用户绑定")
	}

	// 更新用户表的third_id字段
	_, err := g.DB().Model("sys_user").
		Where("id", userId).
		Data(g.Map{
			"third_id":    openid,
			"third_type":  "wechat",
			"update_time": time.Now(),
		}).
		Update()

	return err
}

// UnbindUser 解绑用户和微信
func (s *sOAuthWechat) UnbindUser(ctx context.Context, userId string) error {
	_, err := g.DB().Model("sys_user").
		Where("id", userId).
		Data(g.Map{
			"third_id":    "",
			"third_type":  "",
			"update_time": time.Now(),
		}).
		Update()

	return err
}

// GetUserByOpenid 通过openid获取用户ID
func (s *sOAuthWechat) GetUserByOpenid(ctx context.Context, openid string) (string, error) {
	var userId string
	err := g.DB().Model("sys_user").
		Where("third_id", openid).
		Where("third_type", "wechat").
		Where("del_flag", 0).
		Fields("id").
		Scan(&userId)

	return userId, err
}

// CreateUserFromWechat 从微信信息创建新用户
func (s *sOAuthWechat) CreateUserFromWechat(ctx context.Context, userInfo *WechatUserInfo) (string, error) {
	userId := guid.S()
	username := "wx_" + userInfo.Openid[len(userInfo.Openid)-8:]

	_, err := g.DB().Model("sys_user").Data(g.Map{
		"id":          userId,
		"username":    username,
		"realname":    userInfo.Nickname,
		"avatar":      userInfo.Headimgurl,
		"sex":         userInfo.Sex,
		"third_id":    userInfo.Openid,
		"third_type":  "wechat",
		"status":      1,
		"del_flag":    0,
		"create_time": time.Now(),
		"update_time": time.Now(),
	}).Insert()

	if err != nil {
		return "", err
	}

	return userId, nil
}
