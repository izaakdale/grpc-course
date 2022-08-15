package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("Read blog invoked")
	req := &pb.BlogId{Id: id}

	res, err := c.Read(context.Background(), req)
	if err != nil {
		fmt.Printf("Error while reading %v\n", err)
	} else {
		fmt.Printf("Blog read: %v\n", res)
	}

	return res
}
