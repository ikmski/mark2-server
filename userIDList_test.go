package main

import (
	"testing"

	"github.com/ikmski/mark2-server/proto"
)

func TestUserIDList(t *testing.T) {

	userIDList := getUserIDListInstance()
	userIDList.clear()

	var groupID uint32 = 10

	var userID1 uint32 = 100001
	var userID2 uint32 = 100002
	var userID3 uint32 = 100003

	var status = mark2.UserStatus_Login

	// 空リスト
	list, err := userIDList.get(groupID, status)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 0 {
		t.Errorf("got %v\nwant %v", len(list), 0)
	}

	// 追加
	err = userIDList.add(groupID, status, userID1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = userIDList.add(groupID, status, userID2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = userIDList.add(groupID, status, userID3)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	list, err = userIDList.get(groupID, status)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 3 {
		t.Errorf("got %v\nwant %v", len(list), 3)
	}

	// 削除
	err = userIDList.remove(groupID, status, userID2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// リスト取得
	list, err = userIDList.get(groupID, status)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 2 {
		t.Errorf("got %v\nwant %v", len(list), 2)
	}

	id1 := list[0]
	if id1 != userID1 {
		t.Errorf("got %v\nwant %v", id1, userID1)
	}

	id3 := list[1]
	if id3 != userID3 {
		t.Errorf("got %v\nwant %v", id3, userID3)
	}
}
