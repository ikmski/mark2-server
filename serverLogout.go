package main

import (
	"context"

	"github.com/ikmski/mark2-server/proto"
)

func (s *server) Logout(ctx context.Context, token *mark2.AccessToken) (*mark2.Result, error) {

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

	result.Code = mark2.ResultCode_OK

	return result, nil
}
