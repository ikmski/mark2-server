package main

import (
	"sync"
)

var userIDIssuerMutex sync.Mutex
var initialUserID uint32 = 1000000
var maxUserID = initialUserID

func initializeUserID() {
	maxUserID = initialUserID
}

func issueUserID() uint32 {

	userIDIssuerMutex.Lock()
	defer userIDIssuerMutex.Unlock()

	maxUserID++

	return maxUserID
}
