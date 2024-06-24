package request

// CreateCategoryReq 创建分类入参
type CreateCategoryReq struct {
	Name    string `json:"name" binding:"required" message:"分类名称不能为空"`
	ParenId *int64 `json:"parentId" binding:"required" message:"父分类id不能为空"`
	Icon    string `json:"icon"`
	IsIndex *int64 `json:"isIndex"`
	Sort    int64  `json:"sort" binding:"required,min=1" message:"排序不能为空"`
}

// UpdateCategoryReq 更新分类入参
type UpdateCategoryReq struct {
	Id      int64  `json:"id" binding:"required,min=1" message:"分类id不能为空"`
	Name    string `json:"name" binding:"required" message:"分类名称不能为空"`
	ParenId *int64 `json:"parentId" binding:"required" message:"父分类id不能为空"`
	Icon    string `json:"icon"`
	IsIndex *int64 `json:"isIndex"`
	Sort    int64  `json:"sort" binding:"required,min=1" message:"排序不能为空"`
	Status  *int64 `json:"status" binding:"required,oneof=0 1"  message:"状态不能为空"`
}
