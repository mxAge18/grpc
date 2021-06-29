package main

import (
	"context"
	"fmt"
	"go_code/grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)
const (
	port = ":50051"
)
type server struct {
	pb.UnimplementedTeacherServiceServer
}

func (this *server) SayHi(ctx context.Context, in *pb.Teacher) (*pb.Teacher, error) {
	return &pb.Teacher{Name: "Hi " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterTeacherServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("start listen the tcp conn from client")
}