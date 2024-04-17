package rpc_client

import (
	"fmt"
	"github.com/SliverFlow/ksmall/app/user/http/internal/config"
	"github.com/SliverFlow/ksmall/app/user/rpc/pb"
	"github.com/SliverFlow/ksmall/core/rpc/client"
)

func NewUserRpc(c *config.Possess) pb.UserClient {
	conn, err := client.NewRpcClient(c.UserRpc)
	if err != nil {
		fmt.Println("new user rpc client error")
		panic(err)
		return nil
	}
	return pb.NewUserClient(conn)
}
