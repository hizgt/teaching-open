package vo

// CourseDeptAssignReq 批量分配课程到部门请求
type CourseDeptAssignReq struct {
	CourseId string   `json:"courseId" v:"required" dc:"课程ID"`
	DeptIds  []string `json:"deptIds" v:"required" dc:"部门ID列表"`
}

// CourseDeptListReq 课程部门关联列表请求
type CourseDeptListReq struct {
	CourseId string `json:"courseId" dc:"课程ID"`
	DeptId   string `json:"deptId" dc:"部门ID"`
}

// CourseDeptListRes 课程部门关联列表响应
type CourseDeptListRes struct {
	Records  []CourseDeptItem `json:"records"`
	Total    int              `json:"total"`
	Page     int              `json:"pageNo"`
	PageSize int              `json:"pageSize"`
}

// CourseDeptItem 课程部门关联项
type CourseDeptItem struct {
	Id       string `json:"id"`
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	DeptId   string `json:"deptId"`
	DeptName string `json:"deptName"`
	CreateBy string `json:"createBy"`
	CreateTime string `json:"createTime"`
}

// CourseDeptRemoveReq 移除课程部门关联请求
type CourseDeptRemoveReq struct {
	Id string `json:"id" v:"required" dc:"关联ID"`
}

// CourseDeptBatchAddReq 批量添加课程部门关联请求
type CourseDeptBatchAddReq struct {
	CourseIds []string `json:"courseIds" v:"required" dc:"课程ID列表"`
	DeptIds   []string `json:"deptIds" v:"required" dc:"部门ID列表"`
}
