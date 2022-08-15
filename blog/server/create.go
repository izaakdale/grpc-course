package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Create(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create blog invoked %v\n", in)

	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to oid %v", err),
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil

}
