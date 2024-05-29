package model

// Category 商品分类
type Category struct {
	Id       int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"not null;column:name;type:varchar(255);default:'';comment:'分类名称'"`
	ParentId int64  `gorm:"not null;column:parent_id;type:int(11);default:0;comment:'父分类id'"`
	Level    int64  `gorm:"not null;column:level;type:int(11);default:0;comment:'分类级别'"`
	Icon     string `gorm:"not null;column:icon;type:varchar(255);default:'';comment:'分类图标'"`
	IsIndex  int64  `gorm:"not null;column:is_index;type:int(11);default:0;comment:'是否首页显示 0:否 1:是'"`
	Sort     int64  `gorm:"not null;column:sort;type:int(11);default:0;comment:'排序'"`
	Status   int64  `gorm:"not null;column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Deleted  int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

func (c *Category) TableName() string {
	return "mall_category"
}

const (
	CategoryIsIndex  = 1 // 是
	CategoryNotIndex = 0 // 否
)

var CategoryCol = struct {
	Id       string
	Name     string
	ParentId string
	Level    string
	Icon     string
	IsIndex  string
	Sort     string
	Status   string
	Deleted  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Id:       "id",
	Name:     "name",
	ParentId: "parent_id",
	Level:    "level",
	Icon:     "icon",
	IsIndex:  "is_index",
	Sort:     "sort",
	Status:   "status",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
