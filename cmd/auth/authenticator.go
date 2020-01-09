/*
 * 2848869
 * 8089098
 * 3861852
 */

package auth

import (
	"github.com/estellegraef/Strava_Light/cmd/hashAndSalt"
	"github.com/estellegraef/Strava_Light/cmd/user"
)

type Authenticator interface {
	Authenticate(user, password string) bool
}

type AuthenticatorFunc func(username, password string) bool
type funcWrapper func() []user.User

func (af AuthenticatorFunc) Authenticate(user, password string) bool {
	return af(user, password)
}

func CheckUserIsValid(username, password string) bool {
	//return username == "Rico" && password == "1234"
	//user1: go!Project?2020
	//user2: user2Password
	return checkUserIsValidWrapper(username, password, user.GetUsersFromFile)
}

// Decouples the actual UserCred file from the simple check for better testing
func checkUserIsValidWrapper(username, password string, getUsers funcWrapper) bool {
	var users []user.User
	users = getUsers()
	for _, user := range users {
		if user.GetUserName() == username {
			if hashAndSalt.Match([]byte(password), user.GetPassword(), user.GetSalt()) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
