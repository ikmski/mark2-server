package main

import (
	"testing"

	mark2 "github.com/ikmski/mark2-server/proto"
)

func TestWaitStreamsSet(t *testing.T) {

	ws := getWaitStreamsInstance()
	ws.clear()

	var id1 uint32 = 10001
	var id2 uint32 = 10002

	stream1 := new(mark2.MessageService_WaitMessageServer)
	stream2 := new(mark2.MessageService_WaitMessageServer)

	has := ws.has(id1)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}
	has = ws.has(id2)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	err := ws.set(id1, stream1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = ws.set(id2, stream2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	has = ws.has(id1)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}
	has = ws.has(id2)
	if !has {
		t.Errorf("got %v\nwant %v", has, true)
	}

	ws.clear()
}

func TestWaitStreamsGet(t *testing.T) {

	ws := getWaitStreamsInstance()
	ws.clear()

	var id1 uint32 = 10001
	var id2 uint32 = 10002

	stream1 := new(mark2.MessageService_WaitMessageServer)
	stream2 := new(mark2.MessageService_WaitMessageServer)

	err := ws.set(id1, stream1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = ws.set(id2, stream2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	s1, err := ws.get(id1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if s1 != stream1 {
		t.Errorf("got %v\nwant %v", s1, stream1)
	}

	s2, err := ws.get(id2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if s2 != stream2 {
		t.Errorf("got %v\nwant %v", s2, stream2)
	}

	ws.clear()
}

func TestWaitStreamsDel(t *testing.T) {

	ws := getWaitStreamsInstance()
	ws.clear()

	var id1 uint32 = 10001
	var id2 uint32 = 10002

	stream1 := new(mark2.MessageService_WaitMessageServer)
	stream2 := new(mark2.MessageService_WaitMessageServer)

	err := ws.set(id1, stream1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = ws.set(id2, stream2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	err = ws.del(id1)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	has := ws.has(id1)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	err = ws.del(id2)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	has = ws.has(id2)
	if has {
		t.Errorf("got %v\nwant %v", has, false)
	}

	ws.clear()
}
