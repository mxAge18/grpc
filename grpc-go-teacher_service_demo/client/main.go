package main

import (
	"context"
	"go_code/grpc/pb"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)
func main() {
	conn, err := grpc.Dial(address,grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("net.Dial fail, %v", err)
	}
	defer conn.Close()
	c := pb.NewTeacherServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHi(ctx, &pb.Teacher{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("result: %s", r.GetName())

}