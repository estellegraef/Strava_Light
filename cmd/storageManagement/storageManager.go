/*
 * 2848869
 * 8089098
 * 3861852
 */
package filemanagement

import (
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
	if err != nil {
		log.Println(err)
	}

	defer CheckCloser(xmlFile)

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Println(err)
	}
	return byteValue
}

func ReadReceiveFile(file multipart.File) []byte {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
 	return fileBytes
}

func CreateFile(dir string, filename string, content []byte) bool {
	var success = true
	err := ioutil.WriteFile(filepath.Join(dir, filename), content, 0755)
	if err != nil {
		success = false
		log.Println(err)
	}
	return success
}

func DeleteFile(dir string, filename string) bool {
	var success = true
	err := os.Remove(filepath.Join(dir, filename))
	if err != nil {
		success = false
		log.Println(err)
	}
	return success
}

func SaveFile(dir string, filename string, newcontent[]byte) bool {
	var success = true
	file, err := os.OpenFile(filepath.Join(dir, filename), os.O_RDWR, 0644)

	if err != nil {
		success = false
		log.Println(err)
	}
	defer CheckCloser(file)

	_, err = file.Write(newcontent) // Write at 0 beginning
	if err != nil {
		success = false
		log.Println(err)
	}
	return success
}

func GetAllFilesFromDir(directory string) []string {
	var dirNames []string
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println(err)
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
	if err != nil {
		log.Println(err)
	}
}
