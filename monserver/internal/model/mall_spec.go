package model

// Spec 规格表
type Spec struct {
	Id          int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	CategoryId  int64  `gorm:"not null;column:category_id;type:int(11);default:0;comment:'分类id'"`
	GoodsId     int64  `gorm:"not null;column:goods_id;type:int(11);default:0;comment:'商品id'"`
	Goods       int64  `gorm:"not null;column:goods;type:int(11);default:0;comment:'是否关联商品 0:否 1:是'"`
	Name        string `gorm:"not null;column:name;type:varchar(255);default:'';comment:'规格名称'"`
	Main        int64  `gorm:"not null;column:main;type:int(11);default:0;comment:'是否主规格 0:否 1:是'"`
	EffectPrice int64  `gorm:"not null;column:price;type:int(11);default:0;comment:'是否会影响价格 0:不影响 1:影响'"`
	Sort        int64  `gorm:"not null;column:sort;type:int(11);default:0;comment:'排序'"`
	Remark      string `gorm:"not null;column:remark;type:varchar(255);default:'';comment:'备注'"`
	Status      int64  `gorm:"not null;column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Sorted      int64  `gorm:"not null;column:sorted;type:int(11);default:0;comment:'排序'"`
	Deleted     int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt    int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt    int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt    int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

func (s *Spec) TableName() string {
	return "mall_spec"
}

var SpecCol = struct {
	Id          string
	CategoryId  string
	GoodsId     string
	Goods       string
	Name        string
	Main        string
	EffectPrice string
	Sort        string
	Remark      string
	Status      string
	Sorted      string
	Deleted     string
	CreateAt    string
	UpdateAt    string
	DeleteAt    string
}{
	Id:          "id",
	CategoryId:  "category_id",
	GoodsId:     "goods_id",
	Goods:       "goods",
	Name:        "name",
	Main:        "main",
	EffectPrice: "price",
	Sort:        "sort",
	Remark:      "remark",
	Status:      "status",
	Sorted:      "sorted",
	Deleted:     "deleted",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
	DeleteAt:    "delete_at",
}
