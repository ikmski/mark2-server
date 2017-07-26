package main

import "testing"

func TestUserIDIssuer(t *testing.T) {

	initializeUserID()
	id := initialUserID

	userID := issueUserID()
	id++
	if userID != id {
		t.Errorf("got %v\nwant %v", userID, id)
	}

	userID = issueUserID()
	id++
	if userID != id {
		t.Errorf("got %v\nwant %v", userID, id)
	}

}
