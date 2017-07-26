package main

import (
	"fmt"
	"sync"
)

type rooms struct {
	mutex   sync.Mutex
	roomMap map[uint32]*room
}

var roomsInstance = newRooms()

func newRooms() *rooms {
	r := new(rooms)
	r.roomMap = make(map[uint32]*room)
	return r
}

func getRoomsInstance() *rooms {
	return roomsInstance
}

func (r *rooms) clear() {
	r.roomMap = make(map[uint32]*room)
}

func (r *rooms) has(key uint32) bool {
	_, ok := r.roomMap[key]
	return ok
}

func (r *rooms) get(key uint32) (*room, error) {

	val, ok := r.roomMap[key]
	if !ok {
		err := fmt.Errorf("%d not found", key)
		return nil, err
	}

	return val, nil
}

func (r *rooms) set(key uint32, value *room) error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.roomMap[key] = value

	return nil
}

func (r *rooms) del(key uint32) error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.has(key) {
		return fmt.Errorf("%d not found", key)
	}
	delete(r.roomMap, key)

	return nil
}
