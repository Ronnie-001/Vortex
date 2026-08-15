package main

import (
	"context"
	"log"
	pb "vortex/proto/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	conn, err := grpc.NewClient(":8080", opts...)
	if err != nil {
		log.Fatalf("[ERROR] Unable to create a new gPRC client: %v", err)
	}

	defer conn.Close()
	
	// Create the stub within the client.
	client := pb.NewGreeterClient(conn)

	// Make a gPRC call!
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Ronald"})
	if err != nil {
		log.Fatalf("[ERROR] Failed to make the gPRC call!: %v,", err)
	}
	
	print(response.GetMessage())
}
