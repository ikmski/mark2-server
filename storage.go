package main

import (
	"errors"
	"fmt"
	"sync"
)

type storage struct {
	mutex   sync.Mutex
	mapData map[string]string
	setData map[string]([]string)
}

var storageInstance *storage = newStorage()

func newStorage() *storage {
	s := new(storage)
	s.mapData = make(map[string]string)
	s.setData = make(map[string]([]string))
	return s
}

func getStorageInstance() *storage {
	return storageInstance
}

/* 確認 */
func (s *storage) has(key string) bool {
	return s.hasMap(key) || s.hasSet(key)
}

/* 削除 */
func (s *storage) del(key string) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	var err error

	if s.hasMap(key) {
		err = s.delMap(key)
		if err != nil {
			return err
		}
	}

	if s.hasSet(key) {
		err = s.delSet(key)
		if err != nil {
			return err
		}
	}

	return nil
}

/* 取得 */
func (s *storage) get(key string) (string, error) {

	val, ok := s.mapData[key]
	if !ok {
		err := errors.New(fmt.Sprintf("%s not found\n", key))
		return "", err
	}

	return val, nil
}

/* 設定 */
func (s *storage) set(key string, value string) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.mapData[key] = value

	return nil
}

/* メンバー取得 */
func (s *storage) members(key string) ([]string, error) {

	val, ok := s.setData[key]
	if !ok {
		err := errors.New(fmt.Sprintf("%s not found\n", key))
		return nil, err
	}

	return val, nil
}

/* 要素の追加 */
func (s *storage) add(key string, value string) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	set, ok := s.setData[key]
	if !ok {
		set = make([]string, 0)
	}

	set = append(set, value)
	s.setData[key] = set

	return nil
}

/* 要素の削除 */
func (s *storage) remove(key string, value string) error {

	set, ok := s.setData[key]
	if !ok {
		err := errors.New(fmt.Sprintf("%s not found\n", key))
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	newSet := make([]string, len(set))
	i := 0
	for _, v := range set {
		if value != v {
			newSet[i] = v
			i++
		}
	}

	delete(s.setData, key)
	s.setData[key] = newSet

	return nil
}

func (s *storage) hasMap(key string) bool {
	_, ok := s.mapData[key]
	return ok
}

func (s *storage) hasSet(key string) bool {
	_, ok := s.setData[key]
	return ok
}

func (s *storage) delMap(key string) error {
	has := s.hasMap(key)
	if has {
		delete(s.mapData, key)
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s not found\n", key))
	}
}

func (s *storage) delSet(key string) error {
	has := s.hasSet(key)
	if has {
		delete(s.setData, key)
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s not found\n", key))
	}
}
