package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"qiaosong03.com/simple03_grpc/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4444", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Errorf("连接创建失败！")
	}

	client := pb.NewSimpleMathServiceClient(conn)

	squareReq := &pb.SquareRequest{
		A: 5,
	}
	stream, err := client.ServerStream(context.Background(), squareReq)
	if err != nil {
		fmt.Errorf("请求失败")
	}
	for {
		squareRes, err := stream.Recv()
		if err != nil {
			fmt.Errorf("接收错误")
		}
		if err == io.EOF {
			break
		}
		fmt.Println(squareRes.Str)
	}
}
