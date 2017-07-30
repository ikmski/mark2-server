package main

import (
	"sync"
)

var messageIDIssuerMutex sync.Mutex
var initialMessageID uint32 = 1000000
var maxMessageID = initialMessageID

func initializeMessageID() {
	maxMessageID = initialMessageID
}

func issueMessageID() uint32 {

	messageIDIssuerMutex.Lock()
	defer messageIDIssuerMutex.Unlock()

	maxMessageID++

	return maxMessageID
}
