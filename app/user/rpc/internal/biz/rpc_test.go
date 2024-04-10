package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestRpcServer(t *testing.T) {
	conn, err := grpc.Dial("localhost:6061", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	_, err = c.FindByUsername(context.Background(), &pb.FindByUsernameReq{Username: "ge"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
