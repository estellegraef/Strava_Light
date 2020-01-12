/*
 * 2848869
 * 8089098
 * 3861852
 */

package user

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

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

func TestUser_GetUserName(t *testing.T) {
	username := "username"
	password := []byte("password")
	salt := []byte("salt")
	storageDir := "testdir/test"

	user := User{username, password, salt, storageDir}
	assert.Equal(t, username, user.GetUserName())
}

func TestUser_GetPassword(t *testing.T) {
	username := "username"
	password := []byte("password")
	salt := []byte("salt")
	storageDir := "testdir/test"

	user := User{username, password, salt, storageDir}
	assert.Equal(t, password, user.GetPassword())
}

func TestUser_GetStoragePath(t *testing.T) {
	username := "username"
	password := []byte("password")
	salt := []byte("salt")
	storageDir := "testdir/test"

	user := User{username, password, salt, storageDir}
	assert.Equal(t, storageDir, user.GetStoragePath())
}

func TestUser_GetSalt(t *testing.T) {
	username := "username"
	password := []byte("password")
	salt := []byte("salt")
	storageDir := "testdir/test"

	user := User{username, password, salt, storageDir}
	assert.Equal(t, salt, user.GetSalt())
}

func TestGetUsersFromFile(t *testing.T) {
	username1 := "user1"

	user1password := "DFfsk2pSDCh217moS7rhhEdQLYWoYNrGbt1Ycbz1k2R4FtSUTFYQU6PIJ2vImGWCeMjLuBEDCEvieASj5PeEKw=="
	user1passwordDecode, err1 := base64.StdEncoding.DecodeString(user1password)
	assert.NoError(t, err1)

	user1salt := "4T52AMvYx11VKYTZKuSEu82Gt1m22R8X2zimHwKExmxiTtVqlFi/n2gnvYPuZO/rbZLv9ujl7Qc2XvB3Xcsz2diEZODmRd0V"
	user1saltDecode, err2 := base64.StdEncoding.DecodeString(user1salt)
	assert.NoError(t, err2)

	user1 := User{userName: username1, password: user1passwordDecode, salt: user1saltDecode, storagePath: ""}

	users := GetUsers()
	assert.Equal(t, user1, (*users)[0])
}

func TestCreateStorageForUsersFilePathNotExists(t *testing.T) {
	path, err := os.Getwd()
	assert.NoError(t, err)
	path2 := filepath.Join(path, "storage")
	defer os.RemoveAll(path2)
	CreateStorageForUsers(path)
	users := GetUsers()
	expectedPath := filepath.Join(path, "storage", "user1")
	assert.Equal(t, expectedPath, (*users)[0].GetStoragePath())
}

func TestCreateStorageForUsersFilePathExists(t *testing.T) {
	path, err := os.Getwd()
	assert.NoError(t, err)

	path2 := filepath.Join(path, "test")
	defer os.RemoveAll(path2)

	existingPath := filepath.Join(path, "test", "storage", "user1")

	err2 := os.MkdirAll(existingPath, os.ModePerm)
	assert.NoError(t, err2)

	path = filepath.Join(path, "test")

	CreateStorageForUsers(path)
	users := GetUsers()

	assert.Equal(t, existingPath, (*users)[0].GetStoragePath())
}
