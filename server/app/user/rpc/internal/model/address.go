package model

const (
	AddressStatusNormal = 0 // 正常
	AddressStatusLock   = 1 // 锁定
	AddressIsDefault    = 1 // 默认
	AddressNotDefault   = 0 // 非默认
)

type Address struct {
	Id            int64  // Id
	UserId        int64  // 用户ID
	ReceiverName  string // 收货人姓名
	ReceiverPhone string // 收货人电话
	Street        string // 街道
	City          string // 城市
	Province      string // 省份
	PostalCode    string // 邮政编码
	IsDefault     int64  // 是否为默认地址
	Status        int64  // 状态
	DeleteFlag    int64  // 删除标识
}

var AddressCol = struct {
	Id            string
	UserId        string
	ReceiverName  string
	ReceiverPhone string
	Street        string
	City          string
	Province      string
	PostalCode    string
	IsDefault     string
	Status        string
	DeleteFlag    string
}{
	Id:            "id",
	UserId:        "user_id",
	ReceiverName:  "receiver_name",
	ReceiverPhone: "receiver_phone",
	Street:        "street",
	City:          "city",
	Province:      "province",
	PostalCode:    "postal_code",
	IsDefault:     "is_default",
	Status:        "status",
	DeleteFlag:    "delete_flag",
}
