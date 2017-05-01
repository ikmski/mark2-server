package main

import (
	"testing"
)

func TestTokenValid(t *testing.T) {

	groupId := 1001
	userId := 2001
	uniqueKey := "test_unique_key"
	userName := "test_user_name"

	subject := "test_subject"

	claims01 := newTokenClaims()
	claims01.GroupId = groupId
	claims01.UserId = userId
	claims01.UniqueKey = uniqueKey
	claims01.UserName = userName
	claims01.Subject = subject

	tokenstring, err := claims01.encode()
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	ok, err := tokenVerify(tokenstring)
	if err != nil {
		t.Errorf("got %v\n", err)
	}
	if !ok {
		t.Errorf("got %v\nwant %v", ok, true)
	}

}

func TestTokenDecode(t *testing.T) {

	groupId := 1001
	userId := 2001
	uniqueKey := "test_unique_key"
	userName := "test_user_name"

	subject := "test_subject"

	claims01 := newTokenClaims()
	claims01.GroupId = groupId
	claims01.UserId = userId
	claims01.UniqueKey = uniqueKey
	claims01.UserName = userName
	claims01.Subject = subject

	tokenstring, err := claims01.encode()
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	claims02, err := tokenDecode(tokenstring)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	if claims02.GroupId != groupId {
		t.Errorf("got %v\nwant %v", claims02.GroupId, groupId)
	}
	if claims02.UserId != userId {
		t.Errorf("got %v\nwant %v", claims02.UserId, userId)
	}
	if claims02.UniqueKey != uniqueKey {
		t.Errorf("got %v\nwant %v", claims02.UniqueKey, uniqueKey)
	}
	if claims02.UserName != userName {
		t.Errorf("got %v\nwant %v", claims02.UserName, userName)
	}
	if claims02.Subject != subject {
		t.Errorf("got %v\nwant %v", claims02.Subject, subject)
	}

}
