// =================================================================================
// Logic implementation for sys file module
// =================================================================================
package sys

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/do"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

type sSysFile struct{}

func init() {
	service.RegisterSysFile(&sSysFile{})
}

// getUploadPath 获取上传目录
func (s *sSysFile) getUploadPath() string {
	uploadPath := g.Cfg().MustGet(context.Background(), "upload.path", "resource/upload").String()
	if !gfile.Exists(uploadPath) {
		gfile.Mkdir(uploadPath)
	}
	return uploadPath
}

// getFileType 根据文件扩展名判断文件类型
func (s *sSysFile) getFileType(filename string) int {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".svg":
		return 1 // 图片
	case ".doc", ".docx", ".pdf", ".xls", ".xlsx", ".ppt", ".pptx", ".txt":
		return 2 // 文档
	case ".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv":
		return 3 // 视频
	case ".mp3", ".wav", ".flac", ".aac":
		return 4 // 音频
	case ".zip", ".rar", ".7z", ".tar", ".gz":
		return 5 // 压缩包
	default:
		return 0 // 其他
	}
}

// Upload 上传文件
func (s *sSysFile) Upload(ctx context.Context, file *ghttp.UploadFile, fileTag string, location int) (*v1.FileUploadRes, error) {
	if file == nil {
		return nil, errors.New("请选择要上传的文件")
	}

	// 检查文件大小（最大50MB）
	maxSize := g.Cfg().MustGet(ctx, "upload.maxSize", 50*1024*1024).Int64()
	if file.Size > maxSize {
		return nil, errors.New("文件大小超过限制")
	}

	// 获取上传路径
	uploadPath := s.getUploadPath()

	// 按日期创建子目录
	datePath := time.Now().Format("2006/01/02")
	fullPath := filepath.Join(uploadPath, datePath)
	if !gfile.Exists(fullPath) {
		gfile.Mkdir(fullPath)
	}

	// 生成唯一文件名
	fileId := guid.S()
	ext := filepath.Ext(file.Filename)
	newFileName := fileId + ext
	originalFileName := file.Filename

	// 保存文件
	file.Filename = newFileName
	_, err := file.Save(fullPath)
	if err != nil {
		g.Log().Error(ctx, "保存文件失败:", err)
		return nil, errors.New("保存文件失败")
	}
	filePath := filepath.Join(fullPath, newFileName)

	// 获取相对路径用于存储
	relPath := filepath.Join(datePath, newFileName)
	relPath = strings.ReplaceAll(relPath, "\\", "/")

	// 获取文件类型
	fileType := s.getFileType(originalFileName)

	// 保存文件记录到数据库
	_, err = dao.SysFile.Ctx(ctx).Insert(do.SysFile{
		Id:           fileId,
		FileName:     originalFileName,
		FilePath:     relPath,
		FileType:     fileType,
		FileLocation: location,
		FileTag:      fileTag,
		DelFlag:      0,
		CreateTime:   gtime.Now(),
		UpdateTime:   gtime.Now(),
	})
	if err != nil {
		// 删除已上传的文件
		os.Remove(filePath)
		g.Log().Error(ctx, "保存文件记录失败:", err)
		return nil, errors.New("保存文件记录失败")
	}

	// 返回文件访问路径
	baseUrl := g.Cfg().MustGet(ctx, "upload.baseUrl", "").String()
	var accessPath string
	if baseUrl != "" {
		accessPath = baseUrl + "/" + relPath
	} else {
		accessPath = "/api/v1/sys/file/view/" + fileId
	}

	g.Log().Infof(ctx, "文件上传成功: id=%s, fileName=%s", fileId, originalFileName)

	return &v1.FileUploadRes{
		Id:       fileId,
		FileName: originalFileName,
		FilePath: accessPath,
	}, nil
}

// UploadBatch 批量上传文件
func (s *sSysFile) UploadBatch(ctx context.Context, files []*ghttp.UploadFile, fileTag string, location int) ([]*v1.FileUploadRes, error) {
	if len(files) == 0 {
		return nil, errors.New("请选择要上传的文件")
	}

	results := make([]*v1.FileUploadRes, 0, len(files))
	for _, file := range files {
		result, err := s.Upload(ctx, file, fileTag, location)
		if err != nil {
			g.Log().Warning(ctx, "批量上传中文件失败:", file.Filename, err)
			continue
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, errors.New("所有文件上传失败")
	}

	return results, nil
}

// GetList 获取文件列表
func (s *sSysFile) GetList(ctx context.Context, req *v1.FileListReq) (list []*v1.FileInfo, total int64, err error) {
	// 构建查询
	m := dao.SysFile.Ctx(ctx).Where(dao.SysFile.Columns().DelFlag, 0)

	// 文件名模糊查询
	if req.FileName != "" {
		m = m.WhereLike(dao.SysFile.Columns().FileName, "%"+req.FileName+"%")
	}

	// 文件类型筛选
	if req.FileType > 0 {
		m = m.Where(dao.SysFile.Columns().FileType, req.FileType)
	}

	// 文件标签筛选
	if req.FileTag != "" {
		m = m.Where(dao.SysFile.Columns().FileTag, req.FileTag)
	}

	// 分页查询
	var files []*entity.SysFile
	var totalInt int
	err = m.Page(req.Page, req.PageSize).
		OrderDesc(dao.SysFile.Columns().CreateTime).
		ScanAndCount(&files, &totalInt, false)
	if err != nil {
		g.Log().Error(ctx, "查询文件列表失败:", err)
		return nil, 0, errors.New("查询文件列表失败")
	}
	total = int64(totalInt)

	// 获取基础URL
	baseUrl := g.Cfg().MustGet(ctx, "upload.baseUrl", "").String()

	// 转换为FileInfo
	list = make([]*v1.FileInfo, 0, len(files))
	for _, file := range files {
		createTime := ""
		if file.CreateTime != nil {
			createTime = file.CreateTime.String()
		}

		// 构建文件访问路径
		var accessPath string
		if baseUrl != "" {
			accessPath = baseUrl + "/" + file.FilePath
		} else {
			accessPath = "/api/v1/sys/file/view/" + file.Id
		}

		list = append(list, &v1.FileInfo{
			Id:           file.Id,
			FileName:     file.FileName,
			FilePath:     accessPath,
			FileType:     file.FileType,
			FileLocation: file.FileLocation,
			FileTag:      file.FileTag,
			CreateTime:   createTime,
			CreateBy:     file.CreateBy,
		})
	}

	return list, total, nil
}

// GetById 根据ID获取文件详情
func (s *sSysFile) GetById(ctx context.Context, id string) (*v1.FileInfo, error) {
	var file entity.SysFile
	err := dao.SysFile.Ctx(ctx).
		Where(dao.SysFile.Columns().Id, id).
		Where(dao.SysFile.Columns().DelFlag, 0).
		Scan(&file)
	if err != nil {
		g.Log().Error(ctx, "查询文件失败:", err)
		return nil, errors.New("查询文件失败")
	}

	if file.Id == "" {
		return nil, errors.New("文件不存在")
	}

	createTime := ""
	if file.CreateTime != nil {
		createTime = file.CreateTime.String()
	}

	// 获取基础URL
	baseUrl := g.Cfg().MustGet(ctx, "upload.baseUrl", "").String()
	var accessPath string
	if baseUrl != "" {
		accessPath = baseUrl + "/" + file.FilePath
	} else {
		accessPath = "/api/v1/sys/file/view/" + file.Id
	}

	return &v1.FileInfo{
		Id:           file.Id,
		FileName:     file.FileName,
		FilePath:     accessPath,
		FileType:     file.FileType,
		FileLocation: file.FileLocation,
		FileTag:      file.FileTag,
		CreateTime:   createTime,
		CreateBy:     file.CreateBy,
	}, nil
}

// Delete 删除文件
func (s *sSysFile) Delete(ctx context.Context, id string) error {
	// 查询文件信息
	var file entity.SysFile
	err := dao.SysFile.Ctx(ctx).
		Where(dao.SysFile.Columns().Id, id).
		Where(dao.SysFile.Columns().DelFlag, 0).
		Scan(&file)
	if err != nil {
		g.Log().Error(ctx, "查询文件失败:", err)
		return errors.New("查询文件失败")
	}

	if file.Id == "" {
		return errors.New("文件不存在")
	}

	// 逻辑删除文件记录
	_, err = dao.SysFile.Ctx(ctx).
		Where(dao.SysFile.Columns().Id, id).
		Update(do.SysFile{
			DelFlag:    1,
			UpdateTime: gtime.Now(),
		})
	if err != nil {
		g.Log().Error(ctx, "删除文件记录失败:", err)
		return errors.New("删除文件记录失败")
	}

	// 删除物理文件（本地存储）
	if file.FileLocation == 1 {
		uploadPath := s.getUploadPath()
		fullPath := filepath.Join(uploadPath, file.FilePath)
		if gfile.Exists(fullPath) {
			err = os.Remove(fullPath)
			if err != nil {
				g.Log().Warning(ctx, "删除物理文件失败:", err)
				// 不返回错误，文件记录已删除
			}
		}
	}

	g.Log().Infof(ctx, "删除文件成功: id=%s", id)
	return nil
}

// DeleteBatch 批量删除文件
func (s *sSysFile) DeleteBatch(ctx context.Context, ids string) error {
	if ids == "" {
		return errors.New("文件ID不能为空")
	}

	idList := strings.Split(ids, ",")
	if len(idList) == 0 {
		return errors.New("文件ID不能为空")
	}

	// 查询要删除的文件信息
	var files []*entity.SysFile
	err := dao.SysFile.Ctx(ctx).
		WhereIn(dao.SysFile.Columns().Id, idList).
		Where(dao.SysFile.Columns().DelFlag, 0).
		Scan(&files)
	if err != nil {
		g.Log().Error(ctx, "查询文件列表失败:", err)
		return errors.New("查询文件列表失败")
	}

	// 逻辑删除文件记录
	_, err = dao.SysFile.Ctx(ctx).
		WhereIn(dao.SysFile.Columns().Id, idList).
		Update(do.SysFile{
			DelFlag:    1,
			UpdateTime: gtime.Now(),
		})
	if err != nil {
		g.Log().Error(ctx, "批量删除文件记录失败:", err)
		return errors.New("批量删除文件记录失败")
	}

	// 删除物理文件
	uploadPath := s.getUploadPath()
	for _, file := range files {
		if file.FileLocation == 1 {
			fullPath := filepath.Join(uploadPath, file.FilePath)
			if gfile.Exists(fullPath) {
				os.Remove(fullPath)
			}
		}
	}

	g.Log().Infof(ctx, "批量删除文件成功: ids=%s", ids)
	return nil
}

// GetFilePath 获取文件物理路径
func (s *sSysFile) GetFilePath(ctx context.Context, id string) (string, error) {
	var file entity.SysFile
	err := dao.SysFile.Ctx(ctx).
		Where(dao.SysFile.Columns().Id, id).
		Where(dao.SysFile.Columns().DelFlag, 0).
		Scan(&file)
	if err != nil {
		g.Log().Error(ctx, "查询文件失败:", err)
		return "", errors.New("查询文件失败")
	}

	if file.Id == "" {
		return "", errors.New("文件不存在")
	}

	// 返回完整路径
	uploadPath := s.getUploadPath()
	return filepath.Join(uploadPath, file.FilePath), nil
}
