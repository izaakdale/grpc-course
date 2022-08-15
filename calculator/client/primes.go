package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
)

func requestPrimes(c pb.CalculatorServiceClient) {
	req := &pb.PrimeRequest{
		Num: 129049928,
	}
	stream, err := c.ManyPrimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error sending many greets %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading from stream %v\n", err)
		}
		fmt.Printf("A prime number of %d is %d\n", req.Num, msg.Prime)
	}
}
