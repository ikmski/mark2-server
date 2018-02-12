package main

import "testing"

func TestUint32Set(t *testing.T) {

	set := getUint32SetInstance()
	set.clear()

	var key = "uint32set_test_key"

	var id01 uint32 = 100001
	var id02 uint32 = 100002
	var id03 uint32 = 100003

	// 空リスト
	list, err := set.get(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 0 {
		t.Errorf("got %v\nwant %v", len(list), 0)
	}

	// 追加
	err = set.add(key, id01)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = set.add(key, id02)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	err = set.add(key, id03)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	list, err = set.get(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 3 {
		t.Errorf("got %v\nwant %v", len(list), 3)
	}

	// 削除
	err = set.remove(key, id02)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	// リスト取得
	list, err = set.get(key)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if len(list) != 2 {
		t.Errorf("got %v\nwant %v", len(list), 2)
	}

	if list[0] != id01 {
		t.Errorf("got %v\nwant %v", list[0], id01)
	}

	if list[1] != id03 {
		t.Errorf("got %v\nwant %v", list[1], id03)
	}
}
