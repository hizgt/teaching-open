package system

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
)

// SysDepartController 部门控制器
type SysDepartController struct {
	departService service.SysDepartService
}

// NewSysDepartController 创建部门控制器
func NewSysDepartController() *SysDepartController {
	return &SysDepartController{
		departService: service.NewSysDepartService(),
	}
}

// QueryTreeList 查询部门树列表
func (c *SysDepartController) QueryTreeList(r *ghttp.Request) {
	var req vo.DepartTreeReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	tree, err := c.departService.GetTree(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, tree)
}

// QueryById 根据ID查询部门
func (c *SysDepartController) QueryById(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "部门ID不能为空")
		return
	}

	depart, err := c.departService.GetById(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	if depart == nil {
		response.Error(r, "部门不存在")
		return
	}
	response.Success(r, depart)
}

// Add 添加部门
func (c *SysDepartController) Add(r *ghttp.Request) {
	var req vo.DepartCreateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	result, err := c.departService.Create(r.Context(), &req, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, result)
}

// Edit 编辑部门
func (c *SysDepartController) Edit(r *ghttp.Request) {
	var req vo.DepartUpdateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	err := c.departService.Update(r.Context(), &req, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "修改成功")
}

// Delete 删除部门
func (c *SysDepartController) Delete(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "部门ID不能为空")
		return
	}

	err := c.departService.Delete(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "删除成功")
}

// QueryIdTree 查询部门ID树
func (c *SysDepartController) QueryIdTree(r *ghttp.Request) {
	tree, err := c.departService.GetIdTree(r.Context())
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, tree)
}

// SearchBy 搜索部门
func (c *SysDepartController) SearchBy(r *ghttp.Request) {
	var req vo.DepartSearchReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	departs, err := c.departService.Search(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, departs)
}

// QueryDepartTreeSync 查询部门用户树
func (c *SysDepartController) QueryDepartTreeSync(r *ghttp.Request) {
	tree, err := c.departService.GetDepartUserTree(r.Context())
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, tree)
}

// QueryUserDepart 查询用户部门
func (c *SysDepartController) QueryUserDepart(r *ghttp.Request) {
	userId := r.Get("userId").String()
	if userId == "" {
		response.Error(r, "用户ID不能为空")
		return
	}

	depIds, err := c.departService.GetUserDeparts(r.Context(), userId)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, depIds)
}

// SaveUserDepart 保存用户部门
func (c *SysDepartController) SaveUserDepart(r *ghttp.Request) {
	var req vo.UserDepartReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	err := c.departService.SaveUserDeparts(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "保存成功")
}
