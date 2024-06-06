package model

// Stock 库存
type Stock struct {
	Id       int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	Name     string `gorm:"not null;column:name;type:varchar(255);default:'';comment:'库存名称'"`
	Remark   string `gorm:"not null;column:remark;type:varchar(255);default:'';comment:'备注'"`
	GoodsId  int64  `gorm:"not null;column:goods_id;type:int(11);default:0;index:idx_good_id;comment:'商品id'"`
	Num      int64  `gorm:"not null;column:num;type:int(11);default:0;comment:'库存数量'"`
	Version  int64  `gorm:"not null;column:version;type:int(11);default:0;comment:'版本号'"`
	Active   int64  `gorm:"not null;column:active;type:int(11);default:0;comment:'激活状态 0:未激活 1:已激活'"`
	Status   int64  `gorm:"not null;column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Deleted  int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

func (s *Stock) TableName() string {
	return "mall_stock"
}

const (
	StockActive    = 1 // 正在使用的库存
	StockNotActive = 0 // 未使用的库存
)

var StockCol = struct {
	Id       string
	Name     string
	Remark   string
	GoodsId  string
	Num      string
	Version  string
	Status   string
	Deleted  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Id:       "id",
	Name:     "name",
	Remark:   "remark",
	GoodsId:  "goods_id",
	Num:      "num",
	Version:  "version",
	Status:   "status",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
