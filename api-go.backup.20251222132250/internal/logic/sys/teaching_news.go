package sys

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/dao"
	"teaching-open/internal/model/entity"
	"teaching-open/internal/service"
	"teaching-open/utility/jwt"
)

func init() {
	service.RegisterTeachingNews(NewTeachingNewsLogic())
}

type sTeachingNews struct{}

func NewTeachingNewsLogic() *sTeachingNews {
	return &sTeachingNews{}
}

// List 获取新闻列表
func (s *sTeachingNews) List(ctx context.Context, req *v1.NewsListReq) (list interface{}, total int, err error) {
	m := dao.TeachingNews.Ctx(ctx)

	// 标题查询
	if req.NewsTitle != "" {
		m = m.WhereLike("news_title", "%"+req.NewsTitle+"%")
	}
	// 状态查询 (-1 表示全部)
	if req.NewsStatus >= 0 {
		m = m.Where("news_status", req.NewsStatus)
	}

	// 统计总数
	count, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	total = count

	// 分页查询
	pageNo := req.PageNo
	pageSize := req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var entities []*entity.TeachingNews
	err = m.Order("create_time DESC").Page(pageNo, pageSize).Scan(&entities)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	var newsInfos []*v1.NewsInfo
	for _, e := range entities {
		newsInfos = append(newsInfos, &v1.NewsInfo{
			Id:          e.Id,
			NewsTitle:   e.NewsTitle,
			NewsContent: e.NewsContent,
			NewsStatus:  e.NewsStatus,
			CreateBy:    e.CreateBy,
			CreateTime:  e.CreateTime,
			UpdateBy:    e.UpdateBy,
			UpdateTime:  e.UpdateTime,
		})
	}

	return newsInfos, total, nil
}

// PublicList 获取公开新闻列表（已发布）
func (s *sTeachingNews) PublicList(ctx context.Context, req *v1.NewsPublicListReq) (list interface{}, total int, err error) {
	m := dao.TeachingNews.Ctx(ctx).Where("news_status", 1) // 只查已发布

	// 统计总数
	count, err := m.Count()
	if err != nil {
		return nil, 0, err
	}
	total = count

	// 分页查询
	pageNo := req.PageNo
	pageSize := req.PageSize
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var entities []*entity.TeachingNews
	err = m.Order("create_time DESC").Page(pageNo, pageSize).Scan(&entities)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	var newsInfos []*v1.NewsInfo
	for _, e := range entities {
		newsInfos = append(newsInfos, &v1.NewsInfo{
			Id:          e.Id,
			NewsTitle:   e.NewsTitle,
			NewsContent: e.NewsContent,
			NewsStatus:  e.NewsStatus,
			CreateBy:    e.CreateBy,
			CreateTime:  e.CreateTime,
			UpdateBy:    e.UpdateBy,
			UpdateTime:  e.UpdateTime,
		})
	}

	return newsInfos, total, nil
}

// Add 添加新闻
func (s *sTeachingNews) Add(ctx context.Context, req *v1.NewsAddReq) (id string, err error) {
	username := jwt.GetUsername(ctx)
	newId := g.NewVar(nil).String()
	if newId == "" {
		newId = gconv.String(gtime.TimestampNano())
	}

	data := &entity.TeachingNews{
		Id:          newId,
		NewsTitle:   req.NewsTitle,
		NewsContent: req.NewsContent,
		NewsStatus:  0, // 默认草稿
		CreateBy:    username,
		CreateTime:  gtime.Now(),
		UpdateBy:    username,
		UpdateTime:  gtime.Now(),
	}

	_, err = dao.TeachingNews.Ctx(ctx).Insert(data)
	if err != nil {
		return "", err
	}

	return newId, nil
}

// Edit 编辑新闻
func (s *sTeachingNews) Edit(ctx context.Context, req *v1.NewsEditReq) error {
	username := jwt.GetUsername(ctx)

	data := g.Map{
		"news_title":   req.NewsTitle,
		"news_content": req.NewsContent,
		"update_by":    username,
		"update_time":  gtime.Now(),
	}

	// -1 表示不更新状态
	if req.NewsStatus >= 0 {
		data["news_status"] = req.NewsStatus
	}

	_, err := dao.TeachingNews.Ctx(ctx).Where("id", req.Id).Data(data).Update()
	return err
}

// Delete 删除新闻
func (s *sTeachingNews) Delete(ctx context.Context, id string) error {
	_, err := dao.TeachingNews.Ctx(ctx).Where("id", id).Delete()
	return err
}

// DeleteBatch 批量删除新闻
func (s *sTeachingNews) DeleteBatch(ctx context.Context, ids string) error {
	_, err := dao.TeachingNews.Ctx(ctx).WhereIn("id", ids).Delete()
	return err
}

// GetById 获取新闻详情
func (s *sTeachingNews) GetById(ctx context.Context, id string) (*v1.NewsInfo, error) {
	var e *entity.TeachingNews
	err := dao.TeachingNews.Ctx(ctx).Where("id", id).Scan(&e)
	if err != nil {
		return nil, err
	}
	if e == nil {
		return nil, nil
	}

	return &v1.NewsInfo{
		Id:          e.Id,
		NewsTitle:   e.NewsTitle,
		NewsContent: e.NewsContent,
		NewsStatus:  e.NewsStatus,
		CreateBy:    e.CreateBy,
		CreateTime:  e.CreateTime,
		UpdateBy:    e.UpdateBy,
		UpdateTime:  e.UpdateTime,
	}, nil
}

// Publish 发布新闻
func (s *sTeachingNews) Publish(ctx context.Context, id string) error {
	username := jwt.GetUsername(ctx)

	_, err := dao.TeachingNews.Ctx(ctx).Where("id", id).Data(g.Map{
		"news_status": 1,
		"update_by":   username,
		"update_time": gtime.Now(),
	}).Update()
	return err
}

// Offline 下架新闻
func (s *sTeachingNews) Offline(ctx context.Context, id string) error {
	username := jwt.GetUsername(ctx)

	_, err := dao.TeachingNews.Ctx(ctx).Where("id", id).Data(g.Map{
		"news_status": 0,
		"update_by":   username,
		"update_time": gtime.Now(),
	}).Update()
	return err
}
