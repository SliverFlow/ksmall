package request

// RoleCreateReq 角色创建请求入参
type RoleCreateReq struct {
	Name   string `json:"name" binding:"required" message:"角色名称不能为空"`
	Remark string `json:"remark"`
	Sorted int64  `json:"sorted" binding:"required;min=1;max=1000" message:"排序不能为空,且在1-1000之间"`
	Key    int64  `json:"key" binding:"required;min=1;max=1000" message:"角色编号不能为空,且在1-1000之间"`
}

// RoleUpdateReq 角色更新请求入参
type RoleUpdateReq struct {
	Id     int64  `json:"id" binding:"required;min=1" message:"角色ID不能为空"`
	Name   string `json:"name" binding:"required" message:"角色名称不能为空"`
	Remark string `json:"remark"`
	Sorted int64  `json:"sorted" binding:"required;min=1;max=1000" message:"排序不能为空,且在1-1000之间"`
	Key    int64  `json:"key" binding:"required;min=1;max=1000" message:"角色编号不能为空,且在1-1000之间"`
	Status *int64 `json:"status" binding:"required" message:"状态不能为空"`
}

// RoleListReq 角色列表请求入参
type RoleListReq struct {
	Name   string `json:"name"`
	Status *int64 `json:"status"`
}
