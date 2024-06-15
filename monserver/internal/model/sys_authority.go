package model

// Authority 权限表
type Authority struct {
	Id               int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	UserId           int64  `gorm:"not null;column:user_id;type:int(11);default:0;comment:'创建者用户id'"`
	AuthorityGroupId int64  `gorm:"not null;column:authority_group_id;type:int(11);default:0;comment:'权限组id'"`
	Name             string `gorm:"not null;column:name;type:varchar(255);default:'';comment:'权限名称'"`
	Url              string `gorm:"not null;column:url;type:varchar(255);default:'';comment:'权限url'"`
	Auth             int64  `gorm:"not null;column:auth;type:int(11);default:0;comment:'是否鉴权 0:否 1:是'"`
	Remark           string `gorm:"not null;column:remark;type:varchar(255);default:'';comment:'备注'"`
	Status           int64  `gorm:"not null;column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Sort             int64  `gorm:"not null;column:sort;type:int(11);default:0;comment:'排序'"`
	Deleted          int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt         int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt         int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt         int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

// TableName 数据表名
func (a *Authority) TableName() string {
	return "sys_authority"
}

const (
	AuthorityAuthYes = 1
	AuthorityAuthNo  = 0
)

var AuthorityCol = struct {
	Id       string
	Name     string
	Url      string
	Auth     string
	Remark   string
	Status   string
	Sort     string
	Deleted  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Id:       "id",
	Name:     "name",
	Url:      "url",
	Auth:     "auth",
	Remark:   "remark",
	Status:   "status",
	Sort:     "sort",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
