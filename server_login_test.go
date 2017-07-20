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
	userStorage := getUserStorageInstance()
	userStorage.clear()

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mark2.NewMessageServiceClient(conn)

	var groupId uint32 = 10001
	var uniqueKey string = "test_unique_key"

	request := new(mark2.LoginRequest)
	request.GroupId = groupId
	request.UniqueKey = uniqueKey

	result, err := c.Login(context.Background(), request)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if result.Result == nil || result.Result.Code != mark2.ResultCodes_OK {
		t.Errorf("got %v\n", result)
	}

	if result.AccessToken == nil || result.AccessToken.Token == "" {
		t.Errorf("got %v\n", result)
	}
	ok, err := tokenVerify(result.AccessToken.Token)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

	userStorage.clear()
}
