/*
 run the test:
 1. run the online rpc server
 2. go run test_client.go
*/

package main

import (
	"context"
	"fmt"

	"github.com/wymli/bcsns/app/online_rpc/pb"
	"github.com/wymli/bcsns/common/server_framework/rpc"
	"google.golang.org/grpc"
)

func main() {
	rpc := pb.NewOnlineClient(rpc.Must(grpc.Dial("localhost:9999")))

	rspOnline, err := rpc.BatchOnlineUser(context.TODO(), &pb.BatchOnlineUserReq{UserIdList: []uint64{1, 4}, GatewayAddr: "localhost:1234"})
	fmt.Printf("rsp:%v, err:%v\n", rspOnline, err)

	rspOnline, err = rpc.BatchOnlineUser(context.TODO(), &pb.BatchOnlineUserReq{UserIdList: []uint64{2}, GatewayAddr: "localhost:9994"})
	fmt.Printf("rsp:%v, err:%v\n", rspOnline, err)

	rspAll, err := rpc.GetAllOnlineUser(context.TODO(), &pb.GetAllOnlineUserReq{})
	fmt.Printf("rsp:%v, err:%v\n", rspAll, err)

	rspGet, err := rpc.BatchGetUserGateway(context.TODO(), &pb.BatchGetUserGatewayReq{UserIdList: []uint64{1, 2, 3, 4}})
	fmt.Printf("rsp:%v, err:%v\n", rspGet, err)
}
