package api

import "github.com/gin-gonic/gin"

type Group struct {
}

func NewApiGroup() *Group {
	return &Group{}
}

func (api *Group) InitApi(r *gin.Engine) {
	r.GET("/category", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "order",
		})
	})
}
