// =================================================================================
// Service interface for sys file module
// =================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

	v1 "teaching-open/api/v1/sys"
)

// ISysFile 系统文件服务接口
type ISysFile interface {
	// Upload 上传文件
	Upload(ctx context.Context, file *ghttp.UploadFile, fileTag string, location int) (*v1.FileUploadRes, error)
	// UploadBatch 批量上传文件
	UploadBatch(ctx context.Context, files []*ghttp.UploadFile, fileTag string, location int) ([]*v1.FileUploadRes, error)
	// GetList 获取文件列表
	GetList(ctx context.Context, req *v1.FileListReq) (list []*v1.FileInfo, total int64, err error)
	// GetById 根据ID获取文件详情
	GetById(ctx context.Context, id string) (*v1.FileInfo, error)
	// Delete 删除文件
	Delete(ctx context.Context, id string) error
	// DeleteBatch 批量删除文件
	DeleteBatch(ctx context.Context, ids string) error
	// GetFilePath 获取文件物理路径
	GetFilePath(ctx context.Context, id string) (string, error)
}

var localSysFile ISysFile

// SysFile 获取系统文件服务实例
func SysFile() ISysFile {
	if localSysFile == nil {
		panic("implement not found for interface ISysFile, forgot register?")
	}
	return localSysFile
}

// RegisterSysFile 注册系统文件服务实现
func RegisterSysFile(i ISysFile) {
	localSysFile = i
}