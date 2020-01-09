/*
 * 2848869
 * 8089098
 * 3861852
 */

package user

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)

type User struct {
	userName    string
	password    []byte
	salt        []byte
	storagePath string
}

func NewUser(name string, password []byte, salt []byte, storagePath string) User {
	return User{name, password, salt, storagePath}
}

func GetUsersFromFile() []User {
	data, err := ioutil.ReadFile("./resources/user_credentials/users.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	dataString := string(data)
	s := strings.Fields(dataString)
	var users []User
	for _, user := range s {
		userSplit := strings.Split(user, ";")

		passwordDecode, err1 := base64.StdEncoding.DecodeString(userSplit[1])
		if err1 != nil {
			fmt.Println("Base64 Decoding error", err1)
		}
		saltDecode, err2 := base64.StdEncoding.DecodeString(userSplit[2])
		if err2 != nil {
			fmt.Println("Base64 Decoding error", err2)
		}

		users = append(users, User{userSplit[0], passwordDecode, saltDecode, ""})
	}
	return users
}

func (u User) GetUserName() string {
	return u.userName
}

func (u User) GetPassword() []byte {
	return u.password
}

func (u User) GetSalt() []byte {
	return u.salt
}

func (u User) GetStoragePath() string {
	return u.storagePath
}
