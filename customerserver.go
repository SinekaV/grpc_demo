package main

import (
	"fmt"
	pb "grpc/customer"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type CustomerServiceServer struct {
	mu       sync.Mutex
	customer map[string]*pb.CustomerDetails
	pb.UnimplementedCustomerserviceServer
}

func (c *CustomerServiceServer) AddAccount(ctx context.Context, req *pb.CustomerDetails) (*pb.FetchDetails, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	accID := "1"
	req.AccID = accID
	c.customer[accID] = req
	return &pb.FetchDetails{
		Customer: []*pb.CustomerDetails{},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to lIsten:%v", err)
		return
	}
	s := grpc.NewServer()
	// hw.RegisterGreeterServer(s,&server{})
	pb.CustomerserviceServer(s, &CustomerServiceServer{
		customer: make(map[string]*pb.CustomerDetails),
	})
	fmt.Println("Server listening on:50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to server:%v", err)
	}
}
