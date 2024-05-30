package model

type User struct {
	Id       int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	RoleId   int64  `gorm:"not null;column:role_id;type:int(11);default:0;comment:'角色id'"`
	Uuid     string `gorm:"not null;column:uuid;type:varchar(255);default:'';comment:'uuid'"`
	Male     int64  `gorm:"not null;column:male;type:int(11);default:0;comment:'性别 0:未知 1:男 2:女'"`
	WxOpenid string `gorm:"not null;column:wx_openid;type:varchar(255);default:'';comment:'微信openid'"`
	Username string `gorm:"not null;column:username;type:varchar(255);default:'';comment:'用户名'"`
	Nickname string `gorm:"not null;column:nickname;type:varchar(255);default:'';comment:'昵称'"`
	Password string `gorm:"not null;column:password;type:varchar(255);default:'';comment:'密码'"`
	Email    string `gorm:"not null;column:email;type:varchar(255);default:'';comment:'邮箱'"`
	Phone    string `gorm:"not null;column:phone;type:varchar(255);default:'';comment:'手机号'"`
	Avatar   string `gorm:"not null;column:avatar;type:varchar(255);default:'';comment:'头像'"`
	Status   int64  `gorm:"not null;column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Deleted  int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

func (u *User) TableName() string {
	return "user"
}

var UserCol = struct {
	Id       string
	RoleId   string
	Uuid     string
	Male     string
	WxOpenid string
	Username string
	Nickname string
	Password string
	Email    string
	Phone    string
	Avatar   string
	Status   string
	Deleted  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Id:       "id",
	RoleId:   "role_id",
	Uuid:     "uuid",
	Male:     "male",
	WxOpenid: "wx_openid",
	Username: "username",
	Nickname: "nickname",
	Password: "password",
	Email:    "email",
	Phone:    "phone",
	Avatar:   "avatar",
	Status:   "status",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
