/*
 * 2848869
 * 8089098
 * 3861852
 */
package filemanagement

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Closer interface {
	Close() error
}

func ReadFile(filepath string) []byte {
	xmlFile, err := os.Open(filepath)
	CheckError(err)

	defer CheckCloser(xmlFile)

	byteValue, err := ioutil.ReadAll(xmlFile)
	CheckError(err)
	return byteValue
}

func ReadReceiveFile(file multipart.File, header multipart.FileHeader) []byte {
	//TODO check if header is needed
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
 	return fileBytes
}

func DeleteFile(file string) bool {
	var success = true
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return success
}

func SaveFile(file string) bool {
	return true
}

func GetAllFilesFromDir(directory string) []string {
	var dirNames []string
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files{
		dirNames = append(dirNames, filepath.Join(directory, file.Name()))
	}
	return dirNames
}

//TODO implement UUID generation
func GenerateId() uint32 {
	return 1
}

func CheckCloser(closer Closer) {
	err := closer.Close()
	CheckError(err)
}
func CheckError(err error) {
	if err != nil {
		fmt.Errorf("Fehler: %v ", err)
	}
}
