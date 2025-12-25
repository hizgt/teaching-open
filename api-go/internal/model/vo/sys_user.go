package vo

// UserListReq 用户列表请求
type UserListReq struct {
	Page      int    `json:"page"      d:"1"  v:"min:1"`
	PageSize  int    `json:"pageSize"  d:"10" v:"min:1|max:100"`
	Username  string `json:"username"`
	Realname  string `json:"realname"`
	Phone     string `json:"phone"`
	Status    int    `json:"status"`
	DepartId  string `json:"departId"`
	RoleId    string `json:"roleId"`
}

// UserListRes 用户列表响应
type UserListRes struct {
	Records  []UserItem `json:"records"`
	Total    int64      `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
}

// UserItem 用户列表项
type UserItem struct {
	Id         string   `json:"id"`
	Username   string   `json:"username"`
	Realname   string   `json:"realname"`
	Avatar     string   `json:"avatar"`
	Sex        int      `json:"sex"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Status     int      `json:"status"`
	OrgCode    string   `json:"orgCode"`
	WorkNo     string   `json:"workNo"`
	Post       string   `json:"post"`
	CreateTime string   `json:"createTime"`
	RoleNames  []string `json:"roleNames"`
	DepartName string   `json:"departName"`
}

// UserCreateReq 创建用户请求
type UserCreateReq struct {
	Username  string   `json:"username"  v:"required|length:3,30"`
	Realname  string   `json:"realname"  v:"required|length:2,30"`
	Password  string   `json:"password"  v:"required|length:6,30"`
	Email     string   `json:"email"     v:"email"`
	Phone     string   `json:"phone"     v:"phone"`
	Sex       int      `json:"sex"       v:"in:1,2"`
	Status    int      `json:"status"    d:"1" v:"in:1,2"`
	WorkNo    string   `json:"workNo"`
	Post      string   `json:"post"`
	School    string   `json:"school"`
	RoleIds   []string `json:"roleIds"`
	DepartIds []string `json:"departIds"`
}

// UserCreateRes 创建用户响应
type UserCreateRes struct {
	Id string `json:"id"`
}

// UserUpdateReq 更新用户请求
type UserUpdateReq struct {
	Id        string   `json:"id"        v:"required"`
	Realname  string   `json:"realname"  v:"length:2,30"`
	Email     string   `json:"email"     v:"email"`
	Phone     string   `json:"phone"     v:"phone"`
	Sex       int      `json:"sex"       v:"in:1,2"`
	Status    int      `json:"status"    v:"in:1,2"`
	WorkNo    string   `json:"workNo"`
	Post      string   `json:"post"`
	School    string   `json:"school"`
	RoleIds   []string `json:"roleIds"`
	DepartIds []string `json:"departIds"`
}

// UserDeleteReq 删除用户请求
type UserDeleteReq struct {
	Id string `json:"id" v:"required"`
}

// UserBatchDeleteReq 批量删除用户请求
type UserBatchDeleteReq struct {
	Ids string `json:"ids" v:"required"`
}

// ResetPasswordReq 重置密码请求
type ResetPasswordReq struct {
	Username   string `json:"username"   v:"required"`
	OldPwd     string `json:"oldPwd"     v:"required"`
	NewPwd     string `json:"newPwd"     v:"required|length:6,30"`
	ConfirmPwd string `json:"confirmPwd" v:"required|same:NewPwd"`
}

// UserDetailRes 用户详情响应
type UserDetailRes struct {
	Id           string   `json:"id"`
	Username     string   `json:"username"`
	Realname     string   `json:"realname"`
	Avatar       string   `json:"avatar"`
	Birthday     string   `json:"birthday"`
	Sex          int      `json:"sex"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	OrgCode      string   `json:"orgCode"`
	Status       int      `json:"status"`
	WorkNo       string   `json:"workNo"`
	Post         string   `json:"post"`
	School       string   `json:"school"`
	Telephone    string   `json:"telephone"`
	UserIdentity int      `json:"userIdentity"`
	CreateTime   string   `json:"createTime"`
	RoleIds      []string `json:"roleIds"`
	DepartIds    []string `json:"departIds"`
}
