package main

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	mark2 "github.com/ikmski/mark2-server/proto"
)

type roomStorage struct {
	s     *storage
	mutex sync.Mutex
}

func newRoomStorage() *roomStorage {
	rs := new(roomStorage)
	rs.s = getStorageInstance()
	return rs
}

func (rs *roomStorage) clear() {
	rs.s.clear()
}

func (rs *roomStorage) createNewRoomID() (uint32, error) {

	key := "max_room_id"

	rs.mutex.Lock()
	defer rs.mutex.Unlock()

	var id uint32 = 1000000
	has := rs.s.has(key)
	if has {
		var err error
		id, err = rs.s.getUint32(key)
		if err != nil {
			return 0, err
		}
	}

	id++

	err := rs.s.setUint32(key, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (rs *roomStorage) getRoomInfoByRoomID(roomID uint32) (*mark2.RoomInfo, error) {

	key := fmt.Sprintf("room_info_by_room_id.%d", roomID)

	val, err := rs.s.get(key)
	if err != nil {
		return nil, err
	}

	roomInfo := mark2.NewRoomInfo()
	err = proto.Unmarshal(val, roomInfo)
	if err != nil {
		return nil, err
	}

	return roomInfo, nil
}

func (rs *roomStorage) setRoomInfoByRoomID(roomID uint32, roomInfo *mark2.RoomInfo) error {

	key := fmt.Sprintf("room_info_by_room_id.%d", roomID)

	buf, err := proto.Marshal(roomInfo)
	if err != nil {
		return err
	}

	return rs.s.set(key, buf)
}

func (rs *roomStorage) removeRoomInfoByRoomID(roomID uint32) error {

	key := fmt.Sprintf("room_info_by_room_id.%d", roomID)

	return rs.s.del(key)
}

func (rs *roomStorage) getRoomInfoListByStatus(groupID uint32, status mark2.RoomStatus) (*mark2.RoomInfoList, error) {

	key := fmt.Sprintf("room_info_list_by_group_id.%d_status.%s", groupID, status.String())

	roomInfoList := mark2.NewRoomInfoList()

	has := rs.s.has(key)
	if has {

		list, err := rs.s.members(key)
		if err != nil {
			return nil, err
		}

		for _, v := range list {
			info := mark2.NewRoomInfo()
			err = proto.Unmarshal(v, info)
			if err != nil {
				return nil, err
			}

			roomInfoList.List = append(roomInfoList.List, info)
		}
	}

	return roomInfoList, nil
}

func (rs *roomStorage) addRoomInfoListByStatus(groupID uint32, status mark2.RoomStatus, roomInfo *mark2.RoomInfo) error {

	key := fmt.Sprintf("room_info_list_by_group_id.%d_status.%s", groupID, status.String())

	buf, err := proto.Marshal(roomInfo)
	if err != nil {
		return err
	}

	return rs.s.add(key, buf)
}

func (rs *roomStorage) removeRoomInfoListByStatus(groupID uint32, status mark2.RoomStatus, roomInfo *mark2.RoomInfo) error {

	key := fmt.Sprintf("room_info_list_by_group_id.%d_status.%s", groupID, status.String())

	buf, err := proto.Marshal(roomInfo)
	if err != nil {
		return err
	}

	return rs.s.remove(key, buf)
}
