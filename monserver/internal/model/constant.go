package model

const (
	Deleted    = 1 // 已删除
	NotDeleted = 0 // 未删除
	Enable     = 1 // 启用
	Disable    = 0 // 禁用
)

var StatusMap = map[int64]string{
	Enable:  "启用",
	Disable: "禁用",
}
