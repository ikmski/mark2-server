package main

import (
	"fmt"
	"time"

	"github.com/ikmski/mark2-server/proto"
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

func (s *messageServer) GetUserInfo(ctx context.Context, req *mark2.UserInfoRequest) (*mark2.UserInfoResult, error) {

	result := mark2.NewUserInfoResult()

	claims, err := tokenDecode(req.Token.Token)
	if err != nil {
		return result, err
	}

	if len(req.UserIdList) == 0 {

		user, err := getUsersInstance().get(claims.UserID)
		if err != nil {
			return result, err
		}
		result.UserInfoList = append(result.UserInfoList, user.info)

	} else {

		for _, id := range req.UserIdList {

			user, err := getUsersInstance().get(id)
			if err != nil {
				result.UserInfoList = append(result.UserInfoList, nil)
			} else {
				result.UserInfoList = append(result.UserInfoList, user.info)
			}

		}
	}

	result.Result.Code = mark2.ResultCodes_OK

	return result, nil
}

func (s *messageServer) GetRoomInfo(ctx context.Context, req *mark2.RoomInfoRequest) (*mark2.RoomInfoResult, error) {

	result := mark2.NewRoomInfoResult()

	claims, err := tokenDecode(req.Token.Token)
	if err != nil {
		return result, err
	}

	if len(req.RoomIdList) == 0 {

		user, err := getUsersInstance().get(claims.UserID)
		if err != nil {
			return result, err
		}

		room, err := getRoomsInstance().get(user.roomID)
		if err != nil {
			return result, err
		}
		result.RoomInfoList = append(result.RoomInfoList, room.info)

	} else {

		for _, id := range req.RoomIdList {

			room, err := getRoomsInstance().get(id)
			if err != nil {
				result.RoomInfoList = append(result.RoomInfoList, nil)
			} else {
				result.RoomInfoList = append(result.RoomInfoList, room.info)
			}

		}
	}

	result.Result.Code = mark2.ResultCodes_OK

	return result, nil
}

func (s *messageServer) MatchRandom(ctx context.Context, req *mark2.MatchRequest) (*mark2.RoomInfoResult, error) {

	result := mark2.NewRoomInfoResult()

	claims, err := tokenDecode(req.Token.Token)
	if err != nil {
		return result, err
	}

	own, err := getUsersInstance().get(claims.UserID)
	if err != nil {
		return result, err
	}

	// Check my status
	if own.info.Status != mark2.UserStatus_Login {
		err := fmt.Errorf("invalit user status")
		return result, err
	}

	// get users waiting for match
	userIDList := getUserIDListInstance()
	userIDs, err := userIDList.get(claims.GroupID, mark2.UserStatus_WaitMatch)
	if err != nil {
		return result, err
	}

	if len(userIDs) > 0 { // Match!

		other, err := getUsersInstance().get(userIDs[0])
		if err != nil {
			return result, err
		}

		// Create and Join the room
		newRoom, err := createRoom(claims.GroupID, 2, own.info.Id)
		if err != nil {
			return result, err
		}
		err = newRoom.join(other.info.Id)
		if err != nil {
			return result, err
		}

		// Set RoomID
		// TODO unset when exit the room
		own.roomID = newRoom.info.Id
		other.roomID = newRoom.info.Id

		// Change Status
		err = own.changeStatus(mark2.UserStatus_Matched)
		if err != nil {
			return result, err
		}
		err = other.changeStatus(mark2.UserStatus_Matched)
		if err != nil {
			return result, err
		}

		result.RoomInfoList = append(result.RoomInfoList, newRoom.info)

	} else {

		err = own.changeStatus(mark2.UserStatus_WaitMatch)
		if err != nil {
			return result, err
		}

		isSuccess := false
		for i := 0; i < 10; i++ {

			if own.info.Status == mark2.UserStatus_Matched {

				// find joind room
				roomIDList := getRoomIDListInstance()
				roomIDs, err := roomIDList.get(claims.GroupID, mark2.RoomStatus_CLOSED)
				if err != nil {
					return result, err
				}

				rooms := getRoomsInstance()
				for _, roomID := range roomIDs {
					room, err := rooms.get(roomID)
					if err != nil {
						return result, err
					}

					if room.isJoined(own.info.Id) {

						result.RoomInfoList = append(result.RoomInfoList, room.info)
						isSuccess = true
						break
					}
				}
			}

			if isSuccess {
				break
			}

			time.Sleep(1 * time.Second)
		}

		if !isSuccess {

			// could not match
			err = own.changeStatus(mark2.UserStatus_Login)
			if err != nil {
				return result, err
			}
		}
	}

	result.Result.Code = mark2.ResultCodes_OK
	return result, nil
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
			message.Id = issueMessageID()
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

			default:
			}
		}
	}()

	select {
	case err := <-errChan:
		fmt.Printf("%v\n", err)
		return err
	}

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

		default:
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
