package consts

// 系统常量定义

const (
	// 上下文键
	CtxKeyUserId   = "userId"
	CtxKeyUsername = "username"
	CtxKeyRealname = "realname"
	CtxKeyToken    = "token"
	CtxKeyOrgCode  = "orgCode"

	// 缓存键前缀
	CacheKeyPrefix     = "teaching_open:"
	CacheKeyUser       = CacheKeyPrefix + "user:"
	CacheKeyRole       = CacheKeyPrefix + "role:"
	CacheKeyPermission = CacheKeyPrefix + "permission:"
	CacheKeyDict       = CacheKeyPrefix + "dict:"
	CacheKeyToken      = CacheKeyPrefix + "token:"

	// Token相关
	TokenHeader = "X-Token"
	TokenPrefix = "Bearer "

	// 用户状态
	UserStatusNormal = 1 // 正常
	UserStatusFrozen = 2 // 冻结

	// 删除标记
	DelFlagNormal  = 0 // 正常
	DelFlagDeleted = 1 // 已删除

	// 数据权限类型
	DataRuleSelf    = 1 // 仅本人
	DataRuleDept    = 2 // 本部门
	DataRuleDeptSub = 3 // 本部门及子部门
	DataRuleAll     = 4 // 全部

	// 菜单类型
	MenuTypeMenu   = 0 // 菜单
	MenuTypeButton = 1 // 按钮
	MenuTypeLink   = 2 // 外链

	// 角色类型
	RoleTypeAdmin = 1 // 超级管理员
	RoleTypeUser  = 2 // 普通用户

	// 作品状态
	WorkStatusDraft     = 0 // 草稿
	WorkStatusSubmitted = 1 // 已提交
	WorkStatusCorrected = 2 // 已批改

	// 课程状态
	CourseStatusDraft     = 0 // 草稿
	CourseStatusPublished = 1 // 已发布
	CourseStatusOffline   = 2 // 已下架

	// 默认页码
	DefaultPage     = 1
	DefaultPageSize = 10
	MaxPageSize     = 100
)

// 操作类型
const (
	OperationLogin  = "login"
	OperationLogout = "logout"
	OperationAdd    = "add"
	OperationEdit   = "edit"
	OperationDelete = "delete"
	OperationExport = "export"
	OperationImport = "import"
)
