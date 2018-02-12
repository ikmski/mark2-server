package main

import (
	"context"

	"github.com/ikmski/mark2-server/proto"
)

func (s *server) GetRoomInfo(ctx context.Context, req *mark2.RoomInfoRequest) (*mark2.RoomInfoResult, error) {

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

	result.Result.Code = mark2.ResultCode_OK

	return result, nil
}
