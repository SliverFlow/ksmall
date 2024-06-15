package request

// CreateAuthorityReq 接口权限结构体
type CreateAuthorityReq struct {
	Name string `json:"name"`
	Url  string `json:"url"  binding:"required,min=1" message:"路径不能为空"`
}
