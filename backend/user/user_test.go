/*
 * 2848869
 * 8089098
 * 3861852
 */

package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//TODO write test for the function getUsersFromFile and GetUsers
func TestGetUsers(t *testing.T) {
	userArray := GetUsers()
	assert.NotEqual(t, 0, len(*userArray))
}

func TestUser_ChangeStoragePath(t *testing.T) {
	path := "test"
	user := User{}

	user.ChangeStoragePath(path)
	assert.Equal(t, path, user.GetStoragePath(), "wrong storagePath")
}
