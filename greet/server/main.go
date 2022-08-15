package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/izaakdale/grpc-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	fmt.Printf("Listening on %s\n", addr)

	tls := true //change to false to use without tls
	opts := []grpc.ServerOption{}
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certs %v", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	gs := grpc.NewServer(opts...)
	ls := &Server{}

	pb.RegisterGreetServiceServer(gs, ls)

	if err = gs.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}
