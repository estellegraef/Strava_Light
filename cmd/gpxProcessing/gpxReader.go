/*
 * 2848869
 * 8089098
 * 3861852
 */

package gpxProcessing

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Closer interface {
	Close() error
}

//read any filepath and return contained files converted to GpxFiles in a list
func ReadFile(fileName string) []GpxFile {
	var files []GpxFile
	if CheckFileNonExistent(fileName) == false {
		switch filepath.Ext(fileName) {
		case ".zip":
			files = ReadZip(fileName)
		case ".gpx":
			files = append(files, ReadGpx(fileName))
		default:
			fmt.Errorf("Invalid file extension: " + filepath.Ext(fileName))
		}
	}
	return files
}

//read .zip file
func ReadZip(fileName string) []GpxFile {
	read, err := zip.OpenReader(fileName)
	checkError(err)

	defer checkCloser(read)

	var containedFiles []GpxFile
	//read and convert each file contained in the .zip file
	for _, file := range read.File {
		content := ReadZipContent(file)
		gpx := UnmarshalXML(content)

		//put each files in the .zip file into a list
		containedFiles = append(containedFiles, gpx)
	}
	return containedFiles
}

//read a single file contained in a .zip file and return its byte value
func ReadZipContent(file *zip.File) []byte {
	read, err := file.Open()
	checkError(err)
	defer checkCloser(read)
	content, err := ioutil.ReadAll(read)
	checkError(err)
	return content
}

//read a .gpx file and convert it to a GpxFile object
func ReadGpx(filePath string) GpxFile {
	xmlFile, err := os.Open(filePath)
	checkError(err)

	defer checkCloser(xmlFile)

	byteValue, err := ioutil.ReadAll(xmlFile)
	checkError(err)
	file := UnmarshalXML(byteValue)

	return file
}

//Convert a byte stream to a GpxFile
func UnmarshalXML(byteVal []byte) GpxFile {
	var file GpxFile

	err := xml.Unmarshal(byteVal, &file)
	checkError(err)

	return file
}

//check if file is nonexistent based on OS status
func CheckFileNonExistent(fileName string) bool {
	var nonExistent = false
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		nonExistent = true
	}
	return nonExistent
}

//check if any type of closer throws an error
func checkCloser(closer Closer) {
	err := closer.Close()
	checkError(err)
}

//print error
func checkError(err error) {
	if err != nil {
		fmt.Errorf("Fehler: %v ", err)
	}
}
