package main

import (
	"context"

	"github.com/ikmski/mark2-server/proto"
)

func (s *server) Login(ctx context.Context, req *mark2.LoginRequest) (*mark2.LoginResult, error) {

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
	result.Result.Code = mark2.ResultCode_OK

	return result, nil
}
