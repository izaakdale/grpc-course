package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	fmt.Println("Sqrt invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Num: n,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Fatalf("We probably entered a negative number")
				return
			}
		} else {
			log.Fatalf("Non grpc error %v", err)
		}
	}

	fmt.Printf("Sqrt is %f\n", res.Sqrt)
}
