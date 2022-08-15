package main

import (
	"context"
	"fmt"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	sum := in.X + in.Y
	fmt.Printf("Sum is %d\n", sum)

	return &pb.SumResponse{
		Sum: sum,
	}, nil
}
