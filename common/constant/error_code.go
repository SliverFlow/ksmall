package constant

const (
	ServerErrorCode       = 500    // 服务器错误
	RequestTimeoutCode    = 408    // 请求超时
	RequestFailedCode     = 100000 // 请求失败
	NotFoundCode          = 100001 // 未找到
	InvalidParamsCode     = 100002 // 请求参数错误
	UnauthorizedCode      = 100003 // 未授权
	ForbiddenCode         = 100004 // 禁止访问
	ExpiredCode           = 100005 // 过期
	InternalErrorCode     = 100006 // 内部错误
	InvalidParamsUuidCode = 100007 // 请求参数uuid错误
)

const (
	UserNotFoundCode    = 200000 // 用户未找到
	UserExistCode       = 200001 // 用户已存在
	UserPasswordCode    = 200002 // 用户密码错误
	UserLoginExpCode    = 200003 // 用户登录过期
	UserCreateErrorCode = 200004 // 用户创建失败
)

const (
	TokenInvalidCode = 300000 // token无效
	TokenExpiredCode = 300001 // token过期
)
