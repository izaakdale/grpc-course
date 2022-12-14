package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) List(_ *emptypb.Empty, stream pb.BlogService_ListServer) error {
	log.Println("List blogs invoked")

	cur, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown error %v", err))
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error decoding data from mongo %v", err))
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown error %v", err))
	}
	return nil
}
