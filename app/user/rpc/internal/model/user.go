package model

const (
	UserStatusNormal  = 0 // 正常
	UserStatusLock    = 1 // 锁定
	UserStatusDisable = 2 // 禁用
)

type User struct {
	Id         int64
	Uuid       string
	WxOpenId   string
	Username   string
	Password   string
	Email      string
	Phone      string
	RoleId     int64
	Avatar     string
	Status     int64
	DeleteFlag int64
	CreateTime int64
	UpdateTime int64
	DeleteTime int64
}

func (m *User) TableName() string {
	return "sys_user"
}

var UserCol = struct {
	Id         string
	Uuid       string
	WxOpenId   string
	Username   string
	Password   string
	Email      string
	Phone      string
	RoleId     string
	Avatar     string
	Status     string
	DeleteFlag string
	CreateTime string
	UpdateTime string
	DeleteTime string
}{
	Id:         "id",
	Uuid:       "uuid",
	WxOpenId:   "wx_open_id",
	Username:   "username",
	Password:   "password",
	Email:      "email",
	Phone:      "phone",
	RoleId:     "role_id",
	Avatar:     "avatar",
	Status:     "status",
	DeleteFlag: "delete_flag",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}
