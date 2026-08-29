package main

import (
	"context"
	"log"
	"net"

	"vortex/internal/configs"
	pb "vortex/proto/v1"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(cxt context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Message recieved from vortexd!, Name: " + req.Name}, nil
}

func (s *server) SayHelloAgain(cxt context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Message recieved from vortexd! Saying Hello Again!, Name: " + req.Name}, nil
}

func main() {

	// Create new db session
	session, _ := configs.NewSession()
	
	defer session.Close()

	// Open the connection
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("[ERROR] Unable to start up the TCP server: %v", err)
	}
	
	// Create and start the gPRC server
	serv := grpc.NewServer()
	pb.RegisterGreeterServer(serv, &server{})
	
	log.Printf("vortexd server listening on port: %v", lis.Addr())

	// Send the data
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("[ERROR] Unable to send data to client: %v", err)
	}
}
