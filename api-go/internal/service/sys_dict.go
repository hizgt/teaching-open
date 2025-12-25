package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/model/vo"
	"teaching-open/utility/redis"
)

// SysDictService 字典服务接口
type SysDictService interface {
	List(ctx context.Context, req *DictListReq) (*DictListRes, error)
	GetById(ctx context.Context, id string) (*entity.SysDict, error)
	Create(ctx context.Context, dict *entity.SysDict, createBy string) error
	Update(ctx context.Context, dict *entity.SysDict, updateBy string) error
	Delete(ctx context.Context, id string) error
	GetDictItems(ctx context.Context, dictCode string) ([]*entity.SysDictItem, error)
	QueryAllDictItems(ctx context.Context) (map[string][]vo.DictItem, error)
	QueryTableDictItems(ctx context.Context, table, text, code string) ([]vo.DictItem, error)
}

// DictListReq 字典列表请求
type DictListReq struct {
	Page     int    `json:"page"     d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
	DictName string `json:"dictName"`
	DictCode string `json:"dictCode"`
}

// DictListRes 字典列表响应
type DictListRes struct {
	Records  []entity.SysDict `json:"records"`
	Total    int64            `json:"total"`
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
}

// sysDictServiceImpl 字典服务实现
type sysDictServiceImpl struct {
	redisUtil *redis.Redis
}

// NewSysDictService 创建字典服务实例
func NewSysDictService() SysDictService {
	return &sysDictServiceImpl{
		redisUtil: redis.New(),
	}
}

// List 分页查询字典
func (s *sysDictServiceImpl) List(ctx context.Context, req *DictListReq) (*DictListRes, error) {
	model := dao.SysDict.Ctx(ctx).Where(dao.SysDict.Columns().DelFlag, 0)

	if req.DictName != "" {
		model = model.WhereLike(dao.SysDict.Columns().DictName, "%"+req.DictName+"%")
	}
	if req.DictCode != "" {
		model = model.WhereLike(dao.SysDict.Columns().DictCode, "%"+req.DictCode+"%")
	}

	total, err := model.Count()
	if err != nil {
		return nil, err
	}

	var dicts []entity.SysDict
	err = model.Page(req.Page, req.PageSize).
		OrderDesc(dao.SysDict.Columns().CreateTime).
		Scan(&dicts)
	if err != nil {
		return nil, err
	}

	return &DictListRes{
		Records:  dicts,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// GetById 根据ID获取字典
func (s *sysDictServiceImpl) GetById(ctx context.Context, id string) (*entity.SysDict, error) {
	var dict entity.SysDict
	err := dao.SysDict.Ctx(ctx).Where(dao.SysDict.Columns().Id, id).Scan(&dict)
	if err != nil {
		return nil, err
	}
	if dict.Id == "" {
		return nil, nil
	}
	return &dict, nil
}

// Create 创建字典
func (s *sysDictServiceImpl) Create(ctx context.Context, dict *entity.SysDict, createBy string) error {
	// 检查字典编码是否存在
	count, err := dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().DictCode, dict.DictCode).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("字典编码已存在")
	}

	dict.Id = guid.S()
	dict.DelFlag = 0
	dict.CreateBy = createBy
	dict.CreateTime = gtime.Now()
	dict.UpdateBy = createBy
	dict.UpdateTime = gtime.Now()

	_, err = dao.SysDict.Ctx(ctx).Data(dict).Insert()
	return err
}

// Update 更新字典
func (s *sysDictServiceImpl) Update(ctx context.Context, dict *entity.SysDict, updateBy string) error {
	dict.UpdateBy = updateBy
	dict.UpdateTime = gtime.Now()

	_, err := dao.SysDict.Ctx(ctx).Data(g.Map{
		dao.SysDict.Columns().DictName:    dict.DictName,
		dao.SysDict.Columns().Description: dict.Description,
		dao.SysDict.Columns().UpdateBy:    dict.UpdateBy,
		dao.SysDict.Columns().UpdateTime:  dict.UpdateTime,
	}).Where(dao.SysDict.Columns().Id, dict.Id).Update()

	if err == nil {
		// 清除缓存
		s.redisUtil.InvalidateDictCache(ctx, dict.DictCode)
	}

	return err
}

// Delete 删除字典
func (s *sysDictServiceImpl) Delete(ctx context.Context, id string) error {
	// 获取字典信息
	dict, err := s.GetById(ctx, id)
	if err != nil {
		return err
	}
	if dict == nil {
		return gerror.New("字典不存在")
	}

	// 逻辑删除
	_, err = dao.SysDict.Ctx(ctx).Data(g.Map{
		dao.SysDict.Columns().DelFlag: 1,
	}).Where(dao.SysDict.Columns().Id, id).Update()

	if err == nil {
		// 清除缓存
		s.redisUtil.InvalidateDictCache(ctx, dict.DictCode)
	}

	return err
}

// GetDictItems 获取字典项 (按sortOrder排序)
func (s *sysDictServiceImpl) GetDictItems(ctx context.Context, dictCode string) ([]*entity.SysDictItem, error) {
	// 先从缓存获取
	cached, err := s.redisUtil.GetDictItems(ctx, dictCode)
	if err == nil && cached != "" {
		var items []*entity.SysDictItem
		if json.Unmarshal([]byte(cached), &items) == nil {
			return items, nil
		}
	}

	// 从数据库查询
	var dict entity.SysDict
	err = dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().DictCode, dictCode).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Scan(&dict)
	if err != nil {
		return nil, err
	}
	if dict.Id == "" {
		return nil, nil
	}

	var items []*entity.SysDictItem
	err = dao.SysDictItem.Ctx(ctx).
		Where(dao.SysDictItem.Columns().DictId, dict.Id).
		Where(dao.SysDictItem.Columns().Status, 1).
		OrderAsc(dao.SysDictItem.Columns().SortOrder).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	// 缓存结果
	if len(items) > 0 {
		data, _ := json.Marshal(items)
		s.redisUtil.SetDictItems(ctx, dictCode, string(data), 24*time.Hour)
	}

	return items, nil
}

// QueryAllDictItems 查询所有字典项
func (s *sysDictServiceImpl) QueryAllDictItems(ctx context.Context) (map[string][]vo.DictItem, error) {
	result := make(map[string][]vo.DictItem)

	// 查询所有字典
	var dicts []entity.SysDict
	err := dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Scan(&dicts)
	if err != nil {
		return nil, err
	}

	// 查询每个字典的字典项
	for _, dict := range dicts {
		items, err := s.GetDictItems(ctx, dict.DictCode)
		if err != nil {
			continue
		}

		var voItems []vo.DictItem
		for _, item := range items {
			voItems = append(voItems, vo.DictItem{
				Text:  item.ItemText,
				Value: item.ItemValue,
				Title: item.ItemText,
			})
		}
		result[dict.DictCode] = voItems
	}

	return result, nil
}

// QueryTableDictItems 查询表字典项
func (s *sysDictServiceImpl) QueryTableDictItems(ctx context.Context, table, text, code string) ([]vo.DictItem, error) {
	var items []vo.DictItem

	sql := "SELECT " + text + " as text, " + code + " as value FROM " + table
	result, err := g.DB().GetAll(ctx, sql)
	if err != nil {
		return nil, err
	}

	for _, row := range result {
		items = append(items, vo.DictItem{
			Text:  row["text"].String(),
			Value: row["value"].String(),
		})
	}

	return items, nil
}
