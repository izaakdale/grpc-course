package main

import (
	"log"

	pb "github.com/izaakdale/grpc-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var addr = "localhost:50053"

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %s", addr)
	}
	defer conn.Close()
	a := pb.NewBlogServiceClient(conn)
	id1 := createBlog(a)
	id2 := createBlog(a)

	readBlog(a, id1)
	updateBlog(a, id1)

	ls := listBlogs(a)

	for _, blog := range ls {
		log.Printf("Author : %s\n", blog.AuthorId)
	}

	deleteBlog(a, id2)
}
