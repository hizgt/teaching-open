package vo

// CourseListReq 课程列表请求
type CourseListReq struct {
	Page       int    `json:"page" v:"required|min:1" dc:"页码"`
	PageSize   int    `json:"pageSize" v:"required|min:1|max:100" dc:"每页数量"`
	Name       string `json:"name" dc:"课程名称"`
	Type       string `json:"type" dc:"课程类型"`
	Status     string `json:"status" dc:"状态"`
	CreateBy   string `json:"createBy" dc:"创建人"`
}

// CourseListRes 课程列表响应
type CourseListRes struct {
	Records   []CourseItem `json:"records"`
	Total     int          `json:"total"`
	Page      int          `json:"pageNo"`
	PageSize  int          `json:"pageSize"`
}

// CourseItem 课程列表项
type CourseItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CoverImage  string `json:"coverImage"`
	Status      string `json:"status"`
	CreateBy    string `json:"createBy"`
	CreateTime  string `json:"createTime"`
	UnitCount   int    `json:"unitCount"`
}

// CourseCreateReq 创建课程请求
type CourseCreateReq struct {
	Name        string   `json:"name" v:"required" dc:"课程名称"`
	Type        string   `json:"type" v:"required|in:scratch,python" dc:"课程类型"`
	Description string   `json:"description" dc:"课程描述"`
	CoverImage  string   `json:"coverImage" dc:"封面图"`
	DeptIds     []string `json:"deptIds" dc:"关联部门ID列表"`
}

// CourseCreateRes 创建课程响应
type CourseCreateRes struct {
	Id string `json:"id"`
}

// CourseUpdateReq 更新课程请求
type CourseUpdateReq struct {
	Id          string   `json:"id" v:"required" dc:"课程ID"`
	Name        string   `json:"name" v:"required" dc:"课程名称"`
	Type        string   `json:"type" v:"required|in:scratch,python" dc:"课程类型"`
	Description string   `json:"description" dc:"课程描述"`
	CoverImage  string   `json:"coverImage" dc:"封面图"`
	DeptIds     []string `json:"deptIds" dc:"关联部门ID列表"`
}

// CourseDetailRes 课程详情响应
type CourseDetailRes struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	Description string         `json:"description"`
	CoverImage  string         `json:"coverImage"`
	Status      string         `json:"status"`
	CreateBy    string         `json:"createBy"`
	CreateTime  string         `json:"createTime"`
	UpdateBy    string         `json:"updateBy"`
	UpdateTime  string         `json:"updateTime"`
	Depts       []DeptItem     `json:"depts"`
	Units       []UnitItem     `json:"units"`
}

// DeptItem 部门项
type DeptItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// UnitItem 单元项
type UnitItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	SortOrder   int    `json:"sortOrder"`
	ResourceUrl string `json:"resourceUrl"`
}

// CoursePublishReq 发布/下架课程请求
type CoursePublishReq struct {
	Id     string `json:"id" v:"required" dc:"课程ID"`
	Status string `json:"status" v:"required|in:draft,published" dc:"状态"`
}

// CourseDeleteReq 删除课程请求
type CourseDeleteReq struct {
	Id string `json:"id" v:"required" dc:"课程ID"`
}