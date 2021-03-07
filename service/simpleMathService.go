package service

import (
	"context"
	"qiaosong03.com/simple02_grpc/pb"
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
