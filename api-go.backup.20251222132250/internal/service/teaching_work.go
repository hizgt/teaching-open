package service

import (
	"context"

	v1 "teaching-open/api/v1/sys"
)

// ITeachingWork 作品管理服务接口
type ITeachingWork interface {
	// 作品基础管理
	List(ctx context.Context, req *v1.WorkListReq) (list interface{}, total int, err error)
	Mine(ctx context.Context, req *v1.WorkMineReq) (list interface{}, total int, err error)
	GreatWork(ctx context.Context, req *v1.WorkGreatReq) (list interface{}, total int, err error)
	StarWork(ctx context.Context, req *v1.WorkStarReq) (list interface{}, total int, err error)
	Leaderboard(ctx context.Context, req *v1.WorkLeaderboardReq) (list interface{}, total int, err error)
	Add(ctx context.Context, req *v1.WorkAddReq) (id string, err error)
	Edit(ctx context.Context, req *v1.WorkEditReq) error
	Delete(ctx context.Context, id string) error
	DeleteBatch(ctx context.Context, ids string) error
	GetById(ctx context.Context, id string) (*v1.WorkInfo, error)
	Submit(ctx context.Context, id string) error
	StudentWorkInfo(ctx context.Context, req *v1.WorkStudentInfoReq) (workCount int, works interface{}, err error)
	SendWork(ctx context.Context, req *v1.WorkSendReq) error
	MineAdditional(ctx context.Context, req *v1.WorkMineAdditionalReq) (list interface{}, total int, err error)

	// 点赞和收藏
	StarToggle(ctx context.Context, workId string) (isStared bool, err error)
	CollectToggle(ctx context.Context, workId string) (isCollected bool, err error)

	// 作品批改
	CorrectList(ctx context.Context, workId string) (list interface{}, err error)
	CorrectAdd(ctx context.Context, req *v1.WorkCorrectAddReq) error

	// 作品评论
	CommentList(ctx context.Context, workId string) (list interface{}, err error)
	CommentAdd(ctx context.Context, req *v1.WorkCommentAddReq) error
	CommentDelete(ctx context.Context, id string) error

	// 作品标签
	TagGet(ctx context.Context, workId string) (tags []string, err error)
	TagSet(ctx context.Context, workId, tag string) error
	TagDelete(ctx context.Context, workId, tag string) error
}

var localTeachingWork ITeachingWork

func TeachingWork() ITeachingWork {
	if localTeachingWork == nil {
		panic("ITeachingWork service not registered")
	}
	return localTeachingWork
}

func RegisterTeachingWork(s ITeachingWork) {
	localTeachingWork = s
}