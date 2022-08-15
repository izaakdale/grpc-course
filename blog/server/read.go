package main

import (
	"context"
	"fmt"

	pb "github.com/izaakdale/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Read(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	fmt.Println("Read blog invoked")
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Cannot pass ID")
	}

	res := &BlogItem{}
	filter := bson.M{"_id": oid}

	err = collection.FindOne(ctx, filter).Decode(res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "Could not find Blog")
		} else {
			return nil, status.Error(codes.Internal, "Unexpected error")
		}
	}

	return documentToBlog(res), nil
}
