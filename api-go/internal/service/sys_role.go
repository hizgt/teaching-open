package service

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/model/vo"
)

// SysRoleService 角色服务接口
type SysRoleService interface {
	List(ctx context.Context, req *vo.RoleListReq) (*vo.RoleListRes, error)
	GetById(ctx context.Context, id string) (*entity.SysRole, error)
	GetAll(ctx context.Context) ([]entity.SysRole, error)
	Create(ctx context.Context, req *vo.RoleCreateReq, createBy string) (*vo.RoleCreateRes, error)
	Update(ctx context.Context, req *vo.RoleUpdateReq, updateBy string) error
	Delete(ctx context.Context, ids string) error
	GetUserRoles(ctx context.Context, userId string) ([]string, error)
	SaveUserRoles(ctx context.Context, req *vo.UserRolesReq) error
	GetRolePermissionIds(ctx context.Context, roleId string) ([]string, error)
	SaveRolePermission(ctx context.Context, req *vo.RolePermissionReq) error
}

// sysRoleServiceImpl 角色服务实现
type sysRoleServiceImpl struct{}

// NewSysRoleService 创建角色服务实例
func NewSysRoleService() SysRoleService {
	return &sysRoleServiceImpl{}
}

// List 分页查询角色列表
func (s *sysRoleServiceImpl) List(ctx context.Context, req *vo.RoleListReq) (*vo.RoleListRes, error) {
	model := dao.SysRole.Ctx(ctx)

	if req.RoleName != "" {
		model = model.WhereLike(dao.SysRole.Columns().RoleName, "%"+req.RoleName+"%")
	}
	if req.RoleCode != "" {
		model = model.WhereLike(dao.SysRole.Columns().RoleCode, "%"+req.RoleCode+"%")
	}

	total, err := model.Count()
	if err != nil {
		return nil, err
	}

	var roles []entity.SysRole
	err = model.Page(req.Page, req.PageSize).
		OrderAsc(dao.SysRole.Columns().RoleLevel).
		OrderDesc(dao.SysRole.Columns().CreateTime).
		Scan(&roles)
	if err != nil {
		return nil, err
	}

	var records []vo.RoleItem
	for _, role := range roles {
		records = append(records, vo.RoleItem{
			Id:          role.Id,
			RoleName:    role.RoleName,
			RoleCode:    role.RoleCode,
			RoleLevel:   role.RoleLevel,
			Description: role.Description,
			CreateTime:  role.CreateTime.String(),
		})
	}

	return &vo.RoleListRes{
		Records:  records,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// GetById 根据ID获取角色
func (s *sysRoleServiceImpl) GetById(ctx context.Context, id string) (*entity.SysRole, error) {
	var role entity.SysRole
	err := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Id, id).Scan(&role)
	if err != nil {
		return nil, err
	}
	if role.Id == "" {
		return nil, nil
	}
	return &role, nil
}

// GetAll 获取所有角色
func (s *sysRoleServiceImpl) GetAll(ctx context.Context) ([]entity.SysRole, error) {
	var roles []entity.SysRole
	err := dao.SysRole.Ctx(ctx).
		OrderAsc(dao.SysRole.Columns().RoleLevel).
		Scan(&roles)
	return roles, err
}

// Create 创建角色
func (s *sysRoleServiceImpl) Create(ctx context.Context, req *vo.RoleCreateReq, createBy string) (*vo.RoleCreateRes, error) {
	// 检查角色编码是否存在
	count, err := dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleCode, req.RoleCode).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("角色编码已存在")
	}

	// 检查角色名称是否存在
	count, err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().RoleName, req.RoleName).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("角色名称已存在")
	}

	roleId := guid.S()
	role := entity.SysRole{
		Id:          roleId,
		RoleName:    req.RoleName,
		RoleCode:    req.RoleCode,
		RoleLevel:   req.RoleLevel,
		Description: req.Description,
		CreateBy:    createBy,
		CreateTime:  gtime.Now(),
		UpdateBy:    createBy,
		UpdateTime:  gtime.Now(),
	}

	_, err = dao.SysRole.Ctx(ctx).Data(role).Insert()
	if err != nil {
		return nil, err
	}

	return &vo.RoleCreateRes{Id: roleId}, nil
}

// Update 更新角色
func (s *sysRoleServiceImpl) Update(ctx context.Context, req *vo.RoleUpdateReq, updateBy string) error {
	// 检查角色是否存在
	role, err := s.GetById(ctx, req.Id)
	if err != nil {
		return err
	}
	if role == nil {
		return gerror.New("角色不存在")
	}

	// 检查角色编码是否被其他角色使用
	if req.RoleCode != role.RoleCode {
		count, err := dao.SysRole.Ctx(ctx).
			Where(dao.SysRole.Columns().RoleCode, req.RoleCode).
			WhereNot(dao.SysRole.Columns().Id, req.Id).
			Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New("角色编码已存在")
		}
	}

	// 检查角色名称是否被其他角色使用
	if req.RoleName != role.RoleName {
		count, err := dao.SysRole.Ctx(ctx).
			Where(dao.SysRole.Columns().RoleName, req.RoleName).
			WhereNot(dao.SysRole.Columns().Id, req.Id).
			Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New("角色名称已存在")
		}
	}

	_, err = dao.SysRole.Ctx(ctx).Data(g.Map{
		dao.SysRole.Columns().RoleName:    req.RoleName,
		dao.SysRole.Columns().RoleCode:    req.RoleCode,
		dao.SysRole.Columns().RoleLevel:   req.RoleLevel,
		dao.SysRole.Columns().Description: req.Description,
		dao.SysRole.Columns().UpdateBy:    updateBy,
		dao.SysRole.Columns().UpdateTime:  gtime.Now(),
	}).Where(dao.SysRole.Columns().Id, req.Id).Update()

	return err
}

// Delete 删除角色
func (s *sysRoleServiceImpl) Delete(ctx context.Context, ids string) error {
	idList := strings.Split(ids, ",")

	// 检查角色是否被用户使用
	for _, id := range idList {
		count, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().RoleId, id).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New("角色正在被使用，无法删除")
		}
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除角色
		_, err := dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().Id, idList).Delete()
		if err != nil {
			return err
		}

		// 删除角色权限关联
		_, err = dao.SysRolePermission.Ctx(ctx).WhereIn(dao.SysRolePermission.Columns().RoleId, idList).Delete()
		return err
	})
}

// GetUserRoles 获取用户角色ID列表
func (s *sysRoleServiceImpl) GetUserRoles(ctx context.Context, userId string) ([]string, error) {
	var roleIds []string
	err := dao.SysUserRole.Ctx(ctx).
		Fields(dao.SysUserRole.Columns().RoleId).
		Where(dao.SysUserRole.Columns().UserId, userId).
		Scan(&roleIds)
	return roleIds, err
}

// SaveUserRoles 保存用户角色
func (s *sysRoleServiceImpl) SaveUserRoles(ctx context.Context, req *vo.UserRolesReq) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除原有关联
		_, err := dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, req.UserId).Delete()
		if err != nil {
			return err
		}

		// 添加新关联
		if len(req.RoleIds) > 0 {
			var relations []g.Map
			for _, roleId := range req.RoleIds {
				relations = append(relations, g.Map{
					dao.SysUserRole.Columns().Id:     guid.S(),
					dao.SysUserRole.Columns().UserId: req.UserId,
					dao.SysUserRole.Columns().RoleId: roleId,
				})
			}
			_, err = dao.SysUserRole.Ctx(ctx).Data(relations).Insert()
		}
		return err
	})
}

// GetRolePermissionIds 获取角色权限ID列表
func (s *sysRoleServiceImpl) GetRolePermissionIds(ctx context.Context, roleId string) ([]string, error) {
	var permissionIds []string
	err := dao.SysRolePermission.Ctx(ctx).
		Fields(dao.SysRolePermission.Columns().PermissionId).
		Where(dao.SysRolePermission.Columns().RoleId, roleId).
		Scan(&permissionIds)
	return permissionIds, err
}

// SaveRolePermission 保存角色权限
func (s *sysRoleServiceImpl) SaveRolePermission(ctx context.Context, req *vo.RolePermissionReq) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除原有关联
		_, err := dao.SysRolePermission.Ctx(ctx).Where(dao.SysRolePermission.Columns().RoleId, req.RoleId).Delete()
		if err != nil {
			return err
		}

		// 添加新关联
		if len(req.PermissionIds) > 0 {
			var relations []g.Map
			for _, permId := range req.PermissionIds {
				relations = append(relations, g.Map{
					dao.SysRolePermission.Columns().Id:           guid.S(),
					dao.SysRolePermission.Columns().RoleId:       req.RoleId,
					dao.SysRolePermission.Columns().PermissionId: permId,
					dao.SysRolePermission.Columns().OperateDate:  gtime.Now(),
				})
			}
			_, err = dao.SysRolePermission.Ctx(ctx).Data(relations).Insert()
		}
		return err
	})
}
