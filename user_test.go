package main

import (
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"
)

func TestUserCreateUser(t *testing.T) {

	var groupId uint32 = 1001
	var uniqueKey string = "test_unique_key"

	exists := userExists(uniqueKey)
	if exists {
		t.Errorf("got %v\nwant %v", exists, false)
	}

	user, err := createUser(uniqueKey, groupId)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if user.info.GroupId != groupId {
		t.Errorf("got %v\nwant %v", user.info.GroupId, groupId)
	}
	if user.info.Status != mark2.UserStatus_Logout {
		t.Errorf("got %v\nwant %v", user.info.Status, mark2.UserStatus_Logout)
	}

	err = user.remove()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestUserChangeStatus(t *testing.T) {

	var groupId uint32 = 1001
	var uniqueKey string = "test_unique_key"

	exists := userExists(uniqueKey)
	if exists {
		t.Errorf("got %v\nwant %v", exists, false)
	}

	user, err := createUser(uniqueKey, groupId)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = user.changeStatus(mark2.UserStatus_Login)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if user.info.Status != mark2.UserStatus_Login {
		t.Errorf("got %v\nwant %v", user.info.Status, mark2.UserStatus_Login)
	}

	err = user.changeStatus(mark2.UserStatus_Logout)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if user.info.Status != mark2.UserStatus_Logout {
		t.Errorf("got %v\nwant %v", user.info.Status, mark2.UserStatus_Logout)
	}

	err = user.remove()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}
