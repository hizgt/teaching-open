package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

type IOAuthQQService interface {
	GetAuthURL(ctx context.Context, state string) (string, error)
	GetAccessToken(ctx context.Context, code string) (string, error)
	GetOpenid(ctx context.Context, accessToken string) (string, error)
	GetUserInfo(ctx context.Context, accessToken, openid string) (*QQUserInfo, error)
	BindUser(ctx context.Context, userId, openid string) error
	UnbindUser(ctx context.Context, userId string) error
	GetUserByOpenid(ctx context.Context, openid string) (string, error)
}

type QQUserInfo struct {
	Ret             int    `json:"ret"`
	Msg             string `json:"msg"`
	Nickname        string `json:"nickname"`
	Figureurl       string `json:"figureurl"`
	Figureurl1      string `json:"figureurl_1"`
	Figureurl2      string `json:"figureurl_2"`
	FigureurlQQ1    string `json:"figureurl_qq_1"`
	FigureurlQQ2    string `json:"figureurl_qq_2"`
	Gender          string `json:"gender"`
	IsYellowVip     string `json:"is_yellow_vip"`
	Vip             string `json:"vip"`
	YellowVipLevel  string `json:"yellow_vip_level"`
	Level           string `json:"level"`
	IsYellowYearVip string `json:"is_yellow_year_vip"`
}

type sOAuthQQ struct{}

func OAuthQQ() IOAuthQQService {
	return &sOAuthQQ{}
}

// GetAuthURL 获取QQ授权URL
func (s *sOAuthQQ) GetAuthURL(ctx context.Context, state string) (string, error) {
	appId := g.Cfg().MustGet(ctx, "oauth.qq.appId").String()
	redirectUri := g.Cfg().MustGet(ctx, "oauth.qq.redirectUri").String()

	if appId == "" || redirectUri == "" {
		return "", fmt.Errorf("QQ登录配置不完整")
	}

	authURL := fmt.Sprintf(
		"https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=get_user_info",
		appId, redirectUri, state,
	)

	return authURL, nil
}

// GetAccessToken 通过code获取access_token
func (s *sOAuthQQ) GetAccessToken(ctx context.Context, code string) (string, error) {
	appId := g.Cfg().MustGet(ctx, "oauth.qq.appId").String()
	appKey := g.Cfg().MustGet(ctx, "oauth.qq.appKey").String()
	redirectUri := g.Cfg().MustGet(ctx, "oauth.qq.redirectUri").String()

	url := fmt.Sprintf(
		"https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		appId, appKey, code, redirectUri,
	)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// QQ返回格式: access_token=xxx&expires_in=7776000&refresh_token=xxx
	re := regexp.MustCompile(`access_token=([^&]+)`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", fmt.Errorf("获取access_token失败: %s", string(body))
	}

	return matches[1], nil
}

// GetOpenid 获取用户openid
func (s *sOAuthQQ) GetOpenid(ctx context.Context, accessToken string) (string, error) {
	url := fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", accessToken)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// QQ返回格式: callback( {"client_id":"xxx","openid":"xxx"} );
	re := regexp.MustCompile(`"openid":"([^"]+)"`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", fmt.Errorf("获取openid失败: %s", string(body))
	}

	return matches[1], nil
}

// GetUserInfo 获取QQ用户信息
func (s *sOAuthQQ) GetUserInfo(ctx context.Context, accessToken, openid string) (*QQUserInfo, error) {
	appId := g.Cfg().MustGet(ctx, "oauth.qq.appId").String()

	url := fmt.Sprintf(
		"https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s",
		accessToken, appId, openid,
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

	var userInfo QQUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	if userInfo.Ret != 0 {
		return nil, fmt.Errorf("获取用户信息失败: %s", userInfo.Msg)
	}

	return &userInfo, nil
}

// BindUser 绑定用户和QQ openid
func (s *sOAuthQQ) BindUser(ctx context.Context, userId, openid string) error {
	// 检查openid是否已被绑定
	existUserId, _ := s.GetUserByOpenid(ctx, openid)
	if existUserId != "" && existUserId != userId {
		return fmt.Errorf("该QQ账号已被其他用户绑定")
	}

	// 更新用户表的third_id字段
	_, err := g.DB().Model("sys_user").
		Where("id", userId).
		Data(g.Map{
			"third_id":    openid,
			"third_type":  "qq",
			"update_time": time.Now(),
		}).
		Update()

	return err
}

// UnbindUser 解绑用户和QQ
func (s *sOAuthQQ) UnbindUser(ctx context.Context, userId string) error {
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
func (s *sOAuthQQ) GetUserByOpenid(ctx context.Context, openid string) (string, error) {
	var userId string
	err := g.DB().Model("sys_user").
		Where("third_id", openid).
		Where("third_type", "qq").
		Where("del_flag", 0).
		Fields("id").
		Scan(&userId)

	return userId, err
}

// CreateUserFromQQ 从QQ信息创建新用户
func (s *sOAuthQQ) CreateUserFromQQ(ctx context.Context, openid string, userInfo *QQUserInfo) (string, error) {
	userId := guid.S()
	username := "qq_" + openid[len(openid)-8:]

	sex := 0
	if userInfo.Gender == "男" {
		sex = 1
	} else if userInfo.Gender == "女" {
		sex = 2
	}

	avatar := userInfo.FigureurlQQ2
	if avatar == "" {
		avatar = userInfo.FigureurlQQ1
	}
	if avatar == "" {
		avatar = userInfo.Figureurl2
	}

	_, err := g.DB().Model("sys_user").Data(g.Map{
		"id":          userId,
		"username":    username,
		"realname":    userInfo.Nickname,
		"avatar":      avatar,
		"sex":         sex,
		"third_id":    openid,
		"third_type":  "qq",
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
