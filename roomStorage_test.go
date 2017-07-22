package main

import (
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"
)

func TestRoomStorageCreateNewRoomID(t *testing.T) {

	roomStorage := newRoomStorage()
	roomStorage.clear()

	id, err := roomStorage.createNewRoomID()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if id != 1000001 {
		t.Errorf("got %v\nwant %v", id, 1000001)
	}

	id, err = roomStorage.createNewRoomID()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if id != 1000002 {
		t.Errorf("got %v\nwant %v", id, 1000002)
	}

	id, err = roomStorage.createNewRoomID()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if id != 1000003 {
		t.Errorf("got %v\nwant %v", id, 1000003)
	}
}

func TestRoomStorageRoomInfo(t *testing.T) {

	roomStorage1 := newRoomStorage()
	roomStorage2 := newRoomStorage()
	roomStorage1.clear()
	roomStorage2.clear()

	roomID, err := roomStorage1.createNewRoomID()
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	roomInfo := new(mark2.RoomInfo)

	roomInfo.GroupId = 1111
	roomInfo.Id = 2222
	roomInfo.Capacity = 5

	err = roomStorage1.setRoomInfoByRoomID(roomID, roomInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	info, err := roomStorage2.getRoomInfoByRoomID(roomID)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	if info.GroupId != roomInfo.GroupId {
		t.Errorf("got %v\nwant %v", info.GroupId, roomInfo.GroupId)
	}
	if info.Id != roomInfo.Id {
		t.Errorf("got %v\nwant %v", info.Id, roomInfo.Id)
	}
	if info.Capacity != roomInfo.Capacity {
		t.Errorf("got %v\nwant %v", info.Capacity, roomInfo.Capacity)
	}

	err = roomStorage1.removeRoomInfoByRoomID(roomID)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestRoomStorageRoomInfoList(t *testing.T) {

	roomStorage := newRoomStorage()
	roomStorage.clear()

	roomInfo := new(mark2.RoomInfo)
	var groupId uint32 = 10

	var roomId1 uint32 = 100001
	var roomId2 uint32 = 100002
	var roomId3 uint32 = 100003

	// 空リスト
	list, err := roomStorage.getRoomInfoListByStatus(groupId, mark2.RoomStatus_CLOSED)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list.GetList()) != 0 {
		t.Errorf("got %v\nwant %v", len(list.GetList()), 0)
	}

	// 追加
	roomInfo.Id = roomId1
	err = roomStorage.addRoomInfoListByStatus(groupId, mark2.RoomStatus_CLOSED, roomInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	roomInfo.Id = roomId2
	err = roomStorage.addRoomInfoListByStatus(groupId, mark2.RoomStatus_CLOSED, roomInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	roomInfo.Id = roomId3
	err = roomStorage.addRoomInfoListByStatus(groupId, mark2.RoomStatus_CLOSED, roomInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	list, err = roomStorage.getRoomInfoListByStatus(groupId, mark2.RoomStatus_CLOSED)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list.GetList()) != 3 {
		t.Errorf("got %v\nwant %v", len(list.GetList()), 3)
	}

	// 削除
	roomInfo.Id = roomId2
	err = roomStorage.removeRoomInfoListByStatus(groupId, mark2.RoomStatus_CLOSED, roomInfo)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// リスト取得
	list, err = roomStorage.getRoomInfoListByStatus(groupId, mark2.RoomStatus_CLOSED)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list.GetList()) != 2 {
		t.Errorf("got %v\nwant %v", len(list.GetList()), 2)
	}

	info1 := list.GetList()[0]
	if info1.Id != roomId1 {
		t.Errorf("got %v\nwant %v", info1.Id, roomId1)
	}

	info3 := list.GetList()[1]
	if info3.Id != roomId3 {
		t.Errorf("got %v\nwant %v", info3.Id, roomId3)
	}
}
