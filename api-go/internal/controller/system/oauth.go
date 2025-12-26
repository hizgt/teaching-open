package system

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
	"teaching-open/internal/service"
	"teaching-open/utility/password"
	"teaching-open/utility/response"
)

type OAuthController struct{}

func NewOAuth() *OAuthController {
	return &OAuthController{}
}

// WechatAuthURL 获取微信授权URL
func (c *OAuthController) WechatAuthURL(r *ghttp.Request) {
	ctx := r.Context()
	state := guid.S()

	// 将state存入Redis，5分钟有效
	g.Redis().Set(ctx, "oauth:wechat:state:"+state, "1")
	g.Redis().Expire(ctx, "oauth:wechat:state:"+state, 300)

	authURL, err := service.OAuthWechat().GetAuthURL(ctx, state)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, g.Map{
		"authUrl": authURL,
		"state":   state,
	})
}

// WechatCallback 微信登录回调
func (c *OAuthController) WechatCallback(r *ghttp.Request) {
	ctx := r.Context()
	code := r.Get("code").String()
	state := r.Get("state").String()

	// 验证state
	stateValue, err := g.Redis().Get(ctx, "oauth:wechat:state:"+state)
	if err != nil || stateValue.IsEmpty() {
		response.Error(r, "state验证失败")
		return
	}
	g.Redis().Del(ctx, "oauth:wechat:state:"+state)

	// 获取access_token
	tokenInfo, err := service.OAuthWechat().GetAccessToken(ctx, code)
	if err != nil {
		response.Error(r, "获取access_token失败: "+err.Error())
		return
	}

	// 获取用户信息
	userInfo, err := service.OAuthWechat().GetUserInfo(ctx, tokenInfo.AccessToken, tokenInfo.Openid)
	if err != nil {
		response.Error(r, "获取用户信息失败: "+err.Error())
		return
	}

	// 查找是否已绑定用户
	userId, err := service.OAuthWechat().GetUserByOpenid(ctx, userInfo.Openid)
	if err != nil || userId == "" {
		// 未绑定，创建新用户
		userId, err = service.OAuthWechat().CreateUserFromWechat(ctx, userInfo)
		if err != nil {
			response.Error(r, "创建用户失败: "+err.Error())
			return
		}
	}

	// 生成JWT token
	user, err := service.SysUser().GetUserById(ctx, userId)
	if err != nil {
		response.Error(r, "获取用户信息失败")
		return
	}

	token, err := service.Login().GenerateToken(ctx, user)
	if err != nil {
		response.Error(r, "生成token失败")
		return
	}

	response.Success(r, g.Map{
		"token":    token,
		"userInfo": user,
	})
}

// WechatBind 绑定微信账号
func (c *OAuthController) WechatBind(r *ghttp.Request) {
	ctx := r.Context()
	userId := r.GetCtxVar("userId").String()
	code := r.Get("code").String()

	// 获取access_token
	tokenInfo, err := service.OAuthWechat().GetAccessToken(ctx, code)
	if err != nil {
		response.Error(r, "获取access_token失败: "+err.Error())
		return
	}

	// 绑定用户
	err = service.OAuthWechat().BindUser(ctx, userId, tokenInfo.Openid)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, "绑定成功")
}

// WechatUnbind 解绑微信账号
func (c *OAuthController) WechatUnbind(r *ghttp.Request) {
	ctx := r.Context()
	userId := r.GetCtxVar("userId").String()

	err := service.OAuthWechat().UnbindUser(ctx, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, "解绑成功")
}

// QQAuthURL 获取QQ授权URL
func (c *OAuthController) QQAuthURL(r *ghttp.Request) {
	ctx := r.Context()
	state := guid.S()

	// 将state存入Redis，5分钟有效
	g.Redis().Set(ctx, "oauth:qq:state:"+state, "1")
	g.Redis().Expire(ctx, "oauth:qq:state:"+state, 300)

	authURL, err := service.OAuthQQ().GetAuthURL(ctx, state)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, g.Map{
		"authUrl": authURL,
		"state":   state,
	})
}

// QQCallback QQ登录回调
func (c *OAuthController) QQCallback(r *ghttp.Request) {
	ctx := r.Context()
	code := r.Get("code").String()
	state := r.Get("state").String()

	// 验证state
	stateValue, err := g.Redis().Get(ctx, "oauth:qq:state:"+state)
	if err != nil || stateValue.IsEmpty() {
		response.Error(r, "state验证失败")
		return
	}
	g.Redis().Del(ctx, "oauth:qq:state:"+state)

	// 获取access_token
	accessToken, err := service.OAuthQQ().GetAccessToken(ctx, code)
	if err != nil {
		response.Error(r, "获取access_token失败: "+err.Error())
		return
	}

	// 获取openid
	openid, err := service.OAuthQQ().GetOpenid(ctx, accessToken)
	if err != nil {
		response.Error(r, "获取openid失败: "+err.Error())
		return
	}

	// 获取用户信息
	userInfo, err := service.OAuthQQ().GetUserInfo(ctx, accessToken, openid)
	if err != nil {
		response.Error(r, "获取用户信息失败: "+err.Error())
		return
	}

	// 查找是否已绑定用户
	userId, err := service.OAuthQQ().GetUserByOpenid(ctx, openid)
	if err != nil || userId == "" {
		// 未绑定，创建新用户
		userId, err = service.OAuthQQ().CreateUserFromQQ(ctx, openid, userInfo)
		if err != nil {
			response.Error(r, "创建用户失败: "+err.Error())
			return
		}
	}

	// 生成JWT token
	user, err := service.SysUser().GetUserById(ctx, userId)
	if err != nil {
		response.Error(r, "获取用户信息失败")
		return
	}

	token, err := service.Login().GenerateToken(ctx, user)
	if err != nil {
		response.Error(r, "生成token失败")
		return
	}

	response.Success(r, g.Map{
		"token":    token,
		"userInfo": user,
	})
}

// QQBind 绑定QQ账号
func (c *OAuthController) QQBind(r *ghttp.Request) {
	ctx := r.Context()
	userId := r.GetCtxVar("userId").String()
	code := r.Get("code").String()

	// 获取access_token
	accessToken, err := service.OAuthQQ().GetAccessToken(ctx, code)
	if err != nil {
		response.Error(r, "获取access_token失败: "+err.Error())
		return
	}

	// 获取openid
	openid, err := service.OAuthQQ().GetOpenid(ctx, accessToken)
	if err != nil {
		response.Error(r, "获取openid失败: "+err.Error())
		return
	}

	// 绑定用户
	err = service.OAuthQQ().BindUser(ctx, userId, openid)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, "绑定成功")
}

// QQUnbind 解绑QQ账号
func (c *OAuthController) QQUnbind(r *ghttp.Request) {
	ctx := r.Context()
	userId := r.GetCtxVar("userId").String()

	err := service.OAuthQQ().UnbindUser(ctx, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, "解绑成功")
}
