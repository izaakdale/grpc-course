package main

import (
	"context"
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("Delete blog invoked\n")

	_, err := c.Delete(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error deleting %v", err)
	}

	log.Printf("Blog deleted!\n")
}
