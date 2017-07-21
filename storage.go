package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
)

type storage struct {
	mutex   sync.Mutex
	mapData map[string][]byte
	setData map[string]([][]byte)
}

var storageInstance = newStorage()

func newStorage() *storage {
	s := new(storage)
	s.mapData = make(map[string][]byte)
	s.setData = make(map[string]([][]byte))
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

/* クリア */
func (s *storage) clear() {
	s.mapData = make(map[string][]byte)
	s.setData = make(map[string]([][]byte))
}

/* 取得 */
func (s *storage) get(key string) ([]byte, error) {

	val, ok := s.mapData[key]
	if !ok {
		err := fmt.Errorf("%s not found", key)
		return nil, err
	}

	return val, nil
}

func (s *storage) getUint32(key string) (uint32, error) {

	val, err := s.get(key)
	if err != nil {
		return 0, err
	}

	reader := bytes.NewReader(val)
	var result uint32
	err = binary.Read(reader, binary.LittleEndian, &result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (s *storage) getString(key string) (string, error) {

	val, err := s.get(key)
	if err != nil {
		return "", err
	}

	return string(val), nil
}

/* 設定 */
func (s *storage) set(key string, value []byte) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.mapData[key] = value

	return nil
}

func (s *storage) setUint32(key string, value uint32) error {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	return s.set(key, buf.Bytes())
}

func (s *storage) setString(key string, value string) error {

	return s.set(key, []byte(value))
}

/* メンバー取得 */
func (s *storage) members(key string) ([][]byte, error) {

	val, ok := s.setData[key]
	if !ok {
		err := fmt.Errorf("%s not found", key)
		return nil, err
	}

	return val, nil
}

func (s *storage) membersUint32(key string) ([]uint32, error) {

	list, err := s.members(key)
	if err != nil {
		return nil, err
	}

	result := make([]uint32, len(list))
	for i, v := range list {

		reader := bytes.NewReader(v)
		var item uint32
		err = binary.Read(reader, binary.LittleEndian, &item)
		if err != nil {
			return nil, err
		}

		result[i] = item
	}

	return result, nil
}

func (s *storage) membersString(key string) ([]string, error) {

	list, err := s.members(key)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(list))
	for i, v := range list {
		result[i] = string(v)
	}

	return result, nil
}

/* 要素の追加 */
func (s *storage) add(key string, value []byte) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	set, ok := s.setData[key]
	if !ok {
		set = make([][]byte, 0)
	}

	set = append(set, value)
	s.setData[key] = set

	return nil
}

func (s *storage) addUint32(key string, value uint32) error {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	return s.add(key, buf.Bytes())
}

func (s *storage) addString(key string, value string) error {

	return s.add(key, []byte(value))
}

/* 要素の削除 */
func (s *storage) remove(key string, value []byte) error {

	set, ok := s.setData[key]
	if !ok {
		err := fmt.Errorf("%s not found", key)
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	newSet := make([][]byte, 0, len(set))
	for _, v := range set {
		if !bytes.Equal(value, v) {
			newSet = append(newSet, v)
		}
	}

	delete(s.setData, key)
	s.setData[key] = newSet

	return nil
}

func (s *storage) removeUint32(key string, value uint32) error {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, value)
	if err != nil {
		return err
	}

	return s.remove(key, buf.Bytes())
}

func (s *storage) removeString(key string, value string) error {

	return s.remove(key, []byte(value))
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
	}

	return fmt.Errorf("%s not found", key)
}

func (s *storage) delSet(key string) error {
	has := s.hasSet(key)
	if has {
		delete(s.setData, key)
		return nil
	}

	return fmt.Errorf("%s not found", key)
}
