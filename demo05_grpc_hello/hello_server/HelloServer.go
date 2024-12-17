package main

import (
	pb "GoDistributedAppDevelop/demo05_grpc_hello/protos/hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

// 要监听的地址 ip:端口号
const Address = "127.0.0.1:8001"

type HelloService struct {
	pb.UnimplementedHelloServer
}

func (h HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("收到客户端的请求，请求参数：", req.Name)
	rsp := &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}
	return rsp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println("监听失败,监听地址为:", Address)
	} else {
		fmt.Println("监听成功,监听地址为:", Address)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	helloSvr := HelloService{}
	pb.RegisterHelloServer(s, helloSvr)

	err = s.Serve(listen)
	if err != nil {
		fmt.Println("GRPC 服务启动失败", err)
	} else {
		fmt.Println("GRPC 服务启动成功")
	}
}
