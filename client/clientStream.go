package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"qiaosong03.com/supportHttp/pb"
	"strconv"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4444", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Errorf("连接创建失败！")
	}

	client := pb.NewSimpleMathServiceClient(conn)

	clientStream, err := client.ClientStream(context.Background())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	for i := 0; i < 10; i++ {
		err := clientStream.Send(&pb.HelloRequest{
			Name: "wang_0" + strconv.Itoa(i),
		})
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(res.Hello)
}
