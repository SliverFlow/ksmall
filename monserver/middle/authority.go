package middle

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Authority struct {
	iAuthority IAuthority
}

type IAuthority interface {
	// CheckAccess Check if the user has the right to access the resource
	CheckAccess(ctx context.Context, userId int64, path string) (bool, error)
}

func NewAuthority(iAuthority IAuthority) *Authority {
	return &Authority{iAuthority: iAuthority}
}

func (a *Authority) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userId", int64(2))
		userId, _ := c.Get("userId")
		path := c.Request.URL.Path
		ctx := c.Request.Context()
		access, err := a.iAuthority.CheckAccess(ctx, userId.(int64), path)
		if err != nil {
			if !access {
				c.JSON(403, gin.H{"error": "Forbidden"})
				c.Abort()
				return
			}
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}
		c.Next()
	}
}
