package service

import (
	"context"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/model/vo"
)

// SysPermissionService 权限服务接口
type SysPermissionService interface {
	GetTree(ctx context.Context) ([]*vo.PermissionTreeItem, error)
	GetById(ctx context.Context, id string) (*entity.SysPermission, error)
	Create(ctx context.Context, req *vo.PermissionCreateReq, createBy string) (*vo.PermissionCreateRes, error)
	Update(ctx context.Context, req *vo.PermissionUpdateReq, updateBy string) error
	Delete(ctx context.Context, id string) error
	GetUserPermission(ctx context.Context, userId string) (*vo.UserPermissionRes, error)
	GetUserMenus(ctx context.Context, userId string) ([]*vo.PermissionTreeItem, error)
	GetUserPermCodes(ctx context.Context, userId string) ([]string, error)
}

// sysPermissionServiceImpl 权限服务实现
type sysPermissionServiceImpl struct{}

// NewSysPermissionService 创建权限服务实例
func NewSysPermissionService() SysPermissionService {
	return &sysPermissionServiceImpl{}
}

// GetTree 获取权限树
func (s *sysPermissionServiceImpl) GetTree(ctx context.Context) ([]*vo.PermissionTreeItem, error) {
	var permissions []entity.SysPermission
	err := dao.SysPermission.Ctx(ctx).
		Where(dao.SysPermission.Columns().DelFlag, 0).
		OrderAsc(dao.SysPermission.Columns().SortNo).
		Scan(&permissions)
	if err != nil {
		return nil, err
	}
	return s.buildTree(permissions, ""), nil
}

// GetById 根据ID获取权限
func (s *sysPermissionServiceImpl) GetById(ctx context.Context, id string) (*entity.SysPermission, error) {
	var perm entity.SysPermission
	err := dao.SysPermission.Ctx(ctx).Where(dao.SysPermission.Columns().Id, id).Scan(&perm)
	if err != nil {
		return nil, err
	}
	if perm.Id == "" {
		return nil, nil
	}
	return &perm, nil
}

// Create 创建权限
func (s *sysPermissionServiceImpl) Create(ctx context.Context, req *vo.PermissionCreateReq, createBy string) (*vo.PermissionCreateRes, error) {
	permId := guid.S()
	perm := entity.SysPermission{
		Id:                 permId,
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
		DelFlag:            0,
		CreateBy:           createBy,
		CreateTime:         gtime.Now(),
		UpdateBy:           createBy,
		UpdateTime:         gtime.Now(),
	}

	_, err := dao.SysPermission.Ctx(ctx).Data(perm).Insert()
	if err != nil {
		return nil, err
	}

	// 更新父节点的isLeaf状态
	if req.ParentId != "" {
		_, _ = dao.SysPermission.Ctx(ctx).Data(g.Map{
			dao.SysPermission.Columns().IsLeaf: false,
		}).Where(dao.SysPermission.Columns().Id, req.ParentId).Update()
	}

	return &vo.PermissionCreateRes{Id: permId}, nil
}

// Update 更新权限
func (s *sysPermissionServiceImpl) Update(ctx context.Context, req *vo.PermissionUpdateReq, updateBy string) error {
	perm, err := s.GetById(ctx, req.Id)
	if err != nil {
		return err
	}
	if perm == nil {
		return gerror.New("权限不存在")
	}

	_, err = dao.SysPermission.Ctx(ctx).Data(g.Map{
		dao.SysPermission.Columns().ParentId:           req.ParentId,
		dao.SysPermission.Columns().Name:               req.Name,
		dao.SysPermission.Columns().Url:                req.Url,
		dao.SysPermission.Columns().Component:          req.Component,
		dao.SysPermission.Columns().ComponentName:      req.ComponentName,
		dao.SysPermission.Columns().Redirect:           req.Redirect,
		dao.SysPermission.Columns().MenuType:           req.MenuType,
		dao.SysPermission.Columns().Perms:              req.Perms,
		dao.SysPermission.Columns().PermsType:          req.PermsType,
		dao.SysPermission.Columns().SortNo:             req.SortNo,
		dao.SysPermission.Columns().AlwaysShow:         req.AlwaysShow,
		dao.SysPermission.Columns().Icon:               req.Icon,
		dao.SysPermission.Columns().IsRoute:            req.IsRoute,
		dao.SysPermission.Columns().IsLeaf:             req.IsLeaf,
		dao.SysPermission.Columns().KeepAlive:          req.KeepAlive,
		dao.SysPermission.Columns().Hidden:             req.Hidden,
		dao.SysPermission.Columns().Description:        req.Description,
		dao.SysPermission.Columns().Status:             req.Status,
		dao.SysPermission.Columns().InternalOrExternal: req.InternalOrExternal,
		dao.SysPermission.Columns().UpdateBy:           updateBy,
		dao.SysPermission.Columns().UpdateTime:         gtime.Now(),
	}).Where(dao.SysPermission.Columns().Id, req.Id).Update()

	return err
}

// Delete 删除权限
func (s *sysPermissionServiceImpl) Delete(ctx context.Context, id string) error {
	// 检查是否有子节点
	count, err := dao.SysPermission.Ctx(ctx).
		Where(dao.SysPermission.Columns().ParentId, id).
		Where(dao.SysPermission.Columns().DelFlag, 0).
		Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("存在子菜单，无法删除")
	}

	// 逻辑删除
	_, err = dao.SysPermission.Ctx(ctx).Data(g.Map{
		dao.SysPermission.Columns().DelFlag:    1,
		dao.SysPermission.Columns().UpdateTime: gtime.Now(),
	}).Where(dao.SysPermission.Columns().Id, id).Update()

	return err
}

// GetUserPermission 获取用户权限信息
func (s *sysPermissionServiceImpl) GetUserPermission(ctx context.Context, userId string) (*vo.UserPermissionRes, error) {
	// 获取用户角色
	var roleIds []string
	err := dao.SysUserRole.Ctx(ctx).
		Fields(dao.SysUserRole.Columns().RoleId).
		Where(dao.SysUserRole.Columns().UserId, userId).
		Scan(&roleIds)
	if err != nil {
		return nil, err
	}

	// 获取所有权限
	allPerms, err := s.GetTree(ctx)
	if err != nil {
		return nil, err
	}

	if len(roleIds) == 0 {
		return &vo.UserPermissionRes{
			AllAuth: s.treeToList(allPerms),
			Auth:    []vo.PermissionTreeItem{},
			Menu:    []vo.PermissionTreeItem{},
		}, nil
	}

	// 获取角色权限ID
	var permIds []string
	err = dao.SysRolePermission.Ctx(ctx).
		Fields(dao.SysRolePermission.Columns().PermissionId).
		WhereIn(dao.SysRolePermission.Columns().RoleId, roleIds).
		Scan(&permIds)
	if err != nil {
		return nil, err
	}

	// 获取用户有权限的菜单
	var userPerms []entity.SysPermission
	if len(permIds) > 0 {
		err = dao.SysPermission.Ctx(ctx).
			Where(dao.SysPermission.Columns().DelFlag, 0).
			WhereIn(dao.SysPermission.Columns().Id, permIds).
			OrderAsc(dao.SysPermission.Columns().SortNo).
			Scan(&userPerms)
		if err != nil {
			return nil, err
		}
	}

	// 构建用户权限树
	userPermTree := s.buildTree(userPerms, "")

	// 过滤菜单(menuType=0或1)
	menuTree := s.filterMenus(userPermTree)

	return &vo.UserPermissionRes{
		AllAuth: s.treeToList(allPerms),
		Auth:    s.treeToList(userPermTree),
		Menu:    s.treeToList(menuTree),
	}, nil
}

// GetUserMenus 获取用户菜单
func (s *sysPermissionServiceImpl) GetUserMenus(ctx context.Context, userId string) ([]*vo.PermissionTreeItem, error) {
	// 获取用户角色
	var roleIds []string
	err := dao.SysUserRole.Ctx(ctx).
		Fields(dao.SysUserRole.Columns().RoleId).
		Where(dao.SysUserRole.Columns().UserId, userId).
		Scan(&roleIds)
	if err != nil {
		return nil, err
	}

	if len(roleIds) == 0 {
		return []*vo.PermissionTreeItem{}, nil
	}

	// 获取角色权限ID
	var permIds []string
	err = dao.SysRolePermission.Ctx(ctx).
		Fields(dao.SysRolePermission.Columns().PermissionId).
		WhereIn(dao.SysRolePermission.Columns().RoleId, roleIds).
		Scan(&permIds)
	if err != nil {
		return nil, err
	}

	if len(permIds) == 0 {
		return []*vo.PermissionTreeItem{}, nil
	}

	// 获取菜单权限
	var perms []entity.SysPermission
	err = dao.SysPermission.Ctx(ctx).
		Where(dao.SysPermission.Columns().DelFlag, 0).
		WhereIn(dao.SysPermission.Columns().Id, permIds).
		WhereIn(dao.SysPermission.Columns().MenuType, []int{0, 1}).
		OrderAsc(dao.SysPermission.Columns().SortNo).
		Scan(&perms)
	if err != nil {
		return nil, err
	}

	return s.buildTree(perms, ""), nil
}

// GetUserPermCodes 获取用户权限编码
func (s *sysPermissionServiceImpl) GetUserPermCodes(ctx context.Context, userId string) ([]string, error) {
	// 获取用户角色
	var roleIds []string
	err := dao.SysUserRole.Ctx(ctx).
		Fields(dao.SysUserRole.Columns().RoleId).
		Where(dao.SysUserRole.Columns().UserId, userId).
		Scan(&roleIds)
	if err != nil {
		return nil, err
	}

	if len(roleIds) == 0 {
		return []string{}, nil
	}

	// 获取角色权限ID
	var permIds []string
	err = dao.SysRolePermission.Ctx(ctx).
		Fields(dao.SysRolePermission.Columns().PermissionId).
		WhereIn(dao.SysRolePermission.Columns().RoleId, roleIds).
		Scan(&permIds)
	if err != nil {
		return nil, err
	}

	if len(permIds) == 0 {
		return []string{}, nil
	}

	// 获取权限编码
	var permCodes []string
	err = dao.SysPermission.Ctx(ctx).
		Fields(dao.SysPermission.Columns().Perms).
		Where(dao.SysPermission.Columns().DelFlag, 0).
		WhereIn(dao.SysPermission.Columns().Id, permIds).
		WhereNot(dao.SysPermission.Columns().Perms, "").
		Scan(&permCodes)
	if err != nil {
		return nil, err
	}

	// 去重
	codeMap := make(map[string]bool)
	var result []string
	for _, code := range permCodes {
		codes := strings.Split(code, ",")
		for _, c := range codes {
			c = strings.TrimSpace(c)
			if c != "" && !codeMap[c] {
				codeMap[c] = true
				result = append(result, c)
			}
		}
	}

	return result, nil
}

// buildTree 构建权限树
func (s *sysPermissionServiceImpl) buildTree(perms []entity.SysPermission, parentId string) []*vo.PermissionTreeItem {
	var tree []*vo.PermissionTreeItem

	for _, perm := range perms {
		if perm.ParentId == parentId {
			item := &vo.PermissionTreeItem{
				Id:                 perm.Id,
				ParentId:           perm.ParentId,
				Key:                perm.Id,
				Title:              perm.Name,
				Name:               perm.Name,
				Url:                perm.Url,
				Component:          perm.Component,
				ComponentName:      perm.ComponentName,
				Redirect:           perm.Redirect,
				MenuType:           perm.MenuType,
				Perms:              perm.Perms,
				PermsType:          perm.PermsType,
				SortNo:             perm.SortNo,
				AlwaysShow:         perm.AlwaysShow,
				Icon:               perm.Icon,
				IsRoute:            perm.IsRoute,
				IsLeaf:             perm.IsLeaf,
				KeepAlive:          perm.KeepAlive,
				Hidden:             perm.Hidden,
				Description:        perm.Description,
				Status:             perm.Status,
				InternalOrExternal: perm.InternalOrExternal,
			}
			item.Children = s.buildTree(perms, perm.Id)
			tree = append(tree, item)
		}
	}

	// 按sortNo排序
	sort.Slice(tree, func(i, j int) bool {
		return tree[i].SortNo < tree[j].SortNo
	})

	return tree
}

// filterMenus 过滤菜单(只保留menuType=0或1)
func (s *sysPermissionServiceImpl) filterMenus(tree []*vo.PermissionTreeItem) []*vo.PermissionTreeItem {
	var result []*vo.PermissionTreeItem
	for _, item := range tree {
		if item.MenuType == 0 || item.MenuType == 1 {
			newItem := *item
			newItem.Children = s.filterMenus(item.Children)
			result = append(result, &newItem)
		}
	}
	return result
}

// treeToList 树转列表
func (s *sysPermissionServiceImpl) treeToList(tree []*vo.PermissionTreeItem) []vo.PermissionTreeItem {
	var list []vo.PermissionTreeItem
	for _, item := range tree {
		list = append(list, *item)
	}
	return list
}
