/*
 * 2848869
 * 8089098
 * 3861852
 */

package auth

import (
	"github.com/estellegraef/Strava_Light/backend/hashAndSalt"
	"github.com/estellegraef/Strava_Light/backend/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckUserIsValidWrapperWithWrongPassword(t *testing.T) {
	username := "user1"
	userPassword := "test1"

	passwordRight := "password"
	salt := hashAndSalt.GenerateSalt([]byte(passwordRight))

	userCredRight := user.NewUser("user1", hashAndSalt.GeneratePasswordAndSaltHash(salt, []byte(passwordRight)), salt, "")

	actualBool := checkUserIsValidWrapper(username, userPassword, func() *[]user.User {
		return &[]user.User{userCredRight}
	})
	assert.Equal(t, false, actualBool)

}

func TestCheckUserIsValidWWrapperWithWrongUsername(t *testing.T) {
	username := "user12345"
	userPassword := "test1"

	passwordRight := "password"
	salt := hashAndSalt.GenerateSalt([]byte(passwordRight))

	userCredRight := user.NewUser("user1", hashAndSalt.GeneratePasswordAndSaltHash(salt, []byte(passwordRight)), salt, "")

	actualBool := checkUserIsValidWrapper(username, userPassword, func() *[]user.User {
		return &[]user.User{userCredRight}
	})
	assert.Equal(t, false, actualBool)
}

func TestCheckUserIsValidWrapperWithRightCredentials(t *testing.T) {
	username := "user1"
	passwordRight := "test1"

	salt := hashAndSalt.GenerateSalt([]byte(passwordRight))

	userCredRight := user.NewUser("user1", hashAndSalt.GeneratePasswordAndSaltHash(salt, []byte(passwordRight)), salt, "")

	actualBool := checkUserIsValidWrapper(username, passwordRight, func() *[]user.User {
		return &[]user.User{userCredRight}
	})
	assert.Equal(t, true, actualBool)
}

func TestCheckUserIsValidWithRightCred(t *testing.T) {
	username := "user2"
	passwordRight := "user2Password"

	assert.True(t, CheckUserIsValid(username, passwordRight))
}

func TestCheckUserIsValidWithWrongCred(t *testing.T) {
	username := "user2"
	passwordRight := "test"

	assert.False(t, CheckUserIsValid(username, passwordRight))
}

func TestAuthenticatorFunc_Authenticate(t *testing.T) {
	username := "user2"
	passwordRight := "user2Password"

	assert.True(t, AuthenticatorFunc(CheckUserIsValid).Authenticate(username, passwordRight))
}
