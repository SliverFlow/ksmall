package model

// RoleAuthRef 角色权限关联表
type RoleAuthRef struct {
	Id          int64 `gorm:"not null;column:id;primary_key;AUTO_INCREMENT"`
	RoleId      int64 `gorm:"not null;column:role_id;type:int(11);default:0;comment:'角色id'"`
	AuthorityId int64 `gorm:"not null;column:authority_id;type:int(11);default:0;comment:'权限id'"`
}

// TableName 数据表名
func (r *RoleAuthRef) TableName() string {
	return "sys_role_auth_ref"
}

var RoleAuthRefCol = struct {
	Id          string
	RoleId      string
	AuthorityId string
}{
	Id:          "id",
	RoleId:      "role_id",
	AuthorityId: "authority_id",
}
