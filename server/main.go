package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "token-eg/proto" // Replace with your actual package path
)

type authServiceServer struct{
	ctx context.Context
	
}

func (s *authServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	// Your sign-in logic here
	// Example: Check username and password and generate an access token
	if req.Username == "user123" && req.Password == "password123" {
		return &pb.SignInResponse{
			AccessToken: "your_access_token_here",
			Message:     "Sign-in successful",
		}, nil
	}

	return &pb.SignInResponse{
		AccessToken: "",
		Message:     "Invalid credentials",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &authServiceServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
