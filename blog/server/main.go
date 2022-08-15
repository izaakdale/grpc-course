package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/izaakdale/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	var addr = "localhost:50053"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	fmt.Printf("Listening on %s\n", addr)

	gs := grpc.NewServer()
	pb.RegisterBlogServiceServer(gs, &Server{})
	reflection.Register(gs)

	log.Fatal(gs.Serve(lis))
}
