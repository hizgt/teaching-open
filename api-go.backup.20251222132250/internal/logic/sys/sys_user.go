package sys

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/consts"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/do"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
	"teaching-open/utility/jwt"
)

type sSysUser struct{}

func init() {
	service.SysUser = &sSysUser{}
}

// Login 用户登录
func (s *sSysUser) Login(ctx context.Context, username, password string) (token string, userInfo *v1.UserInfo, err error) {
	// 1. 根据用户名查询用户
	var user entity.SysUser
	err = dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Username, username).
		Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "查询用户失败:", err)
		return "", nil, errors.New(consts.GetErrorMessage(consts.CodeDatabaseError))
	}

	// 2. 检查用户是否存在
	if user.Id == "" {
		return "", nil, errors.New(consts.GetErrorMessage(consts.CodeUserNotFound))
	}

	// 3. 检查用户删除标记
	if user.DelFlag == consts.DelFlagDeleted {
		return "", nil, errors.New(consts.GetErrorMessage(consts.CodeUserNotFound))
	}

	// 4. 检查用户状态
	if user.Status != consts.UserStatusNormal {
		return "", nil, errors.New(consts.GetErrorMessage(consts.CodeUserFrozen))
	}

	// 5. 验证密码 - 使用MD5(password + salt)
	passwordMd5 := s.encryptPassword(password, user.Salt)
	if passwordMd5 != user.Password {
		return "", nil, errors.New(consts.GetErrorMessage(consts.CodePasswordError))
	}

	// 6. 生成JWT token
	token, err = jwt.GenerateToken(user.Id, user.Username, user.Realname)
	if err != nil {
		g.Log().Error(ctx, "生成token失败:", err)
		return "", nil, errors.New("生成token失败")
	}

	// 7. 构建用户信息(不返回密码和盐)
	userInfo = &v1.UserInfo{
		Id:       user.Id,
		Username: user.Username,
		Realname: user.Realname,
		Avatar:   user.Avatar,
		Status:   user.Status,
		School:   user.School,
		Phone:    user.Phone,
		Email:    user.Email,
	}

	g.Log().Infof(ctx, "用户登录成功: username=%s, userId=%s", username, user.Id)

	return token, userInfo, nil
}

// encryptPassword 加密密码 - MD5(password + salt)
func (s *sSysUser) encryptPassword(password, salt string) string {
	h := md5.New()
	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}

// GetList 获取用户列表
func (s *sSysUser) GetList(ctx context.Context, req *v1.UserListReq) (list []*v1.UserInfo, total int64, err error) {
	// 构建查询
	m := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal)

	// 用户名模糊查询
	if req.Username != "" {
		m = m.WhereLike(dao.SysUser.Columns().Username, "%"+req.Username+"%")
	}

	// 真实姓名模糊查询
	if req.Realname != "" {
		m = m.WhereLike(dao.SysUser.Columns().Realname, "%"+req.Realname+"%")
	}

	// 状态筛选
	if req.Status > 0 {
		m = m.Where(dao.SysUser.Columns().Status, req.Status)
	}

	// 分页查询
	var users []*entity.SysUser
	var totalInt int
	err = m.Page(req.Page, req.PageSize).
		OrderDesc(dao.SysUser.Columns().CreateTime).
		ScanAndCount(&users, &totalInt, false)
	if err != nil {
		g.Log().Error(ctx, "查询用户列表失败:", err)
		return nil, 0, errors.New("查询用户列表失败")
	}
	total = int64(totalInt)

	// 转换为UserInfo
	list = make([]*v1.UserInfo, 0, len(users))
	for _, user := range users {
		list = append(list, &v1.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Realname: user.Realname,
			Avatar:   user.Avatar,
			Status:   user.Status,
			School:   user.School,
			Phone:    user.Phone,
			Email:    user.Email,
			Sex:      user.Sex,
		})
	}

	return list, total, nil
}

// Add 新增用户
func (s *sSysUser) Add(ctx context.Context, req *v1.UserAddReq) error {
	// 1. 检查用户名是否已存在
	count, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Username, req.Username).
		Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
		Count()
	if err != nil {
		g.Log().Error(ctx, "检查用户名失败:", err)
		return errors.New("检查用户名失败")
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}

	// 2. 检查手机号是否已存在
	if req.Phone != "" {
		count, err = dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Phone, req.Phone).
			Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查手机号失败:", err)
			return errors.New("检查手机号失败")
		}
		if count > 0 {
			return errors.New("手机号已存在")
		}
	}

	// 3. 检查邮箱是否已存在
	if req.Email != "" {
		count, err = dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Email, req.Email).
			Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查邮箱失败:", err)
			return errors.New("检查邮箱失败")
		}
		if count > 0 {
			return errors.New("邮箱已存在")
		}
	}

	// 4. 生成salt和加密密码
	salt := grand.S(8)
	encryptedPassword := s.encryptPassword(req.Password, salt)

	// 5. 插入用户数据
	_, err = dao.SysUser.Ctx(ctx).Insert(do.SysUser{
		Id:         guid.S(),
		Username:   req.Username,
		Realname:   req.Realname,
		Password:   encryptedPassword,
		Salt:       salt,
		Phone:      req.Phone,
		Email:      req.Email,
		School:     req.School,
		Sex:        req.Sex,
		Status:     consts.UserStatusNormal,
		DelFlag:    consts.DelFlagNormal,
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	})
	if err != nil {
		g.Log().Error(ctx, "新增用户失败:", err)
		return errors.New("新增用户失败")
	}

	g.Log().Infof(ctx, "新增用户成功: username=%s", req.Username)
	return nil
}

// Edit 编辑用户
func (s *sSysUser) Edit(ctx context.Context, req *v1.UserEditReq) error {
	// 1. 检查用户是否存在
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, req.Id).
		Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
		Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "查询用户失败:", err)
		return errors.New("查询用户失败")
	}
	if user.Id == "" {
		return errors.New("用户不存在")
	}

	// 2. 如果修改了用户名，检查是否重复
	if req.Username != "" && req.Username != user.Username {
		count, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Username, req.Username).
			Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
			WhereNot(dao.SysUser.Columns().Id, req.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查用户名失败:", err)
			return errors.New("检查用户名失败")
		}
		if count > 0 {
			return errors.New("用户名已存在")
		}
	}

	// 3. 如果修改了手机号，检查是否重复
	if req.Phone != "" && req.Phone != user.Phone {
		count, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Phone, req.Phone).
			Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
			WhereNot(dao.SysUser.Columns().Id, req.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查手机号失败:", err)
			return errors.New("检查手机号失败")
		}
		if count > 0 {
			return errors.New("手机号已存在")
		}
	}

	// 4. 如果修改了邮箱，检查是否重复
	if req.Email != "" && req.Email != user.Email {
		count, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Email, req.Email).
			Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
			WhereNot(dao.SysUser.Columns().Id, req.Id).
			Count()
		if err != nil {
			g.Log().Error(ctx, "检查邮箱失败:", err)
			return errors.New("检查邮箱失败")
		}
		if count > 0 {
			return errors.New("邮箱已存在")
		}
	}

	// 5. 更新用户数据
	updateData := do.SysUser{
		UpdateTime: gtime.Now(),
	}
	if req.Username != "" {
		updateData.Username = req.Username
	}
	if req.Realname != "" {
		updateData.Realname = req.Realname
	}
	if req.Phone != "" {
		updateData.Phone = req.Phone
	}
	if req.Email != "" {
		updateData.Email = req.Email
	}
	if req.School != "" {
		updateData.School = req.School
	}
	if req.Sex >= 0 {
		updateData.Sex = req.Sex
	}
	if req.Status > 0 {
		updateData.Status = req.Status
	}

	_, err = dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, req.Id).
		Update(updateData)
	if err != nil {
		g.Log().Error(ctx, "更新用户失败:", err)
		return errors.New("更新用户失败")
	}

	g.Log().Infof(ctx, "编辑用户成功: id=%s", req.Id)
	return nil
}

// Delete 删除用户(逻辑删除)
func (s *sSysUser) Delete(ctx context.Context, id string) error {
	// 1. 检查用户是否存在
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, id).
		Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
		Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "查询用户失败:", err)
		return errors.New("查询用户失败")
	}
	if user.Id == "" {
		return errors.New("用户不存在")
	}

	// 2. 逻辑删除
	_, err = dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, id).
		Update(do.SysUser{
			DelFlag:    consts.DelFlagDeleted,
			UpdateTime: gtime.Now(),
		})
	if err != nil {
		g.Log().Error(ctx, "删除用户失败:", err)
		return errors.New("删除用户失败")
	}

	g.Log().Infof(ctx, "删除用户成功: id=%s", id)
	return nil
}

// GetById 根据ID获取用户详情
func (s *sSysUser) GetById(ctx context.Context, id string) (*v1.UserInfo, error) {
	// 查询用户
	var user entity.SysUser
	err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, id).
		Where(dao.SysUser.Columns().DelFlag, consts.DelFlagNormal).
		Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "查询用户失败:", err)
		return nil, errors.New("查询用户失败")
	}
	if user.Id == "" {
		return nil, errors.New("用户不存在")
	}

	// 返回用户信息(不包含密码和salt)
	return &v1.UserInfo{
		Id:       user.Id,
		Username: user.Username,
		Realname: user.Realname,
		Avatar:   user.Avatar,
		Status:   user.Status,
		School:   user.School,
		Phone:    user.Phone,
		Email:    user.Email,
		Sex:      user.Sex,
	}, nil
}
