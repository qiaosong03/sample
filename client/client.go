package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"qiaosong03.com/simple02_grpc/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4444", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Errorf("连接创建失败！")
	}

	client := pb.NewSimpleMathServiceClient(conn)

	addReq := &pb.AddRequest{
		A: 1,
		B: 2,
	}
	addRes, err := client.Add(context.Background(), addReq)
	if err != nil {
		fmt.Errorf("add接口调用失败！")
	}
	fmt.Println(addRes.Res)
}
