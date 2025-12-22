package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
	"teaching-open/utility/jwt"
)

func init() {
	service.RegisterTeachingScratchAsset(NewTeachingScratchAssetLogic())
}

type sTeachingScratchAsset struct{}

func NewTeachingScratchAssetLogic() *sTeachingScratchAsset {
	return &sTeachingScratchAsset{}
}

// List 获取素材列表（后台管理）
func (s *sTeachingScratchAsset) List(ctx context.Context, req *v1.ScratchAssetListReq) (list interface{}, total int, err error) {
	m := dao.TeachingScratchAsset.Ctx(ctx).Where("del_flag", 0)

	// 素材类型查询 (-1 表示全部)
	if req.AssetType > 0 {
		m = m.Where("asset_type", req.AssetType)
	}
	// 名称查询
	if req.AssetName != "" {
		m = m.WhereLike("asset_name", "%"+req.AssetName+"%")
	}
	// 标签查询
	if req.Tags != "" {
		m = m.WhereLike("tags", "%"+req.Tags+"%")
	}

	// 统计总数
	count, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	total = count

	// 分页查询
	pageNo := req.PageNo
	pageSize := req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var entities []*entity.TeachingScratchAsset
	err = m.Order("create_time DESC").Page(pageNo, pageSize).Scan(&entities)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	var assetInfos []*v1.ScratchAssetInfo
	for _, e := range entities {
		assetInfos = append(assetInfos, &v1.ScratchAssetInfo{
			Id:         e.Id,
			AssetType:  e.AssetType,
			AssetName:  e.AssetName,
			AssetData:  e.AssetData,
			Md5Ext:     e.Md5Ext,
			Tags:       e.Tags,
			CreateBy:   e.CreateBy,
			CreateTime: e.CreateTime,
			UpdateBy:   e.UpdateBy,
			UpdateTime: e.UdpateTime,
			DelFlag:    e.DelFlag,
		})
	}

	return assetInfos, total, nil
}

// GetScratchAssets 获取Scratch素材（前端编辑器使用）
func (s *sTeachingScratchAsset) GetScratchAssets(ctx context.Context, assetType int) (list interface{}, err error) {
	m := dao.TeachingScratchAsset.Ctx(ctx).
		Where("del_flag", 0).
		Where("asset_type", assetType)

	var entities []*entity.TeachingScratchAsset
	err = m.Order("create_time DESC").Scan(&entities)
	if err != nil {
		return nil, err
	}

	// 转换为Scratch需要的数据格式
	// 返回素材的JSON数据列表，解析asset_data字段
	var assetList []g.Map
	for _, e := range entities {
		// 解析JSON数据
		item := g.Map{
			"id":        e.Id,
			"name":      e.AssetName,
			"tags":      e.Tags,
			"md5ext":    e.Md5Ext,
			"assetData": e.AssetData,
		}
		assetList = append(assetList, item)
	}

	return assetList, nil
}

// Add 添加素材
func (s *sTeachingScratchAsset) Add(ctx context.Context, req *v1.ScratchAssetAddReq) (id string, err error) {
	username := jwt.GetUsername(ctx)
	newId := g.NewVar(nil).String()
	if newId == "" {
		newId = gconv.String(gtime.TimestampNano())
	}

	data := &entity.TeachingScratchAsset{
		Id:         newId,
		AssetType:  req.AssetType,
		AssetName:  req.AssetName,
		AssetData:  req.AssetData,
		Md5Ext:     req.Md5Ext,
		Tags:       req.Tags,
		CreateBy:   username,
		CreateTime: gtime.Now(),
		UpdateBy:   username,
		UdpateTime: gtime.Now(),
		DelFlag:    0,
	}

	_, err = dao.TeachingScratchAsset.Ctx(ctx).Insert(data)
	if err != nil {
		return "", err
	}

	return newId, nil
}

// Edit 编辑素材
func (s *sTeachingScratchAsset) Edit(ctx context.Context, req *v1.ScratchAssetEditReq) error {
	username := jwt.GetUsername(ctx)

	data := g.Map{
		"update_by":   username,
		"udpate_time": gtime.Now(),
	}

	// -1 表示不更新类型
	if req.AssetType > 0 {
		data["asset_type"] = req.AssetType
	}
	if req.AssetName != "" {
		data["asset_name"] = req.AssetName
	}
	if req.AssetData != "" {
		data["asset_data"] = req.AssetData
	}
	if req.Md5Ext != "" {
		data["md5_ext"] = req.Md5Ext
	}
	if req.Tags != "" {
		data["tags"] = req.Tags
	}

	_, err := dao.TeachingScratchAsset.Ctx(ctx).Where("id", req.Id).Data(data).Update()
	return err
}

// Delete 删除素材（软删除）
func (s *sTeachingScratchAsset) Delete(ctx context.Context, id string) error {
	username := jwt.GetUsername(ctx)

	_, err := dao.TeachingScratchAsset.Ctx(ctx).Where("id", id).Data(g.Map{
		"del_flag":    1,
		"update_by":   username,
		"udpate_time": gtime.Now(),
	}).Update()
	return err
}

// DeleteBatch 批量删除素材（软删除）
func (s *sTeachingScratchAsset) DeleteBatch(ctx context.Context, ids string) error {
	username := jwt.GetUsername(ctx)
	idList := strings.Split(ids, ",")

	_, err := dao.TeachingScratchAsset.Ctx(ctx).WhereIn("id", idList).Data(g.Map{
		"del_flag":    1,
		"update_by":   username,
		"udpate_time": gtime.Now(),
	}).Update()
	return err
}

// GetById 获取素材详情
func (s *sTeachingScratchAsset) GetById(ctx context.Context, id string) (*v1.ScratchAssetInfo, error) {
	var e *entity.TeachingScratchAsset
	err := dao.TeachingScratchAsset.Ctx(ctx).Where("id", id).Where("del_flag", 0).Scan(&e)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, nil
	}

	return &v1.ScratchAssetInfo{
		Id:         e.Id,
		AssetType:  e.AssetType,
		AssetName:  e.AssetName,
		AssetData:  e.AssetData,
		Md5Ext:     e.Md5Ext,
		Tags:       e.Tags,
		CreateBy:   e.CreateBy,
		CreateTime: e.CreateTime,
		UpdateBy:   e.UpdateBy,
		UpdateTime: e.UdpateTime,
		DelFlag:    e.DelFlag,
	}, nil
}
