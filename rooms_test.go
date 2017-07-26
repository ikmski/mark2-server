package main

import "testing"

func TestRooms(t *testing.T) {

	rooms := getRoomsInstance()
	rooms.clear()

	var groupID uint32 = 1001

	newID := issueRoomID()
	room1 := newRoom()
	room1.info.Id = newID
	room1.info.GroupId = groupID

	has := rooms.has(room1.info.Id)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	err := rooms.set(room1.info.Id, room1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	has = rooms.has(room1.info.Id)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	u, err := rooms.get(room1.info.Id)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if u != room1 {
		t.Errorf("got %v\nwant %v", u, room1)
	}

	err = rooms.del(room1.info.Id)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	has = rooms.has(room1.info.Id)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

}
