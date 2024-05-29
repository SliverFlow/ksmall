package model

type Banner struct {
	Id       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'自增主键'"`
	Url      string `gorm:"column:url;type:varchar(255);default:'';comment:'图片点击跳转链接'"`
	ImageUrl string `gorm:"column:image_url;type:varchar(255);default:'';comment:'图片链接'"`
	Position int64  `gorm:"column:position;type:int(11);default:0;comment:'位置 0:首页轮播图 1:首页活动位 2:分类位'"`
	Sort     int64  `gorm:"column:sort;type:int(11);default:0;comment:'排序'"`
	Status   int64  `gorm:"column:status;type:int(11);default:0;comment:'状态 0:禁用 1:启用'"`
	Deleted  int64  `gorm:"column:deleted;type:int(11);default:0;comment:'删除标志 0:未删除 1:已删除'"`
	CreateAt int64  `gorm:"column:create_at;type:int(11);default:0;comment:'创建时间'"`
	UpdateAt int64  `gorm:"column:update_at;type:int(11);default:0;comment:'更新时间'"`
	DeleteAt int64  `gorm:"column:delete_at;type:int(11);default:0;comment:'删除时间'"`
}

func (b *Banner) TableName() string {
	return "mall_banner"
}

const (
	BannerPositionIndex = 0 // 首页轮播图
	BannerPositionHome  = 1 // 首页活动位
	BannerPositionCate  = 2 // 分类位
)

var BannerCol = struct {
	Id       string
	Url      string
	ImageUrl string
	Position string
	Sort     string
	Status   string
	Deleted  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}{
	Id:       "id",
	Url:      "url",
	ImageUrl: "image_url",
	Position: "position",
	Sort:     "sort",
	Status:   "status",
	Deleted:  "deleted",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}
