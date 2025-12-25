package system

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
)

// SysUserController 用户控制器
type SysUserController struct {
	userService service.SysUserService
}

// NewSysUserController 创建用户控制器实例
func NewSysUserController() *SysUserController {
	return &SysUserController{
		userService: service.NewSysUserService(),
	}
}

// List 用户列表
// @Summary 用户列表
// @Tags 系统-用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param username query string false "用户名"
// @Param realname query string false "真实姓名"
// @Param phone query string false "手机号"
// @Param status query int false "状态"
// @Param departId query string false "部门ID"
// @Param roleId query string false "角色ID"
// @Success 200 {object} vo.UserListRes
// @Router /sys/user/list [get]
func (c *SysUserController) List(r *ghttp.Request) {
	var req vo.UserListReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	res, err := c.userService.List(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// Add 添加用户
// @Summary 添加用户
// @Tags 系统-用户管理
// @Accept json
// @Produce json
// @Param data body vo.UserCreateReq true "用户信息"
// @Success 200 {object} vo.UserCreateRes
// @Router /sys/user/add [post]
func (c *SysUserController) Add(r *ghttp.Request) {
	var req vo.UserCreateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.GetCtxVar("userId").String()
	res, err := c.userService.Create(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r, res)
}

// Edit 编辑用户
// @Summary 编辑用户
// @Tags 系统-用户管理
// @Accept json
// @Produce json
// @Param data body vo.UserUpdateReq true "用户信息"
// @Success 200 {object} response.Response
// @Router /sys/user/edit [put]
func (c *SysUserController) Edit(r *ghttp.Request) {
	var req vo.UserUpdateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	userId := r.GetCtxVar("userId").String()
	err := c.userService.Update(r.Context(), &req, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r)
}

// Delete 删除用户
// @Summary 删除用户
// @Tags 系统-用户管理
// @Accept json
// @Produce json
// @Param id query string true "用户ID"
// @Success 200 {object} response.Response
// @Router /sys/user/delete [delete]
func (c *SysUserController) Delete(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "用户ID不能为空")
		return
	}

	userId := r.GetCtxVar("userId").String()
	err := c.userService.Delete(r.Context(), id, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r)
}

// DeleteBatch 批量删除用户
// @Summary 批量删除用户
// @Tags 系统-用户管理
// @Accept json
// @Produce json
// @Param ids query string true "用户ID列表(逗号分隔)"
// @Success 200 {object} response.Response
// @Router /sys/user/deleteBatch [delete]
func (c *SysUserController) DeleteBatch(r *ghttp.Request) {
	ids := r.Get("ids").String()
	if ids == "" {
		response.Error(r, "用户ID列表不能为空")
		return
	}

	userId := r.GetCtxVar("userId").String()
	err := c.userService.DeleteBatch(r.Context(), ids, userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r)
}

// ResetPassword 重置密码
// @Summary 重置密码
// @Tags 系统-用户管理
// @Accept json
// @Produce json
// @Param data body vo.ResetPasswordReq true "重置密码参数"
// @Success 200 {object} response.Response
// @Router /sys/user/resetPassword [put]
func (c *SysUserController) ResetPassword(r *ghttp.Request) {
	var req vo.ResetPasswordReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	err := c.userService.ResetPassword(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.Success(r)
}

// QueryById 根据ID查询用户
// @Summary 根据ID查询用户
// @Tags 系统-用户管理
// @Accept json
// @Produce json
// @Param id query string true "用户ID"
// @Success 200 {object} entity.SysUser
// @Router /sys/user/queryById [get]
func (c *SysUserController) QueryById(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "用户ID不能为空")
		return
	}

	user, err := c.userService.GetUserById(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	if user == nil {
		response.Error(r, "用户不存在")
		return
	}

	response.Success(r, user)
}
