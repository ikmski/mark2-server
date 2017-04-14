package main

import (
	"encoding/json"
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey []byte = []byte("MySigningKey") // TODO

type tokenClaims struct {
	groupId   int    `json:"group_id"`
	userId    int    `json:"user_id"`
	uniqueKey string `json:"unique_key"`
	userName  string `json:"user_name"`
}

func (tc tokenClaims) Valid() error {
	return nil
}

func newTokenClaims() *tokenClaims {
	tc := new(tokenClaims)
	return tc
}

func (tc *tokenClaims) encode() (string, error) {

	jsonValue, err := json.Marshal(tc)
	fmt.Printf("%v, %v\n", jsonValue, err)
	clames := tokenClaims{0, 0, "", ""}
	//	clames.groupId = tc.groupId
	//	clames.userId = tc.userId
	//	clames.uniqueKey = tc.uniqueKey
	//	clames.userName = tc.userName

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clames)

	fmt.Printf("%v\n", token.Claims)

	tokenstring, err := token.SignedString(mySigningKey)

	fmt.Printf("%v\n", token.Claims)
	fmt.Printf("%v\n", tokenstring)
	return tokenstring, err
}

func tokenDecode(str string) (*tokenClaims, error) {

	fmt.Printf("%v\n", str)
	tc := newTokenClaims()
	token, err := jwt.ParseWithClaims(str, tc, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Printf("%v\n", tc)
	fmt.Printf("%v\n", token.Claims)
	tc, ok := token.Claims.(*tokenClaims)
	if ok && token.Valid {
		fmt.Printf("%v\n", tc)
		return tc, nil

	} else {
		err := errors.New(fmt.Sprintf("token is invalid"))
		return nil, err
	}
}

func tokenVerify(str string) (bool, error) {

	token, err := jwt.ParseWithClaims(str, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return false, err
	}

	fmt.Printf("%v\n", token.Claims)
	_, ok := token.Claims.(*tokenClaims)
	if ok && token.Valid {
		return true, nil

	} else {
		err := errors.New(fmt.Sprintf("token is invalid"))
		return false, err
	}
}
