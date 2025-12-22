package sys

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "teaching-open/api/v1/sys"
	"teaching-open/internal/service"
)

// TeachingWorkController 作品管理控制器
type TeachingWorkController struct{}

var TeachingWork = &TeachingWorkController{}

// List 作品列表
func (c *TeachingWorkController) List(ctx context.Context, req *v1.WorkListReq) (res *v1.WorkListRes, err error) {
	list, total, err := service.TeachingWork().List(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询作品列表失败")
	}
	return &v1.WorkListRes{
		List:  list,
		Total: total,
	}, nil
}

// Mine 我的作品
func (c *TeachingWorkController) Mine(ctx context.Context, req *v1.WorkMineReq) (res *v1.WorkMineRes, err error) {
	list, total, err := service.TeachingWork().Mine(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询我的作品失败")
	}
	return &v1.WorkMineRes{
		List:  list,
		Total: total,
	}, nil
}

// GreatWork 优秀作品
func (c *TeachingWorkController) GreatWork(ctx context.Context, req *v1.WorkGreatReq) (res *v1.WorkGreatRes, err error) {
	list, total, err := service.TeachingWork().GreatWork(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询优秀作品失败")
	}
	return &v1.WorkGreatRes{
		List:  list,
		Total: total,
	}, nil
}

// StarWork 收藏作品
func (c *TeachingWorkController) StarWork(ctx context.Context, req *v1.WorkStarReq) (res *v1.WorkStarRes, err error) {
	list, total, err := service.TeachingWork().StarWork(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询收藏作品失败")
	}
	return &v1.WorkStarRes{
		List:  list,
		Total: total,
	}, nil
}

// Leaderboard 作品排行榜
func (c *TeachingWorkController) Leaderboard(ctx context.Context, req *v1.WorkLeaderboardReq) (res *v1.WorkLeaderboardRes, err error) {
	list, total, err := service.TeachingWork().Leaderboard(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询作品排行榜失败")
	}
	return &v1.WorkLeaderboardRes{
		List:  list,
		Total: total,
	}, nil
}

// Add 添加作品
func (c *TeachingWorkController) Add(ctx context.Context, req *v1.WorkAddReq) (res *v1.WorkAddRes, err error) {
	id, err := service.TeachingWork().Add(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "添加作品失败")
	}
	return &v1.WorkAddRes{Id: id}, nil
}

// Edit 编辑作品
func (c *TeachingWorkController) Edit(ctx context.Context, req *v1.WorkEditReq) (res *v1.WorkEditRes, err error) {
	err = service.TeachingWork().Edit(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "编辑作品失败")
	}
	return &v1.WorkEditRes{}, nil
}

// Delete 删除作品
func (c *TeachingWorkController) Delete(ctx context.Context, req *v1.WorkDeleteReq) (res *v1.WorkDeleteRes, err error) {
	err = service.TeachingWork().Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除作品失败")
	}
	return &v1.WorkDeleteRes{}, nil
}

// DeleteBatch 批量删除作品
func (c *TeachingWorkController) DeleteBatch(ctx context.Context, req *v1.WorkDeleteBatchReq) (res *v1.WorkDeleteBatchRes, err error) {
	err = service.TeachingWork().DeleteBatch(ctx, req.Ids)
	if err != nil {
		return nil, gerror.Wrap(err, "批量删除作品失败")
	}
	return &v1.WorkDeleteBatchRes{}, nil
}

// GetById 作品详情
func (c *TeachingWorkController) GetById(ctx context.Context, req *v1.WorkGetByIdReq) (res *v1.WorkGetByIdRes, err error) {
	work, err := service.TeachingWork().GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "查询作品详情失败")
	}
	return &v1.WorkGetByIdRes{WorkInfo: work}, nil
}

// Submit 提交作品
func (c *TeachingWorkController) Submit(ctx context.Context, req *v1.WorkSubmitReq) (res *v1.WorkSubmitRes, err error) {
	err = service.TeachingWork().Submit(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "提交作品失败")
	}
	return &v1.WorkSubmitRes{}, nil
}

// StudentWorkInfo 学生作品信息
func (c *TeachingWorkController) StudentWorkInfo(ctx context.Context, req *v1.WorkStudentInfoReq) (res *v1.WorkStudentInfoRes, err error) {
	workCount, works, err := service.TeachingWork().StudentWorkInfo(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询学生作品信息失败")
	}
	return &v1.WorkStudentInfoRes{
		WorkCount: workCount,
		Works:     works,
	}, nil
}

// SendWork 发送作品给其他用户
func (c *TeachingWorkController) SendWork(ctx context.Context, req *v1.WorkSendReq) (res *v1.WorkSendRes, err error) {
	err = service.TeachingWork().SendWork(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "发送作品失败")
	}
	return &v1.WorkSendRes{}, nil
}

// MineAdditional 我的附加作业作品
func (c *TeachingWorkController) MineAdditional(ctx context.Context, req *v1.WorkMineAdditionalReq) (res *v1.WorkMineAdditionalRes, err error) {
	list, total, err := service.TeachingWork().MineAdditional(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "查询附加作业作品失败")
	}
	return &v1.WorkMineAdditionalRes{
		List:  list,
		Total: total,
	}, nil
}

// StarToggle 点赞/取消点赞
func (c *TeachingWorkController) StarToggle(ctx context.Context, req *v1.WorkStarToggleReq) (res *v1.WorkStarToggleRes, err error) {
	isStared, err := service.TeachingWork().StarToggle(ctx, req.WorkId)
	if err != nil {
		return nil, gerror.Wrap(err, "点赞操作失败")
	}
	return &v1.WorkStarToggleRes{IsStared: isStared}, nil
}

// CollectToggle 收藏/取消收藏
func (c *TeachingWorkController) CollectToggle(ctx context.Context, req *v1.WorkCollectToggleReq) (res *v1.WorkCollectToggleRes, err error) {
	isCollected, err := service.TeachingWork().CollectToggle(ctx, req.WorkId)
	if err != nil {
		return nil, gerror.Wrap(err, "收藏操作失败")
	}
	return &v1.WorkCollectToggleRes{IsCollected: isCollected}, nil
}

// CorrectList 批改记录列表
func (c *TeachingWorkController) CorrectList(ctx context.Context, req *v1.WorkCorrectListReq) (res *v1.WorkCorrectListRes, err error) {
	list, err := service.TeachingWork().CorrectList(ctx, req.WorkId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询批改记录失败")
	}
	return &v1.WorkCorrectListRes{List: list}, nil
}

// CorrectAdd 添加批改记录
func (c *TeachingWorkController) CorrectAdd(ctx context.Context, req *v1.WorkCorrectAddReq) (res *v1.WorkCorrectAddRes, err error) {
	err = service.TeachingWork().CorrectAdd(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "批改作品失败")
	}
	return &v1.WorkCorrectAddRes{}, nil
}

// CommentList 评论列表
func (c *TeachingWorkController) CommentList(ctx context.Context, req *v1.WorkCommentListReq) (res *v1.WorkCommentListRes, err error) {
	list, err := service.TeachingWork().CommentList(ctx, req.WorkId)
	if err != nil {
		return nil, gerror.Wrap(err, "查询评论列表失败")
	}
	return &v1.WorkCommentListRes{List: list}, nil
}

// CommentAdd 添加评论
func (c *TeachingWorkController) CommentAdd(ctx context.Context, req *v1.WorkCommentAddReq) (res *v1.WorkCommentAddRes, err error) {
	err = service.TeachingWork().CommentAdd(ctx, req)
	if err != nil {
		return nil, gerror.Wrap(err, "添加评论失败")
	}
	return &v1.WorkCommentAddRes{}, nil
}

// CommentDelete 删除评论
func (c *TeachingWorkController) CommentDelete(ctx context.Context, req *v1.WorkCommentDeleteReq) (res *v1.WorkCommentDeleteRes, err error) {
	err = service.TeachingWork().CommentDelete(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "删除评论失败")
	}
	return &v1.WorkCommentDeleteRes{}, nil
}

// TagGet 获取作品标签
func (c *TeachingWorkController) TagGet(ctx context.Context, req *v1.WorkTagGetReq) (res *v1.WorkTagGetRes, err error) {
	tags, err := service.TeachingWork().TagGet(ctx, req.WorkId)
	if err != nil {
		return nil, gerror.Wrap(err, "获取标签失败")
	}
	return &v1.WorkTagGetRes{Tags: tags}, nil
}

// TagSet 设置作品标签
func (c *TeachingWorkController) TagSet(ctx context.Context, req *v1.WorkTagSetReq) (res *v1.WorkTagSetRes, err error) {
	err = service.TeachingWork().TagSet(ctx, req.WorkId, req.Tag)
	if err != nil {
		return nil, gerror.Wrap(err, "设置标签失败")
	}
	return &v1.WorkTagSetRes{}, nil
}

// TagDelete 删除作品标签
func (c *TeachingWorkController) TagDelete(ctx context.Context, req *v1.WorkTagDeleteReq) (res *v1.WorkTagDeleteRes, err error) {
	err = service.TeachingWork().TagDelete(ctx, req.WorkId, req.Tag)
	if err != nil {
		return nil, gerror.Wrap(err, "删除标签失败")
	}
	return &v1.WorkTagDeleteRes{}, nil
}
