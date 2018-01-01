package main

import (
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"
)

func TestRoomCreateRoom(t *testing.T) {

	var groupID uint32 = 1001
	var capacity uint32 = 3
	var userID uint32 = 10001

	room, err := createRoom(groupID, capacity, userID)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if room.info.GroupId != groupID {
		t.Errorf("got %v\nwant %v", room.info.GroupId, groupID)
	}
	if room.info.Capacity != capacity {
		t.Errorf("got %v\nwant %v", room.info.Capacity, capacity)
	}
	if room.info.Status != mark2.RoomStatus_OPEN {
		t.Errorf("got %v\nwant %v", room.info.Status, mark2.RoomStatus_OPEN)
	}

	err = room.remove()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestRoomJoinRoom(t *testing.T) {

	var groupID uint32 = 1001
	var capacity uint32 = 3

	var userID1 uint32 = 10001
	var userID2 uint32 = 10002
	var userID3 uint32 = 10003

	room, err := createRoom(groupID, capacity, userID1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	canJoin := room.canJoin()
	if !canJoin {
		t.Errorf("got %v\nwant %v", canJoin, true)
	}

	err = room.join(userID2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	canJoin = room.canJoin()
	if !canJoin {
		t.Errorf("got %v\nwant %v", canJoin, true)
	}

	err = room.join(userID3)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	canJoin = room.canJoin()
	if canJoin {
		t.Errorf("got %v\nwant %v", canJoin, false)
	}

	err = room.remove()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestRoomIsJoined(t *testing.T) {

	var groupID uint32 = 1001
	var capacity uint32 = 3

	var userID1 uint32 = 10001
	var userID2 uint32 = 10002
	var userID3 uint32 = 10003

	room, err := createRoom(groupID, capacity, userID1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = room.join(userID2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	isJoined := room.isJoined(userID1)
	if !isJoined {
		t.Errorf("got %v\nwant %v", isJoined, true)
	}

	isJoined = room.isJoined(userID2)
	if !isJoined {
		t.Errorf("got %v\nwant %v", isJoined, true)
	}

	isJoined = room.isJoined(userID3)
	if isJoined {
		t.Errorf("got %v\nwant %v", isJoined, false)
	}

	err = room.remove()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}

func TestRoomExitRoom(t *testing.T) {

	var groupID uint32 = 1001
	var capacity uint32 = 2

	var userID1 uint32 = 10001
	var userID2 uint32 = 10002

	room, err := createRoom(groupID, capacity, userID1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	canJoin := room.canJoin()
	if !canJoin {
		t.Errorf("got %v\nwant %v", canJoin, true)
	}

	err = room.join(userID2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	canJoin = room.canJoin()
	if canJoin {
		t.Errorf("got %v\nwant %v", canJoin, false)
	}

	err = room.exit(userID2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	canJoin = room.canJoin()
	if !canJoin {
		t.Errorf("got %v\nwant %v", canJoin, true)
	}

	err = room.exit(userID1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	canJoin = room.canJoin()
	if canJoin {
		t.Errorf("got %v\nwant %v", canJoin, false)
	}

}
