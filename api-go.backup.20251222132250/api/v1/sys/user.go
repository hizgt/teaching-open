package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LoginReq 登录请求
type LoginReq struct {
	g.Meta   `path:"/sys/login" method:"post" tags:"系统管理" summary:"用户登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

// LoginRes 登录响应
type LoginRes struct {
	Token    string    `json:"token" dc:"JWT令牌"`
	UserInfo *UserInfo `json:"userInfo" dc:"用户信息"`
}

// UserInfo 用户信息
type UserInfo struct {
	Id       string `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Realname string `json:"realname" dc:"真实姓名"`
	Avatar   string `json:"avatar" dc:"头像"`
	Status   int    `json:"status" dc:"状态(1-正常,2-冻结)"`
	School   string `json:"school" dc:"学校"`
	Phone    string `json:"phone" dc:"电话"`
	Email    string `json:"email" dc:"邮箱"`
	Sex      int    `json:"sex" dc:"性别(0-未知,1-男,2-女)"`
}

// UserListReq 用户列表请求
type UserListReq struct {
	g.Meta   `path:"/sys/user/list" method:"get" tags:"用户管理" summary:"用户列表"`
	Page     int    `json:"page" d:"1" v:"min:1#页码最小为1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" v:"min:1|max:100#每页条数最小为1|每页条数最大为100" dc:"每页条数"`
	Username string `json:"username" dc:"用户名(模糊查询)"`
	Realname string `json:"realname" dc:"真实姓名(模糊查询)"`
	Status   int    `json:"status" dc:"状态(1-正常,2-冻结)"`
}

// UserListRes 用户列表响应
type UserListRes struct {
	List     []*UserInfo `json:"list" dc:"用户列表"`
	Total    int64       `json:"total" dc:"总数"`
	Page     int         `json:"page" dc:"当前页"`
	PageSize int         `json:"pageSize" dc:"每页条数"`
}

// UserAddReq 新增用户请求
type UserAddReq struct {
	g.Meta   `path:"/sys/user" method:"post" tags:"用户管理" summary:"新增用户"`
	Username string `json:"username" v:"required|length:4,20#用户名不能为空|用户名长度为4-20" dc:"用户名"`
	Realname string `json:"realname" v:"required#真实姓名不能为空" dc:"真实姓名"`
	Password string `json:"password" v:"required|length:6,20#密码不能为空|密码长度为6-20" dc:"密码"`
	Phone    string `json:"phone" v:"phone#手机号格式错误" dc:"手机号"`
	Email    string `json:"email" v:"email#邮箱格式错误" dc:"邮箱"`
	School   string `json:"school" dc:"学校"`
	Sex      int    `json:"sex" dc:"性别(0-未知,1-男,2-女)"`
}

// UserEditReq 编辑用户请求
type UserEditReq struct {
	g.Meta   `path:"/sys/user" method:"put" tags:"用户管理" summary:"编辑用户"`
	Id       string `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
	Username string `json:"username" v:"length:4,20#用户名长度为4-20" dc:"用户名"`
	Realname string `json:"realname" dc:"真实姓名"`
	Phone    string `json:"phone" v:"phone#手机号格式错误" dc:"手机号"`
	Email    string `json:"email" v:"email#邮箱格式错误" dc:"邮箱"`
	School   string `json:"school" dc:"学校"`
	Sex      int    `json:"sex" dc:"性别(0-未知,1-男,2-女)"`
	Status   int    `json:"status" dc:"状态(1-正常,2-冻结)"`
}

// UserDeleteReq 删除用户请求
type UserDeleteReq struct {
	g.Meta `path:"/sys/user/:id" method:"delete" tags:"用户管理" summary:"删除用户"`
	Id     string `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}

// UserGetReq 获取用户详情请求
type UserGetReq struct {
	g.Meta `path:"/sys/user/:id" method:"get" tags:"用户管理" summary:"用户详情"`
	Id     string `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}
