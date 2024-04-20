package reply

import "github.com/SliverFlow/ksmall/server/common"

// UserInfoReply  用户信息
type UserInfoReply struct {
	Id         int64  `json:"id"`
	Nickname   string `json:"nickname"`
	RoleId     int64  `json:"roleId"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	Male       int64  `json:"male"`
	Birthday   int64  `json:"birthday"`
	VIPLevel   int64  `json:"vipLevel"`
	Points     int64  `json:"points"`
	CreateTime int64  `json:"createdAt"`
}

type UserListReply struct {
	common.ListPageReply
}
