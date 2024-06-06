package data

// Comment 评论
type Comment struct {
	Id       int64  `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	ParentId int64  `gorm:"not null;column:parent_id;type:int(11);default:0;comment:'父评论id'"`
	UserId   int64  `gorm:"not null;column:user_id;type:int(11);default:0;index:idx_user_id;comment:'用户id'"`
	OrderId  int64  `gorm:"not null;column:order_id;type:int(11);default:0;index:idx_order_id;comment:'订单id'"`
	GoodsId  int64  `gorm:"not null;column:goods_id;type:int(11);default:0;comment:'商品id'"`
	Content  string `gorm:"not null;column:content;type:varchar(255);default:'';comment:'评论内容'"`
	Images   string `gorm:"not null;column:images;type:varchar(255);default:'';comment:'评论图片'"`
	Level    int64  `gorm:"not null;column:level;type:int(11);default:0;comment:'评论等级 0:好评 1:中评 2:差评'"`
	Star     int64  `gorm:"not null;column:star;type:int(11);default:0;comment:'星级'"`
	Deleted  int64  `gorm:"not null;column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64  `gorm:"not null;column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64  `gorm:"not null;column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64  `gorm:"not null;column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

// TableName 表名
func (c *Comment) TableName() string {
	return "mall_comment"
}

var CommentCol = struct {
	Id       string
	ParentId string
	UserId   string
	OrderId  string
	GoodsId  string
	Content  string
	Images   string
	Level    string
	Star     string
	Deleted  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Id:       "id",
	ParentId: "parent_id",
	UserId:   "user_id",
	OrderId:  "order_id",
	GoodsId:  "goods_id",
	Content:  "content",
	Images:   "images",
	Level:    "level",
	Star:     "star",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
