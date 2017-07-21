package main

import mark2 "github.com/ikmski/mark2-server/proto"

type user struct {
	uniqueKey string
	info      *mark2.UserInfo
	roomID    uint32
}

func newUser() *user {
	u := new(user)
	u.info = mark2.NewUserInfo()
	return u
}

func newUserWithUserID(id uint32) *user {
	u := newUser()
	u.info.Id = id
	return u
}

func userExists(uniqueKey string) bool {

	userStorage := newUserStorage()

	_, err := userStorage.getUserIDByUniqueKey(uniqueKey)
	if err != nil {
		return false
	}

	return true
}

func fetchOrCreateUser(uniqueKey string, groupID uint32) (*user, error) {

	exists := userExists(uniqueKey)
	if exists {
		return fetchUser(uniqueKey)
	}
	return createUser(uniqueKey, groupID)
}

func createUser(uniqueKey string, groupID uint32) (*user, error) {

	userStorage := newUserStorage()

	id, err := userStorage.createNewUserID()
	if err != nil {
		return nil, err
	}

	// ユーザ作成
	user := newUser()
	user.uniqueKey = uniqueKey
	user.info.GroupId = groupID
	user.info.Id = id

	// 保存
	err = userStorage.setUserInfoByUserID(id, user.info)
	if err != nil {
		return nil, err
	}
	err = userStorage.setUserIDByUniqueKey(uniqueKey, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func fetchUser(uniqueKey string) (*user, error) {

	userStorage := newUserStorage()

	userID, err := userStorage.getUserIDByUniqueKey(uniqueKey)
	if err != nil {
		return nil, err
	}

	user := newUser()
	user.uniqueKey = uniqueKey

	user.info, err = userStorage.getUserInfoByUserID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) remove() error {

	userStorage := newUserStorage()

	err := userStorage.removeUserInfoByUserID(u.info.Id)
	if err != nil {
		return err
	}

	if u.info.Status != mark2.UserStatus_Logout {
		err = userStorage.removeUserInfoListByStatus(u.info.GroupId, u.info.Status, u.info)
		if err != nil {
			return err
		}
	}

	err = userStorage.removeUserIDByUniqueKey(u.uniqueKey)
	if err != nil {
		return err
	}

	return nil
}

func (u *user) changeStatus(newStatus mark2.UserStatus) error {

	if newStatus != u.info.Status {

		userStorage := newUserStorage()

		if u.info.Status != mark2.UserStatus_Logout {

			err := userStorage.removeUserInfoListByStatus(u.info.GroupId, u.info.Status, u.info)
			if err != nil {
				return err
			}

		}

		u.info.Status = newStatus

		err := userStorage.setUserInfoByUserID(u.info.Id, u.info)
		if err != nil {
			return err
		}

		if newStatus != mark2.UserStatus_Logout {

			err = userStorage.addUserInfoListByStatus(u.info.GroupId, u.info.Status, u.info)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
