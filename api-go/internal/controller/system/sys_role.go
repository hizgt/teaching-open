package system

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
)

// SysRoleController 角色控制器
type SysRoleController struct {
	roleService service.SysRoleService
}

// NewSysRoleController 创建角色控制器
func NewSysRoleController() *SysRoleController {
	return &SysRoleController{
		roleService: service.NewSysRoleService(),
	}
}

// List 角色列表
func (c *SysRoleController) List(r *ghttp.Request) {
	var req vo.RoleListReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	result, err := c.roleService.List(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.PageSuccess(r, result.Records, result.Total, result.Page, result.PageSize)
}

// QueryAll 查询所有角色
func (c *SysRoleController) QueryAll(r *ghttp.Request) {
	roles, err := c.roleService.GetAll(r.Context())
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, roles)
}

// QueryById 根据ID查询角色
func (c *SysRoleController) QueryById(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "角色ID不能为空")
		return
	}

	role, err := c.roleService.GetById(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	if role == nil {
		response.Error(r, "角色不存在")
		return
	}
	response.Success(r, role)
}

// Add 添加角色
func (c *SysRoleController) Add(r *ghttp.Request) {
	var req vo.RoleCreateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	// 获取当前用户
	username := r.Get("username").String()

	result, err := c.roleService.Create(r.Context(), &req, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, result)
}

// Edit 编辑角色
func (c *SysRoleController) Edit(r *ghttp.Request) {
	var req vo.RoleUpdateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	err := c.roleService.Update(r.Context(), &req, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "修改成功")
}

// Delete 删除角色
func (c *SysRoleController) Delete(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "角色ID不能为空")
		return
	}

	err := c.roleService.Delete(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "删除成功")
}

// DeleteBatch 批量删除角色
func (c *SysRoleController) DeleteBatch(r *ghttp.Request) {
	ids := r.Get("ids").String()
	if ids == "" {
		response.Error(r, "角色ID不能为空")
		return
	}

	err := c.roleService.Delete(r.Context(), ids)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "删除成功")
}

// QueryUserRole 查询用户角色
func (c *SysRoleController) QueryUserRole(r *ghttp.Request) {
	userId := r.Get("userId").String()
	if userId == "" {
		response.Error(r, "用户ID不能为空")
		return
	}

	roleIds, err := c.roleService.GetUserRoles(r.Context(), userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, roleIds)
}

// SaveUserRole 保存用户角色
func (c *SysRoleController) SaveUserRole(r *ghttp.Request) {
	var req vo.UserRolesReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	err := c.roleService.SaveUserRoles(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "保存成功")
}

// QueryRolePermission 查询角色权限
func (c *SysRoleController) QueryRolePermission(r *ghttp.Request) {
	roleId := r.Get("roleId").String()
	if roleId == "" {
		response.Error(r, "角色ID不能为空")
		return
	}

	permIds, err := c.roleService.GetRolePermissionIds(r.Context(), roleId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, permIds)
}

// SaveRolePermission 保存角色权限
func (c *SysRoleController) SaveRolePermission(r *ghttp.Request) {
	var req vo.RolePermissionReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	err := c.roleService.SaveRolePermission(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "保存成功")
}
