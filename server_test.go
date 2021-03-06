package main

import (
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/ikmski/mark2-server/proto"
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
	mark2.RegisterMessageServiceServer(s, newServer())

	err = s.Serve(lis)
	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}

func TestMain(m *testing.M) {
	go testServer()
	os.Exit(m.Run())
}
