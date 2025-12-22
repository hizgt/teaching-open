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

type sSysRole struct{}

func init() {
	service.SysRole = &sSysRole{}
}

// GetList 获取角色列表
func (s *sSysRole) GetList(ctx context.Context, req *v1.RoleListReq) (list []*v1.RoleInfo, total int64, err error) {
	// 构建查询
	m := dao.SysRole.Ctx(ctx)

	// 角色名称模糊查询
	if req.RoleName != "" {
		m = m.WhereLike(dao.SysRole.Columns().RoleName, "%"+req.RoleName+"%")
	}

	// 角色编码模糊查询
	if req.RoleCode != "" {
		m = m.WhereLike(dao.SysRole.Columns().RoleCode, "%"+req.RoleCode+"%")
	}

	// 分页查询
	var roles []*entity.SysRole
	var totalInt int
	err = m.Page(req.Page, req.PageSize).
		OrderDesc(dao.SysRole.Columns().CreateTime).
		ScanAndCount(&roles, &totalInt, false)
	if err != nil {
		g.Log().Error(ctx, "查询角色列表失败:", err)
		return nil, 0, errors.New("查询角色列表失败")
	}
	total = int64(totalInt)

	// 转换为RoleInfo
	list = make([]*v1.RoleInfo, 0, len(roles))
	for _, role := range roles {
		createTime := ""
		if role.CreateTime != nil {
			createTime = role.CreateTime.String()
		}
		list = append(list, &v1.RoleInfo{
			Id:          role.Id,
			RoleName:    role.RoleName,
			RoleCode:    role.RoleCode,
			Description: role.Description,
			RoleLevel:   role.RoleLevel,
			CreateTime:  createTime,
		})
	}

	return list, total, nil
}

// Add 新增角色
func (s *sSysRole) Add(ctx context.Context, req *v1.RoleAddReq) error {
	// 1. 检查角色编码是否已存在
	count, err := dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().RoleCode, req.RoleCode).
		Count()
	if err != nil {
		g.Log().Error(ctx, "检查角色编码失败:", err)
		return errors.New("检查角色编码失败")
	}
	if count > 0 {
		return errors.New("角色编码已存在")
	}

	// 2. 检查角色名称是否已存在
	count, err = dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().RoleName, req.RoleName).
		Count()
	if err != nil {
		g.Log().Error(ctx, "检查角色名称失败:", err)
		return errors.New("检查角色名称失败")
	}
	if count > 0 {
		return errors.New("角色名称已存在")
	}

	// 3. 插入角色数据
	_, err = dao.SysRole.Ctx(ctx).Insert(do.SysRole{
		Id:          guid.S(),
		RoleName:    req.RoleName,
		RoleCode:    req.RoleCode,
		Description: req.Description,
		RoleLevel:   req.RoleLevel,
		CreateTime:  gtime.Now(),
		UpdateTime:  gtime.Now(),
	})
	if err != nil {
		g.Log().Error(ctx, "新增角色失败:", err)
		return errors.New("新增角色失败")
	}

	g.Log().Infof(ctx, "新增角色成功: roleName=%s, roleCode=%s", req.RoleName, req.RoleCode)
	return nil
}

// Edit 编辑角色
func (s *sSysRole) Edit(ctx context.Context, req *v1.RoleEditReq) error {
	// 1. 检查角色是否存在
	var role entity.SysRole
	err := dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Id, req.Id).
		Scan(&role)
	if err != nil {
		g.Log().Error(ctx, "查询角色失败:", err)
		return errors.New("查询角色失败")
	}
	if role.Id == "" {
		return errors.New("角色不存在")
	}

	// 2. 如果修改了角色编码，检查是否重复
	if req.RoleCode != "" && req.RoleCode != role.RoleCode {
		count, err := dao.SysRole.Ctx(ctx).
			Where(dao.SysRole.Columns().RoleCode, req.RoleCode).
			WhereNot(dao.SysRole.Columns().Id, req.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查角色编码失败:", err)
			return errors.New("检查角色编码失败")
		}
		if count > 0 {
			return errors.New("角色编码已存在")
		}
	}

	// 3. 如果修改了角色名称，检查是否重复
	if req.RoleName != "" && req.RoleName != role.RoleName {
		count, err := dao.SysRole.Ctx(ctx).
			Where(dao.SysRole.Columns().RoleName, req.RoleName).
			WhereNot(dao.SysRole.Columns().Id, req.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查角色名称失败:", err)
			return errors.New("检查角色名称失败")
		}
		if count > 0 {
			return errors.New("角色名称已存在")
		}
	}

	// 4. 更新角色数据
	updateData := do.SysRole{
		UpdateTime: gtime.Now(),
	}
	if req.RoleName != "" {
		updateData.RoleName = req.RoleName
	}
	if req.RoleCode != "" {
		updateData.RoleCode = req.RoleCode
	}
	if req.Description != "" {
		updateData.Description = req.Description
	}
	updateData.RoleLevel = req.RoleLevel

	_, err = dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Id, req.Id).
		Update(updateData)
	if err != nil {
		g.Log().Error(ctx, "更新角色失败:", err)
		return errors.New("更新角色失败")
	}

	g.Log().Infof(ctx, "编辑角色成功: id=%s", req.Id)
	return nil
}

// Delete 删除角色
func (s *sSysRole) Delete(ctx context.Context, id string) error {
	// 1. 检查角色是否存在
	var role entity.SysRole
	err := dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Id, id).
		Scan(&role)
	if err != nil {
		g.Log().Error(ctx, "查询角色失败:", err)
		return errors.New("查询角色失败")
	}
	if role.Id == "" {
		return errors.New("角色不存在")
	}

	// 2. 检查是否有用户使用该角色
	count, err := dao.SysUserRole.Ctx(ctx).
		Where(dao.SysUserRole.Columns().RoleId, id).
		Count()
	if err != nil {
		g.Log().Error(ctx, "检查角色使用情况失败:", err)
		return errors.New("检查角色使用情况失败")
	}
	if count > 0 {
		return errors.New("该角色已被用户使用，无法删除")
	}

	// 3. 删除角色权限关联
	_, err = dao.SysRolePermission.Ctx(ctx).
		Where(dao.SysRolePermission.Columns().RoleId, id).
		Delete()
	if err != nil {
		g.Log().Error(ctx, "删除角色权限关联失败:", err)
		return errors.New("删除角色权限关联失败")
	}

	// 4. 删除角色
	_, err = dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Id, id).
		Delete()
	if err != nil {
		g.Log().Error(ctx, "删除角色失败:", err)
		return errors.New("删除角色失败")
	}

	g.Log().Infof(ctx, "删除角色成功: id=%s", id)
	return nil
}

// GetById 根据ID获取角色详情
func (s *sSysRole) GetById(ctx context.Context, id string) (*v1.RoleInfo, error) {
	// 查询角色
	var role entity.SysRole
	err := dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Id, id).
		Scan(&role)
	if err != nil {
		g.Log().Error(ctx, "查询角色失败:", err)
		return nil, errors.New("查询角色失败")
	}
	if role.Id == "" {
		return nil, errors.New("角色不存在")
	}

	createTime := ""
	if role.CreateTime != nil {
		createTime = role.CreateTime.String()
	}

	return &v1.RoleInfo{
		Id:          role.Id,
		RoleName:    role.RoleName,
		RoleCode:    role.RoleCode,
		Description: role.Description,
		RoleLevel:   role.RoleLevel,
		CreateTime:  createTime,
	}, nil
}

// GetAll 获取所有角色
func (s *sSysRole) GetAll(ctx context.Context) ([]*v1.RoleInfo, error) {
	var roles []*entity.SysRole
	err := dao.SysRole.Ctx(ctx).
		OrderAsc(dao.SysRole.Columns().RoleLevel).
		OrderAsc(dao.SysRole.Columns().CreateTime).
		Scan(&roles)
	if err != nil {
		g.Log().Error(ctx, "查询所有角色失败:", err)
		return nil, errors.New("查询所有角色失败")
	}

	list := make([]*v1.RoleInfo, 0, len(roles))
	for _, role := range roles {
		createTime := ""
		if role.CreateTime != nil {
			createTime = role.CreateTime.String()
		}
		list = append(list, &v1.RoleInfo{
			Id:          role.Id,
			RoleName:    role.RoleName,
			RoleCode:    role.RoleCode,
			Description: role.Description,
			RoleLevel:   role.RoleLevel,
			CreateTime:  createTime,
		})
	}

	return list, nil
}

// GetUserRoles 获取用户角色
func (s *sSysRole) GetUserRoles(ctx context.Context, userId string) ([]string, []*v1.RoleInfo, error) {
	// 查询用户角色关联
	var userRoles []*entity.SysUserRole
	err := dao.SysUserRole.Ctx(ctx).
		Where(dao.SysUserRole.Columns().UserId, userId).
		Scan(&userRoles)
	if err != nil {
		g.Log().Error(ctx, "查询用户角色失败:", err)
		return nil, nil, errors.New("查询用户角色失败")
	}

	if len(userRoles) == 0 {
		return []string{}, []*v1.RoleInfo{}, nil
	}

	// 获取角色ID列表
	roleIds := make([]string, 0, len(userRoles))
	for _, ur := range userRoles {
		roleIds = append(roleIds, ur.RoleId)
	}

	// 查询角色信息
	var roles []*entity.SysRole
	err = dao.SysRole.Ctx(ctx).
		WhereIn(dao.SysRole.Columns().Id, roleIds).
		Scan(&roles)
	if err != nil {
		g.Log().Error(ctx, "查询角色信息失败:", err)
		return nil, nil, errors.New("查询角色信息失败")
	}

	// 转换为RoleInfo
	list := make([]*v1.RoleInfo, 0, len(roles))
	for _, role := range roles {
		createTime := ""
		if role.CreateTime != nil {
			createTime = role.CreateTime.String()
		}
		list = append(list, &v1.RoleInfo{
			Id:          role.Id,
			RoleName:    role.RoleName,
			RoleCode:    role.RoleCode,
			Description: role.Description,
			RoleLevel:   role.RoleLevel,
			CreateTime:  createTime,
		})
	}

	return roleIds, list, nil
}

// SaveUserRoles 保存用户角色
func (s *sSysRole) SaveUserRoles(ctx context.Context, userId string, roleIds []string) error {
	// 1. 删除原有的用户角色关联
	_, err := dao.SysUserRole.Ctx(ctx).
		Where(dao.SysUserRole.Columns().UserId, userId).
		Delete()
	if err != nil {
		g.Log().Error(ctx, "删除用户角色关联失败:", err)
		return errors.New("删除用户角色关联失败")
	}

	// 2. 如果角色列表为空，直接返回
	if len(roleIds) == 0 {
		g.Log().Infof(ctx, "清空用户角色成功: userId=%s", userId)
		return nil
	}

	// 3. 批量插入新的用户角色关联
	userRoles := make([]do.SysUserRole, 0, len(roleIds))
	for _, roleId := range roleIds {
		userRoles = append(userRoles, do.SysUserRole{
			Id:     guid.S(),
			UserId: userId,
			RoleId: roleId,
		})
	}

	_, err = dao.SysUserRole.Ctx(ctx).Insert(userRoles)
	if err != nil {
		g.Log().Error(ctx, "新增用户角色关联失败:", err)
		return errors.New("新增用户角色关联失败")
	}

	g.Log().Infof(ctx, "保存用户角色成功: userId=%s, roleIds=%v", userId, roleIds)
	return nil
}
