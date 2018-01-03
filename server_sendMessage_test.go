package main

import (
	"context"
	"sync"
	"testing"
	"time"

	mark2 "github.com/ikmski/mark2-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func TestServerSendMessage(t *testing.T) {

	// Set up a connection to the Server.
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mark2.NewMessageServiceClient(conn)

	var groupID uint32 = 10001

	// Login
	loginRequest := new(mark2.LoginRequest)
	loginRequest.GroupId = groupID

	loginResult, err := c.Login(context.Background(), loginRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	token1 := loginResult.AccessToken

	loginResult, err = c.Login(context.Background(), loginRequest)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	token2 := loginResult.AccessToken

	// Matching
	{
		ch := make(chan struct{})
		go func() {

			matchRequest := new(mark2.MatchRequest)
			matchRequest.Token = token1
			_, err := c.MatchRandom(context.Background(), matchRequest)
			if err != nil {
				t.Errorf("got %v\n", err)
			}

			close(ch)
		}()

		time.Sleep(1 * time.Second)

		matchRequest := new(mark2.MatchRequest)
		matchRequest.Token = token2
		_, err := c.MatchRandom(context.Background(), matchRequest)
		if err != nil {
			t.Errorf("got %v\n", err)
		}

		<-ch
	}

	// WaitMessage
	wg := sync.WaitGroup{}
	{
		// WaitMessage
		wg.Add(1)
		go func() {

			c1, err := c.WaitMessage(context.Background(), token1)
			if err != nil {
				t.Errorf("got %v\n", err)
			}

			message, err := c1.Recv()
			if err != nil {
				t.Errorf("got %v\n", err)
			}
			if message.Content == "" {
				t.Errorf("got %v\n", message)
			}

			wg.Done()
		}()

		wg.Add(1)
		go func() {

			c2, err := c.WaitMessage(context.Background(), token2)
			if err != nil {
				t.Errorf("got %v\n", err)
			}

			message, err := c2.Recv()
			if err != nil {
				t.Errorf("got %v\n", err)
			}
			if message.Content == "" {
				t.Errorf("got %v\n", message)
			}

			wg.Done()
		}()
	}

	// Send Message
	{
		time.Sleep(1 * time.Second)

		c2, err := c.SendStream(context.Background())
		if err != nil {
			t.Errorf("got %v\n", err)
		}

		messageReqest := new(mark2.MessageRequest)
		messageReqest.Token = token2
		messageReqest.Content = "hoge"
		err = c2.Send(messageReqest)
		if err != nil {
			t.Errorf("got %v\n", err)
		}

		wg.Wait()
	}

	// Logout
	{
		_, err = c.Logout(context.Background(), token1)
		if err != nil {
			t.Errorf("got %v\n", err)
		}
		_, err = c.Logout(context.Background(), token2)
		if err != nil {
			t.Errorf("got %v\n", err)
		}
	}

}
