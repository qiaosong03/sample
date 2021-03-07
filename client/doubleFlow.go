package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"qiaosong03.com/simple03_grpc/pb"
	"strconv"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4444", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Errorf("连接创建失败！")
	}
	client := pb.NewSimpleMathServiceClient(conn)

	stream, err := client.DoubleStream(context.Background())
	if err != nil {
		fmt.Errorf("请求失败")
	}

	for i := 0; i < 10; i++ {
		err = stream.Send(&pb.HelloRequest{
			Name: "chen_" + strconv.Itoa(rand.Int()),
		})
		if err != nil {
			break
		}

		res, err := stream.Recv()
		if err != nil {
			fmt.Errorf("接收错误")
		}
		if err == io.EOF {
			break
		}
		fmt.Println(res.Hello)
	}
}
