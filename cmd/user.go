/**
*2848869
*8089098
*3861852
 */

package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type User struct {
	userName string
	password string
	salt     string
}

func NewUser(name string, password string, salt string) User {
	return User{name, password, salt}
}

func GetUsersFromFile() []User {
	data, err := ioutil.ReadFile("./resources/user_credentials/users.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	dataString := string(data)
	s := strings.Fields(dataString)
	users := make([]User, 0)
	for _, user := range s {
		userSplit := strings.Split(user, ";")
		users = append(users, User{userSplit[0], userSplit[1], userSplit[2]})
	}
	return users
}

func (u User) GetUserName() string {
	return u.userName
}

func (u User) GetPassword() string {
	return u.password
}

func (u User) GetSalt() string {
	return u.salt
}
