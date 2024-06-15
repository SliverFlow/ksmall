package request

// CreateAuthorityReq 接口权限结构体
type CreateAuthorityReq struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
