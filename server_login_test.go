package main

import (
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/ikmski/mark2-server/proto"
)

const (
	testServerPort = ":50051"
)

func testServer() {

	lis, err := net.Listen("tcp", testServerPort)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, newServer())

	err = s.Serve(lis)
	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}

func TestServerLogin(t *testing.T) {

	go testServer()

}
