package main

import (
	"context"
	"testing"
	"time"

	mark2 "github.com/ikmski/mark2-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func TestServerMatchRandomTimeOut(t *testing.T) {

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mark2.NewMessageServiceClient(conn)

	var groupID uint32 = 10001

	// Login
	loginRequest := new(mark2.LoginRequest)
	loginRequest.GroupId = groupID

	loginResult, err := c.Login(context.Background(), loginRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	token1 := loginResult.AccessToken

	// Match Request
	matchRequest := new(mark2.MatchRequest)
	matchRequest.Token = token1
	roomInfoResult, err := c.MatchRandom(context.Background(), matchRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(roomInfoResult.RoomInfoList) > 0 {
		t.Errorf("got %v\nwant %v", len(roomInfoResult.RoomInfoList), 0)
	}

	userInfoRequest := new(mark2.UserInfoRequest)
	userInfoRequest.Token = token1

	userInfoResult, err := c.GetUserInfo(context.Background(), userInfoRequest)
	if userInfoResult.UserInfoList[0].Status != mark2.UserStatus_Login {
		t.Errorf("got %v\nwant %v", userInfoResult.UserInfoList[0].Status, mark2.UserStatus_Login)
	}

	// Logout
	_, err = c.Logout(context.Background(), token1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestServerMatchRandomSuccess(t *testing.T) {

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mark2.NewMessageServiceClient(conn)

	var groupID uint32 = 10001

	// Login
	loginRequest := new(mark2.LoginRequest)
	loginRequest.GroupId = groupID

	loginResult, err := c.Login(context.Background(), loginRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	token1 := loginResult.AccessToken

	loginResult, err = c.Login(context.Background(), loginRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	token2 := loginResult.AccessToken

	// Match Request
	ch := make(chan struct{})
	go func() {

		matchRequest := new(mark2.MatchRequest)
		matchRequest.Token = token1
		roomInfoResult, err := c.MatchRandom(context.Background(), matchRequest)
		if err != nil {
			t.Errorf("got %v\n", err)
		}
		if len(roomInfoResult.RoomInfoList) != 1 {
			t.Errorf("got %v\nwant %v", len(roomInfoResult.RoomInfoList), 1)
		}

		userInfoRequest := new(mark2.UserInfoRequest)
		userInfoRequest.Token = token1
		userInfoResult, err := c.GetUserInfo(context.Background(), userInfoRequest)
		if userInfoResult.UserInfoList[0].Status != mark2.UserStatus_Matched {
			t.Errorf("got %v\nwant %v", userInfoResult.UserInfoList[0].Status, mark2.UserStatus_Matched)
		}

		_, err = c.Logout(context.Background(), token1)
		if err != nil {
			t.Errorf("got %v\n", err)
		}

		close(ch)
	}()

	time.Sleep(1 * time.Second)

	matchRequest := new(mark2.MatchRequest)
	matchRequest.Token = token2
	roomInfoResult, err := c.MatchRandom(context.Background(), matchRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(roomInfoResult.RoomInfoList) != 1 {
		t.Errorf("got %v\nwant %v", len(roomInfoResult.RoomInfoList), 1)
	}

	userInfoRequest := new(mark2.UserInfoRequest)
	userInfoRequest.Token = token2
	userInfoResult, err := c.GetUserInfo(context.Background(), userInfoRequest)
	if userInfoResult.UserInfoList[0].Status != mark2.UserStatus_Matched {
		t.Errorf("got %v\nwant %v", userInfoResult.UserInfoList[0].Status, mark2.UserStatus_Matched)
	}

	// Logout
	_, err = c.Logout(context.Background(), token2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	<-ch
}
