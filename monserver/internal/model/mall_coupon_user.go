package model

// CouponUser 优惠券用户
type CouponUser struct {
	Id         int64   `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'自增主键'"`
	UserId     int64   `gorm:"column:user_id;type:int(11);default:0;comment:'用户id'"`
	CouponId   int64   `gorm:"column:coupon_id;type:int(11);default:0;comment:'优惠券id'"`
	CouponCode string  `gorm:"column:coupon_code;type:varchar(255);default:'';comment:'优惠券码'"`
	OrderId    int64   `gorm:"column:order_id;type:int(11);default:0;comment:'订单id'"`
	OrderSn    string  `gorm:"column:order_sn;type:varchar(255);default:'';comment:'订单号'"`
	UsedTime   int64   `gorm:"column:used_time;type:int(11);default:0;comment:'使用时间'"`
	UseStatus  int64   `gorm:"column:use_status;type:int(11);default:0;comment:'使用状态 0:未使用 1:已使用 2:已过期'"`
	Discount   float64 `gorm:"column:discount;decimal(10,2);default:0.00;comment:'折扣金额'"`
	Deleted    int64   `gorm:"column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt   int64   `gorm:"column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt   int64   `gorm:"column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt   int64   `gorm:"column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}
