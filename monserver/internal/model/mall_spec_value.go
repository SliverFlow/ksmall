package model

// SpecValue 规格值表
type SpecValue struct {
	Id       int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	SpecId   int64  `gorm:"not null;column:spec_id;type:int(11);default:0;comment:'规格id'"`
	GoodsId  int64  `gorm:"not null;column:goods_id;type:int(11);default:0;index:idx_goods_id;comment:'商品id'"`
	Name     string `gorm:"not null;column:name;type:varchar(255);default:'';comment:'规格值名称'"`
	Default  int64  `gorm:"not null;column:default;type:int(11);default:0;comment:'是否默认 0:否 1:是'"`
	Remark   string `gorm:"not null;column:remark;type:varchar(255);default:'';comment:'备注'"`
	Status   int64  `gorm:"not null;column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Sort     int64  `gorm:"not null;column:sort;type:int(11);default:0;comment:'排序'"`
	Deleted  int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

func (s *SpecValue) TableName() string {
	return "mall_spec_value"
}

var SpecValueCol = struct {
	Id       string
	SpecId   string
	GoodsId  string
	Name     string
	Default  string
	Remark   string
	Status   string
	Sort     string
	Deleted  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Id:       "id",
	SpecId:   "spec_id",
	GoodsId:  "goods_id",
	Name:     "name",
	Default:  "default",
	Remark:   "remark",
	Status:   "status",
	Sort:     "sort",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
