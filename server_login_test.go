package main

import (
	"context"
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func TestServerLogin(t *testing.T) {

	// clear Storage
	userStorage := newUserStorage()
	userStorage.clear()

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mark2.NewMessageServiceClient(conn)

	var groupID uint32 = 10001
	var uniqueKey string = "test_unique_key"

	loginRequest := new(mark2.LoginRequest)
	loginRequest.GroupId = groupID
	loginRequest.UniqueKey = uniqueKey

	// Login
	loginResult, err := c.Login(context.Background(), loginRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if loginResult == nil {
		t.Errorf("got %v\n", loginResult)
	}
	if loginResult.Result == nil {
		t.Errorf("got %v\n", loginResult.Result)
	}
	if loginResult.Result.Code != mark2.ResultCodes_OK {
		t.Errorf("got %v\n", loginResult.Result.Code)
	}
	if loginResult.AccessToken == nil {
		t.Errorf("got %v\n", loginResult.AccessToken)
	}
	if loginResult.AccessToken.Token == "" {
		t.Errorf("got %v\n", loginResult.AccessToken.Token)
	}

	ok, err := tokenVerify(loginResult.AccessToken.Token)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

	// Logout
	logoutResult, err := c.Logout(context.Background(), loginResult.AccessToken)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if logoutResult == nil {
		t.Errorf("got %v\n", logoutResult)
	}
	if logoutResult.Code != mark2.ResultCodes_OK {
		t.Errorf("got %v\n", logoutResult.Code)
	}

	userStorage.clear()
}
