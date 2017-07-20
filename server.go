package main

import (
	"fmt"

	mark2 "github.com/ikmski/mark2-server/proto"
	"golang.org/x/net/context"
)

type messageServer struct {
}

func newServer() *messageServer {
	s := new(messageServer)
	return s
}

func (s *messageServer) Login(ctx context.Context, req *mark2.LoginRequest) (*mark2.LoginResult, error) {

	u, err := fetchOrCreateUser(req.UniqueKey, req.GroupId)
	if err != nil {
		return nil, err
	}

	// Status をログインに変更
	fmt.Printf("%v\n", u)
	u.changeStatus(mark2.UserStatus_Login)

	result := mark2.NewLoginResult()
	result.Result.Code = mark2.ResultCodes_OK

	return result, nil
}

func (s *messageServer) Logout(ctx context.Context, token *mark2.AccessToken) (*mark2.Result, error) {

	return new(mark2.Result), nil
}

func (s *messageServer) GetUserInfoList(ctx context.Context, token *mark2.AccessToken) (*mark2.UserInfoListResult, error) {

	return new(mark2.UserInfoListResult), nil
}

func (s *messageServer) GetRoomInfoList(ctx context.Context, token *mark2.AccessToken) (*mark2.RoomInfoListResult, error) {

	return new(mark2.RoomInfoListResult), nil
}

func (s *messageServer) GetOwnUserInfo(ctx context.Context, token *mark2.AccessToken) (*mark2.UserInfoResult, error) {

	return new(mark2.UserInfoResult), nil
}

func (s *messageServer) GetOwnRoomInfo(ctx context.Context, token *mark2.AccessToken) (*mark2.RoomInfoResult, error) {

	return new(mark2.RoomInfoResult), nil
}

func (s *messageServer) CreateRoom(ctx context.Context, req *mark2.CreateRoomRequest) (*mark2.RoomInfoResult, error) {

	return new(mark2.RoomInfoResult), nil
}

func (s *messageServer) JoinRoom(ctx context.Context, req *mark2.JoinRoomRequest) (*mark2.RoomInfoResult, error) {

	return new(mark2.RoomInfoResult), nil
}

func (s *messageServer) MatchRandom(ctx context.Context, req *mark2.MatchRequest) (*mark2.RoomInfoResult, error) {

	return new(mark2.RoomInfoResult), nil
}

func (s *messageServer) ExitRoom(ctx context.Context, token *mark2.AccessToken) (*mark2.Result, error) {

	return new(mark2.Result), nil
}

func (s *messageServer) SendMessage(ctx context.Context, req *mark2.MessageRequest) (*mark2.Result, error) {

	return new(mark2.Result), nil
}

func (s *messageServer) SendStream(srv mark2.MessageService_SendStreamServer) error {

	return nil
}

func (s *messageServer) WaitMessage(token *mark2.AccessToken, srv mark2.MessageService_WaitMessageServer) error {

	return nil
}
