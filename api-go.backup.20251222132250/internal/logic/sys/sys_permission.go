// =================================================================================
// Logic implementation for sys permission module
// =================================================================================

package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

func init() {
	service.RegisterSysPermission(&sSysPermission{})
}

// sSysPermission 权限服务实现
type sSysPermission struct{}

// GetList 获取权限列表（平铺）
func (s *sSysPermission) GetList(ctx context.Context) ([]*v1.PermissionInfo, error) {
	var list []*entity.SysPermission
	err := dao.SysPermission.Ctx(ctx).
		Where("del_flag", 0).
		OrderAsc("sort_no").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	result := make([]*v1.PermissionInfo, 0, len(list))
	for _, item := range list {
		result = append(result, s.entityToInfo(item))
	}
	return result, nil
}

// GetTree 获取权限树
func (s *sSysPermission) GetTree(ctx context.Context) ([]*v1.PermissionInfo, error) {
	list, err := s.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return s.buildTree(list, ""), nil
}

// Add 添加权限
func (s *sSysPermission) Add(ctx context.Context, req *v1.PermissionAddReq) error {
	permission := &entity.SysPermission{
		Id:                 guid.S(),
		ParentId:           req.ParentId,
		Name:               req.Name,
		Url:                req.Url,
		Component:          req.Component,
		ComponentName:      req.ComponentName,
		Redirect:           req.Redirect,
		MenuType:           req.MenuType,
		Perms:              req.Perms,
		PermsType:          req.PermsType,
		SortNo:             req.SortNo,
		AlwaysShow:         req.AlwaysShow,
		Icon:               req.Icon,
		IsRoute:            req.IsRoute,
		IsLeaf:             req.IsLeaf,
		KeepAlive:          req.KeepAlive,
		Hidden:             req.Hidden,
		Description:        req.Description,
		Status:             req.Status,
		InternalOrExternal: req.InternalOrExternal,
		CreateTime:         gtime.Now(),
		DelFlag:            0,
	}
	_, err := dao.SysPermission.Ctx(ctx).Insert(permission)
	return err
}

// Edit 编辑权限
func (s *sSysPermission) Edit(ctx context.Context, req *v1.PermissionEditReq) error {
	count, err := dao.SysPermission.Ctx(ctx).Where("id", req.Id).Where("del_flag", 0).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New(g.I18n().T(ctx, "权限不存在"))
	}

	_, err = dao.SysPermission.Ctx(ctx).
		Where("id", req.Id).
		Data(g.Map{
			"parent_id":            req.ParentId,
			"name":                 req.Name,
			"url":                  req.Url,
			"component":            req.Component,
			"component_name":       req.ComponentName,
			"redirect":             req.Redirect,
			"menu_type":            req.MenuType,
			"perms":                req.Perms,
			"perms_type":           req.PermsType,
			"sort_no":              req.SortNo,
			"always_show":          req.AlwaysShow,
			"icon":                 req.Icon,
			"is_route":             req.IsRoute,
			"is_leaf":              req.IsLeaf,
			"keep_alive":           req.KeepAlive,
			"hidden":               req.Hidden,
			"description":          req.Description,
			"status":               req.Status,
			"internal_or_external": req.InternalOrExternal,
			"update_time":          gtime.Now(),
		}).
		Update()
	return err
}

// Delete 删除权限
func (s *sSysPermission) Delete(ctx context.Context, id string) error {
	ids := strings.Split(id, ",")
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, permId := range ids {
			permId = strings.TrimSpace(permId)
			if permId == "" {
				continue
			}
			childCount, err := dao.SysPermission.Ctx(ctx).
				Where("parent_id", permId).
				Where("del_flag", 0).
				Count()
			if err != nil {
				return err
			}
			if childCount > 0 {
				return gerror.New(g.I18n().T(ctx, "存在子菜单，无法删除"))
			}
			_, err = dao.SysPermission.Ctx(ctx).
				Where("id", permId).
				Data(g.Map{
					"del_flag":    1,
					"update_time": gtime.Now(),
				}).
				Update()
			if err != nil {
				return err
			}
			_, err = dao.SysRolePermission.Ctx(ctx).
				Where("permission_id", permId).
				Delete()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// GetById 根据ID获取权限详情
func (s *sSysPermission) GetById(ctx context.Context, id string) (*v1.PermissionInfo, error) {
	var permission *entity.SysPermission
	err := dao.SysPermission.Ctx(ctx).
		Where("id", id).
		Where("del_flag", 0).
		Scan(&permission)
	if err != nil {
		return nil, err
	}
	if permission == nil {
		return nil, nil
	}
	return s.entityToInfo(permission), nil
}

// GetRolePermissions 获取角色的权限ID列表
func (s *sSysPermission) GetRolePermissions(ctx context.Context, roleId string) ([]string, error) {
	var rolePermissions []*entity.SysRolePermission
	err := dao.SysRolePermission.Ctx(ctx).
		Where("role_id", roleId).
		Scan(&rolePermissions)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(rolePermissions))
	for _, rp := range rolePermissions {
		ids = append(ids, rp.PermissionId)
	}
	return ids, nil
}

// SaveRolePermissions 保存角色权限
func (s *sSysPermission) SaveRolePermissions(ctx context.Context, roleId string, permissionIds []string) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.SysRolePermission.Ctx(ctx).
			Where("role_id", roleId).
			Delete()
		if err != nil {
			return err
		}
		if len(permissionIds) > 0 {
			records := make([]g.Map, 0, len(permissionIds))
			for _, permId := range permissionIds {
				records = append(records, g.Map{
					"id":            guid.S(),
					"role_id":       roleId,
					"permission_id": permId,
				})
			}
			_, err = dao.SysRolePermission.Ctx(ctx).Insert(records)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// GetUserPermissions 获取用户权限（菜单和权限标识）
func (s *sSysPermission) GetUserPermissions(ctx context.Context, userId string) ([]*v1.PermissionInfo, []string, error) {
	var userRoles []*entity.SysUserRole
	err := dao.SysUserRole.Ctx(ctx).
		Where("user_id", userId).
		Scan(&userRoles)
	if err != nil {
		return nil, nil, err
	}
	if len(userRoles) == 0 {
		return []*v1.PermissionInfo{}, []string{}, nil
	}
	roleIds := make([]string, 0, len(userRoles))
	for _, ur := range userRoles {
		roleIds = append(roleIds, ur.RoleId)
	}
	var rolePermissions []*entity.SysRolePermission
	err = dao.SysRolePermission.Ctx(ctx).
		WhereIn("role_id", roleIds).
		Scan(&rolePermissions)
	if err != nil {
		return nil, nil, err
	}
	if len(rolePermissions) == 0 {
		return []*v1.PermissionInfo{}, []string{}, nil
	}
	permIdMap := make(map[string]bool)
	for _, rp := range rolePermissions {
		permIdMap[rp.PermissionId] = true
	}
	permIds := make([]string, 0, len(permIdMap))
	for id := range permIdMap {
		permIds = append(permIds, id)
	}
	var permissions []*entity.SysPermission
	err = dao.SysPermission.Ctx(ctx).
		WhereIn("id", permIds).
		Where("del_flag", 0).
		Where("status", "1").
		OrderAsc("sort_no").
		Scan(&permissions)
	if err != nil {
		return nil, nil, err
	}
	menuList := make([]*v1.PermissionInfo, 0)
	authList := make([]string, 0)
	for _, perm := range permissions {
		if perm.MenuType == 2 {
			if perm.Perms != "" {
				authList = append(authList, perm.Perms)
			}
		} else {
			menuList = append(menuList, s.entityToInfo(perm))
		}
	}
	menuTree := s.buildTree(menuList, "")
	return menuTree, authList, nil
}

// entityToInfo 实体转信息结构
func (s *sSysPermission) entityToInfo(e *entity.SysPermission) *v1.PermissionInfo {
	info := &v1.PermissionInfo{
		Id:                 e.Id,
		ParentId:           e.ParentId,
		Name:               e.Name,
		Url:                e.Url,
		Component:          e.Component,
		ComponentName:      e.ComponentName,
		Redirect:           e.Redirect,
		MenuType:           e.MenuType,
		Perms:              e.Perms,
		PermsType:          e.PermsType,
		SortNo:             e.SortNo,
		AlwaysShow:         e.AlwaysShow,
		Icon:               e.Icon,
		IsRoute:            e.IsRoute,
		IsLeaf:             e.IsLeaf,
		KeepAlive:          e.KeepAlive,
		Hidden:             e.Hidden,
		Description:        e.Description,
		Status:             e.Status,
		InternalOrExternal: e.InternalOrExternal,
	}
	if e.CreateTime != nil {
		info.CreateTime = e.CreateTime.Format("Y-m-d H:i:s")
	}
	return info
}

// buildTree 构建权限树
func (s *sSysPermission) buildTree(list []*v1.PermissionInfo, parentId string) []*v1.PermissionInfo {
	result := make([]*v1.PermissionInfo, 0)
	for _, item := range list {
		if item.ParentId == parentId {
			item.Children = s.buildTree(list, item.Id)
			result = append(result, item)
		}
	}
	return result
}
