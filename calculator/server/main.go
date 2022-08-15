package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	var addr = "localhost:50052"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	fmt.Printf("Listening on %s\n", addr)

	gs := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(gs, &Server{})
	reflection.Register(gs)

	log.Fatal(gs.Serve(lis))
}
