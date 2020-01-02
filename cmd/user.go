/**
*2848869
*8089098
 */

package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type User struct {
	UserName string
	Password string
	Salt     string
}

func newUser(name string, password string, salt string) User {
	return User{name, password, salt}
}

func getUserfromFile() []User {
	data, err := ioutil.ReadFile("../resources/user_credentials/users.txt")
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
	return u.UserName
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetSalt() string {
	return u.Salt
}
