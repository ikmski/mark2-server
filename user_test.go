package main

import (
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"
)

func TestUserCreateUser(t *testing.T) {

	var groupId uint32 = 1001

	user, err := createUser(groupId)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if user.info.GroupId != groupId {
		t.Errorf("got %v\nwant %v", user.info.GroupId, groupId)
	}
	if user.info.Status != mark2.UserStatus_Login {
		t.Errorf("got %v\nwant %v", user.info.Status, mark2.UserStatus_Login)
	}

	err = user.remove()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestUserChangeStatus(t *testing.T) {

	var groupId uint32 = 1001

	user, err := createUser(groupId)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = user.changeStatus(mark2.UserStatus_WaitMatch)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if user.info.Status != mark2.UserStatus_WaitMatch {
		t.Errorf("got %v\nwant %v", user.info.Status, mark2.UserStatus_WaitMatch)
	}

	err = user.changeStatus(mark2.UserStatus_Matched)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if user.info.Status != mark2.UserStatus_Matched {
		t.Errorf("got %v\nwant %v", user.info.Status, mark2.UserStatus_Matched)
	}

	err = user.remove()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}
