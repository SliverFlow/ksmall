package client

import (
	"fmt"
	"github.com/SliverFlow/ksmall/core/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ConnTypeDirect = "direct"
const ConnTypeNacos = "nacos"
const ConnTypeEtcd = "etcd"

// NewRpcClient 创建 rpc client
func NewRpcClient(c *config.RpcClient) (*grpc.ClientConn, error) {
	switch c.Type {
	case ConnTypeDirect:
		return newClientWithIp(c.Ip)
	case "etcd":
	default:
		return newClientWithIp(c.Ip)
	}
	return nil, fmt.Errorf("rpc client not support")
}

// newClientWithIp 创建直连的 rpc client
func newClientWithIp(ip string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
