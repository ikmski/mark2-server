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

func TestUserStorageUserInfoList(t *testing.T) {

	userStorage := getUserStorageInstance()

	userInfo := new(mark2.UserInfo)
	groupId := 10

	var userId1 int32 = 100001
	var userId2 int32 = 100002
	var userId3 int32 = 100003

	userName1 := "test_user_name_1"
	userName2 := "test_user_name_2"
	userName3 := "test_user_name_3"

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
	userInfo.Name = userName1
	err = userStorage.addUserInfoListByStatus(groupId, mark2.UserStatus_Login, userInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	userInfo.Id = userId2
	userInfo.Name = userName2
	err = userStorage.addUserInfoListByStatus(groupId, mark2.UserStatus_Login, userInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	userInfo.Id = userId3
	userInfo.Name = userName3
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
	userInfo.Name = userName2
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
	if info1.Name != userName1 {
		t.Errorf("got %v\nwant %v", info1.Name, userName1)
	}

	info3 := list.GetList()[1]
	if info3.Id != userId3 {
		t.Errorf("got %v\nwant %v", info3.Id, userId3)
	}
	if info3.Name != userName3 {
		t.Errorf("got %v\nwant %v", info3.Name, userName3)
	}
}
