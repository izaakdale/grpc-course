package main

import (
	"context"
	"io"
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) []*pb.Blog {

	stream, err := c.List(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while listing blogs %v", err)
	}

	var list []*pb.Blog

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unexpected error %v", err)
		}

		list = append(list, res)
	}

	return list
}
