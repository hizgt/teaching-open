package service

import (
	"context"
	"fmt"
	"sort"
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

// SysDepartService 部门服务接口
type SysDepartService interface {
	GetTree(ctx context.Context, req *vo.DepartTreeReq) ([]*vo.DepartTreeNode, error)
	GetById(ctx context.Context, id string) (*entity.SysDepart, error)
	Create(ctx context.Context, req *vo.DepartCreateReq, createBy string) (*vo.DepartCreateRes, error)
	Update(ctx context.Context, req *vo.DepartUpdateReq, updateBy string) error
	Delete(ctx context.Context, id string) error
	GetIdTree(ctx context.Context) ([]*vo.DepartIdTreeNode, error)
	Search(ctx context.Context, req *vo.DepartSearchReq) ([]entity.SysDepart, error)
	GetUserDeparts(ctx context.Context, userId string) ([]string, error)
	SaveUserDeparts(ctx context.Context, req *vo.UserDepartReq) error
	GetDepartUserTree(ctx context.Context) ([]*vo.DepartUserTreeNode, error)
}

// sysDepartServiceImpl 部门服务实现
type sysDepartServiceImpl struct{}

// NewSysDepartService 创建部门服务实例
func NewSysDepartService() SysDepartService {
	return &sysDepartServiceImpl{}
}

// GetTree 获取部门树
func (s *sysDepartServiceImpl) GetTree(ctx context.Context, req *vo.DepartTreeReq) ([]*vo.DepartTreeNode, error) {
	model := dao.SysDepart.Ctx(ctx).Where(dao.SysDepart.Columns().DelFlag, "0")

	if req != nil && req.DepartName != "" {
		model = model.WhereLike(dao.SysDepart.Columns().DepartName, "%"+req.DepartName+"%")
	}
	if req != nil && req.OrgCode != "" {
		model = model.WhereLike(dao.SysDepart.Columns().OrgCode, req.OrgCode+"%")
	}

	var departs []entity.SysDepart
	err := model.OrderAsc(dao.SysDepart.Columns().DepartOrder).Scan(&departs)
	if err != nil {
		return nil, err
	}

	return s.buildTree(departs, ""), nil
}

// buildTree 构建部门树
func (s *sysDepartServiceImpl) buildTree(departs []entity.SysDepart, parentId string) []*vo.DepartTreeNode {
	var nodes []*vo.DepartTreeNode
	for _, depart := range departs {
		if depart.ParentId == parentId {
			node := &vo.DepartTreeNode{
				Id:             depart.Id,
				ParentId:       depart.ParentId,
				DepartName:     depart.DepartName,
				DepartNameEn:   depart.DepartNameEn,
				DepartNameAbbr: depart.DepartNameAbbr,
				DepartOrder:    depart.DepartOrder,
				Description:    depart.Description,
				OrgCategory:    depart.OrgCategory,
				OrgType:        depart.OrgType,
				OrgCode:        depart.OrgCode,
				Mobile:         depart.Mobile,
				Address:        depart.Address,
				Status:         depart.Status,
				Key:            depart.Id,
				Value:          depart.Id,
				Title:          depart.DepartName,
			}
			children := s.buildTree(departs, depart.Id)
			if len(children) > 0 {
				node.Children = children
				node.IsLeaf = false
			} else {
				node.IsLeaf = true
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

// GetById 根据ID获取部门
func (s *sysDepartServiceImpl) GetById(ctx context.Context, id string) (*entity.SysDepart, error) {
	var depart entity.SysDepart
	err := dao.SysDepart.Ctx(ctx).
		Where(dao.SysDepart.Columns().Id, id).
		Where(dao.SysDepart.Columns().DelFlag, "0").
		Scan(&depart)
	if err != nil {
		return nil, err
	}
	if depart.Id == "" {
		return nil, nil
	}
	return &depart, nil
}

// Create 创建部门
func (s *sysDepartServiceImpl) Create(ctx context.Context, req *vo.DepartCreateReq, createBy string) (*vo.DepartCreateRes, error) {
	// 生成OrgCode
	orgCode, err := s.generateOrgCode(ctx, req.ParentId)
	if err != nil {
		return nil, err
	}

	departId := guid.S()
	depart := entity.SysDepart{
		Id:             departId,
		ParentId:       req.ParentId,
		DepartName:     req.DepartName,
		DepartNameEn:   req.DepartNameEn,
		DepartNameAbbr: req.DepartNameAbbr,
		DepartOrder:    req.DepartOrder,
		Description:    req.Description,
		OrgCategory:    req.OrgCategory,
		OrgType:        req.OrgType,
		OrgCode:        orgCode,
		Mobile:         req.Mobile,
		Fax:            req.Fax,
		Address:        req.Address,
		Memo:           req.Memo,
		Status:         "1",
		DelFlag:        "0",
		CreateBy:       createBy,
		CreateTime:     gtime.Now(),
		UpdateBy:       createBy,
		UpdateTime:     gtime.Now(),
	}

	_, err = dao.SysDepart.Ctx(ctx).Data(depart).Insert()
	if err != nil {
		return nil, err
	}

	return &vo.DepartCreateRes{Id: departId}, nil
}

// generateOrgCode 生成机构编码
func (s *sysDepartServiceImpl) generateOrgCode(ctx context.Context, parentId string) (string, error) {
	if parentId == "" {
		// 顶级部门，查找最大的顶级编码
		var maxCode string
		err := dao.SysDepart.Ctx(ctx).
			Fields("MAX(org_code) as org_code").
			Where(dao.SysDepart.Columns().ParentId, "").
			Where(dao.SysDepart.Columns().DelFlag, "0").
			Scan(&maxCode)
		if err != nil {
			return "", err
		}
		if maxCode == "" {
			return "A01", nil
		}
		return s.nextCode(maxCode), nil
	}

	// 子部门，先获取父部门编码
	parent, err := s.GetById(ctx, parentId)
	if err != nil {
		return "", err
	}
	if parent == nil {
		return "", gerror.New("父部门不存在")
	}

	// 查找该父部门下最大的子编码
	var maxCode string
	err = dao.SysDepart.Ctx(ctx).
		Fields("MAX(org_code) as org_code").
		Where(dao.SysDepart.Columns().ParentId, parentId).
		Where(dao.SysDepart.Columns().DelFlag, "0").
		Scan(&maxCode)
	if err != nil {
		return "", err
	}

	if maxCode == "" {
		return parent.OrgCode + "A01", nil
	}
	return s.nextCode(maxCode), nil
}

// nextCode 生成下一个编码
func (s *sysDepartServiceImpl) nextCode(code string) string {
	if len(code) < 3 {
		return "A01"
	}
	// 取最后3位
	suffix := code[len(code)-3:]
	letter := suffix[0]
	num := (int(suffix[1]-'0') * 10) + int(suffix[2]-'0')

	num++
	if num > 99 {
		num = 1
		letter++
		if letter > 'Z' {
			letter = 'A'
		}
	}

	prefix := code[:len(code)-3]
	return fmt.Sprintf("%s%c%02d", prefix, letter, num)
}

// Update 更新部门
func (s *sysDepartServiceImpl) Update(ctx context.Context, req *vo.DepartUpdateReq, updateBy string) error {
	depart, err := s.GetById(ctx, req.Id)
	if err != nil {
		return err
	}
	if depart == nil {
		return gerror.New("部门不存在")
	}

	// 不能将部门设为自己的子部门
	if req.ParentId == req.Id {
		return gerror.New("不能将部门设为自己的子部门")
	}

	// 检查是否将部门移动到自己的子部门下
	if req.ParentId != "" && req.ParentId != depart.ParentId {
		children, err := s.getAllChildIds(ctx, req.Id)
		if err != nil {
			return err
		}
		for _, childId := range children {
			if childId == req.ParentId {
				return gerror.New("不能将部门移动到自己的子部门下")
			}
		}
	}

	_, err = dao.SysDepart.Ctx(ctx).Data(g.Map{
		dao.SysDepart.Columns().ParentId:       req.ParentId,
		dao.SysDepart.Columns().DepartName:     req.DepartName,
		dao.SysDepart.Columns().DepartNameEn:   req.DepartNameEn,
		dao.SysDepart.Columns().DepartNameAbbr: req.DepartNameAbbr,
		dao.SysDepart.Columns().DepartOrder:    req.DepartOrder,
		dao.SysDepart.Columns().Description:    req.Description,
		dao.SysDepart.Columns().OrgCategory:    req.OrgCategory,
		dao.SysDepart.Columns().OrgType:        req.OrgType,
		dao.SysDepart.Columns().Mobile:         req.Mobile,
		dao.SysDepart.Columns().Fax:            req.Fax,
		dao.SysDepart.Columns().Address:        req.Address,
		dao.SysDepart.Columns().Memo:           req.Memo,
		dao.SysDepart.Columns().UpdateBy:       updateBy,
		dao.SysDepart.Columns().UpdateTime:     gtime.Now(),
	}).Where(dao.SysDepart.Columns().Id, req.Id).Update()

	return err
}

// getAllChildIds 获取所有子部门ID
func (s *sysDepartServiceImpl) getAllChildIds(ctx context.Context, parentId string) ([]string, error) {
	var ids []string
	var children []entity.SysDepart
	err := dao.SysDepart.Ctx(ctx).
		Where(dao.SysDepart.Columns().ParentId, parentId).
		Where(dao.SysDepart.Columns().DelFlag, "0").
		Scan(&children)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		ids = append(ids, child.Id)
		childIds, err := s.getAllChildIds(ctx, child.Id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, childIds...)
	}
	return ids, nil
}

// Delete 删除部门（递归删除子部门）
func (s *sysDepartServiceImpl) Delete(ctx context.Context, id string) error {
	// 获取所有子部门ID
	childIds, err := s.getAllChildIds(ctx, id)
	if err != nil {
		return err
	}

	allIds := append([]string{id}, childIds...)

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 逻辑删除部门
		_, err := dao.SysDepart.Ctx(ctx).Data(g.Map{
			dao.SysDepart.Columns().DelFlag:    "1",
			dao.SysDepart.Columns().UpdateTime: gtime.Now(),
		}).WhereIn(dao.SysDepart.Columns().Id, allIds).Update()
		if err != nil {
			return err
		}

		// 删除用户部门关联
		_, err = dao.SysUserDepart.Ctx(ctx).WhereIn(dao.SysUserDepart.Columns().DepId, allIds).Delete()
		return err
	})
}

// GetIdTree 获取部门ID树
func (s *sysDepartServiceImpl) GetIdTree(ctx context.Context) ([]*vo.DepartIdTreeNode, error) {
	var departs []entity.SysDepart
	err := dao.SysDepart.Ctx(ctx).
		Where(dao.SysDepart.Columns().DelFlag, "0").
		OrderAsc(dao.SysDepart.Columns().DepartOrder).
		Scan(&departs)
	if err != nil {
		return nil, err
	}

	return s.buildIdTree(departs, ""), nil
}

// buildIdTree 构建部门ID树
func (s *sysDepartServiceImpl) buildIdTree(departs []entity.SysDepart, parentId string) []*vo.DepartIdTreeNode {
	var nodes []*vo.DepartIdTreeNode
	for _, depart := range departs {
		if depart.ParentId == parentId {
			node := &vo.DepartIdTreeNode{
				Key:   depart.Id,
				Value: depart.Id,
				Title: depart.DepartName,
			}
			children := s.buildIdTree(departs, depart.Id)
			if len(children) > 0 {
				node.Children = children
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

// Search 搜索部门
func (s *sysDepartServiceImpl) Search(ctx context.Context, req *vo.DepartSearchReq) ([]entity.SysDepart, error) {
	model := dao.SysDepart.Ctx(ctx).Where(dao.SysDepart.Columns().DelFlag, "0")

	if req.DepartName != "" {
		model = model.WhereLike(dao.SysDepart.Columns().DepartName, "%"+req.DepartName+"%")
	}
	if req.OrgCode != "" {
		model = model.WhereLike(dao.SysDepart.Columns().OrgCode, "%"+req.OrgCode+"%")
	}
	if req.OrgType != "" {
		model = model.Where(dao.SysDepart.Columns().OrgType, req.OrgType)
	}

	var departs []entity.SysDepart
	err := model.OrderAsc(dao.SysDepart.Columns().DepartOrder).Scan(&departs)
	return departs, err
}

// GetUserDeparts 获取用户部门ID列表
func (s *sysDepartServiceImpl) GetUserDeparts(ctx context.Context, userId string) ([]string, error) {
	var depIds []string
	err := dao.SysUserDepart.Ctx(ctx).
		Fields(dao.SysUserDepart.Columns().DepId).
		Where(dao.SysUserDepart.Columns().UserId, userId).
		Scan(&depIds)
	return depIds, err
}

// SaveUserDeparts 保存用户部门关联
func (s *sysDepartServiceImpl) SaveUserDeparts(ctx context.Context, req *vo.UserDepartReq) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除原有关联
		_, err := dao.SysUserDepart.Ctx(ctx).Where(dao.SysUserDepart.Columns().UserId, req.UserId).Delete()
		if err != nil {
			return err
		}

		// 添加新关联
		if len(req.DepIds) > 0 {
			var relations []g.Map
			for _, depId := range req.DepIds {
				relations = append(relations, g.Map{
					dao.SysUserDepart.Columns().Id:     guid.S(),
					dao.SysUserDepart.Columns().UserId: req.UserId,
					dao.SysUserDepart.Columns().DepId:  depId,
				})
			}
			_, err = dao.SysUserDepart.Ctx(ctx).Data(relations).Insert()
		}
		return err
	})
}

// GetDepartUserTree 获取部门用户树（仅部门，不含用户）
func (s *sysDepartServiceImpl) GetDepartUserTree(ctx context.Context) ([]*vo.DepartUserTreeNode, error) {
	// 获取所有部门
	var departs []entity.SysDepart
	err := dao.SysDepart.Ctx(ctx).
		Where(dao.SysDepart.Columns().DelFlag, "0").
		OrderAsc(dao.SysDepart.Columns().DepartOrder).
		Scan(&departs)
	if err != nil {
		return nil, err
	}

	return s.buildDepartUserTreeNodes(departs, ""), nil
}

// buildDepartUserTreeNodes 构建部门用户树节点
func (s *sysDepartServiceImpl) buildDepartUserTreeNodes(departs []entity.SysDepart, parentId string) []*vo.DepartUserTreeNode {
	var nodes []*vo.DepartUserTreeNode
	for _, depart := range departs {
		if depart.ParentId == parentId {
			node := &vo.DepartUserTreeNode{
				Key:   depart.Id,
				Value: depart.Id,
				Title: depart.DepartName,
				Type:  "depart",
			}
			children := s.buildDepartUserTreeNodes(departs, depart.Id)
			if len(children) > 0 {
				node.Children = children
			}
			nodes = append(nodes, node)
		}
	}
	// 按部门名称排序
	sort.Slice(nodes, func(i, j int) bool {
		return strings.Compare(nodes[i].Title, nodes[j].Title) < 0
	})
	return nodes
}
