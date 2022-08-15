package main

import (
	"context"

	pb "github.com/izaakdale/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Update(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Could not get object id")
	}

	req := &BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": req})
	if err != nil {
		return nil, status.Error(codes.Internal, "Could not update")
	}
	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Could not find blog with id %s", in.Id)
	}

	return &emptypb.Empty{}, nil
}
