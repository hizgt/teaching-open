package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 作品管理 ====================

// WorkListReq 作品列表请求
type WorkListReq struct {
	g.Meta     `path:"/teaching/teachingWork/list" method:"get" tags:"作品管理" summary:"作品列表"`
	WorkName   string `json:"workName" dc:"作品名称"`
	WorkType   string `json:"workType" dc:"作品类型"`
	WorkStatus int    `json:"workStatus" dc:"作品状态" d:"-1"`
	UserId     string `json:"userId" dc:"用户ID"`
	CourseId   string `json:"courseId" dc:"课程ID"`
	DepartId   string `json:"departId" dc:"班级ID"`
	WorkScene  string `json:"workScene" dc:"来源场景"`
	PageNo     int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize   int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// WorkListRes 作品列表响应
type WorkListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}

// WorkMineReq 我的作品请求
type WorkMineReq struct {
	g.Meta     `path:"/teaching/teachingWork/mine" method:"get" tags:"作品管理" summary:"我的作品"`
	WorkName   string `json:"workName" dc:"作品名称"`
	WorkType   string `json:"workType" dc:"作品类型"`
	WorkStatus int    `json:"workStatus" dc:"作品状态" d:"-1"`
	WorkScene  string `json:"workScene" dc:"来源场景"`
	PageNo     int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize   int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// WorkMineRes 我的作品响应
type WorkMineRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}

// WorkGreatReq 优秀作品请求
type WorkGreatReq struct {
	g.Meta   `path:"/teaching/teachingWork/greatWork" method:"get" tags:"作品管理" summary:"优秀作品"`
	WorkType string `json:"workType" dc:"作品类型"`
	PageNo   int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// WorkGreatRes 优秀作品响应
type WorkGreatRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}

// WorkStarReq 收藏作品请求
type WorkStarReq struct {
	g.Meta   `path:"/teaching/teachingWork/starWork" method:"get" tags:"作品管理" summary:"收藏作品"`
	WorkType string `json:"workType" dc:"作品类型"`
	PageNo   int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// WorkStarRes 收藏作品响应
type WorkStarRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}

// WorkLeaderboardReq 作品排行榜请求
type WorkLeaderboardReq struct {
	g.Meta   `path:"/teaching/teachingWork/leaderboard" method:"get" tags:"作品管理" summary:"作品排行榜"`
	WorkType string `json:"workType" dc:"作品类型"`
	SortBy   string `json:"sortBy" dc:"排序方式" d:"star_num"`
	PageNo   int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// WorkLeaderboardRes 作品排行榜响应
type WorkLeaderboardRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}

// WorkAddReq 添加作品请求
type WorkAddReq struct {
	g.Meta       `path:"/teaching/teachingWork/add" method:"post" tags:"作品管理" summary:"添加作品"`
	WorkName     string `json:"workName" v:"required#作品名称不能为空" dc:"作品名称"`
	WorkType     string `json:"workType" v:"required#作品类型不能为空" dc:"作品类型"`
	WorkFile     string `json:"workFile" v:"required#作品文件不能为空" dc:"作品文件"`
	WorkCover    string `json:"workCover" dc:"作品封面"`
	CourseId     string `json:"courseId" dc:"课程ID"`
	DepartId     string `json:"departId" dc:"班级ID"`
	AdditionalId string `json:"additionalId" dc:"附加作业ID"`
	WorkScene    string `json:"workScene" dc:"来源场景"`
	HasCloudData int    `json:"hasCloudData" dc:"是否包含云变量" d:"0"`
}

// WorkAddRes 添加作品响应
type WorkAddRes struct {
	g.Meta `mime:"application/json"`
	Id     string `json:"id"`
}

// WorkEditReq 编辑作品请求
type WorkEditReq struct {
	g.Meta       `path:"/teaching/teachingWork/edit" method:"put" tags:"作品管理" summary:"编辑作品"`
	Id           string `json:"id" v:"required#ID不能为空" dc:"作品ID"`
	WorkName     string `json:"workName" dc:"作品名称"`
	WorkType     string `json:"workType" dc:"作品类型"`
	WorkFile     string `json:"workFile" dc:"作品文件"`
	WorkCover    string `json:"workCover" dc:"作品封面"`
	WorkStatus   int    `json:"workStatus" dc:"作品状态"`
	HasCloudData int    `json:"hasCloudData" dc:"是否包含云变量"`
}

// WorkEditRes 编辑作品响应
type WorkEditRes struct {
	g.Meta `mime:"application/json"`
}

// WorkDeleteReq 删除作品请求
type WorkDeleteReq struct {
	g.Meta `path:"/teaching/teachingWork/delete" method:"delete" tags:"作品管理" summary:"删除作品"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"作品ID"`
}

// WorkDeleteRes 删除作品响应
type WorkDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// WorkDeleteBatchReq 批量删除作品请求
type WorkDeleteBatchReq struct {
	g.Meta `path:"/teaching/teachingWork/deleteBatch" method:"delete" tags:"作品管理" summary:"批量删除作品"`
	Ids    string `json:"ids" v:"required#IDs不能为空" dc:"作品ID列表，逗号分隔"`
}

// WorkDeleteBatchRes 批量删除作品响应
type WorkDeleteBatchRes struct {
	g.Meta `mime:"application/json"`
}

// WorkGetByIdReq 作品详情请求
type WorkGetByIdReq struct {
	g.Meta `path:"/teaching/teachingWork/queryById" method:"get" tags:"作品管理" summary:"作品详情"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"作品ID"`
}

// WorkGetByIdRes 作品详情响应
type WorkGetByIdRes struct {
	g.Meta `mime:"application/json"`
	*WorkInfo
}

// WorkInfo 作品信息
type WorkInfo struct {
	Id           string      `json:"id"`
	CreateBy     string      `json:"createBy"`
	CreateTime   interface{} `json:"createTime"`
	UpdateBy     string      `json:"updateBy"`
	UpdateTime   interface{} `json:"updateTime"`
	SysOrgCode   string      `json:"sysOrgCode"`
	UserId       string      `json:"userId"`
	DepartId     string      `json:"departId"`
	CourseId     string      `json:"courseId"`
	WorkName     string      `json:"workName"`
	WorkType     string      `json:"workType"`
	WorkFile     string      `json:"workFile"`
	WorkCover    string      `json:"workCover"`
	WorkStatus   int         `json:"workStatus"`
	StarNum      int         `json:"starNum"`
	CollectNum   int         `json:"collectNum"`
	ViewNum      int         `json:"viewNum"`
	AdditionalId string      `json:"additionalId"`
	WorkScene    string      `json:"workScene"`
	HasCloudData int         `json:"hasCloudData"`
	// 扩展字段
	UserName    string `json:"userName"`
	RealName    string `json:"realname"`
	CourseName  string `json:"courseName"`
	DepartName  string `json:"departName"`
	IsStarred   bool   `json:"isStarred"`
	IsCollected bool   `json:"isCollected"`
}

// WorkSubmitReq 提交作品请求
type WorkSubmitReq struct {
	g.Meta `path:"/teaching/teachingWork/submit" method:"post" tags:"作品管理" summary:"提交作品"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"作品ID"`
}

// WorkSubmitRes 提交作品响应
type WorkSubmitRes struct {
	g.Meta `mime:"application/json"`
}

// WorkStudentInfoReq 学生作品信息请求
type WorkStudentInfoReq struct {
	g.Meta   `path:"/teaching/teachingWork/studentWorkInfo" method:"get" tags:"作品管理" summary:"学生作品信息"`
	UserId   string `json:"userId" dc:"学生ID"`
	CourseId string `json:"courseId" dc:"课程ID"`
}

// WorkStudentInfoRes 学生作品信息响应
type WorkStudentInfoRes struct {
	g.Meta    `mime:"application/json"`
	WorkCount int         `json:"workCount"`
	Works     interface{} `json:"works"`
}

// WorkSendReq 发送作品请求
type WorkSendReq struct {
	g.Meta `path:"/teaching/teachingWork/sendWork" method:"post" tags:"作品管理" summary:"发送作品给其他用户"`
	WorkId string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
	UserId string `json:"userId" v:"required#目标用户ID不能为空" dc:"目标用户ID"`
}

// WorkSendRes 发送作品响应
type WorkSendRes struct {
	g.Meta `mime:"application/json"`
}

// WorkStarToggleReq 点赞/取消点赞请求
type WorkStarToggleReq struct {
	g.Meta `path:"/teaching/teachingWork/star" method:"post" tags:"作品管理" summary:"点赞/取消点赞"`
	WorkId string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
}

// WorkStarToggleRes 点赞/取消点赞响应
type WorkStarToggleRes struct {
	g.Meta   `mime:"application/json"`
	IsStared bool `json:"isStared"`
}

// WorkCollectToggleReq 收藏/取消收藏请求
type WorkCollectToggleReq struct {
	g.Meta `path:"/teaching/teachingWork/collect" method:"post" tags:"作品管理" summary:"收藏/取消收藏"`
	WorkId string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
}

// WorkCollectToggleRes 收藏/取消收藏响应
type WorkCollectToggleRes struct {
	g.Meta      `mime:"application/json"`
	IsCollected bool `json:"isCollected"`
}

// ==================== 作品批改 ====================

// WorkCorrectListReq 批改记录列表请求
type WorkCorrectListReq struct {
	g.Meta `path:"/teaching/teachingWork/queryTeachingWorkCorrectByMainId" method:"get" tags:"作品批改" summary:"批改记录列表"`
	WorkId string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
}

// WorkCorrectListRes 批改记录列表响应
type WorkCorrectListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
}

// WorkCorrectAddReq 添加批改记录请求
type WorkCorrectAddReq struct {
	g.Meta  `path:"/teaching/teachingWork/correct" method:"post" tags:"作品批改" summary:"批改作品"`
	WorkId  string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
	Score   int    `json:"score" dc:"评分"`
	Comment string `json:"comment" dc:"评语"`
}

// WorkCorrectAddRes 添加批改记录响应
type WorkCorrectAddRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 作品评论 ====================

// WorkCommentListReq 评论列表请求
type WorkCommentListReq struct {
	g.Meta `path:"/teaching/teachingWork/getWorkComments" method:"get" tags:"作品评论" summary:"获取评论列表"`
	WorkId string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
}

// WorkCommentListRes 评论列表响应
type WorkCommentListRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
}

// WorkCommentAddReq 添加评论请求
type WorkCommentAddReq struct {
	g.Meta  `path:"/teaching/teachingWork/saveComment" method:"post" tags:"作品评论" summary:"添加评论"`
	WorkId  string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
	Comment string `json:"comment" v:"required#评论内容不能为空" dc:"评论内容"`
}

// WorkCommentAddRes 添加评论响应
type WorkCommentAddRes struct {
	g.Meta `mime:"application/json"`
}

// WorkCommentDeleteReq 删除评论请求
type WorkCommentDeleteReq struct {
	g.Meta `path:"/teaching/teachingWork/deleteComment" method:"delete" tags:"作品评论" summary:"删除评论"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"评论ID"`
}

// WorkCommentDeleteRes 删除评论响应
type WorkCommentDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// ==================== 作品标签 ====================

// WorkTagGetReq 获取标签请求
type WorkTagGetReq struct {
	g.Meta `path:"/teaching/teachingWork/getWorkTags" method:"get" tags:"作品标签" summary:"获取作品标签"`
	WorkId string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
}

// WorkTagGetRes 获取标签响应
type WorkTagGetRes struct {
	g.Meta `mime:"application/json"`
	Tags   []string `json:"tags"`
}

// WorkTagSetReq 设置标签请求
type WorkTagSetReq struct {
	g.Meta `path:"/teaching/teachingWork/setWorkTag" method:"post" tags:"作品标签" summary:"设置作品标签"`
	WorkId string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
	Tag    string `json:"tag" v:"required#标签不能为空" dc:"标签"`
}

// WorkTagSetRes 设置标签响应
type WorkTagSetRes struct {
	g.Meta `mime:"application/json"`
}

// WorkTagDeleteReq 删除标签请求
type WorkTagDeleteReq struct {
	g.Meta `path:"/teaching/teachingWork/delWorkTag" method:"delete" tags:"作品标签" summary:"删除作品标签"`
	WorkId string `json:"workId" v:"required#作品ID不能为空" dc:"作品ID"`
	Tag    string `json:"tag" v:"required#标签不能为空" dc:"标签"`
}

// WorkTagDeleteRes 删除标签响应
type WorkTagDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// WorkMineAdditionalReq 我的附加作业作品请求
type WorkMineAdditionalReq struct {
	g.Meta       `path:"/teaching/teachingWork/mineAdditionalWork" method:"get" tags:"作品管理" summary:"我的附加作业作品"`
	AdditionalId string `json:"additionalId" dc:"附加作业ID"`
	PageNo       int    `json:"pageNo" dc:"页码" d:"1"`
	PageSize     int    `json:"pageSize" dc:"每页数量" d:"10"`
}

// WorkMineAdditionalRes 我的附加作业作品响应
type WorkMineAdditionalRes struct {
	g.Meta `mime:"application/json"`
	List   interface{} `json:"records"`
	Total  int         `json:"total"`
}