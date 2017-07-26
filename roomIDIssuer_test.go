package main

import "testing"

func TestRoomIDIssuer(t *testing.T) {

	initializeRoomID()
	id := initialRoomID

	roomID := issueRoomID()
	id++
	if roomID != id {
		t.Errorf("got %v\nwant %v", roomID, id)
	}

	roomID = issueRoomID()
	id++
	if roomID != id {
		t.Errorf("got %v\nwant %v", roomID, id)
	}

}
