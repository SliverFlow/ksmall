package model

// CouponUseCondition 优惠券使用条件
type CouponUseCondition struct {
	Id       int64   `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'自增主键'"`
	CouponId int64   `gorm:"column:coupon_id;type:int(11);default:0;comment:'优惠券id'"`
	Min      float64 `gorm:"column:min;decimal(10,2);default:0.00;comment:'最小使用金额'"`
	Type     int64   `gorm:"column:type;type:int(11);default:0;comment:'使用类型 0:全场通用 1:指定商品可用 2:指定分类商品可用'"`
	Status   int64   `gorm:"column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Deleted  int64   `gorm:"column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64   `gorm:"column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64   `gorm:"column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64   `gorm:"column:delete_at;type:int(11);default:0;comment:'删除时间'"`
	Ref      `gorm:"embedded;comment:'关联信息'"`
}

type Ref struct {
	Id []int64
}

// TableName 表名
func (c *CouponUseCondition) TableName() string {
	return "mall_coupon_use_condition"
}

const (
	CouponUseConditionTypeAll      = 0
	CouponUseConditionTypeGoods    = 1
	CouponUseConditionTypeCategory = 2
)

var CouponUseConditionCol = struct {
	Id       string
	CouponId string
	Min      string
	Type     string
	Status   string
	Deleted  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Id:       "id",
	CouponId: "coupon_id",
	Min:      "min",
	Type:     "type",
	Status:   "status",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
