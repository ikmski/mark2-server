package main

import (
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"
)

func TestRoomCreateRoom(t *testing.T) {

	var groupID uint32 = 1001
	var capacity uint32 = 3

	room, err := createRoom(groupID, capacity)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if room.info.GroupId != groupID {
		t.Errorf("got %v\nwant %v", room.info.GroupId, groupID)
	}
	if room.info.Capacity != capacity {
		t.Errorf("got %v\nwant %v", room.info.Capacity, capacity)
	}
	if room.info.Status != mark2.RoomStatus_CLOSED {
		t.Errorf("got %v\nwant %v", room.info.Status, mark2.RoomStatus_CLOSED)
	}

	err = room.remove()
	if err != nil {
		t.Errorf("got %v\n", err)
	}
}
