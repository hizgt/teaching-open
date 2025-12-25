package system

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
)

// SysPermissionController 权限控制器
type SysPermissionController struct {
	permService service.SysPermissionService
}

// NewSysPermissionController 创建权限控制器
func NewSysPermissionController() *SysPermissionController {
	return &SysPermissionController{
		permService: service.NewSysPermissionService(),
	}
}

// List 权限树列表
func (c *SysPermissionController) List(r *ghttp.Request) {
	tree, err := c.permService.GetTree(r.Context())
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, tree)
}

// QueryById 根据ID查询权限
func (c *SysPermissionController) QueryById(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "权限ID不能为空")
		return
	}

	perm, err := c.permService.GetById(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	if perm == nil {
		response.Error(r, "权限不存在")
		return
	}
	response.Success(r, perm)
}

// Add 添加权限
func (c *SysPermissionController) Add(r *ghttp.Request) {
	var req vo.PermissionCreateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	result, err := c.permService.Create(r.Context(), &req, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, result)
}

// Edit 编辑权限
func (c *SysPermissionController) Edit(r *ghttp.Request) {
	var req vo.PermissionUpdateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	err := c.permService.Update(r.Context(), &req, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "修改成功")
}

// Delete 删除权限
func (c *SysPermissionController) Delete(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "权限ID不能为空")
		return
	}

	err := c.permService.Delete(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "删除成功")
}

// GetUserPermission 获取用户权限
func (c *SysPermissionController) GetUserPermission(r *ghttp.Request) {
	userId := r.Get("userId").String()
	if userId == "" {
		response.Error(r, "用户ID不能为空")
		return
	}

	result, err := c.permService.GetUserPermission(r.Context(), userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, result)
}

// GetUserMenus 获取用户菜单
func (c *SysPermissionController) GetUserMenus(r *ghttp.Request) {
	userId := r.Get("userId").String()
	if userId == "" {
		response.Error(r, "用户ID不能为空")
		return
	}

	menus, err := c.permService.GetUserMenus(r.Context(), userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, menus)
}

// GetUserPermCodes 获取用户权限编码
func (c *SysPermissionController) GetUserPermCodes(r *ghttp.Request) {
	userId := r.Get("userId").String()
	if userId == "" {
		response.Error(r, "用户ID不能为空")
		return
	}

	codes, err := c.permService.GetUserPermCodes(r.Context(), userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, codes)
}
