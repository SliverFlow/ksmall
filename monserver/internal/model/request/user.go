package request

// UserCreateReq 用户创建请求
type UserCreateReq struct {
	Username string `json:"username" binding:"required" message:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
	RoleId   string `json:"roleId" binding:"required" message:"角色ID不能为空"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Male     *int64 `json:"male" binding:"required,max=2" message:"性别不能为空"`
	Avatar   string `json:"avatar"`
}

// UpdateUserReq 用户更新入参
type UpdateUserReq struct {
	Id       int64  `json:"id" binding:"required,min=2" message:"用户id不能为空"`
	Username string `json:"username" binding:"required" message:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
	RoleId   string `json:"roleId" binding:"required" message:"角色ID不能为空"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Male     *int64 `json:"male" binding:"required,max=2" message:"性别不能为空"`
	Avatar   string `json:"avatar"`
}
