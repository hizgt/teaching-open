// =================================================================================
// Logic implementation for sys dict module
// =================================================================================

package sys

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/do"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

type sSysDict struct{}

func init() {
	service.RegisterSysDict(&sSysDict{})
}

// GetList 获取字典列表
func (s *sSysDict) GetList(ctx context.Context, req *v1.DictListReq) (list []*v1.DictInfo, total int64, err error) {
	// 构建查询
	m := dao.SysDict.Ctx(ctx).Where(dao.SysDict.Columns().DelFlag, 0)

	// 字典名称模糊查询
	if req.DictName != "" {
		m = m.WhereLike(dao.SysDict.Columns().DictName, "%"+req.DictName+"%")
	}

	// 字典编码模糊查询
	if req.DictCode != "" {
		m = m.WhereLike(dao.SysDict.Columns().DictCode, "%"+req.DictCode+"%")
	}

	// 分页查询
	var dicts []*entity.SysDict
	var totalInt int
	err = m.Page(req.Page, req.PageSize).
		OrderDesc(dao.SysDict.Columns().CreateTime).
		ScanAndCount(&dicts, &totalInt, false)
	if err != nil {
		g.Log().Error(ctx, "查询字典列表失败:", err)
		return nil, 0, errors.New("查询字典列表失败")
	}
	total = int64(totalInt)

	// 转换为DictInfo
	list = make([]*v1.DictInfo, 0, len(dicts))
	for _, dict := range dicts {
		createTime := ""
		if dict.CreateTime != nil {
			createTime = dict.CreateTime.String()
		}
		list = append(list, &v1.DictInfo{
			Id:          dict.Id,
			DictName:    dict.DictName,
			DictCode:    dict.DictCode,
			Description: dict.Description,
			Type:        dict.Type,
			CreateTime:  createTime,
		})
	}

	return list, total, nil
}

// Add 添加字典
func (s *sSysDict) Add(ctx context.Context, req *v1.DictAddReq) error {
	// 检查字典编码是否已存在
	count, err := dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().DictCode, req.DictCode).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Count()
	if err != nil {
		g.Log().Error(ctx, "检查字典编码失败:", err)
		return errors.New("检查字典编码失败")
	}
	if count > 0 {
		return errors.New("字典编码已存在")
	}

	// 检查字典名称是否已存在
	count, err = dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().DictName, req.DictName).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Count()
	if err != nil {
		g.Log().Error(ctx, "检查字典名称失败:", err)
		return errors.New("检查字典名称失败")
	}
	if count > 0 {
		return errors.New("字典名称已存在")
	}

	// 插入字典数据
	_, err = dao.SysDict.Ctx(ctx).Insert(do.SysDict{
		Id:          guid.S(),
		DictName:    req.DictName,
		DictCode:    req.DictCode,
		Description: req.Description,
		Type:        req.Type,
		DelFlag:     0,
		CreateTime:  gtime.Now(),
		UpdateTime:  gtime.Now(),
	})
	if err != nil {
		g.Log().Error(ctx, "添加字典失败:", err)
		return errors.New("添加字典失败")
	}

	g.Log().Infof(ctx, "添加字典成功: dictName=%s, dictCode=%s", req.DictName, req.DictCode)
	return nil
}

// Edit 编辑字典
func (s *sSysDict) Edit(ctx context.Context, req *v1.DictEditReq) error {
	// 检查字典是否存在
	var dict entity.SysDict
	err := dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().Id, req.Id).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Scan(&dict)
	if err != nil {
		g.Log().Error(ctx, "查询字典失败:", err)
		return errors.New("查询字典失败")
	}
	if dict.Id == "" {
		return errors.New("字典不存在")
	}

	// 如果修改了字典编码，检查是否重复
	if req.DictCode != "" && req.DictCode != dict.DictCode {
		count, err := dao.SysDict.Ctx(ctx).
			Where(dao.SysDict.Columns().DictCode, req.DictCode).
			Where(dao.SysDict.Columns().DelFlag, 0).
			WhereNot(dao.SysDict.Columns().Id, req.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查字典编码失败:", err)
			return errors.New("检查字典编码失败")
		}
		if count > 0 {
			return errors.New("字典编码已存在")
		}
	}

	// 如果修改了字典名称，检查是否重复
	if req.DictName != "" && req.DictName != dict.DictName {
		count, err := dao.SysDict.Ctx(ctx).
			Where(dao.SysDict.Columns().DictName, req.DictName).
			Where(dao.SysDict.Columns().DelFlag, 0).
			WhereNot(dao.SysDict.Columns().Id, req.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查字典名称失败:", err)
			return errors.New("检查字典名称失败")
		}
		if count > 0 {
			return errors.New("字典名称已存在")
		}
	}

	// 更新字典数据
	updateData := do.SysDict{
		UpdateTime: gtime.Now(),
	}
	if req.DictName != "" {
		updateData.DictName = req.DictName
	}
	if req.DictCode != "" {
		updateData.DictCode = req.DictCode
	}
	if req.Description != "" {
		updateData.Description = req.Description
	}
	updateData.Type = req.Type

	_, err = dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().Id, req.Id).
		Update(updateData)
	if err != nil {
		g.Log().Error(ctx, "更新字典失败:", err)
		return errors.New("更新字典失败")
	}

	g.Log().Infof(ctx, "编辑字典成功: id=%s", req.Id)
	return nil
}

// Delete 删除字典
func (s *sSysDict) Delete(ctx context.Context, id string) error {
	// 检查字典是否存在
	var dict entity.SysDict
	err := dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().Id, id).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Scan(&dict)
	if err != nil {
		g.Log().Error(ctx, "查询字典失败:", err)
		return errors.New("查询字典失败")
	}
	if dict.Id == "" {
		return errors.New("字典不存在")
	}

	// 逻辑删除字典
	_, err = dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().Id, id).
		Update(do.SysDict{
			DelFlag:    1,
			UpdateTime: gtime.Now(),
		})
	if err != nil {
		g.Log().Error(ctx, "删除字典失败:", err)
		return errors.New("删除字典失败")
	}

	// 删除关联的字典项
	_, err = dao.SysDictItem.Ctx(ctx).
		Where(dao.SysDictItem.Columns().DictId, id).
		Delete()
	if err != nil {
		g.Log().Error(ctx, "删除字典项失败:", err)
		// 不返回错误，字典项删除失败不影响主逻辑
	}

	g.Log().Infof(ctx, "删除字典成功: id=%s", id)
	return nil
}

// GetById 根据ID获取字典详情
func (s *sSysDict) GetById(ctx context.Context, id string) (*v1.DictInfo, error) {
	var dict entity.SysDict
	err := dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().Id, id).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Scan(&dict)
	if err != nil {
		g.Log().Error(ctx, "查询字典失败:", err)
		return nil, errors.New("查询字典失败")
	}
	if dict.Id == "" {
		return nil, errors.New("字典不存在")
	}

	createTime := ""
	if dict.CreateTime != nil {
		createTime = dict.CreateTime.String()
	}

	return &v1.DictInfo{
		Id:          dict.Id,
		DictName:    dict.DictName,
		DictCode:    dict.DictCode,
		Description: dict.Description,
		Type:        dict.Type,
		CreateTime:  createTime,
	}, nil
}

// GetItemsByDictId 根据字典ID获取字典项列表
func (s *sSysDict) GetItemsByDictId(ctx context.Context, dictId string) ([]*v1.DictItemInfo, error) {
	var items []*entity.SysDictItem
	err := dao.SysDictItem.Ctx(ctx).
		Where(dao.SysDictItem.Columns().DictId, dictId).
		Where(dao.SysDictItem.Columns().Status, 1).
		OrderAsc(dao.SysDictItem.Columns().SortOrder).
		Scan(&items)
	if err != nil {
		g.Log().Error(ctx, "查询字典项失败:", err)
		return nil, errors.New("查询字典项失败")
	}

	list := make([]*v1.DictItemInfo, 0, len(items))
	for _, item := range items {
		list = append(list, &v1.DictItemInfo{
			Id:          item.Id,
			DictId:      item.DictId,
			ItemText:    item.ItemText,
			ItemValue:   item.ItemValue,
			Description: item.Description,
			SortOrder:   item.SortOrder,
			Status:      item.Status,
		})
	}

	return list, nil
}

// GetItemsByDictCode 根据字典编码获取字典项列表
func (s *sSysDict) GetItemsByDictCode(ctx context.Context, dictCode string) ([]*v1.DictItemInfo, error) {
	// 先查询字典
	var dict entity.SysDict
	err := dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().DictCode, dictCode).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Scan(&dict)
	if err != nil {
		g.Log().Error(ctx, "查询字典失败:", err)
		return nil, errors.New("查询字典失败")
	}
	if dict.Id == "" {
		return nil, errors.New("字典不存在")
	}

	// 查询字典项
	return s.GetItemsByDictId(ctx, dict.Id)
}

// AddItem 添加字典项
func (s *sSysDict) AddItem(ctx context.Context, req *v1.DictItemAddReq) error {
	// 检查字典是否存在
	count, err := dao.SysDict.Ctx(ctx).
		Where(dao.SysDict.Columns().Id, req.DictId).
		Where(dao.SysDict.Columns().DelFlag, 0).
		Count()
	if err != nil {
		g.Log().Error(ctx, "检查字典失败:", err)
		return errors.New("检查字典失败")
	}
	if count == 0 {
		return errors.New("字典不存在")
	}

	// 检查字典项值是否已存在
	count, err = dao.SysDictItem.Ctx(ctx).
		Where(dao.SysDictItem.Columns().DictId, req.DictId).
		Where(dao.SysDictItem.Columns().ItemValue, req.ItemValue).
		Count()
	if err != nil {
		g.Log().Error(ctx, "检查字典项失败:", err)
		return errors.New("检查字典项失败")
	}
	if count > 0 {
		return errors.New("字典项值已存在")
	}

	// 插入字典项
	_, err = dao.SysDictItem.Ctx(ctx).Insert(do.SysDictItem{
		Id:          guid.S(),
		DictId:      req.DictId,
		ItemText:    req.ItemText,
		ItemValue:   req.ItemValue,
		Description: req.Description,
		SortOrder:   req.SortOrder,
		Status:      req.Status,
		CreateTime:  gtime.Now(),
		UpdateTime:  gtime.Now(),
	})
	if err != nil {
		g.Log().Error(ctx, "添加字典项失败:", err)
		return errors.New("添加字典项失败")
	}

	g.Log().Infof(ctx, "添加字典项成功: dictId=%s, itemText=%s", req.DictId, req.ItemText)
	return nil
}

// EditItem 编辑字典项
func (s *sSysDict) EditItem(ctx context.Context, req *v1.DictItemEditReq) error {
	// 检查字典项是否存在
	var item entity.SysDictItem
	err := dao.SysDictItem.Ctx(ctx).
		Where(dao.SysDictItem.Columns().Id, req.Id).
		Scan(&item)
	if err != nil {
		g.Log().Error(ctx, "查询字典项失败:", err)
		return errors.New("查询字典项失败")
	}
	if item.Id == "" {
		return errors.New("字典项不存在")
	}

	// 如果修改了字典项值，检查是否重复
	if req.ItemValue != "" && req.ItemValue != item.ItemValue {
		count, err := dao.SysDictItem.Ctx(ctx).
			Where(dao.SysDictItem.Columns().DictId, item.DictId).
			Where(dao.SysDictItem.Columns().ItemValue, req.ItemValue).
			WhereNot(dao.SysDictItem.Columns().Id, req.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查字典项失败:", err)
			return errors.New("检查字典项失败")
		}
		if count > 0 {
			return errors.New("字典项值已存在")
		}
	}

	// 更新字典项
	updateData := do.SysDictItem{
		UpdateTime: gtime.Now(),
	}
	if req.ItemText != "" {
		updateData.ItemText = req.ItemText
	}
	if req.ItemValue != "" {
		updateData.ItemValue = req.ItemValue
	}
	if req.Description != "" {
		updateData.Description = req.Description
	}
	if req.SortOrder > 0 {
		updateData.SortOrder = req.SortOrder
	}
	if req.Status >= 0 {
		updateData.Status = req.Status
	}

	_, err = dao.SysDictItem.Ctx(ctx).
		Where(dao.SysDictItem.Columns().Id, req.Id).
		Update(updateData)
	if err != nil {
		g.Log().Error(ctx, "更新字典项失败:", err)
		return errors.New("更新字典项失败")
	}

	g.Log().Infof(ctx, "编辑字典项成功: id=%s", req.Id)
	return nil
}

// DeleteItem 删除字典项
func (s *sSysDict) DeleteItem(ctx context.Context, id string) error {
	// 检查字典项是否存在
	count, err := dao.SysDictItem.Ctx(ctx).
		Where(dao.SysDictItem.Columns().Id, id).
		Count()
	if err != nil {
		g.Log().Error(ctx, "查询字典项失败:", err)
		return errors.New("查询字典项失败")
	}
	if count == 0 {
		return errors.New("字典项不存在")
	}

	// 删除字典项
	_, err = dao.SysDictItem.Ctx(ctx).
		Where(dao.SysDictItem.Columns().Id, id).
		Delete()
	if err != nil {
		g.Log().Error(ctx, "删除字典项失败:", err)
		return errors.New("删除字典项失败")
	}

	g.Log().Infof(ctx, "删除字典项成功: id=%s", id)
	return nil
}
