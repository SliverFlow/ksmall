package model

// CouponTemplate 优惠券模板
type CouponTemplate struct {
	Id           int64   `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'自增主键'"`
	Title        string  `gorm:"column:title;type:varchar(255);default:'';comment:'优惠券标题'"`
	Code         string  `gorm:"column:code;type:varchar(255);default:'';comment:'优惠券兑换码'"`
	Info         string  `gorm:"column:info;type:varchar(255);default:'';comment:'优惠券描述'"`
	Image        string  `gorm:"column:image;type:varchar(255);default:'';comment:'优惠券图片'"`
	ValidStartAt int64   `gorm:"column:valid_start_at;type:int(11);default:0;comment:'有效期开始时间'"`
	ValidEndAt   int64   `gorm:"column:valid_end_at;type:int(11);default:0;comment:'有效期结束时间'"`
	ValidDays    int64   `gorm:"column:valid_days;type:int(11);default:0;comment:'领取后有效天数'"`
	Total        int64   `gorm:"column:total;type:int(11);default:0;comment:'优惠券总数'"`
	Remain       int64   `gorm:"column:remain;type:int(11);default:0;comment:'剩余优惠券数量'"`
	Discount     int64   `gorm:"column:discount;type:int(11);default:0;comment:'优惠券折扣'"`
	Min          float64 `gorm:"column:min;decimal(10,2);default:0.00;comment:'最小折扣金额'"`
	Max          float64 `gorm:"column:max;decimal(10,2);default:0.00;comment:'最大折扣金额'"`
	Limit        int64   `gorm:"column:limit;type:int(11);default:0;comment:'每人限领张数'"`
	Status       int64   `gorm:"column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Deleted      int64   `gorm:"column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt     int64   `gorm:"column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt     int64   `gorm:"column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt     int64   `gorm:"column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

// TableName 表名
func (c *CouponTemplate) TableName() string {
	return "mall_coupon_template"
}

var CouponTemplateCol = struct {
	Id           string
	Title        string
	Info         string
	Image        string
	ValidStartAt string
	ValidEndAt   string
	ValidDays    string
	Total        string
	Remain       string
	Discount     string
	Min          string
	Max          string
	Limit        string
	Status       string
	Deleted      string
	CreateAt     string
	UpdateAt     string
	DeleteAt     string
}{
	Id:           "id",
	Title:        "title",
	Info:         "info",
	Image:        "image",
	ValidStartAt: "valid_start_at",
	ValidEndAt:   "valid_end_at",
	ValidDays:    "valid_days",
	Total:        "total",
	Remain:       "remain",
	Discount:     "discount",
	Min:          "min",
	Max:          "max",
	Limit:        "limit",
	Status:       "status",
	Deleted:      "deleted",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeleteAt:     "delete_at",
}
