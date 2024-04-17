package request

// LoginByAccountReq  账号登录入参
type LoginByAccountReq struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
