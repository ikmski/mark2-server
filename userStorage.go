package main

import (
	"fmt"
	"strconv"
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

func (_ *userStorage) getUserIdByUniqueKey(uniqueKey string) (int, error) {

	storage := getStorageInstance()
	key := "user_id_by_unique_key." + uniqueKey
	val, err := storage.get(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

func (_ *userStorage) setUserIdByUniqueKey(uniqueKey string, userId int) error {

	storage := getStorageInstance()
	key := "user_id_by_unique_key." + uniqueKey
	return storage.set(key, strconv.Itoa(userId))
}

func (_ *userStorage) removeUserIdByUniqueKey(uniqueKey string) error {

	storage := getStorageInstance()
	key := "user_id_by_unique_key." + uniqueKey
	return storage.del(key)
}

func (us *userStorage) createNewUserId() (int, error) {

	storage := getStorageInstance()
	key := "max_user_id"

	us.mutex.Lock()
	defer us.mutex.Unlock()

	id := 1000000
	has := storage.has(key)
	if has {
		val, err := storage.get(key)
		if err != nil {
			return 0, err
		}

		id, err = strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
	}

	id++

	err := storage.set(key, strconv.Itoa(id))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (_ *userStorage) getUserInfoByUserId(userId int) (*mark2.UserInfo, error) {

	storage := getStorageInstance()
	key := "user_info_by_user_id." + strconv.Itoa(userId)

	val, err := storage.get(key)
	if err != nil {
		return nil, err
	}

	userInfo := new(mark2.UserInfo)
	err = proto.UnmarshalText(val, userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (_ *userStorage) setUserInfoByUserId(userId int, userInfo *mark2.UserInfo) error {

	storage := getStorageInstance()
	key := "user_info_by_user_id." + strconv.Itoa(userId)

	return storage.set(key, proto.MarshalTextString(userInfo))
}

func (_ *userStorage) removeUserInfoByUserId(userId int) error {

	storage := getStorageInstance()
	key := "user_info_by_user_id." + strconv.Itoa(userId)

	return storage.del(key)
}

func (_ *userStorage) getUserInfoListByStatus(groupId int, status mark2.UserStatus) (*mark2.UserInfoList, error) {

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
			info := new(mark2.UserInfo)
			err = proto.UnmarshalText(v, info)
			if err != nil {
				return nil, err
			}

			userInfoList.List = append(userInfoList.List, info)
		}
	}

	return userInfoList, nil
}

func (_ *userStorage) addUserInfoListByStatus(groupId int, status mark2.UserStatus, userInfo *mark2.UserInfo) error {

	storage := getStorageInstance()
	key := fmt.Sprintf("user_info_list_by_group_id.%d_status.%s", groupId, status.String())

	return storage.add(key, proto.MarshalTextString(userInfo))
}

func (_ *userStorage) removeUserInfoListByStatus(groupId int, status mark2.UserStatus, userInfo *mark2.UserInfo) error {

	storage := getStorageInstance()
	key := fmt.Sprintf("user_info_list_by_group_id.%d_status.%s", groupId, status.String())

	return storage.remove(key, proto.MarshalTextString(userInfo))
}
