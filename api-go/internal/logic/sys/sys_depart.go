// =================================================================================
// Logic implementation for sys depart module
// =================================================================================

package sys

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
)

func init() {
	service.RegisterSysDepart(&sSysDepart{})
}

// sSysDepart 部门服务实现
type sSysDepart struct{}

// GetTree 获取部门树
func (s *sSysDepart) GetTree(ctx context.Context) ([]*v1.DepartInfo, error) {
	var list []*entity.SysDepart
	err := dao.SysDepart.Ctx(ctx).
		Where("del_flag", "0").
		OrderAsc("depart_order").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 转换为 DepartInfo
	infoList := make([]*v1.DepartInfo, 0, len(list))
	for _, item := range list {
		infoList = append(infoList, s.entityToInfo(item))
	}

	// 构建树
	return s.buildTree(infoList, ""), nil
}

// Add 添加部门
func (s *sSysDepart) Add(ctx context.Context, req *v1.DepartAddReq) error {
	// 生成机构编码
	orgCode := req.OrgCode
	if orgCode == "" {
		orgCode = s.generateOrgCode(ctx, req.ParentId)
	}

	// 确定机构类型
	orgType := req.OrgType
	if orgType == "" {
		if req.ParentId == "" {
			orgType = "1" // 一级部门
		} else {
			orgType = "2" // 子部门
		}
	}

	depart := &entity.SysDepart{
		Id:             guid.S(),
		ParentId:       req.ParentId,
		DepartName:     req.DepartName,
		DepartNameEn:   req.DepartNameEn,
		DepartNameAbbr: req.DepartNameAbbr,
		DepartOrder:    req.DepartOrder,
		Description:    req.Description,
		OrgCategory:    req.OrgCategory,
		OrgType:        orgType,
		OrgCode:        orgCode,
		Mobile:         req.Mobile,
		Fax:            req.Fax,
		Address:        req.Address,
		Memo:           req.Memo,
		Status:         req.Status,
		DelFlag:        "0",
		CreateTime:     gtime.Now(),
	}

	_, err := dao.SysDepart.Ctx(ctx).Insert(depart)
	return err
}

// Edit 编辑部门
func (s *sSysDepart) Edit(ctx context.Context, req *v1.DepartEditReq) error {
	// 检查部门是否存在
	count, err := dao.SysDepart.Ctx(ctx).Where("id", req.Id).Where("del_flag", "0").Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New(g.I18n().T(ctx, "部门不存在"))
	}

	_, err = dao.SysDepart.Ctx(ctx).
		Where("id", req.Id).
		Data(g.Map{
			"parent_id":        req.ParentId,
			"depart_name":      req.DepartName,
			"depart_name_en":   req.DepartNameEn,
			"depart_name_abbr": req.DepartNameAbbr,
			"depart_order":     req.DepartOrder,
			"description":      req.Description,
			"org_category":     req.OrgCategory,
			"org_type":         req.OrgType,
			"org_code":         req.OrgCode,
			"mobile":           req.Mobile,
			"fax":              req.Fax,
			"address":          req.Address,
			"memo":             req.Memo,
			"status":           req.Status,
			"update_time":      gtime.Now(),
		}).
		Update()
	return err
}

// Delete 删除部门
func (s *sSysDepart) Delete(ctx context.Context, id string) error {
	ids := strings.Split(id, ",")
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, depId := range ids {
			depId = strings.TrimSpace(depId)
			if depId == "" {
				continue
			}

			// 检查是否有子部门
			childCount, err := dao.SysDepart.Ctx(ctx).
				Where("parent_id", depId).
				Where("del_flag", "0").
				Count()
			if err != nil {
				return err
			}
			if childCount > 0 {
				return gerror.New(g.I18n().T(ctx, "存在子部门，无法删除"))
			}

			// 检查是否有用户关联
			userCount, err := dao.SysUserDepart.Ctx(ctx).
				Where("dep_id", depId).
				Count()
			if err != nil {
				return err
			}
			if userCount > 0 {
				return gerror.New(g.I18n().T(ctx, "部门下存在用户，无法删除"))
			}

			// 软删除部门
			_, err = dao.SysDepart.Ctx(ctx).
				Where("id", depId).
				Data(g.Map{
					"del_flag":    "1",
					"update_time": gtime.Now(),
				}).
				Update()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// GetById 根据ID获取部门详情
func (s *sSysDepart) GetById(ctx context.Context, id string) (*v1.DepartInfo, error) {
	var depart *entity.SysDepart
	err := dao.SysDepart.Ctx(ctx).
		Where("id", id).
		Where("del_flag", "0").
		Scan(&depart)
	if err != nil {
		return nil, err
	}
	if depart == nil {
		return nil, nil
	}
	return s.entityToInfo(depart), nil
}

// GetIdTree 获取部门ID树（用于下拉选择）
func (s *sSysDepart) GetIdTree(ctx context.Context) ([]*v1.DepartIdNode, error) {
	var list []*entity.SysDepart
	err := dao.SysDepart.Ctx(ctx).
		Where("del_flag", "0").
		OrderAsc("depart_order").
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 转换为 IdNode
	nodeList := make([]*v1.DepartIdNode, 0, len(list))
	for _, item := range list {
		nodeList = append(nodeList, &v1.DepartIdNode{
			Key:   item.Id,
			Value: item.Id,
			Title: item.DepartName,
		})
	}

	// 构建树
	return s.buildIdTree(nodeList, list, ""), nil
}

// SearchByKeyword 关键字搜索部门
func (s *sSysDepart) SearchByKeyword(ctx context.Context, keyword string) ([]*v1.DepartInfo, error) {
	var list []*entity.SysDepart
	query := dao.SysDepart.Ctx(ctx).Where("del_flag", "0")
	if keyword != "" {
		query = query.WhereLike("depart_name", "%"+keyword+"%")
	}
	err := query.OrderAsc("depart_order").Scan(&list)
	if err != nil {
		return nil, err
	}

	result := make([]*v1.DepartInfo, 0, len(list))
	for _, item := range list {
		result = append(result, s.entityToInfo(item))
	}
	return result, nil
}

// GetUserDeparts 获取用户部门
func (s *sSysDepart) GetUserDeparts(ctx context.Context, userId string) ([]string, []*v1.DepartInfo, error) {
	var userDeparts []*entity.SysUserDepart
	err := dao.SysUserDepart.Ctx(ctx).
		Where("user_id", userId).
		Scan(&userDeparts)
	if err != nil {
		return nil, nil, err
	}

	if len(userDeparts) == 0 {
		return []string{}, []*v1.DepartInfo{}, nil
	}

	// 获取部门ID列表
	depIds := make([]string, 0, len(userDeparts))
	for _, ud := range userDeparts {
		depIds = append(depIds, ud.DepId)
	}

	// 获取部门详情
	var departs []*entity.SysDepart
	err = dao.SysDepart.Ctx(ctx).
		WhereIn("id", depIds).
		Where("del_flag", "0").
		Scan(&departs)
	if err != nil {
		return nil, nil, err
	}

	departInfos := make([]*v1.DepartInfo, 0, len(departs))
	for _, d := range departs {
		departInfos = append(departInfos, s.entityToInfo(d))
	}

	return depIds, departInfos, nil
}

// SaveUserDeparts 保存用户部门
func (s *sSysDepart) SaveUserDeparts(ctx context.Context, userId string, departIds []string) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 先删除原有关联
		_, err := dao.SysUserDepart.Ctx(ctx).
			Where("user_id", userId).
			Delete()
		if err != nil {
			return err
		}

		// 批量插入新关联
		if len(departIds) > 0 {
			records := make([]g.Map, 0, len(departIds))
			for _, depId := range departIds {
				records = append(records, g.Map{
					"ID":      guid.S(),
					"user_id": userId,
					"dep_id":  depId,
				})
			}
			_, err = dao.SysUserDepart.Ctx(ctx).Insert(records)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// entityToInfo 实体转信息结构
func (s *sSysDepart) entityToInfo(e *entity.SysDepart) *v1.DepartInfo {
	info := &v1.DepartInfo{
		Id:             e.Id,
		ParentId:       e.ParentId,
		DepartName:     e.DepartName,
		DepartNameEn:   e.DepartNameEn,
		DepartNameAbbr: e.DepartNameAbbr,
		DepartOrder:    e.DepartOrder,
		Description:    e.Description,
		OrgCategory:    e.OrgCategory,
		OrgType:        e.OrgType,
		OrgCode:        e.OrgCode,
		Mobile:         e.Mobile,
		Fax:            e.Fax,
		Address:        e.Address,
		Memo:           e.Memo,
		Status:         e.Status,
	}
	if e.CreateTime != nil {
		info.CreateTime = e.CreateTime.Format("Y-m-d H:i:s")
	}
	return info
}

// buildTree 构建部门树
func (s *sSysDepart) buildTree(list []*v1.DepartInfo, parentId string) []*v1.DepartInfo {
	result := make([]*v1.DepartInfo, 0)
	for _, item := range list {
		if item.ParentId == parentId {
			item.Children = s.buildTree(list, item.Id)
			result = append(result, item)
		}
	}
	return result
}

// buildIdTree 构建部门ID树
func (s *sSysDepart) buildIdTree(nodeList []*v1.DepartIdNode, entityList []*entity.SysDepart, parentId string) []*v1.DepartIdNode {
	result := make([]*v1.DepartIdNode, 0)
	for i, item := range entityList {
		if item.ParentId == parentId {
			node := nodeList[i]
			node.Children = s.buildIdTree(nodeList, entityList, item.Id)
			result = append(result, node)
		}
	}
	return result
}

// generateOrgCode 生成机构编码
func (s *sSysDepart) generateOrgCode(ctx context.Context, parentId string) string {
	if parentId == "" {
		// 一级部门，获取最大编码
		var lastDepart *entity.SysDepart
		err := dao.SysDepart.Ctx(ctx).
			Where("parent_id", "").
			Where("del_flag", "0").
			OrderDesc("org_code").
			Limit(1).
			Scan(&lastDepart)
		if err != nil || lastDepart == nil || lastDepart.OrgCode == "" {
			return "A01"
		}
		// 简单递增
		return s.incrementCode(lastDepart.OrgCode)
	}

	// 子部门
	var parent *entity.SysDepart
	err := dao.SysDepart.Ctx(ctx).Where("id", parentId).Scan(&parent)
	if err != nil || parent == nil {
		return guid.S()[:8]
	}

	var lastChild *entity.SysDepart
	err = dao.SysDepart.Ctx(ctx).
		Where("parent_id", parentId).
		Where("del_flag", "0").
		OrderDesc("org_code").
		Limit(1).
		Scan(&lastChild)
	if err != nil || lastChild == nil || lastChild.OrgCode == "" {
		return parent.OrgCode + "01"
	}

	return s.incrementCode(lastChild.OrgCode)
}

// incrementCode 递增编码
func (s *sSysDepart) incrementCode(code string) string {
	if len(code) < 2 {
		return code + "01"
	}
	// 尝试解析最后两位数字
	suffix := code[len(code)-2:]
	prefix := code[:len(code)-2]

	var num int
	err := g.NewVar(suffix).Scan(&num)
	if err != nil {
		return code + "01"
	}

	num++
	if num > 99 {
		return code + "01"
	}

	return prefix + g.NewVar(num).String()
}
