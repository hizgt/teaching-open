package vo

// RoleListReq 角色列表请求
type RoleListReq struct {
	Page     int    `json:"page"     d:"1"  v:"min:1"`
	PageSize int    `json:"pageSize" d:"10" v:"min:1|max:100"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
}

// RoleListRes 角色列表响应
type RoleListRes struct {
	Records  []RoleItem `json:"records"`
	Total    int        `json:"total"`
	Page     int        `json:"page"`
	PageSize int        `json:"pageSize"`
}

// RoleItem 角色项
type RoleItem struct {
	Id          string `json:"id"`
	RoleName    string `json:"roleName"`
	RoleCode    string `json:"roleCode"`
	RoleLevel   int    `json:"roleLevel"`
	Description string `json:"description"`
	CreateTime  string `json:"createTime"`
}

// RoleCreateReq 创建角色请求
type RoleCreateReq struct {
	RoleName    string `json:"roleName"    v:"required#角色名称不能为空"`
	RoleCode    string `json:"roleCode"    v:"required#角色编码不能为空"`
	RoleLevel   int    `json:"roleLevel"   d:"0"`
	Description string `json:"description"`
}

// RoleCreateRes 创建角色响应
type RoleCreateRes struct {
	Id string `json:"id"`
}

// RoleUpdateReq 更新角色请求
type RoleUpdateReq struct {
	Id          string `json:"id"          v:"required#角色ID不能为空"`
	RoleName    string `json:"roleName"    v:"required#角色名称不能为空"`
	RoleCode    string `json:"roleCode"    v:"required#角色编码不能为空"`
	RoleLevel   int    `json:"roleLevel"`
	Description string `json:"description"`
}

// UserRolesReq 用户角色请求
type UserRolesReq struct {
	UserId  string   `json:"userId"  v:"required#用户ID不能为空"`
	RoleIds []string `json:"roleIds"`
}

// RolePermissionReq 角色权限请求
type RolePermissionReq struct {
	RoleId        string   `json:"roleId"        v:"required#角色ID不能为空"`
	PermissionIds []string `json:"permissionIds"`
	LastPermIds   string   `json:"lastPermIds"`
}
