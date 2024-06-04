package request

// CreateGoodReq defines the request payload for CreateGood
type CreateGoodReq struct {
	Name         string   `json:"name"`
	Free         *int64   `json:"free"`
	Price        float64  `json:"price"`
	CategoryId   int64    `json:"categoryId"`
	Desc         string   `json:"desc"`
	New          *int64   `json:"new"`
	Hot          *int64   `json:"hot"`
	OriginPrice  float64  `json:"originPrice"`
	Cover        string   `json:"cover"`
	Image        []string `json:"image"`
	ContentImage []string `json:"contentImage"`
}
