package main

import "sync"

var roomIDIssuerMutex sync.Mutex
var initialRoomID uint32 = 1000000
var maxRoomID = initialRoomID

func initializeRoomID() {
	maxRoomID = initialRoomID
}

func issueRoomID() uint32 {

	roomIDIssuerMutex.Lock()
	defer roomIDIssuerMutex.Unlock()

	maxRoomID++

	return maxRoomID
}
