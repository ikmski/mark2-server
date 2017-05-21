package main

import (
	"bytes"
	"testing"
)

func TestStorageMapDataBytes(t *testing.T) {

	storage := getStorageInstance()
	storage.clear()

	var key = "test_key"
	val := []byte{0x01, 0x02, 0x03}

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
	if !bytes.Equal(v, val) {
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

func TestStorageMapDataUint32(t *testing.T) {

	storage := getStorageInstance()
	storage.clear()

	var key = "test_key"
	var val uint32 = 100001

	// Has
	has := storage.has(key)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	// Set
	err := storage.setUint32(key, val)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Has
	has = storage.has(key)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	// Get
	v, err := storage.getUint32(key)
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

func TestStorageMapDataString(t *testing.T) {

	storage := getStorageInstance()
	storage.clear()

	var key = "test_key"
	var val = "test_val"

	// Has
	has := storage.has(key)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	// Set
	err := storage.setString(key, val)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Has
	has = storage.has(key)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	// Get
	v, err := storage.getString(key)
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

func TestStorageSetDataBytes(t *testing.T) {

	storage := getStorageInstance()
	storage.clear()

	var key = "test_key"
	val01 := []byte{0x01, 0x02, 0x03}
	val02 := []byte{0x04, 0x05, 0x06}
	val03 := []byte{0x07, 0x08, 0x09}

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

	if !bytes.Equal(v[0], val01) {
		t.Errorf("got %v\nwant %v", v[0], val01)
	}
	if !bytes.Equal(v[1], val02) {
		t.Errorf("got %v\nwant %v", v[1], val02)
	}
	if !bytes.Equal(v[2], val03) {
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

	if !bytes.Equal(v[0], val01) {
		t.Errorf("got %v\nwant %v", v[0], val01)
	}
	if !bytes.Equal(v[1], val03) {
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

func TestStorageSetDataUint32(t *testing.T) {

	storage := getStorageInstance()
	storage.clear()

	var key = "test_key"
	var val01 uint32 = 100001
	var val02 uint32 = 100002
	var val03 uint32 = 100003

	// Has
	has := storage.has(key)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	// Add
	err := storage.addUint32(key, val01)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = storage.addUint32(key, val02)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = storage.addUint32(key, val03)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Has
	has = storage.has(key)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	// Members
	v, err := storage.membersUint32(key)
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
	err = storage.removeUint32(key, val02)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	v, err = storage.membersUint32(key)
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

func TestStorageSetDataString(t *testing.T) {

	storage := getStorageInstance()
	storage.clear()

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
	err := storage.addString(key, val01)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = storage.addString(key, val02)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = storage.addString(key, val03)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Has
	has = storage.has(key)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	// Members
	v, err := storage.membersString(key)
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
	err = storage.removeString(key, val02)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	v, err = storage.membersString(key)
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
	storage1.clear()
	storage2.clear()

	var key = "test_key"
	var val = "test_val"

	// Set
	err := storage1.setString(key, val)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// Get
	v, err := storage2.getString(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if v != val {
		t.Errorf("got %v\nwant %v", v, val)
	}
}
