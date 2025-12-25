package vo

// PermissionListReq 权限列表请求
type PermissionListReq struct {
	Page     int    `json:"page"     d:"1"  v:"min:1"`
	PageSize int    `json:"pageSize" d:"10" v:"min:1|max:100"`
	Name     string `json:"name"`
	MenuType int    `json:"menuType" d:"-1"`
}

// PermissionTreeItem 权限树项
type PermissionTreeItem struct {
	Id                 string                `json:"id"`
	ParentId           string                `json:"parentId"`
	Key                string                `json:"key"`
	Title              string                `json:"title"`
	Name               string                `json:"name"`
	Url                string                `json:"url"`
	Component          string                `json:"component"`
	ComponentName      string                `json:"componentName"`
	Redirect           string                `json:"redirect"`
	MenuType           int                   `json:"menuType"`
	Perms              string                `json:"perms"`
	PermsType          string                `json:"permsType"`
	SortNo             float64               `json:"sortNo"`
	AlwaysShow         bool                  `json:"alwaysShow"`
	Icon               string                `json:"icon"`
	IsRoute            bool                  `json:"isRoute"`
	IsLeaf             bool                  `json:"isLeaf"`
	KeepAlive          bool                  `json:"keepAlive"`
	Hidden             int                   `json:"hidden"`
	Description        string                `json:"description"`
	Status             string                `json:"status"`
	InternalOrExternal bool                  `json:"internalOrExternal"`
	Children           []*PermissionTreeItem `json:"children,omitempty"`
}

// PermissionCreateReq 创建权限请求
type PermissionCreateReq struct {
	ParentId           string  `json:"parentId"`
	Name               string  `json:"name"               v:"required#菜单名称不能为空"`
	Url                string  `json:"url"`
	Component          string  `json:"component"`
	ComponentName      string  `json:"componentName"`
	Redirect           string  `json:"redirect"`
	MenuType           int     `json:"menuType"           d:"0"`
	Perms              string  `json:"perms"`
	PermsType          string  `json:"permsType"`
	SortNo             float64 `json:"sortNo"             d:"1.0"`
	AlwaysShow         bool    `json:"alwaysShow"`
	Icon               string  `json:"icon"`
	IsRoute            bool    `json:"isRoute"            d:"true"`
	IsLeaf             bool    `json:"isLeaf"`
	KeepAlive          bool    `json:"keepAlive"`
	Hidden             int     `json:"hidden"             d:"0"`
	Description        string  `json:"description"`
	Status             string  `json:"status"             d:"1"`
	InternalOrExternal bool    `json:"internalOrExternal"`
}

// PermissionCreateRes 创建权限响应
type PermissionCreateRes struct {
	Id string `json:"id"`
}

// PermissionUpdateReq 更新权限请求
type PermissionUpdateReq struct {
	Id                 string  `json:"id"                 v:"required#权限ID不能为空"`
	ParentId           string  `json:"parentId"`
	Name               string  `json:"name"               v:"required#菜单名称不能为空"`
	Url                string  `json:"url"`
	Component          string  `json:"component"`
	ComponentName      string  `json:"componentName"`
	Redirect           string  `json:"redirect"`
	MenuType           int     `json:"menuType"`
	Perms              string  `json:"perms"`
	PermsType          string  `json:"permsType"`
	SortNo             float64 `json:"sortNo"`
	AlwaysShow         bool    `json:"alwaysShow"`
	Icon               string  `json:"icon"`
	IsRoute            bool    `json:"isRoute"`
	IsLeaf             bool    `json:"isLeaf"`
	KeepAlive          bool    `json:"keepAlive"`
	Hidden             int     `json:"hidden"`
	Description        string  `json:"description"`
	Status             string  `json:"status"`
	InternalOrExternal bool    `json:"internalOrExternal"`
}

// UserPermissionRes 用户权限响应
type UserPermissionRes struct {
	AllAuth []PermissionTreeItem `json:"allAuth"`
	Auth    []PermissionTreeItem `json:"auth"`
	Menu    []PermissionTreeItem `json:"menu"`
}
