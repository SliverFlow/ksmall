package api

import (
	"github.com/SliverFlow/ksmall/app/user/http/internal/service"
	"github.com/SliverFlow/ksmall/core/middleware"
	"github.com/gin-gonic/gin"
	"time"
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

func (api *Group) InitApi(r *gin.Engine) {

	router := r.Use(api.cors.Handle(), api.timeout.Handle())

	router.GET("/user", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.JSON(200, gin.H{
			"status": "order",
		})
	})

}
