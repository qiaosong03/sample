package main

import (
	"google.golang.org/grpc"
	"net"
	"qiaosong03.com/simple02_grpc/pb"
	"qiaosong03.com/simple02_grpc/service"
)

func main() {

	listener, _ := net.Listen("tcp", ":4444")

	ser := service.MySimpleMathService{}
	grpcSer := grpc.NewServer()
	pb.RegisterSimpleMathServiceServer(grpcSer, &ser)

	grpcSer.Serve(listener)
}
