package request

// LoginReq
// @Author: [github.com/SliverFlow]
// @Desc: 登录请求入参
type LoginReq struct {
	Account string `json:"account" binding:"required" msg:"账号不能为空"`
	Captcha string `json:"captcha" binding:"required" msg:"验证码不能为空"`
	Type    int64  `json:"type" binding:"required" msg:"类型不能为空"`
}
