package main

import (
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"
)

func TestUserStorageUserIdByUniquekey(t *testing.T) {

	uniqueKey := "test_unique_key"
	userId := 1001

	userStorage := getUserStorageInstance()

	_, err := userStorage.getUserIdByUniqueKey(uniqueKey)
	if err == nil {
		t.Errorf("got %v\n", err)
	}

	err = userStorage.setUserIdByUniqueKey(uniqueKey, userId)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	v, err := userStorage.getUserIdByUniqueKey(uniqueKey)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if v != userId {
		t.Errorf("got %v\nwant %v", v, userId)
	}

	err = userStorage.removeUserIdByUniqueKey(uniqueKey)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	_, err = userStorage.getUserIdByUniqueKey(uniqueKey)
	if err == nil {
		t.Errorf("got %v\n", err)
	}
}

func TestUserStorageCreateNewUserId(t *testing.T) {

	userStorage := getUserStorageInstance()

	id, err := userStorage.createNewUserId()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if id != 1000001 {
		t.Errorf("got %v\nwant %v", id, 1000001)
	}

	id, err = userStorage.createNewUserId()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if id != 1000002 {
		t.Errorf("got %v\nwant %v", id, 1000002)
	}

	id, err = userStorage.createNewUserId()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if id != 1000003 {
		t.Errorf("got %v\nwant %v", id, 1000003)
	}
}

func TestUserStorageUserInfo(t *testing.T) {

	userStorage1 := getUserStorageInstance()
	userStorage2 := getUserStorageInstance()

	userId, err := userStorage1.createNewUserId()
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	userInfo := new(mark2.UserInfo)

	userInfo.GroupId = 1111
	userInfo.Id = 2222
	userInfo.Name = "test_user_name"

	err = userStorage1.setUserInfoByUserId(userId, userInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	info, err := userStorage2.getUserInfoByUserId(userId)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	if info.GroupId != userInfo.GroupId {
		t.Errorf("got %v\nwant %v", info.GroupId, userInfo.GroupId)
	}
	if info.Id != userInfo.Id {
		t.Errorf("got %v\nwant %v", info.Id, userInfo.Id)
	}
	if info.Name != userInfo.Name {
		t.Errorf("got %v\nwant %v", info.Name, userInfo.Name)
	}

	err = userStorage1.removeUserInfoByUserId(userId)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}
