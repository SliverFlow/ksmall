package data

// Address 收货地址
type Address struct {
	Id         int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	UserId     int64  `gorm:"not null;column:user_id;type:int(11);default:0;index:idx_user_id;comment:'用户id'"`
	Username   string `gorm:"not null;column:user_name;type:varchar(255);default:'';comment:'收货人姓名'"`
	UserMobile string `gorm:"not null;column:user_mobile;type:varchar(255);default:'';comment:'收货人手机号'"`
	Province   string `gorm:"not null;column:province;type:varchar(255);default:'';comment:'省'"`
	City       string `gorm:"not null;column:city;type:varchar(255);default:'';comment:'市'"`
	Area       string `gorm:"not null;column:area;type:varchar(255);default:'';comment:'区'"`
	Address    string `gorm:"not null;column:address;type:varchar(255);default:'';comment:'详细地址'"`
	Default    int64  `gorm:"not null;column:default;type:int(11);default:0;comment:'是否默认 0:否 1:是'"`
	Deleted    int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt   int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt   int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt   int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

// TableName 表名
func (a *Address) TableName() string {
	return "mall_address"
}
