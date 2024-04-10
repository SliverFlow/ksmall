package client

import (
	"github.com/SliverFlow/ksmall/app/user/http/internal/config"
	"github.com/SliverFlow/ksmall/app/user/rpc/pb"
	"github.com/SliverFlow/ksmall/core/rpc/client"
)

func NewUserRpcClient(c *config.Possess) pb.UserClient {
	conn, err := client.NewRpcClient(c.UserRpc)
	if err != nil {
		panic(err)
		return nil
	}
	return pb.NewUserClient(conn)
}
