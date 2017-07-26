package main

import mark2 "github.com/ikmski/mark2-server/proto"

type user struct {
	info   *mark2.UserInfo
	roomID uint32
}

func newUser() *user {
	u := new(user)
	u.info = mark2.NewUserInfo()
	return u
}

func createUser(groupID uint32) (*user, error) {

	newID := issueUserID()

	user := newUser()
	user.info.GroupId = groupID
	user.info.Id = newID

	userIDList := getUserIDListInstance()
	err := userIDList.add(user.info.GroupId, user.info.Status, user.info.Id)
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

	userIDList := getUserIDListInstance()
	err := userIDList.remove(u.info.GroupId, u.info.Status, u.info.Id)
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

		userIDList := getUserIDListInstance()

		err := userIDList.remove(u.info.GroupId, u.info.Status, u.info.Id)
		if err != nil {
			return err
		}

		u.info.Status = newStatus

		err = userIDList.add(u.info.GroupId, u.info.Status, u.info.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
