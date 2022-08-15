package main

import (
	"context"
	"fmt"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
)

func requestSum(c pb.CalculatorServiceClient) {
	x, y := int32(5), int32(10)
	fmt.Printf("Requesting sumation of %d & %d", x, y)
	c.Sum(context.Background(), &pb.SumRequest{
		X: x,
		Y: y,
	})
}
