package main

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	mark2 "github.com/ikmski/mark2-server/proto"
)

type userStorage struct {
	s     *storage
	mutex sync.Mutex
}

func newUserStorage() *userStorage {
	us := new(userStorage)
	us.s = getStorageInstance()
	return us
}

func (us *userStorage) clear() {
	us.s.clear()
}

func (us *userStorage) getUserIdByUniqueKey(uniqueKey string) (uint32, error) {

	key := "user_id_by_unique_key." + uniqueKey
	return us.s.getUint32(key)
}

func (us *userStorage) setUserIdByUniqueKey(uniqueKey string, userId uint32) error {

	key := "user_id_by_unique_key." + uniqueKey
	return us.s.setUint32(key, userId)
}

func (us *userStorage) removeUserIdByUniqueKey(uniqueKey string) error {

	key := "user_id_by_unique_key." + uniqueKey
	return us.s.del(key)
}

func (us *userStorage) createNewUserId() (uint32, error) {

	key := "max_user_id"

	us.mutex.Lock()
	defer us.mutex.Unlock()

	var id uint32 = 1000000
	has := us.s.has(key)
	if has {
		var err error
		id, err = us.s.getUint32(key)
		if err != nil {
			return 0, err
		}
	}

	id++

	err := us.s.setUint32(key, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (us *userStorage) getUserInfoByUserId(userId uint32) (*mark2.UserInfo, error) {

	key := fmt.Sprintf("user_info_by_user_id.%d", userId)

	val, err := us.s.get(key)
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

func (us *userStorage) setUserInfoByUserId(userId uint32, userInfo *mark2.UserInfo) error {

	key := fmt.Sprintf("user_info_by_user_id.%d", userId)

	buf, err := proto.Marshal(userInfo)
	if err != nil {
		return err
	}

	return us.s.set(key, buf)
}

func (us *userStorage) removeUserInfoByUserId(userId uint32) error {

	key := fmt.Sprintf("user_info_by_user_id.%d", userId)

	return us.s.del(key)
}

func (us *userStorage) getUserInfoListByStatus(groupId uint32, status mark2.UserStatus) (*mark2.UserInfoList, error) {

	key := fmt.Sprintf("user_info_list_by_group_id.%d_status.%s", groupId, status.String())

	userInfoList := new(mark2.UserInfoList)

	has := us.s.has(key)
	if has {

		list, err := us.s.members(key)
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

func (us *userStorage) addUserInfoListByStatus(groupId uint32, status mark2.UserStatus, userInfo *mark2.UserInfo) error {

	key := fmt.Sprintf("user_info_list_by_group_id.%d_status.%s", groupId, status.String())

	buf, err := proto.Marshal(userInfo)
	if err != nil {
		return err
	}

	return us.s.add(key, buf)
}

func (us *userStorage) removeUserInfoListByStatus(groupId uint32, status mark2.UserStatus, userInfo *mark2.UserInfo) error {

	key := fmt.Sprintf("user_info_list_by_group_id.%d_status.%s", groupId, status.String())

	buf, err := proto.Marshal(userInfo)
	if err != nil {
		return err
	}

	return us.s.remove(key, buf)
}
