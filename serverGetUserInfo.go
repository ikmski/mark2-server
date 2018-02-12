package main

import (
	"context"

	"github.com/ikmski/mark2-server/proto"
)

func (s *server) GetUserInfo(ctx context.Context, req *mark2.UserInfoRequest) (*mark2.UserInfoResult, error) {

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

	result.Result.Code = mark2.ResultCode_OK

	return result, nil
}
