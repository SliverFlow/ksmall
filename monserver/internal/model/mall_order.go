package model

// Order 订单
type Order struct {
	Id          int64   `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	UserId      int64   `gorm:"not null;column:user_id;type:int(11);default:0;index:idx_user_id;comment:'用户id'"`
	OrderSn     string  `gorm:"not null;column:order_sn;type:varchar(255);default:'';comment:'订单号'"`
	OrderPrice  float64 `gorm:"not null;column:order_price;type:decimal(10,2);default:0.00;comment:'订单价格'"`
	OrderStatus int64   `gorm:"not null;column:order_status;type:int(11);default:0;comment:'订单状态 0:待付款 1:已付款 2:已发货 3:已完成 4:已关闭'"`

	// 支付信息
	PaySn     string  `gorm:"not null;column:pay_sn;type:varchar(255);default:'';comment:'支付单号'"`
	PayPrice  float64 `gorm:"not null;column:pay_price;type:decimal(10,2);default:0.00;comment:'支付价格'"`
	PayStatus int64   `gorm:"not null;column:pay_status;type:int(11);default:0;comment:'支付状态 0:未支付 1:已支付'"`
	PayType   int64   `gorm:"not null;column:pay_type;type:int(11);default:0;comment:'支付方式 0:未支付 1:支付宝 2:微信 3:银行卡'"`

	// 收货信息
	ReceiverName     string `gorm:"not null;column:receiver_name;type:varchar(255);default:'';comment:'收货人姓名'"`
	ReceiverMobile   string `gorm:"not null;column:receiver_mobile;type:varchar(255);default:'';comment:'收货人手机号'"`
	ReceiverAddress  string `gorm:"not null;column:receiver_address;type:varchar(255);default:'';comment:'收货人详细地址地址'"`
	ReceiverProvince string `gorm:"not null;column:receiver_province;type:varchar(255);default:'';comment:'收货人省'"`
	ReceiverCity     string `gorm:"not null;column:receiver_city;type:varchar(255);default:'';comment:'收货人市'"`
	ReceiverArea     string `gorm:"not null;column:receiver_area;type:varchar(255);default:'';comment:'收货人区'"`

	// 优惠券信息
	CouponId int64  `gorm:"not null;column:coupon_id;type:int(11);default:0;comment:'优惠券id'"`
	CouponSn string `gorm:"not null;column:coupon_sn;type:varchar(255);default:'';comment:'优惠券码'"`

	Deleted  int64 `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64 `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64 `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64 `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

// TableName 表名
func (o *Order) TableName() string {
	return "mall_order"
}

var OrderCol = struct {
	Id               string
	UserId           string
	OrderSn          string
	OrderPrice       string
	OrderStatus      string
	PaySn            string
	PayPrice         string
	PayStatus        string
	PayType          string
	ReceiverName     string
	ReceiverMobile   string
	ReceiverAddress  string
	ReceiverProvince string
	ReceiverCity     string
	ReceiverArea     string
	Deleted          string
	CreateAt         string
	UpdateAt         string
	DeleteAt         string
}{
	Id:               "id",
	UserId:           "user_id",
	OrderSn:          "order_sn",
	OrderPrice:       "order_price",
	OrderStatus:      "order_status",
	PaySn:            "pay_sn",
	PayPrice:         "pay_price",
	PayStatus:        "pay_status",
	PayType:          "pay_type",
	ReceiverName:     "receiver_name",
	ReceiverMobile:   "receiver_mobile",
	ReceiverAddress:  "receiver_address",
	ReceiverProvince: "receiver_province",
	ReceiverCity:     "receiver_city",
	ReceiverArea:     "receiver_area",
	Deleted:          "deleted",
	CreateAt:         "create_at",
	UpdateAt:         "update_at",
	DeleteAt:         "delete_at",
}
