package service

import (
	"context"
	"fmt"
	"io"
	"qiaosong03.com/simple03_grpc/pb"
	"strconv"
	"time"
)

type MySimpleMathService struct {
	pb.UnimplementedSimpleMathServiceServer
}

func (sMath *MySimpleMathService) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	res := &pb.AddResponse{
		Res: req.A + req.B,
	}
	return res, nil
}

func (sMath *MySimpleMathService) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	res := &pb.SubtractResponse{
		Res: req.A - req.B,
	}
	return res, nil
}

func (sMath *MySimpleMathService) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	res := &pb.MultiplyResponse{
		Res: req.A * req.B,
	}
	return res, nil
}

func (sMath *MySimpleMathService) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.DivideResponse, error) {
	res := &pb.DivideResponse{
		Res: req.A / req.B,
	}
	return res, nil
}

func (sMath *MySimpleMathService) ServerStream(req *pb.SquareRequest, stream pb.SimpleMathService_ServerStreamServer) error {
	for i := 1; i < int(req.A); i++ {
		err := stream.Send(&pb.SquareResponse{
			Str: strconv.Itoa(int(i)) + "*" + strconv.Itoa(int(i)) + "=" + strconv.Itoa(int(i*i)),
		})
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

func (sMath *MySimpleMathService) ClientStream(stream pb.SimpleMathService_ClientStreamServer) error {
	i := 0
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloResponse{Hello: "服务端总共接收到了" + strconv.Itoa(i) + "个数据！",})
		}
		if err != nil {
			return err
		}
		if err == nil {
			i++
		}
		fmt.Println("我接收到了：", r.Name)
	}
	return nil
}

func (sMath *MySimpleMathService) DoubleStream(stream pb.SimpleMathService_DoubleStreamServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}
		fmt.Println("我接收到了：", r.Name)

		time.Sleep(time.Second * 3)
		err = stream.Send(&pb.HelloResponse{
			Hello: "hello" + r.Name,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
