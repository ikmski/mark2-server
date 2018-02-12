package main

import (
	"fmt"
	"sync"
)

type uint32Set struct {
	mutex sync.Mutex
	set   map[string]([]uint32)
}

var uint32SetInstance = newUint32Set()

func newUint32Set() *uint32Set {
	s := new(uint32Set)
	s.set = make(map[string]([]uint32))
	return s
}

func getUint32SetInstance() *uint32Set {
	return uint32SetInstance
}

func (s *uint32Set) clear() {
	s.set = make(map[string]([]uint32))
}

func (s *uint32Set) get(key string) ([]uint32, error) {

	list, ok := s.set[key]
	if !ok {
		list = make([]uint32, 0)
	}

	return list, nil
}

func (s *uint32Set) add(key string, val uint32) error {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	list, ok := s.set[key]
	if !ok {
		list = make([]uint32, 0)
	}

	list = append(list, val)
	s.set[key] = list

	return nil
}

func (s *uint32Set) remove(key string, val uint32) error {

	list, ok := s.set[key]
	if !ok {
		err := fmt.Errorf("%s not found", key)
		return err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	newList := make([]uint32, 0, len(list))
	for _, v := range list {
		if val != v {
			newList = append(newList, v)
		}
	}

	delete(s.set, key)
	s.set[key] = newList

	return nil
}
