// =================================================================================
// Service interface for sys dict module
// =================================================================================

package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ISysDict 字典服务接口
type ISysDict interface {
	// GetList 获取字典列表
	GetList(ctx context.Context, req *v1.DictListReq) (list []*v1.DictInfo, total int64, err error)
	// Add 添加字典
	Add(ctx context.Context, req *v1.DictAddReq) error
	// Edit 编辑字典
	Edit(ctx context.Context, req *v1.DictEditReq) error
	// Delete 删除字典
	Delete(ctx context.Context, id string) error
	// GetById 根据ID获取字典详情
	GetById(ctx context.Context, id string) (*v1.DictInfo, error)
	// GetItemsByDictId 根据字典ID获取字典项列表
	GetItemsByDictId(ctx context.Context, dictId string) ([]*v1.DictItemInfo, error)
	// GetItemsByDictCode 根据字典编码获取字典项列表
	GetItemsByDictCode(ctx context.Context, dictCode string) ([]*v1.DictItemInfo, error)
	// AddItem 添加字典项
	AddItem(ctx context.Context, req *v1.DictItemAddReq) error
	// EditItem 编辑字典项
	EditItem(ctx context.Context, req *v1.DictItemEditReq) error
	// DeleteItem 删除字典项
	DeleteItem(ctx context.Context, id string) error
}

var localSysDict ISysDict

// SysDict 获取字典服务实例
func SysDict() ISysDict {
	if localSysDict == nil {
		panic("implement not found for interface ISysDict, forgot register?")
	}
	return localSysDict
}

// RegisterSysDict 注册字典服务实现
func RegisterSysDict(i ISysDict) {
	localSysDict = i
}
