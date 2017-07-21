package main

import (
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"
)

func TestUserStorageUserIdByUniquekey(t *testing.T) {

	uniqueKey := "test_unique_key"
	var userId uint32 = 1001

	userStorage := newUserStorage()
	userStorage.clear()

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

	userStorage := newUserStorage()
	userStorage.clear()

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

	userStorage1 := newUserStorage()
	userStorage2 := newUserStorage()
	userStorage1.clear()
	userStorage2.clear()

	userId, err := userStorage1.createNewUserId()
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	userInfo := new(mark2.UserInfo)

	userInfo.GroupId = 1111
	userInfo.Id = 2222

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

	err = userStorage1.removeUserInfoByUserId(userId)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestUserStorageUserInfoList(t *testing.T) {

	userStorage := newUserStorage()
	userStorage.clear()

	userInfo := new(mark2.UserInfo)
	var groupId uint32 = 10

	var userId1 uint32 = 100001
	var userId2 uint32 = 100002
	var userId3 uint32 = 100003

	// 空リスト
	list, err := userStorage.getUserInfoListByStatus(groupId, mark2.UserStatus_Login)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list.GetList()) != 0 {
		t.Errorf("got %v\nwant %v", len(list.GetList()), 0)
	}

	// 追加
	userInfo.Id = userId1
	err = userStorage.addUserInfoListByStatus(groupId, mark2.UserStatus_Login, userInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	userInfo.Id = userId2
	err = userStorage.addUserInfoListByStatus(groupId, mark2.UserStatus_Login, userInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	userInfo.Id = userId3
	err = userStorage.addUserInfoListByStatus(groupId, mark2.UserStatus_Login, userInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	list, err = userStorage.getUserInfoListByStatus(groupId, mark2.UserStatus_Login)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list.GetList()) != 3 {
		t.Errorf("got %v\nwant %v", len(list.GetList()), 3)
	}

	// 削除
	userInfo.Id = userId2
	err = userStorage.removeUserInfoListByStatus(groupId, mark2.UserStatus_Login, userInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// リスト取得
	list, err = userStorage.getUserInfoListByStatus(groupId, mark2.UserStatus_Login)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list.GetList()) != 2 {
		t.Errorf("got %v\nwant %v", len(list.GetList()), 2)
	}

	info1 := list.GetList()[0]
	if info1.Id != userId1 {
		t.Errorf("got %v\nwant %v", info1.Id, userId1)
	}

	info3 := list.GetList()[1]
	if info3.Id != userId3 {
		t.Errorf("got %v\nwant %v", info3.Id, userId3)
	}
}
