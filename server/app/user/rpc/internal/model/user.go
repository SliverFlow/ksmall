package model

const (
	UserStatusNormal  = 0 // 正常
	UserStatusLock    = 1 // 锁定
	UserStatusDisable = 2 // 禁用
	UserMaleMan       = 1 // 男
	UserMaleWoman     = 2 // 女
)

type User struct {
	Id            int64  // id
	Uuid          string // uuid
	WxOpenId      string // 微信 open id
	Nickname      string // 昵称
	Male          int64  // 1 男 2 女
	Email         string // 邮箱
	Phone         string //
	RoleId        int64  // 角色 id
	Avatar        string // 头像
	Birthday      int64  // 生日
	VIPLevel      int64  // 会员等级
	Points        int64  // 积分
	LastLoginTime int64  // 最后登录时间
	Status        int64  // 0 正常 1 锁定 2 禁用
	DeleteFlag    int64  // 0 未删除 1 已删除
	CreateTime    int64  // 创建时间
	UpdateTime    int64  // 更新时间
	DeleteTime    int64  // 删除时间
}

func (m *User) TableName() string {
	return "sys_user"
}

var UserCol = struct {
	Id            string
	Uuid          string
	WxOpenId      string
	Nickname      string
	Password      string
	Male          string
	Email         string
	Phone         string
	RoleId        string
	Avatar        string
	Birthday      string
	VIPLevel      string
	Points        string
	LastLoginTime string
	Status        string
	DeleteFlag    string
	CreateTime    string
	UpdateTime    string
	DeleteTime    string
}{
	Id:            "id",
	Uuid:          "uuid",
	WxOpenId:      "wx_open_id",
	Nickname:      "nickname",
	Password:      "password",
	Male:          "mail",
	Email:         "email",
	Phone:         "phone",
	RoleId:        "role_id",
	Avatar:        "avatar",
	Birthday:      "birthday",
	VIPLevel:      "vip_level",
	Points:        "points",
	LastLoginTime: "last_login_time",
	Status:        "status",
	DeleteFlag:    "delete_flag",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	DeleteTime:    "delete_time",
}
