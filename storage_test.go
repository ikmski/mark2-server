package main

import (
	"testing"
)

func TestStorageMapData(t *testing.T) {

	storage := getStorageInstance()

	var key = "test_key"
	var val = "test_val"

	// Has
	has := storage.has(key)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	// Set
	err := storage.set(key, val)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Has
	has = storage.has(key)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	// Get
	v, err := storage.get(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if v != val {
		t.Errorf("got %v\nwant %v", v, val)
	}

	// Del
	err = storage.del(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Has
	has = storage.has(key)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}
}

func TestStorageSetData(t *testing.T) {

	storage := getStorageInstance()

	var key = "test_key"
	var val01 = "test_val_01"
	var val02 = "test_val_02"
	var val03 = "test_val_03"

	// Has
	has := storage.has(key)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	// Add
	err := storage.add(key, val01)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = storage.add(key, val02)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = storage.add(key, val03)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Has
	has = storage.has(key)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	// Members
	v, err := storage.members(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	if v[0] != val01 {
		t.Errorf("got %v\nwant %v", v[0], val01)
	}
	if v[1] != val02 {
		t.Errorf("got %v\nwant %v", v[1], val02)
	}
	if v[2] != val03 {
		t.Errorf("got %v\nwant %v", v[2], val03)
	}

	// Remove
	err = storage.remove(key, val02)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	v, err = storage.members(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(v) != 2 {
		t.Errorf("got %v\nwant %v", len(v), 2)
	}

	if v[0] != val01 {
		t.Errorf("got %v\nwant %v", v[0], val01)
	}
	if v[1] != val03 {
		t.Errorf("got %v\nwant %v", v[1], val03)
	}

	// Del
	err = storage.del(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Has
	has = storage.has(key)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}
}

func TestStorageSingleton(t *testing.T) {

	storage1 := getStorageInstance()
	storage2 := getStorageInstance()

	var key = "test_key"
	var val = "test_val"

	// Set
	err := storage1.set(key, val)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Get
	v, err := storage2.get(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if v != val {
		t.Errorf("got %v\nwant %v", v, val)
	}
}
