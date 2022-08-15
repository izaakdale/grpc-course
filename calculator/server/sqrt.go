package main

import (
	"context"
	"log"
	"math"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt envoked with %d", in.Num)

	num := in.Num
	if num < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Negative number %d", num)
	}

	return &pb.SqrtResponse{
		Sqrt: math.Sqrt(float64(num)),
	}, nil
}
