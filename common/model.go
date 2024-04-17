package common

// IdReq
// @Author: [github.com/SliverFlow]
// @Desc: id 入参
type IdReq struct {
	Id string `json:"id"`
}

// ListPageReq
// @Author: [github.com/SliverFlow]
// @Desc: 分页列表入参
type ListPageReq struct {
	Page     int64  `json:"page"`
	PageSize int64  `json:"pageSize"`
	KeyWord  string `json:"keyWord"`
}

// GetLimitAndOffset
// @Author: [github.com/SliverFlow]
// @Desc: 获取分页列表的 limit 和 offset
func (l *ListPageReq) GetLimitAndOffset() (limit, offset int64) {
	if l.PageSize <= 0 {
		l.PageSize = 10
	}
	if l.Page <= 0 {
		l.Page = 1
	}
	limit = l.PageSize
	offset = (l.Page - 1) * l.PageSize
	return
}

// ListPageReply
// @Author: [github.com/SliverFlow]
// @Desc: 分页列表出参
type ListPageReply struct {
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
	Page     int64       `json:"page"`
	PageSize int64       `json:"pageSize"`
}
