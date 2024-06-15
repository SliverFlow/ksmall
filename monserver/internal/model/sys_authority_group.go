package model

type AuthorityGroup struct {
	Id       int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"not null;column:name;type:varchar(255);default:'';comment:'权限名称'"`
	Remark   string `gorm:"not null;column:remark;type:varchar(255);default:'';comment:'备注'"`
	Status   int64  `gorm:"not null;column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Sort     int64  `gorm:"not null;column:sort;type:int(11);default:0;comment:'排序'"`
	Deleted  int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

func (m *AuthorityGroup) TableName() string {
	return "sys_authority_group"
}

var AuthorityGroupCol = struct {
	Id       string
	Name     string
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
	Remark:   "remark",
	Status:   "status",
	Sort:     "sort",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
