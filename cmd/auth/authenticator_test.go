package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckUserIsValidWithWrongPassword(t *testing.T) {
	username := "user1"
	userPassword := "test1"

	actualBool := CheckUserIsValid(username, userPassword)
	assert.Equal(t, false, actualBool)

}

func TestCheckUserIsValidWithWrongUsername(t *testing.T) {
	username := "user12345"
	userPassword := "test1"

	actualBool := CheckUserIsValid(username, userPassword)
	assert.Equal(t, false, actualBool)
}

/*func TestCheckUserIsValidWithRightCredentials(t *testing.T) {
	users := GetUsersFromFile()
	username := users[0].GetUserName()
	salt := users[0].GetSalt()
	password := users[0].GetPassword()

	passwordDecode, err1 := base64.StdEncoding.DecodeString(password)
	assert.NoError(t, err1)
	saltDecode, err2 := base64.StdEncoding.DecodeString(salt)
	assert.NoError(t, err2)

	actualBool := CheckUserIsValid(username, password)
	assert.Equal(t, true, actualBool, "wrong credentials")
}*/
