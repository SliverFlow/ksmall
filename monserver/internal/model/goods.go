package model

// Goods 商品
type Goods struct {
	Id           int64   `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	CategoryId   int64   `gorm:"not null;column:category_id;type:bigint(20);index:idx_category_id;comment:分类id"`
	Sn           string  `gorm:"not null;column:sn;type:varchar(255);comment:商品编号"`
	Name         string  `gorm:"not null;column:name;type:varchar(255);comment:商品名称"`
	Free         int64   `gorm:"not null;column:free;type:int(10);comment:是否免费"`
	New          int64   `gorm:"not null;column:new;type:int(10);comment:是否新品"`
	Hot          int64   `gorm:"not null;column:hot;type:int(10);comment:是否热销"`
	Desc         string  `gorm:"not null;column:desc;type:varchar(255);comment:商品描述"`
	ViewCount    int64   `gorm:"not null;column:view_count;type:int(10);comment:浏览次数"`
	BuyCount     int64   `gorm:"not null;column:buy_count;type:int(10);comment:购买次数"`
	FavCount     int64   `gorm:"not null;column:fav_count;type:int(10);comment:收藏次数"`
	Price        float64 `gorm:"not null;column:price;type:decimal(10,2);comment:售卖价格"`
	OriginPrice  float64 `gorm:"not null;column:origin_price;type:decimal(10,2);comment:原价"`
	Cover        string  `gorm:"not null;column:cover;type:varchar(255);comment:封面"`
	Image        string  `gorm:"not null;column:image;type:varchar(255);comment:图片列表"`
	ContentImage string  `gorm:"not null;column:content_image;type:varchar(255);comment:内容图片"`
	Status       int64   `gorm:"column:status;type:int(11);default:0;comment:'状态 0:禁用 1:暂存 2:上架 3:下架"`
	Deleted      int64   `gorm:"column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt     int64   `gorm:"column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt     int64   `gorm:"column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt     int64   `gorm:"column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

// TableName 表名
func (Goods) TableName() string {
	return "mall_goods"
}

const (
	GoodStatusDisable  = 0 // 禁用
	GoodStatusDraft    = 1 // 暂存
	GoodStatusOnShelf  = 2 // 上架
	GoodStatusOffShelf = 3 // 下架
)

var GoodCol = struct {
	Id           string
	CategoryId   string
	CategoryName string
	Sn           string
	Name         string
	Free         string
	New          string
	Hot          string
	Desc         string
	ViewCount    string
	BuyCount     string
	FavCount     string
	Price        string
	OriginPrice  string
	Cover        string
	Image        string
	ContentImage string
	Status       string
	Deleted      string
	CreateAt     string
	UpdateAt     string
	DeleteAt     string
}{
	Id:           "id",
	CategoryId:   "category_id",
	CategoryName: "category_name",
	Sn:           "sn",
	Name:         "name",
	Free:         "free",
	New:          "new",
	Hot:          "hot",
	Desc:         "desc",
	ViewCount:    "view_count",
	BuyCount:     "buy_count",
	FavCount:     "fav_count",
	Price:        "price",
	OriginPrice:  "origin_price",
	Cover:        "cover",
	Image:        "image",
	ContentImage: "content_image",
	Status:       "status",
	Deleted:      "deleted",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeleteAt:     "delete_at",
}
