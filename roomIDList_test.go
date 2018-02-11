package main

import (
	"testing"

	"github.com/ikmski/mark2-server/proto"
)

func TestRoomIDList(t *testing.T) {

	roomIDList := getRoomIDListInstance()
	roomIDList.clear()

	var groupID uint32 = 10

	var roomID1 uint32 = 100001
	var roomID2 uint32 = 100002
	var roomID3 uint32 = 100003

	var status = mark2.RoomStatus_OPEN

	// 空リスト
	list, err := roomIDList.get(groupID, status)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 0 {
		t.Errorf("got %v\nwant %v", len(list), 0)
	}

	// 追加
	err = roomIDList.add(groupID, status, roomID1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = roomIDList.add(groupID, status, roomID2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = roomIDList.add(groupID, status, roomID3)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	list, err = roomIDList.get(groupID, status)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 3 {
		t.Errorf("got %v\nwant %v", len(list), 3)
	}

	// 削除
	err = roomIDList.remove(groupID, status, roomID2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// リスト取得
	list, err = roomIDList.get(groupID, status)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 2 {
		t.Errorf("got %v\nwant %v", len(list), 2)
	}

	id1 := list[0]
	if id1 != roomID1 {
		t.Errorf("got %v\nwant %v", id1, roomID1)
	}

	id3 := list[1]
	if id3 != roomID3 {
		t.Errorf("got %v\nwant %v", id3, roomID3)
	}
}
