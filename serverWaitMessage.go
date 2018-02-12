package main

import (
	"fmt"
	"time"

	"github.com/ikmski/mark2-server/proto"
)

func (s *server) WaitMessage(token *mark2.AccessToken, srv mark2.MessageService_WaitMessageServer) error {

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
