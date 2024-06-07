package model

// OrderItem 订单子表
type OrderItem struct {
	Id         int64   `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	OrderId    int64   `gorm:"not null;column:order_id;type:int(11);default:0;index:idx_order_id;comment:'订单id'"`
	GoodsId    int64   `gorm:"not null;column:goods_id;type:int(11);default:0;comment:'商品id'"`
	GoodsNum   int64   `gorm:"not null;column:goods_num;type:int(11);default:0;comment:'商品数量'"`
	GoodsCover string  `gorm:"not null;column:goods_cover;type:varchar(255);default:'';comment:'商品封面'"`
	GoodsPrice float64 `gorm:"not null;column:goods_price;type:decimal(10,2);default:0.00;comment:'商品价格'"`
	Deleted    int64   `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt   int64   `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt   int64   `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt   int64   `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

// TableName 表名
func (o *OrderItem) TableName() string {
	return "mall_order_item"
}

var OrderItemCol = struct {
	Id         string
	OrderId    string
	GoodsId    string
	GoodsNum   string
	GoodsCover string
	GoodsPrice string
	Deleted    string
	CreateAt   string
	UpdateAt   string
	DeleteAt   string
}{
	Id:         "id",
	OrderId:    "order_id",
	GoodsId:    "goods_id",
	GoodsNum:   "goods_num",
	GoodsCover: "goods_cover",
	GoodsPrice: "goods_price",
	Deleted:    "deleted",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
}
