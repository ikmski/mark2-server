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

func newRoomWithRoomID(id uint32) *room {
	r := newRoom()
	r.info.Id = id
	return r
}

func createRoom(groupID uint32, capacity uint32) (*room, error) {

	roomStorage := newRoomStorage()

	id, err := roomStorage.createNewRoomID()
	if err != nil {
		return nil, err
	}

	// ルーム作成
	room := newRoomWithRoomID(id)
	room.info.GroupId = groupID
	room.info.Capacity = capacity
	room.info.UserIdList = make([]uint32, 0, capacity)

	// 保存
	err = roomStorage.setRoomInfoByRoomID(id, room.info)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *room) remove() error {

	roomStorage := newRoomStorage()

	err := roomStorage.removeRoomInfoByRoomID(r.info.Id)
	if err != nil {
		return err
	}

	if r.info.Status != mark2.RoomStatus_CLOSED {
		err = roomStorage.removeRoomInfoListByStatus(r.info.GroupId, r.info.Status, r.info)
		if err != nil {
			return err
		}
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
		err := fmt.Errorf("cannot join the room [roomID: %d]", r.info.Id)
		return err
	}

	roomStorage := newRoomStorage()

	// 一旦リストから抜ける
	err := roomStorage.removeRoomInfoListByStatus(r.info.GroupId, r.info.Status, r.info)
	if err != nil {
		return err
	}

	r.info.UserIdList = append(r.info.UserIdList, userID)
	listSize := len(r.info.UserIdList)
	if uint32(listSize) >= r.info.Capacity {
		r.info.Status = mark2.RoomStatus_CLOSED
	} else {
		r.info.Status = mark2.RoomStatus_OPEN
		err := roomStorage.addRoomInfoListByStatus(r.info.GroupId, r.info.Status, r.info)
		if err != nil {
			return err
		}
	}

	err = roomStorage.setRoomInfoByRoomID(r.info.Id, r.info)
	if err != nil {
		return err
	}

	return nil
}

func (r *room) exit(userID uint32) error {

	roomStorage := newRoomStorage()

	// 一旦リストから抜ける
	err := roomStorage.removeRoomInfoListByStatus(r.info.GroupId, r.info.Status, r.info)
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
	} else {
		r.info.Status = mark2.RoomStatus_OPEN
		err := roomStorage.addRoomInfoListByStatus(r.info.GroupId, r.info.Status, r.info)
		if err != nil {
			return err
		}
	}

	return nil
}
