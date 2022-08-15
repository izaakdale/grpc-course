package main

import (
	"context"
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Create blog invoked")
	blog := &pb.Blog{
		AuthorId: "Izaak",
		Title:    "First",
		Content:  "Hello!!!!!",
	}

	res, err := c.Create(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error %v", err)
	}

	log.Printf("Blog created, ID %s", res.Id)
	return res.Id
}
