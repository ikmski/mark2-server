package main

import (
	"testing"
)

func TestTokenValid(t *testing.T) {

	var groupID uint32 = 1001
	var userID uint32 = 2001
	userName := "test_user_name"

	subject := "test_subject"

	claims01 := newTokenClaims()
	claims01.GroupID = groupID
	claims01.UserID = userID
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

	var groupID uint32 = 1001
	var userID uint32 = 2001
	userName := "test_user_name"

	subject := "test_subject"

	claims01 := newTokenClaims()
	claims01.GroupID = groupID
	claims01.UserID = userID
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

	if claims02.GroupID != groupID {
		t.Errorf("got %v\nwant %v", claims02.GroupID, groupID)
	}
	if claims02.UserID != userID {
		t.Errorf("got %v\nwant %v", claims02.UserID, userID)
	}
	if claims02.UserName != userName {
		t.Errorf("got %v\nwant %v", claims02.UserName, userName)
	}
	if claims02.Subject != subject {
		t.Errorf("got %v\nwant %v", claims02.Subject, subject)
	}

}
