package auth

import (
	"github.com/estellegraef/Strava_Light/cmd/hashAndSalt"
	"github.com/estellegraef/Strava_Light/cmd/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckUserIsValidWrapperWithWrongPassword(t *testing.T) {
	username := "user1"
	userPassword := "test1"

	passwordRight := "password"
	salt := hashAndSalt.GenerateSalt([]byte(passwordRight))

	userCredRight := user.NewUser("user1", hashAndSalt.GeneratePasswordAndSaltHash(salt, []byte(passwordRight)), salt, "")

	actualBool := checkUserIsValidWrapper(username, userPassword, func() []user.User {
		return []user.User{userCredRight}
	})
	assert.Equal(t, false, actualBool)

}

func TestCheckUserIsValidWithWrongUsername(t *testing.T) {
	username := "user12345"
	userPassword := "test1"

	passwordRight := "password"
	salt := hashAndSalt.GenerateSalt([]byte(passwordRight))

	userCredRight := user.NewUser("user1", hashAndSalt.GeneratePasswordAndSaltHash(salt, []byte(passwordRight)), salt, "")

	actualBool := checkUserIsValidWrapper(username, userPassword, func() []user.User {
		return []user.User{userCredRight}
	})
	assert.Equal(t, false, actualBool)
}

func TestCheckUserIsValidWithRightCredentials(t *testing.T) {
	username := "user1"
	passwordRight := "test1"

	salt := hashAndSalt.GenerateSalt([]byte(passwordRight))

	userCredRight := user.NewUser("user1", hashAndSalt.GeneratePasswordAndSaltHash(salt, []byte(passwordRight)), salt, "")

	actualBool := checkUserIsValidWrapper(username, passwordRight, func() []user.User {
		return []user.User{userCredRight}
	})
	assert.Equal(t, true, actualBool)

}
