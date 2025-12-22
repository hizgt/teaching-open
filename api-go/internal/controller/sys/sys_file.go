// =================================================================================
// Controller for sys file module
// =================================================================================
package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

type cFile struct{}

var SysFile = &cFile{}

// Upload 上传文件
func (c *cFile) Upload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	result, err := service.SysFile().Upload(ctx, req.File, req.FileTag, req.Location)
	if err != nil {
		return nil, gerror.Wrap(err, "上传文件失败")
	}
	return result, nil
}

// UploadBatch 批量上传文件
func (c *cFile) UploadBatch(ctx context.Context, req *v1.FileBatchUploadReq) (res *v1.FileBatchUploadRes, err error) {
	list, err := service.SysFile().UploadBatch(ctx, req.Files, req.FileTag, req.Location)
	if err != nil {
		return nil, gerror.Wrap(err, "批量上传文件失败")
	}
	return &v1.FileBatchUploadRes{
		List: list,
	}, nil
}

// GetList 获取文件列表
func (c *cFile) GetList(ctx context.Context, req *v1.FileListReq) (res *v1.FileListRes, err error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	list, total, err := service.SysFile().GetList(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询文件列表失败")
	}
	return &v1.FileListRes{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// GetById 获取文件详情
func (c *cFile) GetById(ctx context.Context, req *v1.FileGetByIdReq) (res *v1.FileGetByIdRes, err error) {
	info, err := service.SysFile().GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询文件详情失败")
	}
	return &v1.FileGetByIdRes{
		FileInfo: info,
	}, nil
}

// Delete 删除文件
func (c *cFile) Delete(ctx context.Context, req *v1.FileDeleteReq) (res *v1.FileDeleteRes, err error) {
	err = service.SysFile().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除文件失败")
	}
	return &v1.FileDeleteRes{}, nil
}

// DeleteBatch 批量删除文件
func (c *cFile) DeleteBatch(ctx context.Context, req *v1.FileDeleteBatchReq) (res *v1.FileDeleteBatchRes, err error) {
	err = service.SysFile().DeleteBatch(ctx, req.Ids)
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除文件失败")
	}
	return &v1.FileDeleteBatchRes{}, nil
}

// View 预览文件
func (c *cFile) View(ctx context.Context, req *v1.FileViewReq) (res *v1.FileViewRes, err error) {
	// 获取文件路径
	filePath, err := service.SysFile().GetFilePath(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "获取文件失败")
	}

	// 检查文件是否存在
	if !gfile.Exists(filePath) {
		return nil, gerror.New("文件不存在")
	}

	// 获取请求对象并返回文件
	r := g.RequestFromCtx(ctx)
	r.Response.ServeFile(filePath)
	return nil, nil
}

// Download 下载文件
func (c *cFile) Download(ctx context.Context, req *v1.FileDownloadReq) (res *v1.FileDownloadRes, err error) {
	// 获取文件信息
	info, err := service.SysFile().GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "获取文件失败")
	}

	// 获取文件路径
	filePath, err := service.SysFile().GetFilePath(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "获取文件失败")
	}

	// 检查文件是否存在
	if !gfile.Exists(filePath) {
		return nil, gerror.New("文件不存在")
	}

	// 获取请求对象并返回文件下载
	r := g.RequestFromCtx(ctx)
	r.Response.ServeFileDownload(filePath, info.FileName)
	return nil, nil
}

// HandleUpload 处理通用上传（兼容旧接口）
func (c *cFile) HandleUpload(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	if file == nil {
		r.Response.WriteJson(g.Map{
			"code":    400,
			"message": "请选择要上传的文件",
			"success": false,
		})
		return
	}

	fileTag := r.Get("fileTag", "").String()
	location := r.Get("location", 1).Int()

	result, err := service.SysFile().Upload(r.Context(), file, fileTag, location)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code":    500,
			"message": err.Error(),
			"success": false,
		})
		return
	}

	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "上传成功",
		"success": true,
		"result":  result,
	})
}
