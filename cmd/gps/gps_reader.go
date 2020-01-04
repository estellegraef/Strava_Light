/*
 * 2848869
 * 8089098
 * 3861852
 */

package gps

import (
	"fmt"
	"io/ioutil"
)

//TODO make it work for ZIP files
func ReadFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("an error occurred: ", err)
		return ""
	}
	return string(data)
}
