/*
 * 2848869
 * 8089098
 * 3861852
 */

package user

import (
	"encoding/base64"
	"fmt"
	"github.com/estellegraef/Strava_Light/resources"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type User struct {
	userName    string
	password    []byte
	salt        []byte
	storagePath string
}

var (
	users []User
)

func NewUser(name string, password []byte, salt []byte, storagePath string) User {
	return User{name, password, salt, storagePath}
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

func (u *User) ChangeStoragePath(path string) {
	u.storagePath = path
}

// Reads the user credentials from the file and saves them in the variable "users"
func getUsersFromFile() {
	//data, err := ioutil.ReadFile("resources/user_credentials/users.txt")
	data, err := ioutil.ReadFile(resources.GetUserCredsPath())
	if err != nil {
		fmt.Println("File reading error: ", err)
	}
	dataString := string(data)
	s := strings.Fields(dataString)
	users = nil
	for _, user := range s {
		userSplit := strings.Split(user, ";")

		passwordDecode, err1 := base64.StdEncoding.DecodeString(userSplit[1])
		if err1 != nil {
			fmt.Println("Base64 Decoding error: ", err1)
		}
		saltDecode, err2 := base64.StdEncoding.DecodeString(userSplit[2])
		if err2 != nil {
			fmt.Println("Base64 Decoding error: ", err2)
		}

		users = append(users, User{userSplit[0], passwordDecode, saltDecode, ""})
	}
}

// returns the user information in the form of a variable, so that the user file does not have to be read out every time
func GetUsers() *[]User {
	if len(users) == 0 {
		getUsersFromFile()
	}
	return &users
}

// Create the user storage location and return the actual basePath
func CreateStorageForUsers(basePath string) {
	userSlice := GetUsers()
	actualBasePath := filepath.Join(basePath, "storage")
	resources.SetBasePathStorage(actualBasePath)
	absPath, err := filepath.Abs(actualBasePath)
	if err != nil {
		fmt.Println("Can't get absolute Path: ", err)
	}
	for i := 0; i < len(*userSlice); i++ {
		path := filepath.Join(absPath, (*userSlice)[i].userName)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating user directory: ", err)
			} else {
				(*userSlice)[i].ChangeStoragePath(path)
			}
		} else {
			(*userSlice)[i].ChangeStoragePath(path)
		}
	}
}
