package main

import (
	"context"
	"log"
	"time"

	pb "github.com/faizalnurrozi/grpc-client/proto/user"
	"google.golang.org/grpc"
)

func main() {
	// Dial ke gRPC server di localhost:9090
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Siapkan request CreateUser
	req := &pb.CreateUserRequest{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "securepassword",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call CreateUser
	// TODO: Implement all handlers for gRPC methods
	res, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("CreateUser failed: %v", err)
	}

	log.Printf("User created: ID=%d, Name=%s, Email=%s", res.Id, res.Name, res.Email)
}
