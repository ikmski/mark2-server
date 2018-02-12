package main

import (
	"fmt"
	"sync"

	"github.com/ikmski/mark2-server/proto"
)

var messageIDMutex sync.Mutex
var initialMessageID uint32 = 1000000
var currentMessageID = initialMessageID

func initializeMessageID() {
	currentMessageID = initialMessageID
}

func issueMessageID() uint32 {

	messageIDMutex.Lock()
	defer messageIDMutex.Unlock()

	currentMessageID++

	return currentMessageID
}

func (s *server) SendMessage(srv mark2.MessageService_SendMessageServer) error {

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

			result.Code = mark2.ResultCode_OK
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
