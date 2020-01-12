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

//read any file based on its path and return the content and its filename
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

//read a multipart file
func ReadReceiveFile(file multipart.File) []byte {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
 	return fileBytes
}

//create a new file in a given directory, with a given name and content
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

//delete a file from a certain directory
func DeleteFile(dir string, filename string) bool {
	var success = true
	err := os.Remove(filepath.Join(dir, filename))
	if err != nil {
		success = false
		log.Println(err)
	}
	return success
}

//overwrite a file with a new content
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

//receive a single file based on a specific name and file extension
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

//get all files from a goven directory
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

//generate a unique id containing the original filename
func GenerateId(fileName string) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//generate a 5 digit random substring based on the letters
	b := make([]byte, 5)
	for i := range b {
		b[i] = letters[rand.Int63() % int64(len(letters))]
	}
	var values []string
	values = append(values, string(b))
	ext := filepath.Ext(fileName)
	trimmed := strings.TrimSuffix(fileName, ext)
	//append the filename without extension
	values = append(values, trimmed)
	//merge substring and filename with delimiter
	id := strings.Join(values, "§")
	return id
}

//receive the original filename from a generated unique id
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

//function to close any type of closer with error check
func CheckCloser(closer Closer) {
	err := closer.Close()
	if err != nil {
		log.Println(err)
	}
}
