package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
)

func requestAvg(c pb.CalculatorServiceClient) {

	nums := []int32{
		1, 2,
	}

	stream, err := c.RequestAverage(context.Background())
	if err != nil {
		log.Fatalf("Error when setting average stream in client %v", err)
	}

	for _, n := range nums {
		stream.Send(&pb.AvgRequest{
			N: n,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error recieving response from server %v", err)
	}

	fmt.Println(res)
}
