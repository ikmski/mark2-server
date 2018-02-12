package main

import (
	"context"
	"testing"
	"time"

	"github.com/ikmski/mark2-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func TestServerWaitMessageConnection(t *testing.T) {

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mark2.NewMessageServiceClient(conn)

	var groupID uint32 = 10001

	loginRequest := mark2.NewLoginRequest()
	loginRequest.GroupId = groupID

	ctx1 := context.Background()

	// Login
	loginResult, err := c.Login(ctx1, loginRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	token := loginResult.AccessToken

	// WaitMessage
	ctx2, cancel := context.WithCancel(context.Background())
	_, err = c.WaitMessage(ctx2, token)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	time.Sleep(1 * time.Second)

	cancel()

	// Logout
	_, err = c.Logout(ctx1, token)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestServerWaitMessageLogout(t *testing.T) {

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mark2.NewMessageServiceClient(conn)

	var groupID uint32 = 10001

	loginRequest := mark2.NewLoginRequest()
	loginRequest.GroupId = groupID

	ctx1 := context.Background()
	ctx2 := context.Background()

	// Login
	loginResult, err := c.Login(ctx1, loginRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	token := loginResult.AccessToken

	// WaitMessage
	_, err = c.WaitMessage(ctx2, token)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	time.Sleep(1 * time.Second)

	// Logout
	_, err = c.Logout(ctx1, token)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}
