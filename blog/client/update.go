package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("DoUpdate invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Izaak",
		Title:    "New Title",
		Content:  "New content to the blog",
	}

	_, err := c.Update(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error updating %v", id)
	}
	log.Println("Blog updated!")
}
