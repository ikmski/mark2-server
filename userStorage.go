package main

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	mark2 "github.com/ikmski/mark2-server/proto"
)

type userStorage struct {
	mutex sync.Mutex
}

var userStorageInstance *userStorage = newUserStorage()

func newUserStorage() *userStorage {
	us := new(userStorage)
	return us
}

func getUserStorageInstance() *userStorage {
	return userStorageInstance
}

func (_ *userStorage) clear() {
	storage := getStorageInstance()
	storage.clear()
}

func (_ *userStorage) getUserIdByUniqueKey(uniqueKey string) (uint32, error) {

	storage := getStorageInstance()
	key := "user_id_by_unique_key." + uniqueKey
	return storage.getUint32(key)
}

func (_ *userStorage) setUserIdByUniqueKey(uniqueKey string, userId uint32) error {

	storage := getStorageInstance()
	key := "user_id_by_unique_key." + uniqueKey
	return storage.setUint32(key, userId)
}

func (_ *userStorage) removeUserIdByUniqueKey(uniqueKey string) error {

	storage := getStorageInstance()
	key := "user_id_by_unique_key." + uniqueKey
	return storage.del(key)
}

func (us *userStorage) createNewUserId() (uint32, error) {

	storage := getStorageInstance()
	key := "max_user_id"

	us.mutex.Lock()
	defer us.mutex.Unlock()

	var id uint32 = 1000000
	has := storage.has(key)
	if has {
		var err error
		id, err = storage.getUint32(key)
		if err != nil {
			return 0, err
		}
	}

	id++

	err := storage.setUint32(key, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (_ *userStorage) getUserInfoByUserId(userId uint32) (*mark2.UserInfo, error) {

	storage := getStorageInstance()
	key := fmt.Sprintf("user_info_by_user_id.%d", userId)

	val, err := storage.get(key)
	if err != nil {
		return nil, err
	}

	userInfo := mark2.NewUserInfo()
	err = proto.Unmarshal(val, userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (_ *userStorage) setUserInfoByUserId(userId uint32, userInfo *mark2.UserInfo) error {

	storage := getStorageInstance()
	key := fmt.Sprintf("user_info_by_user_id.%d", userId)

	buf, err := proto.Marshal(userInfo)
	if err != nil {
		return err
	}

	return storage.set(key, buf)
}

func (_ *userStorage) removeUserInfoByUserId(userId uint32) error {

	storage := getStorageInstance()
	key := fmt.Sprintf("user_info_by_user_id.%d", userId)

	return storage.del(key)
}

func (_ *userStorage) getUserInfoListByStatus(groupId uint32, status mark2.UserStatus) (*mark2.UserInfoList, error) {

	storage := getStorageInstance()
	key := fmt.Sprintf("user_info_list_by_group_id.%d_status.%s", groupId, status.String())

	userInfoList := new(mark2.UserInfoList)

	has := storage.has(key)
	if has {

		list, err := storage.members(key)
		if err != nil {
			return nil, err
		}

		for _, v := range list {
			info := mark2.NewUserInfo()
			err = proto.Unmarshal(v, info)
			if err != nil {
				return nil, err
			}

			userInfoList.List = append(userInfoList.List, info)
		}
	}

	return userInfoList, nil
}

func (_ *userStorage) addUserInfoListByStatus(groupId uint32, status mark2.UserStatus, userInfo *mark2.UserInfo) error {

	storage := getStorageInstance()
	key := fmt.Sprintf("user_info_list_by_group_id.%d_status.%s", groupId, status.String())

	buf, err := proto.Marshal(userInfo)
	if err != nil {
		return err
	}

	return storage.add(key, buf)
}

func (_ *userStorage) removeUserInfoListByStatus(groupId uint32, status mark2.UserStatus, userInfo *mark2.UserInfo) error {

	storage := getStorageInstance()
	key := fmt.Sprintf("user_info_list_by_group_id.%d_status.%s", groupId, status.String())

	buf, err := proto.Marshal(userInfo)
	if err != nil {
		return err
	}

	return storage.remove(key, buf)
}
