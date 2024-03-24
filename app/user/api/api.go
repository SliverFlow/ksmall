package api

import (
	"github.com/SliverFlow/ksmall/app/user/internal/service"
	"github.com/gin-gonic/gin"
)

type Group struct {
	userService *service.UserService
}

func NewApiGroup(userService *service.UserService) *Group {
	return &Group{userService: userService}
}

func (api *Group) InitApi(r *gin.Engine) {
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "order",
		})
	})
}
