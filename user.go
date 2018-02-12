package main

import (
	"fmt"
	"sync"

	"github.com/ikmski/mark2-server/proto"
)

var userIDMutex sync.Mutex
var initialUserID uint32 = 1000000
var currentUserID = initialUserID

type user struct {
	info   *mark2.UserInfo
	roomID uint32
}

func initializeUserID() {
	currentUserID = initialUserID
}

func issueUserID() uint32 {

	userIDMutex.Lock()
	defer userIDMutex.Unlock()

	currentUserID++

	return currentUserID
}

func createUserIdListKey(groupID uint32, status mark2.UserStatus) string {

	key := fmt.Sprintf("user_id_list_by_group_id.%d_status.%s", groupID, status.String())
	return key
}

func newUser() *user {
	u := new(user)
	u.info = mark2.NewUserInfo()
	u.roomID = 0
	return u
}

func createUser(groupID uint32) (*user, error) {

	newID := issueUserID()

	user := newUser()
	user.info.GroupId = groupID
	user.info.Id = newID

	set := getUint32SetInstance()
	err := set.add(createUserIdListKey(user.info.GroupId, user.info.Status), user.info.Id)
	if err != nil {
		return nil, err
	}

	users := getUsersInstance()
	err = users.set(user.info.Id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) remove() error {

	set := getUint32SetInstance()
	err := set.remove(createUserIdListKey(u.info.GroupId, u.info.Status), u.info.Id)
	if err != nil {
		return err
	}

	users := getUsersInstance()
	err = users.del(u.info.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *user) changeStatus(newStatus mark2.UserStatus) error {

	if newStatus != u.info.Status {

		set := getUint32SetInstance()

		err := set.remove(createUserIdListKey(u.info.GroupId, u.info.Status), u.info.Id)
		if err != nil {
			return err
		}

		u.info.Status = newStatus

		err = set.add(createUserIdListKey(u.info.GroupId, u.info.Status), u.info.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
