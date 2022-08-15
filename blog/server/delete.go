package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Delete(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("Delete invoked\n")

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot parse id")
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Mongo error when trying to delete %v", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Blog was not found")
	}

	return &emptypb.Empty{}, nil
}
