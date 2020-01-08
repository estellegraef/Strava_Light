/*
 * 2848869
 * 8089098
 * 3861852
 */

package cmd

import (
	"encoding/base64"
	"fmt"
)

type Authenticator interface {
	Authenticate(user, password string) bool
}

type AuthenticatorFunc func(username, password string) bool

func (af AuthenticatorFunc) Authenticate(user, password string) bool {
	return af(user, password)
}

func CheckUserIsValid(username, password string) bool {
	//return username == "Rico" && password == "1234"
	//user1: go!Project?2020
	//user2: user2Password
	var users []User
	//if users ==  {
	users = GetUsersFromFile()
	//}
	for _, user := range users {
		if user.GetUserName() == username {
			passwordDecode, err1 := base64.StdEncoding.DecodeString(user.GetPassword())
			if err1 != nil {
				fmt.Println("Base64 Decoding error", err1)
				return false
			}
			saltDecode, err2 := base64.StdEncoding.DecodeString(user.GetSalt())
			if err2 != nil {
				fmt.Println("Base64 Decoding error", err2)
				return false
			}
			if Match([]byte(password), passwordDecode, saltDecode) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
