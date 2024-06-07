package model

// Menu 菜单表
type Menu struct {
	Id        int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	RoleId    int64  `gorm:"not null;column:role_id;type:int(11);default:0;comment:'角色id'"`
	ParentId  int64  `gorm:"not null;column:parent_id;type:int(11);default:0;comment:'父级菜单id'"`
	Path      string `gorm:"not null;column:path;type:varchar(255);default:'';comment:'菜单路径'"`
	Name      string `gorm:"not null;column:name;type:varchar(255);default:'';comment:'菜单名称'"`
	Hidden    int64  `gorm:"not null;column:hidden;type:int(11);default:0;comment:'是否隐藏 0:否 1:是'"`
	Component string `gorm:"not null;column:component;type:varchar(255);default:'';comment:'组件路径'"`
	Sort      int64  `gorm:"not null;column:sort;type:int(11);default:0;comment:'排序'"`
	Icon      string `gorm:"not null;column:icon;type:varchar(255);default:'';comment:'图标'"`
	Deleted   int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt  int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt  int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt  int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
	Meta      `gorm:"embedded;comment:附加属性"`
}

// Meta 菜单元信息
type Meta struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// TableName 数据表名
func (m *Menu) TableName() string {
	return "sys_menu"
}

var MenuCol = struct {
	Id        string
	RoleId    string
	ParentId  string
	Path      string
	Name      string
	Hidden    string
	Component string
	Sort      string
	Icon      string
	Deleted   string
	CreateAt  string
	UpdateAt  string
	DeleteAt  string
}{
	Id:        "id",
	RoleId:    "role_id",
	ParentId:  "parent_id",
	Path:      "path",
	Name:      "name",
	Hidden:    "hidden",
	Component: "component",
	Sort:      "sort",
	Icon:      "icon",
	Deleted:   "deleted",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeleteAt:  "delete_at",
}
