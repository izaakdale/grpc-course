package main

import (
	pb "github.com/izaakdale/grpc-course/calculator/proto"
)

func (s *Server) ManyPrimes(in *pb.PrimeRequest, stream pb.CalculatorService_ManyPrimesServer) error {

	k := uint32(2)
	N := in.GetNum()

	for N > 1 {
		if N%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Prime: k,
			})
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil
}
