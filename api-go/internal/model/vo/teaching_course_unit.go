package vo

// CourseUnitListReq 课程单元列表请求
type CourseUnitListReq struct {
	CourseId string `json:"courseId" v:"required" dc:"课程ID"`
}

// CourseUnitListRes 课程单元列表响应
type CourseUnitListRes struct {
	Records []UnitItem `json:"records"`
}

// CourseUnitCreateReq 创建课程单元请求
type CourseUnitCreateReq struct {
	CourseId    string `json:"courseId" v:"required" dc:"课程ID"`
	Name        string `json:"name" v:"required" dc:"单元名称"`
	Content     string `json:"content" dc:"单元内容"`
	SortOrder   int    `json:"sortOrder" dc:"排序"`
	ResourceUrl string `json:"resourceUrl" dc:"资源链接"`
}

// CourseUnitCreateRes 创建课程单元响应
type CourseUnitCreateRes struct {
	Id string `json:"id"`
}

// CourseUnitUpdateReq 更新课程单元请求
type CourseUnitUpdateReq struct {
	Id          string `json:"id" v:"required" dc:"单元ID"`
	Name        string `json:"name" v:"required" dc:"单元名称"`
	Content     string `json:"content" dc:"单元内容"`
	SortOrder   int    `json:"sortOrder" dc:"排序"`
	ResourceUrl string `json:"resourceUrl" dc:"资源链接"`
}

// CourseUnitSortReq 课程单元排序请求
type CourseUnitSortReq struct {
	CourseId string       `json:"courseId" v:"required" dc:"课程ID"`
	Units    []UnitSortItem `json:"units" v:"required" dc:"单元排序列表"`
}

// UnitSortItem 单元排序项
type UnitSortItem struct {
	Id        string `json:"id" v:"required" dc:"单元ID"`
	SortOrder int    `json:"sortOrder" v:"required|min:0" dc:"排序"`
}

// CourseUnitDeleteReq 删除课程单元请求
type CourseUnitDeleteReq struct {
	Id string `json:"id" v:"required" dc:"单元ID"`
}