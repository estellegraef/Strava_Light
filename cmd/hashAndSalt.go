/*
 * 2848869
 * 8089098
 * 3861852
 */

package cmd

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

// Aus: https://socketloop.com/tutorials/golang-securing-password-with-salt

const saltSize = 8

func GenerateSalt(secret []byte) []byte {
	buf := make([]byte, saltSize, saltSize+sha512.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		fmt.Printf("random read failed: %v", err)
		os.Exit(1)
	}

	hash := sha512.New()
	hash.Write(buf)
	hash.Write(secret)
	return hash.Sum(buf)
}

func GeneratePasswordAndSaltHash(salt, password []byte) []byte { // generate password + salt hash to store into the file
	combination := string(salt) + string(password)
	passwordHash := sha512.New()
	_, err := io.WriteString(passwordHash, combination)
	if err != nil {
		fmt.Printf("passwordHash could not be generated: %v", err)
	}
	return passwordHash.Sum(nil)
}

func Match(enteredPassword, hashedPassword, salt []byte) bool {
	return bytes.Equal(GeneratePasswordAndSaltHash(salt, enteredPassword), hashedPassword)
}
