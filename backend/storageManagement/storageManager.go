/*
 * 2848869
 * 8089098
 * 3861852
 */
package filemanagement

import (
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type Closer interface {
	Close() error
}

func ReadFile(filePath string) (content []byte, fileName string) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}

	defer CheckCloser(xmlFile)

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Println(err)
	}
	return byteValue, filepath.Base(filePath)
}

func ReadReceiveFile(file multipart.File) []byte {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
 	return fileBytes
}

func CreateFile(dir string, filename string, content []byte) (isCreated bool, file string) {
	var success = true
	var createdFilepath = filepath.Join(dir, filename)
	err := ioutil.WriteFile(createdFilepath, content, 0755)
	if err != nil {
		success = false
		createdFilepath = ""
		log.Println(err)
	}
	return success, createdFilepath
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

func UpdateFile(dir string, filename string, newContent []byte) bool {
	var success = true
	file, err := os.OpenFile(filepath.Join(dir, filename), os.O_RDWR, 0644)

	if err != nil {
		success = false
		log.Println(err)
	}
	defer CheckCloser(file)

	_, err = file.Write(newContent)
	if err != nil {
		success = false
		log.Println(err)
	}
	return success
}
func GetSingleFileFromDir(directory string, fileName string, extension string) string {
	files := GetAllFilesFromDir(directory)
	var searchedFile = ""
	for _, file := range files {
		if strings.TrimSuffix(filepath.Base(file), extension) == fileName {
			searchedFile =  file
			break
		}
	}
	return searchedFile
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

func GenerateId(fileName string) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 5)
	for i := range b {
		b[i] = letters[rand.Int63() % int64(len(letters))]
	}
	var values []string
	values = append(values, string(b))
	ext := filepath.Ext(fileName)
	trimmed := strings.TrimSuffix(fileName, ext)
	values = append(values, trimmed)
	id := strings.Join(values, "§")
	return id
}

func GetOriginal(id string) string {
	var original = id
	if strings.Contains(id, "§"){
	pos := strings.LastIndex(id, "§")
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len("§")
	if adjustedPos >= len(id) {
		return ""
	}
	original = id[adjustedPos:len(id)]
	}
	return original
}

func CheckCloser(closer Closer) {
	err := closer.Close()
	if err != nil {
		log.Println(err)
	}
}
