package request

// IdReq id请求入参
type IdReq struct {
	Id int64 `json:"id" binding:"required,min=1" message:"id不能为空"`
}

// IdsReq ids请求入参
type IdsReq struct {
	Ids []int64 `json:"ids" binding:"required,min=1" message:"ids不能为空"`
}

// PageReq 分页请求入参
type PageReq struct {
	Page     int64  `json:"page" binding:"required,min=1" message:"page不能为空,且大于0"`
	PageSize int64  `json:"page_size" binding:"required,min=1" message:"page_size不能为空,且大于0"`
	Keyword  string `json:"keyword"`
}
