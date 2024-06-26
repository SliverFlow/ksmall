package constant

func init() {
	MessageCodeMap = make(map[uint32]string)
	MessageCodeMap[ServerErrorCode] = "服务器错误"
	MessageCodeMap[RequestFailedCode] = "请求失败"
	MessageCodeMap[NotFoundCode] = "未找到"
	MessageCodeMap[InvalidParamsCode] = "请求参数错误"
	MessageCodeMap[InvalidParamsUuidCode] = "请求参数错误"
	MessageCodeMap[UnauthorizedCode] = "未授权"
	MessageCodeMap[ForbiddenCode] = "禁止访问"
	MessageCodeMap[ExpiredCode] = "过期"
	MessageCodeMap[InternalErrorCode] = "内部错误"
	MessageCodeMap[UserNotFoundCode] = "用户未找到"
	MessageCodeMap[UserExistCode] = "用户已存在"
	MessageCodeMap[UserPasswordCode] = "用户密码错误"
	MessageCodeMap[TokenInvalidCode] = "token无效"
	MessageCodeMap[TokenExpiredCode] = "token过期"
	MessageCodeMap[UserLoginExpCode] = "用户登录过期"
	MessageCodeMap[UserCreateErrorCode] = "用户创建失败"
	MessageCodeMap[RequestTimeoutCode] = "请求超时，请刷新再使用。。。"
	MessageCodeMap[CaptchaTypeCode] = "验证码类型错误"
	MessageCodeMap[UserLoginTypeCode] = "用户登录类型错误"
	MessageCodeMap[ServerInternalCode] = "服务器内部错误"

}

var MessageCodeMap map[uint32]string
