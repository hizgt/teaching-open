// =================================================================================
// Service interface for sys depart module
// =================================================================================

package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ISysDepart 部门服务接口
type ISysDepart interface {
	// GetTree 获取部门树
	GetTree(ctx context.Context) ([]*v1.DepartInfo, error)

	// Add 添加部门
	Add(ctx context.Context, req *v1.DepartAddReq) error

	// Edit 编辑部门
	Edit(ctx context.Context, req *v1.DepartEditReq) error

	// Delete 删除部门
	Delete(ctx context.Context, id string) error

	// GetById 根据ID获取部门详情
	GetById(ctx context.Context, id string) (*v1.DepartInfo, error)

	// GetIdTree 获取部门ID树
	GetIdTree(ctx context.Context) ([]*v1.DepartIdNode, error)

	// SearchByKeyword 关键字搜索部门
	SearchByKeyword(ctx context.Context, keyword string) ([]*v1.DepartInfo, error)

	// GetUserDeparts 获取用户部门
	GetUserDeparts(ctx context.Context, userId string) ([]string, []*v1.DepartInfo, error)

	// SaveUserDeparts 保存用户部门
	SaveUserDeparts(ctx context.Context, userId string, departIds []string) error
}

var localSysDepart ISysDepart

// SysDepart 获取部门服务实例
func SysDepart() ISysDepart {
	if localSysDepart == nil {
		panic("implement not found for interface ISysDepart, forgot register?")
	}
	return localSysDepart
}

// RegisterSysDepart 注册部门服务实现
func RegisterSysDepart(i ISysDepart) {
	localSysDepart = i
}
