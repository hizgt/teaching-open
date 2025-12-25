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
	"teaching-open/utility/password"
)

// SysUserService 用户服务接口
type SysUserService interface {
	GetUserByName(ctx context.Context, username string) (*entity.SysUser, error)
	GetUserById(ctx context.Context, id string) (*entity.SysUser, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.SysUser, error)
	List(ctx context.Context, req *vo.UserListReq) (*vo.UserListRes, error)
	Create(ctx context.Context, req *vo.UserCreateReq, createBy string) (*vo.UserCreateRes, error)
	Update(ctx context.Context, req *vo.UserUpdateReq, updateBy string) error
	Delete(ctx context.Context, id string, deleteBy string) error
	DeleteBatch(ctx context.Context, ids string, deleteBy string) error
	ResetPassword(ctx context.Context, req *vo.ResetPasswordReq) error
	GetUserRolesSet(ctx context.Context, username string) ([]string, error)
	GetUserPermissionsSet(ctx context.Context, username string) ([]string, error)
	CheckUserIsEffective(ctx context.Context, user *entity.SysUser) error
	UpdateUserDepart(ctx context.Context, username, orgCode string) error
}

// sysUserServiceImpl 用户服务实现
type sysUserServiceImpl struct{}

// NewSysUserService 创建用户服务实例
func NewSysUserService() SysUserService {
	return &sysUserServiceImpl{}
}

// GetUserByName 根据用户名获取用户
func (s *sysUserServiceImpl) GetUserByName(ctx context.Context, username string) (*entity.SysUser, error) {
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Username, username).
		Where(dao.SysUser.Columns().DelFlag, 0).
		Scan(&user)
	if err != nil {
		return nil, err
	}
	if user.Id == "" {
		return nil, nil
	}
	return &user, nil
}

// GetUserById 根据ID获取用户
func (s *sysUserServiceImpl) GetUserById(ctx context.Context, id string) (*entity.SysUser, error) {
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, id).
		Where(dao.SysUser.Columns().DelFlag, 0).
		Scan(&user)
	if err != nil {
		return nil, err
	}
	if user.Id == "" {
		return nil, nil
	}
	return &user, nil
}

// GetUserByPhone 根据手机号获取用户
func (s *sysUserServiceImpl) GetUserByPhone(ctx context.Context, phone string) (*entity.SysUser, error) {
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Phone, phone).
		Where(dao.SysUser.Columns().DelFlag, 0).
		Scan(&user)
	if err != nil {
		return nil, err
	}
	if user.Id == "" {
		return nil, nil
	}
	return &user, nil
}

// List 分页查询用户
func (s *sysUserServiceImpl) List(ctx context.Context, req *vo.UserListReq) (*vo.UserListRes, error) {
	model := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().DelFlag, 0)

	// 构建查询条件
	if req.Username != "" {
		model = model.WhereLike(dao.SysUser.Columns().Username, "%"+req.Username+"%")
	}
	if req.Realname != "" {
		model = model.WhereLike(dao.SysUser.Columns().Realname, "%"+req.Realname+"%")
	}
	if req.Phone != "" {
		model = model.WhereLike(dao.SysUser.Columns().Phone, "%"+req.Phone+"%")
	}
	if req.Status > 0 {
		model = model.Where(dao.SysUser.Columns().Status, req.Status)
	}

	// 部门过滤
	if req.DepartId != "" {
		model = model.WhereLike(dao.SysUser.Columns().DepartIds, "%"+req.DepartId+"%")
	}

	// 角色过滤
	if req.RoleId != "" {
		userIds, err := s.getUserIdsByRoleId(ctx, req.RoleId)
		if err != nil {
			return nil, err
		}
		if len(userIds) > 0 {
			model = model.WhereIn(dao.SysUser.Columns().Id, userIds)
		} else {
			// 没有匹配的用户
			return &vo.UserListRes{
				Records:  []vo.UserItem{},
				Total:    0,
				Page:     req.Page,
				PageSize: req.PageSize,
			}, nil
		}
	}

	// 分页查询
	total, err := model.Count()
	if err != nil {
		return nil, err
	}

	var users []entity.SysUser
	err = model.Page(req.Page, req.PageSize).
		OrderDesc(dao.SysUser.Columns().CreateTime).
		Scan(&users)
	if err != nil {
		return nil, err
	}

	// 转换为VO
	var records []vo.UserItem
	for _, user := range users {
		item := vo.UserItem{
			Id:         user.Id,
			Username:   user.Username,
			Realname:   user.Realname,
			Avatar:     user.Avatar,
			Sex:        user.Sex,
			Email:      user.Email,
			Phone:      user.Phone,
			Status:     user.Status,
			OrgCode:    user.OrgCode,
			WorkNo:     user.WorkNo,
			Post:       user.Post,
			CreateTime: user.CreateTime.String(),
		}

		// 获取角色名称
		roleNames, _ := s.getUserRoleNames(ctx, user.Id)
		item.RoleNames = roleNames

		// 获取部门名称
		departName, _ := s.getUserDepartName(ctx, user.DepartIds)
		item.DepartName = departName

		records = append(records, item)
	}

	return &vo.UserListRes{
		Records:  records,
		Total:    int64(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// Create 创建用户
func (s *sysUserServiceImpl) Create(ctx context.Context, req *vo.UserCreateReq, createBy string) (*vo.UserCreateRes, error) {
	// 检查用户名是否存在
	existUser, err := s.GetUserByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if existUser != nil {
		return nil, gerror.New("用户名已存在")
	}

	// 检查手机号是否存在
	if req.Phone != "" {
		existUser, err = s.GetUserByPhone(ctx, req.Phone)
		if err != nil {
			return nil, err
		}
		if existUser != nil {
			return nil, gerror.New("手机号已存在")
		}
	}

	// 生成盐值和加密密码
	salt := password.GenerateSalt()
	encryptedPwd := password.Encrypt(req.Username, req.Password, salt)

	userId := guid.S()
	user := entity.SysUser{
		Id:         userId,
		Username:   req.Username,
		Realname:   req.Realname,
		Password:   encryptedPwd,
		Salt:       salt,
		Email:      req.Email,
		Phone:      req.Phone,
		Sex:        req.Sex,
		Status:     req.Status,
		WorkNo:     req.WorkNo,
		Post:       req.Post,
		School:     req.School,
		DepartIds:  strings.Join(req.DepartIds, ","),
		DelFlag:    0,
		CreateBy:   createBy,
		CreateTime: gtime.Now(),
		UpdateBy:   createBy,
		UpdateTime: gtime.Now(),
	}

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 插入用户
		_, err := dao.SysUser.Ctx(ctx).Data(user).Insert()
		if err != nil {
			return err
		}

		// 插入用户角色关联
		if len(req.RoleIds) > 0 {
			err = s.saveUserRoles(ctx, userId, req.RoleIds)
			if err != nil {
				return err
			}
		}

		// 插入用户部门关联
		if len(req.DepartIds) > 0 {
			err = s.saveUserDeparts(ctx, userId, req.DepartIds)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &vo.UserCreateRes{Id: userId}, nil
}

// Update 更新用户
func (s *sysUserServiceImpl) Update(ctx context.Context, req *vo.UserUpdateReq, updateBy string) error {
	// 检查用户是否存在
	user, err := s.GetUserById(ctx, req.Id)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New("用户不存在")
	}

	// 检查手机号是否被其他用户使用
	if req.Phone != "" && req.Phone != user.Phone {
		existUser, err := s.GetUserByPhone(ctx, req.Phone)
		if err != nil {
			return err
		}
		if existUser != nil && existUser.Id != req.Id {
			return gerror.New("手机号已被使用")
		}
	}

	// 更新用户信息
	data := g.Map{
		dao.SysUser.Columns().Realname:   req.Realname,
		dao.SysUser.Columns().Email:      req.Email,
		dao.SysUser.Columns().Phone:      req.Phone,
		dao.SysUser.Columns().Sex:        req.Sex,
		dao.SysUser.Columns().Status:     req.Status,
		dao.SysUser.Columns().WorkNo:     req.WorkNo,
		dao.SysUser.Columns().Post:       req.Post,
		dao.SysUser.Columns().School:     req.School,
		dao.SysUser.Columns().DepartIds:  strings.Join(req.DepartIds, ","),
		dao.SysUser.Columns().UpdateBy:   updateBy,
		dao.SysUser.Columns().UpdateTime: gtime.Now(),
	}

	// 开启事务
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新用户
		_, err := dao.SysUser.Ctx(ctx).Data(data).Where(dao.SysUser.Columns().Id, req.Id).Update()
		if err != nil {
			return err
		}

		// 更新用户角色关联
		if len(req.RoleIds) > 0 {
			// 删除原有关联
			_, err = dao.SysUserRole.Ctx(ctx).Where(dao.SysUserRole.Columns().UserId, req.Id).Delete()
			if err != nil {
				return err
			}
			// 添加新关联
			err = s.saveUserRoles(ctx, req.Id, req.RoleIds)
			if err != nil {
				return err
			}
		}

		// 更新用户部门关联
		if len(req.DepartIds) > 0 {
			// 删除原有关联
			_, err = dao.SysUserDepart.Ctx(ctx).Where(dao.SysUserDepart.Columns().UserId, req.Id).Delete()
			if err != nil {
				return err
			}
			// 添加新关联
			err = s.saveUserDeparts(ctx, req.Id, req.DepartIds)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// Delete 删除用户 (逻辑删除)
func (s *sysUserServiceImpl) Delete(ctx context.Context, id string, deleteBy string) error {
	_, err := dao.SysUser.Ctx(ctx).Data(g.Map{
		dao.SysUser.Columns().DelFlag:    1,
		dao.SysUser.Columns().UpdateBy:   deleteBy,
		dao.SysUser.Columns().UpdateTime: gtime.Now(),
	}).Where(dao.SysUser.Columns().Id, id).Update()
	return err
}

// DeleteBatch 批量删除用户
func (s *sysUserServiceImpl) DeleteBatch(ctx context.Context, ids string, deleteBy string) error {
	idList := strings.Split(ids, ",")
	_, err := dao.SysUser.Ctx(ctx).Data(g.Map{
		dao.SysUser.Columns().DelFlag:    1,
		dao.SysUser.Columns().UpdateBy:   deleteBy,
		dao.SysUser.Columns().UpdateTime: gtime.Now(),
	}).WhereIn(dao.SysUser.Columns().Id, idList).Update()
	return err
}

// ResetPassword 重置密码
func (s *sysUserServiceImpl) ResetPassword(ctx context.Context, req *vo.ResetPasswordReq) error {
	user, err := s.GetUserByName(ctx, req.Username)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New("用户不存在")
	}

	// 验证旧密码
	if !password.Verify(user.Username, req.OldPwd, user.Salt, user.Password) {
		return gerror.New("原密码错误")
	}

	// 生成新盐值和加密新密码
	newSalt := password.GenerateSalt()
	newEncryptedPwd := password.Encrypt(user.Username, req.NewPwd, newSalt)

	_, err = dao.SysUser.Ctx(ctx).Data(g.Map{
		dao.SysUser.Columns().Password:   newEncryptedPwd,
		dao.SysUser.Columns().Salt:       newSalt,
		dao.SysUser.Columns().UpdateTime: gtime.Now(),
	}).Where(dao.SysUser.Columns().Id, user.Id).Update()

	return err
}

// GetUserRolesSet 获取用户角色编码集合
func (s *sysUserServiceImpl) GetUserRolesSet(ctx context.Context, username string) ([]string, error) {
	user, err := s.GetUserByName(ctx, username)
	if err != nil || user == nil {
		return nil, err
	}

	var roleCodes []string
	err = dao.SysUserRole.Ctx(ctx).
		Fields("r.role_code").
		As("ur").
		LeftJoin("sys_role r", "ur.role_id = r.id").
		Where("ur.user_id", user.Id).
		Scan(&roleCodes)

	return roleCodes, err
}

// GetUserPermissionsSet 获取用户权限标识集合
func (s *sysUserServiceImpl) GetUserPermissionsSet(ctx context.Context, username string) ([]string, error) {
	user, err := s.GetUserByName(ctx, username)
	if err != nil || user == nil {
		return nil, err
	}

	var perms []string
	err = dao.SysUserRole.Ctx(ctx).
		Fields("DISTINCT p.perms").
		As("ur").
		LeftJoin("sys_role_permission rp", "ur.role_id = rp.role_id").
		LeftJoin("sys_permission p", "rp.permission_id = p.id").
		Where("ur.user_id", user.Id).
		Where("p.perms IS NOT NULL").
		Where("p.perms != ''").
		Scan(&perms)

	return perms, err
}

// CheckUserIsEffective 检查用户有效性
func (s *sysUserServiceImpl) CheckUserIsEffective(ctx context.Context, user *entity.SysUser) error {
	if user == nil {
		return gerror.New("用户不存在")
	}
	if user.DelFlag == 1 {
		return gerror.New("用户已被删除")
	}
	if user.Status == 2 {
		return gerror.New("用户已被冻结")
	}
	return nil
}

// UpdateUserDepart 更新用户部门
func (s *sysUserServiceImpl) UpdateUserDepart(ctx context.Context, username, orgCode string) error {
	_, err := dao.SysUser.Ctx(ctx).Data(g.Map{
		dao.SysUser.Columns().OrgCode:    orgCode,
		dao.SysUser.Columns().UpdateTime: gtime.Now(),
	}).Where(dao.SysUser.Columns().Username, username).Update()
	return err
}

// 辅助方法

func (s *sysUserServiceImpl) getUserIdsByRoleId(ctx context.Context, roleId string) ([]string, error) {
	var userIds []string
	err := dao.SysUserRole.Ctx(ctx).
		Fields(dao.SysUserRole.Columns().UserId).
		Where(dao.SysUserRole.Columns().RoleId, roleId).
		Scan(&userIds)
	return userIds, err
}

func (s *sysUserServiceImpl) getUserRoleNames(ctx context.Context, userId string) ([]string, error) {
	var roleNames []string
	err := dao.SysUserRole.Ctx(ctx).
		Fields("r.role_name").
		As("ur").
		LeftJoin("sys_role r", "ur.role_id = r.id").
		Where("ur.user_id", userId).
		Scan(&roleNames)
	return roleNames, err
}

func (s *sysUserServiceImpl) getUserDepartName(ctx context.Context, departIds string) (string, error) {
	if departIds == "" {
		return "", nil
	}
	ids := strings.Split(departIds, ",")
	if len(ids) == 0 {
		return "", nil
	}

	var departName string
	err := dao.SysDepart.Ctx(ctx).
		Fields(dao.SysDepart.Columns().DepartName).
		Where(dao.SysDepart.Columns().Id, ids[0]).
		Scan(&departName)
	return departName, err
}

func (s *sysUserServiceImpl) saveUserRoles(ctx context.Context, userId string, roleIds []string) error {
	var relations []g.Map
	for _, roleId := range roleIds {
		relations = append(relations, g.Map{
			dao.SysUserRole.Columns().Id:     guid.S(),
			dao.SysUserRole.Columns().UserId: userId,
			dao.SysUserRole.Columns().RoleId: roleId,
		})
	}
	if len(relations) > 0 {
		_, err := dao.SysUserRole.Ctx(ctx).Data(relations).Insert()
		return err
	}
	return nil
}

func (s *sysUserServiceImpl) saveUserDeparts(ctx context.Context, userId string, departIds []string) error {
	var relations []g.Map
	for _, departId := range departIds {
		relations = append(relations, g.Map{
			dao.SysUserDepart.Columns().Id:     guid.S(),
			dao.SysUserDepart.Columns().UserId: userId,
			dao.SysUserDepart.Columns().DepId:  departId,
		})
	}
	if len(relations) > 0 {
		_, err := dao.SysUserDepart.Ctx(ctx).Data(relations).Insert()
		return err
	}
	return nil
}
