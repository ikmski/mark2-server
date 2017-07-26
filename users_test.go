package main

import "testing"

func TestUsers(t *testing.T) {

	users := getUsersInstance()
	users.clear()

	var groupID uint32 = 1001

	newID := issueUserID()
	user1 := newUser()
	user1.info.Id = newID
	user1.info.GroupId = groupID

	has := users.has(user1.info.Id)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	err := users.set(user1.info.Id, user1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	has = users.has(user1.info.Id)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	u, err := users.get(user1.info.Id)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if u != user1 {
		t.Errorf("got %v\nwant %v", u, user1)
	}

	err = users.del(user1.info.Id)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	has = users.has(user1.info.Id)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

}
