package main

import (
	"log"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var addr = "localhost:50052"

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %s", addr)
	}
	defer conn.Close()
	a := pb.NewCalculatorServiceClient(conn)
	// requestSum(a)

	doSqrt(a, -6)
}
