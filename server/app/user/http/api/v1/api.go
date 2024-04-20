package v1

import (
	"github.com/SliverFlow/ksmall/server/app/user/http/internal/service"
	"github.com/SliverFlow/ksmall/server/core/middleware"
	"github.com/gin-gonic/gin"
)

type Group struct {
	userService *service.UserService

	cors    *middleware.Cors
	timeout *middleware.Timeout
}

func NewApiGroup(
	userService *service.UserService,
	cors *middleware.Cors,
	timeout *middleware.Timeout,
) *Group {
	return &Group{
		userService: userService,
		cors:        cors,
		timeout:     timeout,
	}
}

// InitApi 初始化api
func (api *Group) InitApi(r *gin.Engine) {
	v1 := r.Group("api/v1")
	userGroup := v1.Group("user", api.cors.Handle(), api.timeout.Handle())
	api.userService.InitUserRouter(userGroup)
}
