package main

import (
	"fmt"
	"sync"

	"github.com/ikmski/mark2-server/proto"
)

type roomIDList struct {
	mutex sync.Mutex
	idSet map[string]([]uint32)
}

var roomIDListInstance = newRoomIDList()
var roomIDListKeyFormat = "room_list_by_group_id.%d_status.%s"

func newRoomIDList() *roomIDList {
	r := new(roomIDList)
	r.idSet = make(map[string]([]uint32))
	return r
}

func getRoomIDListInstance() *roomIDList {
	return roomIDListInstance
}

func (r *roomIDList) clear() {
	r.idSet = make(map[string]([]uint32))
}

func (r *roomIDList) get(groupID uint32, status mark2.RoomStatus) ([]uint32, error) {

	key := fmt.Sprintf(roomIDListKeyFormat, groupID, status.String())

	list, ok := r.idSet[key]
	if !ok {
		list = make([]uint32, 0)
	}

	return list, nil
}

func (r *roomIDList) add(groupID uint32, status mark2.RoomStatus, roomID uint32) error {

	key := fmt.Sprintf(roomIDListKeyFormat, groupID, status.String())

	r.mutex.Lock()
	defer r.mutex.Unlock()

	list, ok := r.idSet[key]
	if !ok {
		list = make([]uint32, 0)
	}

	list = append(list, roomID)
	r.idSet[key] = list

	return nil
}

func (r *roomIDList) remove(groupID uint32, status mark2.RoomStatus, roomID uint32) error {

	key := fmt.Sprintf(roomIDListKeyFormat, groupID, status.String())

	list, ok := r.idSet[key]
	if !ok {
		err := fmt.Errorf("%s not found", key)
		return err
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	newList := make([]uint32, 0, len(list))
	for _, v := range list {
		if roomID != v {
			newList = append(newList, v)
		}
	}

	delete(r.idSet, key)
	r.idSet[key] = newList

	return nil
}
