package main

import (
	"fmt"
	"sync"
)

type users struct {
	mutex   sync.Mutex
	userMap map[uint32]*user
}

var usersInstance = newUsers()

func newUsers() *users {
	u := new(users)
	u.userMap = make(map[uint32]*user)
	return u
}

func getUsersInstance() *users {
	return usersInstance
}

func (u *users) clear() {
	u.userMap = make(map[uint32]*user)
}

func (u *users) has(key uint32) bool {
	_, ok := u.userMap[key]
	return ok
}

func (u *users) get(key uint32) (*user, error) {

	val, ok := u.userMap[key]
	if !ok {
		err := fmt.Errorf("%d not found", key)
		return nil, err
	}

	return val, nil
}

func (u *users) set(key uint32, value *user) error {

	u.mutex.Lock()
	defer u.mutex.Unlock()

	u.userMap[key] = value

	return nil
}

func (u *users) del(key uint32) error {

	u.mutex.Lock()
	defer u.mutex.Unlock()

	if !u.has(key) {
		return fmt.Errorf("%d not found", key)
	}
	delete(u.userMap, key)

	return nil
}
