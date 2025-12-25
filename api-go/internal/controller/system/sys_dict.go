package system

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"teaching-open/internal/model/entity"
	"teaching-open/internal/model/vo"
	"teaching-open/internal/service"
	"teaching-open/utility/response"
)

// SysDictController 字典控制器
type SysDictController struct {
	dictService service.SysDictService
}

// NewSysDictController 创建字典控制器
func NewSysDictController() *SysDictController {
	return &SysDictController{
		dictService: service.NewSysDictService(),
	}
}

// List 字典列表
func (c *SysDictController) List(r *ghttp.Request) {
	var req service.DictListReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	result, err := c.dictService.List(r.Context(), &req)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.PageSuccess(r, result.Records, int(result.Total), result.Page, result.PageSize)
}

// QueryById 根据ID查询字典
func (c *SysDictController) QueryById(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "字典ID不能为空")
		return
	}

	dict, err := c.dictService.GetById(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	if dict == nil {
		response.Error(r, "字典不存在")
		return
	}
	response.Success(r, dict)
}

// Add 添加字典
func (c *SysDictController) Add(r *ghttp.Request) {
	var req vo.DictCreateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	dict := &entity.SysDict{
		DictName:    req.DictName,
		DictCode:    req.DictCode,
		Description: req.Description,
		Type:        req.Type,
	}

	err := c.dictService.Create(r.Context(), dict, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "添加成功")
}

// Edit 编辑字典
func (c *SysDictController) Edit(r *ghttp.Request) {
	var req vo.DictUpdateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	dict := &entity.SysDict{
		Id:          req.Id,
		DictName:    req.DictName,
		DictCode:    req.DictCode,
		Description: req.Description,
		Type:        req.Type,
	}

	err := c.dictService.Update(r.Context(), dict, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "修改成功")
}

// Delete 删除字典
func (c *SysDictController) Delete(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "字典ID不能为空")
		return
	}

	err := c.dictService.Delete(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "删除成功")
}

// GetDictItems 根据字典编码获取字典项
func (c *SysDictController) GetDictItems(r *ghttp.Request) {
	dictCode := r.Get("dictCode").String()
	if dictCode == "" {
		response.Error(r, "字典编码不能为空")
		return
	}

	items, err := c.dictService.GetDictItems(r.Context(), dictCode)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, items)
}

// QueryAllDictItems 查询所有字典项
func (c *SysDictController) QueryAllDictItems(r *ghttp.Request) {
	items, err := c.dictService.QueryAllDictItems(r.Context())
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, items)
}

// QueryTableDictItems 查询表字典项
func (c *SysDictController) QueryTableDictItems(r *ghttp.Request) {
	table := r.Get("table").String()
	text := r.Get("text").String()
	code := r.Get("code").String()

	if table == "" || text == "" || code == "" {
		response.Error(r, "参数不完整")
		return
	}

	items, err := c.dictService.QueryTableDictItems(r.Context(), table, text, code)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.Success(r, items)
}

// ListItems 字典项列表
func (c *SysDictController) ListItems(r *ghttp.Request) {
	dictId := r.Get("dictId").String()
	page := r.Get("page").Int()
	pageSize := r.Get("pageSize").Int()

	if dictId == "" {
		response.Error(r, "字典ID不能为空")
		return
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	items, total, err := c.dictService.ListItems(r.Context(), dictId, page, pageSize)
	if err != nil {
		response.Error(r, err.Error())
		return
	}

	response.PageSuccess(r, items, total, page, pageSize)
}

// AddItem 添加字典项
func (c *SysDictController) AddItem(r *ghttp.Request) {
	var req vo.DictItemCreateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	item := &entity.SysDictItem{
		DictId:      req.DictId,
		ItemText:    req.ItemText,
		ItemValue:   req.ItemValue,
		Description: req.Description,
		SortOrder:   req.SortOrder,
		Status:      req.Status,
	}

	err := c.dictService.CreateItem(r.Context(), item, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "添加成功")
}

// EditItem 编辑字典项
func (c *SysDictController) EditItem(r *ghttp.Request) {
	var req vo.DictItemUpdateReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err.Error())
		return
	}

	username := r.Get("username").String()

	item := &entity.SysDictItem{
		Id:          req.Id,
		ItemText:    req.ItemText,
		ItemValue:   req.ItemValue,
		Description: req.Description,
		SortOrder:   req.SortOrder,
		Status:      req.Status,
	}

	err := c.dictService.UpdateItem(r.Context(), item, username)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "修改成功")
}

// DeleteItem 删除字典项
func (c *SysDictController) DeleteItem(r *ghttp.Request) {
	id := r.Get("id").String()
	if id == "" {
		response.Error(r, "字典项ID不能为空")
		return
	}

	err := c.dictService.DeleteItem(r.Context(), id)
	if err != nil {
		response.Error(r, err.Error())
		return
	}
	response.SuccessMsg(r, "删除成功")
}
