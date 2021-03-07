package test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"qiaosong03.com/simple02_grpc/pb"
	"testing"
)

func Test_Add_BothGreaterThenZero(t *testing.T) {
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
	assert.Nil(t, err)
	assert.Equal(t, int32(3), addRes.Res)
}

func Test_Add_OneLessThanZero(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:4444", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Errorf("连接创建失败！")
	}

	client := pb.NewSimpleMathServiceClient(conn)

	addReq := &pb.AddRequest{
		A: 1,
		B: -123,
	}
	addRes, err := client.Add(context.Background(), addReq)
	assert.Nil(t, err)
	assert.Equal(t, int32(-122), addRes.Res)
}