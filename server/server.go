package main

import (
	"google.golang.org/grpc"
	"net"
	"qiaosong03.com/supportHttp/pb"
	"qiaosong03.com/supportHttp/service"
)

func main() {

	listener, _ := net.Listen("tcp", ":9090")

	ser := service.MySimpleMathService{}
	grpcSer := grpc.NewServer()
	pb.RegisterSimpleMathServiceServer(grpcSer, &ser)

	grpcSer.Serve(listener)
}
