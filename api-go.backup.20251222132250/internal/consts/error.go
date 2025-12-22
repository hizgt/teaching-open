package consts

// 错误码定义

const (
	// 成功
	CodeSuccess = 0

	// 系统错误 1xxx
	CodeError            = 1000
	CodeInternalError    = 1001
	CodeInvalidParameter = 1002
	CodeDatabaseError    = 1003
	CodeCacheError       = 1004
	CodeFileError        = 1005

	// 认证授权错误 2xxx
	CodeUnauthorized      = 2001
	CodeTokenExpired      = 2002
	CodeTokenInvalid      = 2003
	CodePermissionDenied  = 2004
	CodeLoginFailed       = 2005
	CodeUserNotFound      = 2006
	CodePasswordError     = 2007
	CodeUserFrozen        = 2008
	CodeUserAlreadyExists = 2009

	// 业务错误 3xxx
	CodeRecordNotFound   = 3001
	CodeRecordExists     = 3002
	CodeDataInvalid      = 3003
	CodeOperationFailed  = 3004
	CodeFileUploadFailed = 3005
	CodeFileNotAllowed   = 3006
	CodeFileTooLarge     = 3007

	// 教学业务错误 4xxx
	CodeCourseNotFound     = 4001
	CodeCourseNotAvailable = 4002
	CodeWorkNotFound       = 4003
	CodeWorkAlreadySubmit  = 4004
	CodeWorkNotSubmit      = 4005
)

// ErrorMessages 错误消息映射
var ErrorMessages = map[int]string{
	CodeSuccess: "操作成功",

	CodeError:            "操作失败",
	CodeInternalError:    "系统内部错误",
	CodeInvalidParameter: "参数错误",
	CodeDatabaseError:    "数据库错误",
	CodeCacheError:       "缓存错误",
	CodeFileError:        "文件操作失败",

	CodeUnauthorized:      "未授权",
	CodeTokenExpired:      "Token已过期",
	CodeTokenInvalid:      "Token无效",
	CodePermissionDenied:  "权限不足",
	CodeLoginFailed:       "登录失败",
	CodeUserNotFound:      "用户不存在",
	CodePasswordError:     "密码错误",
	CodeUserFrozen:        "用户已被冻结",
	CodeUserAlreadyExists: "用户已存在",

	CodeRecordNotFound:   "记录不存在",
	CodeRecordExists:     "记录已存在",
	CodeDataInvalid:      "数据无效",
	CodeOperationFailed:  "操作失败",
	CodeFileUploadFailed: "文件上传失败",
	CodeFileNotAllowed:   "文件类型不允许",
	CodeFileTooLarge:     "文件大小超出限制",

	CodeCourseNotFound:     "课程不存在",
	CodeCourseNotAvailable: "课程不可用",
	CodeWorkNotFound:       "作品不存在",
	CodeWorkAlreadySubmit:  "作品已提交",
	CodeWorkNotSubmit:      "作品未提交",
}

// GetErrorMessage 获取错误消息
func GetErrorMessage(code int) string {
	if msg, ok := ErrorMessages[code]; ok {
		return msg
	}
	return ErrorMessages[CodeError]
}
