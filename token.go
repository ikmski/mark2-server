package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("MySigningKey") // TODO

type tokenClaims struct {
	GroupID   uint32 `json:"group_id"`
	UserID    uint32 `json:"user_id"`
	UniqueKey string `json:"unique_key"`
	UserName  string `json:"user_name"`
	jwt.StandardClaims
}

func (tc tokenClaims) Valid() error {
	return nil
}

func newTokenClaims() *tokenClaims {
	tc := new(tokenClaims)
	return tc
}

func (tc *tokenClaims) encode() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tc)
	tokenstring, err := token.SignedString(mySigningKey)

	return tokenstring, err
}

func tokenDecode(str string) (*tokenClaims, error) {

	tc := newTokenClaims()
	token, err := jwt.ParseWithClaims(str, tc, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	tc, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		err := fmt.Errorf("token is invalid")
		return nil, err
	}

	return tc, nil
}

func tokenVerify(str string) (bool, error) {

	token, err := jwt.ParseWithClaims(str, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return false, err
	}

	_, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		err := fmt.Errorf("token is invalid")
		return false, err
	}

	return true, nil
}
