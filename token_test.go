package main

import (
	"testing"
)

func TestTokenValid(t *testing.T) {

	groupId := 10001
	userId := 90001
	uniqueKey := "test_unique_key"
	userName := "test_user_name"

	claims01 := newTokenClaims()
	claims01.groupId = groupId
	claims01.userId = userId
	claims01.uniqueKey = uniqueKey
	claims01.userName = userName

	tokenstring, err := claims01.encode()
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	//ok, err := tokenVerify(tokenstring)
	//if err != nil {
	//	t.Errorf("got %v\n", err)
	//}
	//if !ok {
	//	t.Errorf("got %v\nwant %v", ok, true)
	//}

	claims02, err := tokenDecode(tokenstring)
	if err != nil {
		t.Errorf("got %v\n", err)
	}

	if claims02.groupId != groupId {
		t.Errorf("got %v\nwant %v", claims02.groupId, groupId)
	}
	if claims02.userId != userId {
		t.Errorf("got %v\nwant %v", claims02.userId, userId)
	}
	if claims02.uniqueKey != uniqueKey {
		t.Errorf("got %v\nwant %v", claims02.uniqueKey, uniqueKey)
	}
	if claims02.userName != userName {
		t.Errorf("got %v\nwant %v", claims02.userName, userName)
	}

}
