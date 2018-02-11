package main

import (
	"context"
	"testing"

	"github.com/ikmski/mark2-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func TestServerGetOwnUserInfo(t *testing.T) {

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mark2.NewMessageServiceClient(conn)

	var groupID uint32 = 10001

	var token *mark2.AccessToken

	// Login
	{
		loginRequest := new(mark2.LoginRequest)
		loginRequest.GroupId = groupID

		// Login
		loginResult, err := c.Login(context.Background(), loginRequest)
		if err != nil {
			t.Errorf("got %v\n", err)
		}

		token = loginResult.AccessToken
	}

	// Get Own UserInfo
	{
		userInfoRequest := new(mark2.UserInfoRequest)
		userInfoRequest.Token = token

		userInfoResult, err := c.GetUserInfo(context.Background(), userInfoRequest)
		if err != nil {
			t.Errorf("got %v\n", err)
		}

		if len(userInfoResult.UserInfoList) != 1 {
			t.Errorf("got %v\nwant %v", len(userInfoResult.UserInfoList), 1)
		}

		if userInfoResult.UserInfoList[0].Id != currentUserID {
			t.Errorf("got %v\nwant %v", userInfoResult.UserInfoList[0].Id, currentUserID)
		}
	}

	// Logout
	{
		_, err = c.Logout(context.Background(), token)
		if err != nil {
			t.Errorf("got %v\n", err)
		}
	}
}
