package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
)

func requestMax(c pb.CalculatorServiceClient) {

	nums := []int32{
		1, 2, 4, 3, 5, 7, 6, 8, 9,
	}

	stream, err := c.GetMax(context.Background())
	if err != nil {
		log.Fatalf("Error establishing connection to stream %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, n := range nums {
			log.Printf("Sending: %d\n", n)
			stream.Send(&pb.MaxRequest{
				N: n,
			})
			time.Sleep(2 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving from stream %v", err)
			}

			fmt.Printf("The current max is: %d\n", res.Max)
		}

		close(waitc)
	}()

	<-waitc
}
