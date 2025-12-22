// =================================================================================
// API definitions for sys file module
// =================================================================================

package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// FileInfo 文件信息
type FileInfo struct {
	Id           string `json:"id"`
	FileName     string `json:"fileName"`     // 文件名
	FilePath     string `json:"filePath"`     // 文件路径/URL
	FileType     int    `json:"fileType"`     // 文件类型
	FileLocation int    `json:"fileLocation"` // 存储位置 1:本地 2:阿里云 3:七牛云
	FileTag      string `json:"fileTag"`      // 文件标签
	CreateTime   string `json:"createTime"`   // 创建时间
	CreateBy     string `json:"createBy"`     // 创建人
}

// ========== 文件上传管理 ==========

// FileUploadReq 文件上传请求
type FileUploadReq struct {
	g.Meta   `path:"/sys/upload/file" method:"post" mime:"multipart/form-data" tags:"文件管理" summary:"上传文件"`
	File     *ghttp.UploadFile `json:"file" type:"file" v:"required#请选择要上传的文件" dc:"上传的文件"`
	FileTag  string            `json:"fileTag" dc:"文件标签"`
	Location int               `json:"location" d:"1" dc:"存储位置 1:本地 2:阿里云 3:七牛云"`
}

// FileUploadRes 文件上传响应
type FileUploadRes struct {
	Id       string `json:"id"`       // 文件ID
	FileName string `json:"fileName"` // 文件名
	FilePath string `json:"filePath"` // 文件路径/URL
}

// FileBatchUploadReq 批量文件上传请求
type FileBatchUploadReq struct {
	g.Meta   `path:"/sys/upload/batch" method:"post" mime:"multipart/form-data" tags:"文件管理" summary:"批量上传文件"`
	Files    []*ghttp.UploadFile `json:"files" type:"file" v:"required#请选择要上传的文件" dc:"上传的文件列表"`
	FileTag  string              `json:"fileTag" dc:"文件标签"`
	Location int                 `json:"location" d:"1" dc:"存储位置"`
}

// FileBatchUploadRes 批量文件上传响应
type FileBatchUploadRes struct {
	List []*FileUploadRes `json:"list"` // 上传结果列表
}

// FileListReq 文件列表请求
type FileListReq struct {
	g.Meta   `path:"/sys/file/list" method:"get" tags:"文件管理" summary:"文件列表"`
	Page     int    `json:"page" d:"1" v:"min:1" dc:"页码"`
	PageSize int    `json:"pageSize" d:"10" v:"min:1|max:100" dc:"每页数量"`
	FileName string `json:"fileName" dc:"文件名（模糊查询）"`
	FileType int    `json:"fileType" dc:"文件类型"`
	FileTag  string `json:"fileTag" dc:"文件标签"`
}

// FileListRes 文件列表响应
type FileListRes struct {
	List     []*FileInfo `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// FileGetByIdReq 获取文件详情请求
type FileGetByIdReq struct {
	g.Meta `path:"/sys/file/queryById" method:"get" tags:"文件管理" summary:"文件详情"`
	Id     string `json:"id" v:"required#文件ID不能为空"`
}

// FileGetByIdRes 获取文件详情响应
type FileGetByIdRes struct {
	*FileInfo
}

// FileDeleteReq 删除文件请求
type FileDeleteReq struct {
	g.Meta `path:"/sys/file/delete" method:"delete" tags:"文件管理" summary:"删除文件"`
	Id     string `json:"id" v:"required#文件ID不能为空"`
}

// FileDeleteRes 删除文件响应
type FileDeleteRes struct{}

// FileDeleteBatchReq 批量删除文件请求
type FileDeleteBatchReq struct {
	g.Meta `path:"/sys/file/deleteBatch" method:"delete" tags:"文件管理" summary:"批量删除文件"`
	Ids    string `json:"ids" v:"required#文件ID不能为空"`
}

// FileDeleteBatchRes 批量删除文件响应
type FileDeleteBatchRes struct{}

// FileViewReq 预览/下载文件请求
type FileViewReq struct {
	g.Meta `path:"/sys/file/view/:id" method:"get" tags:"文件管理" summary:"预览文件"`
	Id     string `json:"id" in:"path" v:"required#文件ID不能为空"`
}

// FileViewRes 预览文件响应（返回文件流，无特定结构）
type FileViewRes struct {
	g.Meta `mime:"application/octet-stream"`
}

// FileDownloadReq 下载文件请求
type FileDownloadReq struct {
	g.Meta `path:"/sys/file/download/:id" method:"get" tags:"文件管理" summary:"下载文件"`
	Id     string `json:"id" in:"path" v:"required#文件ID不能为空"`
}

// FileDownloadRes 下载文件响应（返回文件流）
type FileDownloadRes struct {
	g.Meta `mime:"application/octet-stream"`
}
