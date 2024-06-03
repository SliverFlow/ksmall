package request

// CreateCategoryReq 创建分类入参
type CreateCategoryReq struct {
	Name    string `json:"name" binding:"required" message:"分类名称不能为空"`
	ParenId *int64 `json:"parent_id" binding:"required" message:"父分类id不能为空"`
	Icon    string `json:"icon"`
	IsIndex *int64 `json:"is_index"`
	Sort    int64  `json:"sort" binding:"required,min=1" message:"排序不能为空"`
}
