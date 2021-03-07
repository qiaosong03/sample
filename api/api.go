package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"qiaosong03.com/simple02_grpc/pb"
)

var conn *grpc.ClientConn

func init() {
	conn2, err := grpc.Dial("127.0.0.1:4444", grpc.WithInsecure())
	//defer conn.Close()  //取消注释会panic
	if err != nil {
		fmt.Errorf("连接创建失败！")
	}
	conn = conn2
}

func NewSimpleMathClient() pb.SimpleMathServiceClient {
	client := pb.NewSimpleMathServiceClient(conn)
	return client
}

func Add(a int32, b int32) (*pb.AddResponse, error) {

	client := NewSimpleMathClient()
	req := &pb.AddRequest{
		A: a,
		B: b,
	}
	addRes, err := client.Add(context.Background(), req)
	return addRes, err
}

func Subtract(a int32, b int32) (*pb.SubtractResponse, error) {
	client := NewSimpleMathClient()
	req := &pb.SubtractRequest{
		A: a,
		B: b,
	}
	addRes, err := client.Subtract(context.Background(), req)
	return addRes, err
}

func Multiply(a int32, b int32) (*pb.MultiplyResponse, error) {
	client := NewSimpleMathClient()
	req := &pb.MultiplyRequest{
		A: a,
		B: b,
	}
	addRes, err := client.Multiply(context.Background(), req)
	return addRes, err
}

func Divide(a int32, b int32) (*pb.DivideResponse, error) {
	client := NewSimpleMathClient()
	req := &pb.DivideRequest{
		A: a,
		B: b,
	}
	addRes, err := client.Divide(context.Background(), req)
	return addRes, err
}
