package reply

import "github.com/SliverFlow/ksmall/common"

// UserInfoReply  用户信息
type UserInfoReply struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Male      int64  `json:"male"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserListReply struct {
	common.ListPageReply
}
