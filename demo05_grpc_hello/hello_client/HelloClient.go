package main

import (
	pb "GoDistributedAppDevelop/demo05_grpc_hello/protos/hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 要通信的地址 ip:端口
const Address = "127.0.0.1:8001"

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接GRPC服务端失败: %v", err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "GRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		fmt.Println("调用GRPC服务端方法失败: %v", err)
	}
	fmt.Println("接收到的信息:")
	fmt.Println(res.Message)
}
