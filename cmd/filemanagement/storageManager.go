package filemanagement

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
)

type Closer interface {
	Close() error
}

func ReadFileContent(file string) []byte {
	xmlFile, err := os.Open(file)
	checkError(err)

	defer checkCloser(xmlFile)

	byteValue, err := ioutil.ReadAll(xmlFile)
	checkError(err)
	return byteValue
}

func ReadReceiveFile(file multipart.File, header multipart.FileHeader) []byte {
 return nil
}

func DeleteFile(file string) bool {
	return true
}

func SaveFile(file string) bool {
	return true
}

func GetAllFiles(path string) []string {
	var dirNames []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range files{
		dirNames = append(dirNames, dir.Name())
	}
	return dirNames
}

//TODO implement UUID generation
func GenerateId() uint32 {
	return 1
}

func checkError(err error) {
	if err != nil {
		fmt.Errorf("Fehler: %v ", err)
	}
}

func checkCloser(closer Closer) {
	err := closer.Close()
	checkError(err)
}
