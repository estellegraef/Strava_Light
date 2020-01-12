/*
 * 2848869
 * 8089098
 * 3861852
 */

package auth

import (
	"github.com/estellegraef/Strava_Light/backend/hashAndSalt"
	"github.com/estellegraef/Strava_Light/backend/user"
)

type Authenticator interface {
	Authenticate(user, password string) bool
}

type AuthenticatorFunc func(username, password string) bool
type funcWrapper func() *[]user.User

func (af AuthenticatorFunc) Authenticate(user, password string) bool {
	return af(user, password)
}

func CheckUserIsValid(username, password string) bool {
	return checkUserIsValidWrapper(username, password, user.GetUsers)
}

// Decouples the actual UserCred file from the simple check for better testing
func checkUserIsValidWrapper(username, password string, getUsers funcWrapper) bool {
	//var users []user.User
	users := getUsers()
	for _, user := range *users {
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
