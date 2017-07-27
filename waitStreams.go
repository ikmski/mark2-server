package main

import (
	"fmt"
	"sync"

	mark2 "github.com/ikmski/mark2-server/proto"
)

type waitStreams struct {
	mutex     sync.Mutex
	streamMap map[uint32]mark2.MessageService_WaitMessageServer
}

var waitStreamsInstance = newWaitStreams()

func newWaitStreams() *waitStreams {
	ws := new(waitStreams)
	ws.streamMap = make(map[uint32]mark2.MessageService_WaitMessageServer)
	return ws
}

func getWaitStreamsInstance() *waitStreams {
	return waitStreamsInstance
}

func (ws *waitStreams) has(key uint32) bool {
	_, ok := ws.streamMap[key]
	return ok
}

func (ws *waitStreams) get(key uint32) (mark2.MessageService_WaitMessageServer, error) {

	val, ok := ws.streamMap[key]
	if !ok {
		err := fmt.Errorf("%d not found", key)
		return nil, err
	}

	return val, nil
}

func (ws *waitStreams) set(key uint32, value mark2.MessageService_WaitMessageServer) error {

	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	ws.streamMap[key] = value

	return nil
}

func (ws *waitStreams) del(key uint32) error {

	ws.mutex.Lock()
	defer ws.mutex.Unlock()

	if !ws.has(key) {
		return fmt.Errorf("%d not found", key)
	}
	delete(ws.streamMap, key)

	return nil
}

func (ws *waitStreams) clear() {
	ws.streamMap = make(map[uint32]mark2.MessageService_WaitMessageServer)
}
