package reply

// AuthGroupHasAuthorityReply 权限组含有权限列表返回
type AuthGroupHasAuthorityReply struct {
	Id     int64  `json:"id"`     // 权限组ID
	Name   string `json:"name"`   // 权限组名称
	Remark string `json:"remark"` // 备注
	Status int64  `json:"status"` // 状态
	Sort   int64  `json:"sort"`   // 排序
	// RoleName      string            `json:"rolename"`      // 创建人角色名称
	// Username      string            `json:"username"`      // 创建人名称
	AuthorityList []*AuthorityReply `json:"authorityList"` // 权限列表
}

// AuthorityReply 权限列表返回
type AuthorityReply struct {
	Id     int64  `json:"id"`     // 权限ID
	Name   string `json:"name"`   // 权限名称
	Remark string `json:"remark"` // 备注
	Url    string `json:"url"`    // URL
	Auth   string `json:"auth"`   // 权限
	Status int64  `json:"status"` // 状态
	Sort   int64  `json:"sort"`   // 排序
	// RoleName string `json:"roleName"` // 创建人角色名称
	// Username string `json:"username"` // 创建人名称
}

// AuthorityFindReply 权限查询返回
type AuthorityFindReply struct {
	Id       int64  `json:"id"`       // 权限ID
	Name     string `json:"name"`     // 权限名称
	Remark   string `json:"remark"`   // 备注
	Url      string `json:"url"`      // URL
	Auth     string `json:"auth"`     // 权限
	Status   int64  `json:"status"`   // 状态
	Sort     int64  `json:"sort"`     // 排序
	RoleName string `json:"roleName"` // 创建人角色名称
	Username string `json:"username"` // 创建人名称
	CreateAt int64  `json:"createAt"`
}

// AuthorityGroupFindReply 权限组查询返回
type AuthorityGroupFindReply struct {
	Id       int64  `json:"id"`       // 权限ID
	Name     string `json:"name"`     // 权限名称
	Remark   string `json:"remark"`   // 备注
	Status   int64  `json:"status"`   // 状态
	Sort     int64  `json:"sort"`     // 排序
	RoleName string `json:"roleName"` // 创建人角色名称
	Username string `json:"username"` // 创建人名称
	CreateAt int64  `json:"createAt"`
}
