package vo

// DictListReq 字典列表请求
type DictListReq struct {
	Page     int    `json:"page"     d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	DictName string `json:"dictName"`
	DictCode string `json:"dictCode"`
}

// DictCreateReq 创建字典请求
type DictCreateReq struct {
	DictName    string `json:"dictName" v:"required#字典名称不能为空"`
	DictCode    string `json:"dictCode" v:"required#字典编码不能为空"`
	Description string `json:"description"`
	Type        int    `json:"type" d:"0"`
}

// DictUpdateReq 更新字典请求
type DictUpdateReq struct {
	Id          string `json:"id" v:"required#字典ID不能为空"`
	DictName    string `json:"dictName" v:"required#字典名称不能为空"`
	DictCode    string `json:"dictCode"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}

// DictItemListReq 字典项列表请求
type DictItemListReq struct {
	Page     int    `json:"page"     d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	DictId   string `json:"dictId" v:"required#字典ID不能为空"`
}

// DictItemCreateReq 创建字典项请求
type DictItemCreateReq struct {
	DictId      string `json:"dictId" v:"required#字典ID不能为空"`
	ItemText    string `json:"itemText" v:"required#字典项文本不能为空"`
	ItemValue   string `json:"itemValue" v:"required#字典项值不能为空"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder" d:"0"`
	Status      int    `json:"status" d:"1"`
}

// DictItemUpdateReq 更新字典项请求
type DictItemUpdateReq struct {
	Id          string `json:"id" v:"required#字典项ID不能为空"`
	ItemText    string `json:"itemText" v:"required#字典项文本不能为空"`
	ItemValue   string `json:"itemValue" v:"required#字典项值不能为空"`
	Description string `json:"description"`
	SortOrder   int    `json:"sortOrder"`
	Status      int    `json:"status"`
}
