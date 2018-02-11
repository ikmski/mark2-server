package main

import (
	"fmt"
	"sync"

	"github.com/ikmski/mark2-server/proto"
)

type userIDList struct {
	mutex sync.Mutex
	idSet map[string]([]uint32)
}

var userIDListInstance = newUserIDList()
var userIDListKeyFormat = "user_list_by_group_id.%d_status.%s"

func newUserIDList() *userIDList {
	u := new(userIDList)
	u.idSet = make(map[string]([]uint32))
	return u
}

func getUserIDListInstance() *userIDList {
	return userIDListInstance
}

func (u *userIDList) clear() {
	u.idSet = make(map[string]([]uint32))
}

func (u *userIDList) get(groupID uint32, status mark2.UserStatus) ([]uint32, error) {

	key := fmt.Sprintf(userIDListKeyFormat, groupID, status.String())

	list, ok := u.idSet[key]
	if !ok {
		list = make([]uint32, 0)
	}

	return list, nil
}

func (u *userIDList) add(groupID uint32, status mark2.UserStatus, userID uint32) error {

	key := fmt.Sprintf(userIDListKeyFormat, groupID, status.String())

	u.mutex.Lock()
	defer u.mutex.Unlock()

	list, ok := u.idSet[key]
	if !ok {
		list = make([]uint32, 0)
	}

	list = append(list, userID)
	u.idSet[key] = list

	return nil
}

func (u *userIDList) remove(groupID uint32, status mark2.UserStatus, userID uint32) error {

	key := fmt.Sprintf(userIDListKeyFormat, groupID, status.String())

	list, ok := u.idSet[key]
	if !ok {
		err := fmt.Errorf("%s not found", key)
		return err
	}

	u.mutex.Lock()
	defer u.mutex.Unlock()

	newList := make([]uint32, 0, len(list))
	for _, v := range list {
		if userID != v {
			newList = append(newList, v)
		}
	}

	delete(u.idSet, key)
	u.idSet[key] = newList

	return nil
}
