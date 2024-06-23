package reply

// CategoryTreeListReply 分类树列表响应
type CategoryTreeListReply struct {
	Id       int64                    `json:"id"`
	Name     string                   `json:"name"`
	PatentId int64                    `json:"patentId"`
	Level    int64                    `json:"level"`
	Icon     string                   `json:"icon"`
	IsIndex  int64                    `json:"isIndex"`
	Sort     int64                    `json:"sort"`
	Children []*CategoryTreeListReply `json:"children"`
}
