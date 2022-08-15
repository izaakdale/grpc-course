package main

import (
	"io"
	"log"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
)

func (s *Server) RequestAverage(stream pb.CalculatorService_RequestAverageServer) error {
	log.Println("Requesing average")
	n, total := int32(0), int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := float64(total) / float64(n)
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: avg,
			})
		}
		if err != nil {
			log.Fatalf("Error recieving on stream %v", err)
		}
		log.Printf("Recieved %d to add to average calculation", req.GetN())
		total += req.GetN()
		n++
	}

}
