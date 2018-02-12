package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ikmski/mark2-server/proto"
)

func (s *server) MatchRandom(ctx context.Context, req *mark2.MatchRequest) (*mark2.RoomInfoResult, error) {

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
	set := getUint32SetInstance()
	userIDs, err := set.get(createUserIdListKey(claims.GroupID, mark2.UserStatus_WaitMatch))
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
				roomIDs, err := set.get(createRoomIdListKey(claims.GroupID, mark2.RoomStatus_CLOSED))
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

	result.Result.Code = mark2.ResultCode_OK
	return result, nil
}
