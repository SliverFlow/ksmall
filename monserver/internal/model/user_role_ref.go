package model

// UserRoleRef 用户角色关联表
type UserRoleRef struct {
	Id     int64 `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	UserId int64 `gorm:"not null;column:user_id;type:int(11);default:0;comment:'用户id'"`
	RoleId int64 `gorm:"not null;column:role_id;type:int(11);default:0;comment:'角色id'"`
}

// TableName 数据表名
func (u *UserRoleRef) TableName() string {
	return "sys_user_role_ref"
}

var UserRoleRefCol = struct {
	Id     string
	UserId string
	RoleId string
}{
	Id:     "id",
	UserId: "user_id",
	RoleId: "role_id",
}
