package main

import mark2 "github.com/ikmski/mark2-server/proto"

type user struct {
	uniqueKey string
	info      *mark2.UserInfo
	roomId    uint32
}

func newUser() *user {
	u := new(user)
	u.info = mark2.NewUserInfo()
	return u
}

func userExists(uniqueKey string) bool {

	userStorage := getUserStorageInstance()

	_, err := userStorage.getUserIdByUniqueKey(uniqueKey)
	if err != nil {
		return false
	}

	return true
}

func fetchOrCreateUser(uniqueKey string, groupId uint32) (*user, error) {

	exists := userExists(uniqueKey)
	if exists {
		return fetchUser(uniqueKey)
	}
	return createUser(uniqueKey, groupId)
}

func createUser(uniqueKey string, groupId uint32) (*user, error) {

	userStorage := getUserStorageInstance()

	id, err := userStorage.createNewUserId()
	if err != nil {
		return nil, err
	}

	// ユーザ作成
	user := newUser()
	user.uniqueKey = uniqueKey
	user.info.GroupId = groupId
	user.info.Id = id

	// 保存
	err = userStorage.setUserInfoByUserId(id, user.info)
	if err != nil {
		return nil, err
	}
	err = userStorage.setUserIdByUniqueKey(uniqueKey, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func fetchUser(uniqueKey string) (*user, error) {

	userStorage := getUserStorageInstance()

	userId, err := userStorage.getUserIdByUniqueKey(uniqueKey)
	if err != nil {
		return nil, err
	}

	user := newUser()
	user.uniqueKey = uniqueKey

	user.info, err = userStorage.getUserInfoByUserId(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *user) remove() error {

	userStorage := getUserStorageInstance()

	err := userStorage.removeUserInfoByUserId(u.info.Id)
	if err != nil {
		return err
	}

	if u.info.Status != mark2.UserStatus_Logout {
		err = userStorage.removeUserInfoListByStatus(u.info.GroupId, u.info.Status, u.info)
		if err != nil {
			return err
		}
	}

	err = userStorage.removeUserIdByUniqueKey(u.uniqueKey)
	if err != nil {
		return err
	}

	return nil
}

func (u *user) changeStatus(newStatus mark2.UserStatus) error {

	if newStatus != u.info.Status {

		userStorage := getUserStorageInstance()

		if u.info.Status != mark2.UserStatus_Logout {

			err := userStorage.removeUserInfoListByStatus(u.info.GroupId, u.info.Status, u.info)
			if err != nil {
				return err
			}

		}

		u.info.Status = newStatus

		err := userStorage.setUserInfoByUserId(u.info.Id, u.info)
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
