package main

import (
	"fmt"

	mark2 "github.com/ikmski/mark2-server/proto"
)

type room struct {
	info *mark2.RoomInfo
}

func newRoom() *room {
	r := new(room)
	r.info = mark2.NewRoomInfo()
	return r
}

func createRoom(groupID uint32, capacity uint32, userID uint32) (*room, error) {

	if capacity == 0 {
		return nil, fmt.Errorf("capacity must be greater than 0")
	}

	newID := issueRoomID()

	room := newRoom()
	room.info.GroupId = groupID
	room.info.Id = newID
	room.info.Capacity = capacity
	room.info.UserIdList = make([]uint32, 0, capacity)
	room.info.UserIdList = append(room.info.UserIdList, userID)

	roomIDList := getRoomIDListInstance()
	err := roomIDList.add(room.info.GroupId, room.info.Status, room.info.Id)
	if err != nil {
		return nil, err
	}

	rooms := getRoomsInstance()
	err = rooms.set(room.info.Id, room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *room) remove() error {

	roomIDList := getRoomIDListInstance()
	err := roomIDList.remove(r.info.GroupId, r.info.Status, r.info.Id)
	if err != nil {
		return err
	}

	rooms := getRoomsInstance()
	err = rooms.del(r.info.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *room) canJoin() bool {

	if r.info == nil {
		return false
	}

	return r.info.Status == mark2.RoomStatus_OPEN
}

func (r *room) join(userID uint32) error {

	if !r.canJoin() {
		return fmt.Errorf("cannot join the room [%d]", r.info.Id)
	}

	roomIDList := getRoomIDListInstance()
	err := roomIDList.remove(r.info.GroupId, r.info.Status, r.info.Id)
	if err != nil {
		return err
	}

	r.info.UserIdList = append(r.info.UserIdList, userID)
	listSize := len(r.info.UserIdList)
	if uint32(listSize) >= r.info.Capacity {
		r.info.Status = mark2.RoomStatus_CLOSED
	} else {
		r.info.Status = mark2.RoomStatus_OPEN
	}
	err = roomIDList.add(r.info.GroupId, r.info.Status, r.info.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *room) exit(userID uint32) error {

	roomIDList := getRoomIDListInstance()

	err := roomIDList.remove(r.info.GroupId, r.info.Status, r.info.Id)
	if err != nil {
		return err
	}

	newList := make([]uint32, 0, r.info.Capacity)
	for _, id := range r.info.UserIdList {
		if id != userID {
			newList = append(newList, id)
		}
	}
	r.info.UserIdList = newList

	listSize := len(r.info.UserIdList)
	if uint32(listSize) <= 0 {
		r.info.Status = mark2.RoomStatus_CLOSED

		rooms := getRoomsInstance()
		err = rooms.del(r.info.Id)
		if err != nil {
			return err
		}

	} else {
		r.info.Status = mark2.RoomStatus_OPEN

		err := roomIDList.add(r.info.GroupId, r.info.Status, r.info.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
