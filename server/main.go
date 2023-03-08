package main

import (
	"context"
	pb "example.com/grpc/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":54321"
)

type Server struct {
	pb.UnimplementedAddServiceServer
}

func (s *Server) Chat(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %s", in.A)
	var res = in.GetA() + " -from Server"
	log.Printf("Sending Result: %s", res)
	return &pb.Response{Result: res}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port 54321: %s", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAddServiceServer(grpcServer, &Server{})
	log.Printf("server listening at %s", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server over port 5432 : %s", err)
	}
}
