package request

// CreateAuthorityGroupReq 接口权限结构体
type CreateAuthorityGroupReq struct {
	Name   string `json:"name"  binding:"required" message:"权限组名称不能为空"`
	Remark string `json:"remark"`
	Sort   int64  `json:"sort"  binding:"required,min=1" message:"排序不能为空"`
}

// UpdateAuthorityGroupReq 更新接口权限结构体
type UpdateAuthorityGroupReq struct {
	Id     int64  `json:"id" binding:"required" message:"权限组id不能为空"`
	Name   string `json:"name"  binding:"required" message:"权限组名称不能为空"`
	Remark string `json:"remark"`
	Sort   int64  `json:"sort"  binding:"required,min=1" message:"排序不能为空"`
	Status *int64 `json:"status" binding:"omitempty,oneof=0 1" message:"状态值不正确"`
}
