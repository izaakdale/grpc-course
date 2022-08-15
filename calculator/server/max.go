package main

import (
	"io"
	"log"

	pb "github.com/izaakdale/grpc-course/calculator/proto"
)

func (s *Server) GetMax(stream pb.CalculatorService_GetMaxServer) error {
	log.Println("GetMax envoked")
	currentMax := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving data from stream %v", err)
		}
		if req.N > currentMax {
			currentMax = req.N
		}

		log.Println("Sending max")
		err = stream.Send(&pb.MaxReponse{
			Max: currentMax,
		})
		if err != nil {
			log.Fatalf("Error sending data %v", err)
		}
	}
	return nil
}
