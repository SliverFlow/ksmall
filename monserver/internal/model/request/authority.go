package request

// CreateAuthorityReq 接口权限结构体
type CreateAuthorityReq struct {
	AuthorityGroupId int64  `json:"authority_group_id" binding:"required,min=1" message:"权限组id不能为空"`
	Name             string `json:"name" binding:"required" message:"权限名称不能为空"`
	Url              string `json:"url"  binding:"required" message:"路径不能为空"`
	Auth             *int64 `json:"auth" binding:"required,oneof=0 1" message:"是否鉴权不能为空"`
	Remark           string `json:"remark"`
	Sort             int64  `json:"sort"  binding:"required,min=1" message:"排序不能为空"`
}
