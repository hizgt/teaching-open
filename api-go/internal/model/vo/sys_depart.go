package vo

// DepartTreeReq 部门树查询请求
type DepartTreeReq struct {
	DepartName string `json:"departName" v:""`
	OrgCode    string `json:"orgCode" v:""`
}

// DepartTreeNode 部门树节点
type DepartTreeNode struct {
	Id             string            `json:"id"`
	ParentId       string            `json:"parentId"`
	DepartName     string            `json:"departName"`
	DepartNameEn   string            `json:"departNameEn"`
	DepartNameAbbr string            `json:"departNameAbbr"`
	DepartOrder    int               `json:"departOrder"`
	Description    string            `json:"description"`
	OrgCategory    string            `json:"orgCategory"`
	OrgType        string            `json:"orgType"`
	OrgCode        string            `json:"orgCode"`
	Mobile         string            `json:"mobile"`
	Address        string            `json:"address"`
	Status         string            `json:"status"`
	Key            string            `json:"key"`
	Value          string            `json:"value"`
	Title          string            `json:"title"`
	IsLeaf         bool              `json:"isLeaf"`
	Children       []*DepartTreeNode `json:"children"`
}

// DepartIdTreeNode 部门ID树节点
type DepartIdTreeNode struct {
	Key      string              `json:"key"`
	Value    string              `json:"value"`
	Title    string              `json:"title"`
	Children []*DepartIdTreeNode `json:"children"`
}

// DepartCreateReq 创建部门请求
type DepartCreateReq struct {
	ParentId       string `json:"parentId" v:""`
	DepartName     string `json:"departName" v:"required#部门名称不能为空"`
	DepartNameEn   string `json:"departNameEn" v:""`
	DepartNameAbbr string `json:"departNameAbbr" v:""`
	DepartOrder    int    `json:"departOrder" v:""`
	Description    string `json:"description" v:""`
	OrgCategory    string `json:"orgCategory" v:""`
	OrgType        string `json:"orgType" v:""`
	Mobile         string `json:"mobile" v:""`
	Fax            string `json:"fax" v:""`
	Address        string `json:"address" v:""`
	Memo           string `json:"memo" v:""`
}

// DepartCreateRes 创建部门响应
type DepartCreateRes struct {
	Id string `json:"id"`
}

// DepartUpdateReq 更新部门请求
type DepartUpdateReq struct {
	Id             string `json:"id" v:"required#部门ID不能为空"`
	ParentId       string `json:"parentId" v:""`
	DepartName     string `json:"departName" v:"required#部门名称不能为空"`
	DepartNameEn   string `json:"departNameEn" v:""`
	DepartNameAbbr string `json:"departNameAbbr" v:""`
	DepartOrder    int    `json:"departOrder" v:""`
	Description    string `json:"description" v:""`
	OrgCategory    string `json:"orgCategory" v:""`
	OrgType        string `json:"orgType" v:""`
	Mobile         string `json:"mobile" v:""`
	Fax            string `json:"fax" v:""`
	Address        string `json:"address" v:""`
	Memo           string `json:"memo" v:""`
}

// DepartSearchReq 部门搜索请求
type DepartSearchReq struct {
	DepartName string `json:"departName" v:""`
	OrgCode    string `json:"orgCode" v:""`
	OrgType    string `json:"orgType" v:""`
}

// UserDepartReq 用户部门关联请求
type UserDepartReq struct {
	UserId  string   `json:"userId" v:"required#用户ID不能为空"`
	DepIds  []string `json:"depIds" v:""`
}

// DepartUserTreeNode 部门用户树节点
type DepartUserTreeNode struct {
	Key      string                `json:"key"`
	Value    string                `json:"value"`
	Title    string                `json:"title"`
	Type     string                `json:"type"` // depart 或 user
	Children []*DepartUserTreeNode `json:"children"`
}
