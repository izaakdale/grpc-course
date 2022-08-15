package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/izaakdale/grpc-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Do greet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Izaak",
	})
	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("My name in %s\n", res.Result)
}

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("Receiving many greets")
	req := &pb.GreetRequest{
		FirstName: "izaak",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error receiving many greets %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading from stream %v\n", err)
		}

		log.Printf("Greet many times %s\n", msg.Result)
	}
}

func doSendManyGreets(c pb.GreetServiceClient) {
	names := []string{
		"Izaak",
		"Mahtab",
		"Thistle",
		"Woody",
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error setting long greet stream %v", err)
	}

	for _, name := range names {
		stream.Send(&pb.GreetRequest{
			FirstName: name,
		})
		time.Sleep(time.Second)
		fmt.Printf("Sending %s\n", name)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error recv and closing stream %v", err)
	}

	fmt.Printf("Reponse: \n%s\n", res.Result)
}

func doGreetEveryone(c pb.GreetServiceClient) {
	fmt.Println("Greet everyone envoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while connecting to stream %v", err)
	}

	names := []string{
		"Izaak",
		"Mahtab",
		"Thistle",
		"Woody",
	}

	waitc := make(chan struct{})

	go func() {
		for _, name := range names {
			fmt.Println("Sending greet request")
			stream.Send(&pb.GreetRequest{
				FirstName: name,
			})
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error from stream %v", err)
				break
			}

			fmt.Println(res.Result)
		}

		close(waitc)
	}()

	<-waitc
}

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := pb.GreetRequest{
		FirstName: "Izaak",
	}
	res, err := c.GreetWithDeadline(ctx, &req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline was exceeded")
				return
			} else {
				log.Fatalf("Unexpected grpc error %v", err)
			}

		} else {
			log.Fatalf("Non grpc error %v", err)
		}
	}

	log.Printf("GreetWithDeadline: %s", res.Result)
}
