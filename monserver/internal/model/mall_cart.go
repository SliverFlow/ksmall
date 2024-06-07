package model

// Cart 购物车
type Cart struct {
	Id         int64   `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	UserId     int64   `gorm:"not null;column:user_id;type:int(11);default:0;index:idx_user_id;comment:'用户id'"`
	GoodsId    int64   `gorm:"not null;column:goods_id;type:int(11);default:0;comment:'商品id'"`
	GoodsNum   int64   `gorm:"not null;column:goods_num;type:int(11);default:0;comment:'商品数量'"`
	PriceCount float64 `gorm:"not null;column:price_count;type:int(11);default:0;comment:'价格小计'"`
	Checked    int64   `gorm:"not null;column:checked;type:int(11);default:0;comment:'是否选中 0:否 1:是'"`
	Deleted    int64   `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt   int64   `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt   int64   `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt   int64   `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

func (c *Cart) TableName() string {
	return "mall_cart"
}

const (
	CartChecked  = 1 // 是
	CartNotCheck = 0 // 否
)

var CartCol = struct {
	Id         string
	UserId     string
	GoodsId    string
	GoodsNum   string
	PriceCount string
	Checked    string
	Deleted    string
	CreateAt   string
	UpdateAt   string
	DeleteAt   string
}{
	Id:         "id",
	UserId:     "user_id",
	GoodsId:    "goods_id",
	GoodsNum:   "goods_num",
	PriceCount: "price_count",
	Checked:    "checked",
	Deleted:    "deleted",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
}
