package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "token-eg/proto" // Replace with your actual package path
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)

	// Prepare a sign-in request
	req := &pb.SignInRequest{
		Username: "user123",
		Password: "password123",
	}

	// Call the sign-in service
	res, err := c.SignIn(context.Background(), req)
	if err != nil {
		log.Fatalf("Sign-in failed: %v", err)
	}

	// Handle the response
	fmt.Printf("Access Token: %s\n", res.AccessToken)
	fmt.Printf("Message: %s\n", res.Message)
}
