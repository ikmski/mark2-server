package main

import (
	"fmt"
	"time"

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

	result := mark2.NewLoginResult()

	user, err := createUser(req.GroupId)
	if err != nil {
		return result, err
	}

	// Create access token
	claim := newTokenClaims()
	claim.GroupID = user.info.GroupId
	claim.UserID = user.info.Id
	token, err := claim.encode()
	if err != nil {
		return result, err
	}

	result.AccessToken.Token = token
	result.Result.Code = mark2.ResultCodes_OK

	return result, nil
}

func (s *messageServer) Logout(ctx context.Context, token *mark2.AccessToken) (*mark2.Result, error) {

	result := mark2.NewResult()

	claims, err := tokenDecode(token.Token)
	if err != nil {
		return result, err
	}

	// User
	user, err := getUsersInstance().get(claims.UserID)
	if err != nil {
		return result, err
	}

	// Remove user
	err = user.remove()
	if err != nil {
		return result, err
	}
	user = nil

	result.Code = mark2.ResultCodes_OK

	return result, nil
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

	type receivedMessage struct {
		userID  uint32
		roomID  uint32
		content string
	}

	result := mark2.NewResult()

	ctx := srv.Context()
	errChan := make(chan error, 1)
	messageChan := make(chan receivedMessage, 1)

	// Receive Message
	go func() {
		for {
			messageReq, err := srv.Recv()
			if err != nil {
				errChan <- err
				break
			}

			claims, err := tokenDecode(messageReq.GetToken().Token)
			if err != nil {
				errChan <- err
				break
			}

			user, err := getUsersInstance().get(claims.UserID)
			if err != nil {
				errChan <- err
				break
			}

			rm := receivedMessage{
				user.info.Id,
				user.roomID,
				messageReq.Content,
			}

			messageChan <- rm
		}
	}()

	// Send Message
	go func() {
		for {
			rm := <-messageChan

			message := mark2.NewMessage()
			message.UserId = rm.userID
			message.Content = rm.content

			room, err := getRoomsInstance().get(rm.roomID)
			for _, uid := range room.info.UserIdList {
				stream, err := getWaitStreamsInstance().get(uid)
				if err != nil {
					errChan <- err
					break
				}
				err = stream.Send(message)
				if err != nil {
					errChan <- err
					break
				}
			}

			result.Code = mark2.ResultCodes_OK
			err = srv.Send(result)
			if err != nil {
				errChan <- err
				break
			}
		}
	}()

	// Check Context
	go func() {
		for {
			select {
			case <-ctx.Done():
				err := ctx.Err()
				fmt.Printf("Connection broken: %v\n", err)
				errChan <- err
				break
			}
		}
	}()

	select {
	case err := <-errChan:
		fmt.Printf("%v\n", err)
		return err
	}

	return nil
}

func (s *messageServer) WaitMessage(token *mark2.AccessToken, srv mark2.MessageService_WaitMessageServer) error {

	claims, err := tokenDecode(token.Token)
	if err != nil {
		return err
	}

	waitStreams := getWaitStreamsInstance()
	waitStreams.set(claims.UserID, srv)

	for {

		ctx := srv.Context()
		select {
		case <-ctx.Done():
			err = ctx.Err()
			fmt.Printf("Connection broken: %v\n", err)
			break
		}

		has := getUsersInstance().has(claims.UserID)
		if !has {
			fmt.Printf("User was loged out\n")
			break
		}

		time.Sleep(1 * time.Second)
	}

	waitStreams.del(claims.UserID)
	return err
}
