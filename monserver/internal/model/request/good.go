package request

// CreateGoodReq defines the request payload for CreateGood
type CreateGoodReq struct {
	Name         string   `json:"name" binding:"required" message:"商品名称不能为空"`
	Free         *int64   `json:"free"`
	Price        float64  `json:"price" binding:"required,min=1" message:"商品售卖价格不能为空"`
	CategoryId   int64    `json:"categoryId" binding:"required,min=1" message:"分类id不能为空"`
	Desc         string   `json:"desc"`
	New          *int64   `json:"new"`
	Hot          *int64   `json:"hot"`
	OriginPrice  float64  `json:"originPrice" binding:"required,min=1" message:"商品原价不能为空"`
	Cover        string   `json:"cover"`
	Image        []string `json:"image"`
	ContentImage []string `json:"contentImage"`
}
