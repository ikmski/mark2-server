package main

import mark2 "github.com/ikmski/mark2-server/proto"

type user struct {
	info   mark2.UserInfo
	status mark2.UserStatus
	roomId uint32
}

func newUser() *user {
	u := new(user)
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

func createUser(uniqueKey string, groupId uint32) (*user, error) {

	userStorage := getUserStorageInstance()

	id, err := userStorage.createNewUserId()
	if err != nil {
		return nil, err
	}

	// ユーザ作成
	user := newUser()
	user.info.GroupId = groupId
	user.info.Id = id

	// 保存
	err = userStorage.setUserInfoByUserId(id, &user.info)
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

	userInfo, err := userStorage.getUserInfoByUserId(userId)
	if err != nil {
		return nil, err
	}

	user := newUser()
	user.info = *userInfo

	return user, nil
}
