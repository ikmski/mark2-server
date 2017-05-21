package main

import (
	pb "github.com/ikmski/mark2-server/proto"
	"golang.org/x/net/context"
)

type messageServer struct {
}

func newServer() *messageServer {
	s := new(messageServer)
	return s
}

func (s *messageServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResult, error) {

	// ユーザが存在しているか
	user, _ := fetchUser(req.UniqueKey)
	if user == nil {

		// ユーザを作成
		_, err := createUser(req.UniqueKey, req.GroupId)
		if err != nil {
			return nil, err
		}
	}

	result := pb.NewLoginResult()
	result.Result.Code = pb.ResultCodes_OK

	return result, nil
}

func (s *messageServer) Logout(ctx context.Context, token *pb.AccessToken) (*pb.Result, error) {

	return new(pb.Result), nil
}

func (s *messageServer) GetUserInfoList(ctx context.Context, token *pb.AccessToken) (*pb.UserInfoListResult, error) {

	return new(pb.UserInfoListResult), nil
}

func (s *messageServer) GetRoomInfoList(ctx context.Context, token *pb.AccessToken) (*pb.RoomInfoListResult, error) {

	return new(pb.RoomInfoListResult), nil
}

func (s *messageServer) GetOwnUserInfo(ctx context.Context, token *pb.AccessToken) (*pb.UserInfoResult, error) {

	return new(pb.UserInfoResult), nil
}

func (s *messageServer) GetOwnRoomInfo(ctx context.Context, token *pb.AccessToken) (*pb.RoomInfoResult, error) {

	return new(pb.RoomInfoResult), nil
}

func (s *messageServer) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.RoomInfoResult, error) {

	return new(pb.RoomInfoResult), nil
}

func (s *messageServer) JoinRoom(ctx context.Context, req *pb.JoinRoomRequest) (*pb.RoomInfoResult, error) {

	return new(pb.RoomInfoResult), nil
}

func (s *messageServer) MatchRandom(ctx context.Context, req *pb.MatchRequest) (*pb.RoomInfoResult, error) {

	return new(pb.RoomInfoResult), nil
}

func (s *messageServer) ExitRoom(ctx context.Context, token *pb.AccessToken) (*pb.Result, error) {

	return new(pb.Result), nil
}

func (s *messageServer) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.Result, error) {

	return new(pb.Result), nil
}

func (s *messageServer) SendStream(srv pb.MessageService_SendStreamServer) error {

	return nil
}

func (s *messageServer) WaitMessage(token *pb.AccessToken, srv pb.MessageService_WaitMessageServer) error {

	return nil
}
