/*
 * 2848869
 * 8089098
 * 3861852
 */

package gpxProcessing

import (
	"archive/zip"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
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
			log.Println("Invalid file extension: " + filepath.Ext(fileName))
		}
	}
	return files
}

//read .zip file
func ReadZip(fileName string) []GpxFile {
	read, err := zip.OpenReader(fileName)
	if err != nil{
		log.Println(err)
	}

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
	if err != nil{
		log.Println(err)
	}

	defer checkCloser(read)
	content, err := ioutil.ReadAll(read)
	if err != nil{
		log.Println(err)
	}

	return content
}

//read a .gpx file and convert it to a GpxFile object
func ReadGpx(filePath string) GpxFile {
	xmlFile, err := os.Open(filePath)
	if err != nil{
		log.Println(err)
	}

	defer checkCloser(xmlFile)

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil{
		log.Println(err)
	}
	// for bytearray only unmarshalgpx
	file := UnmarshalXML(byteValue)

	return file
}

//Convert a byte stream to a GpxFile
func UnmarshalXML(byteVal []byte) GpxFile {
	var file GpxFile

	err := xml.Unmarshal(byteVal, &file)
	if err != nil{
		log.Println(err)
	}

	fileTime := file.GetMeta().GetTime()
	if fileTime.IsZero() {
		file.Meta.Time = time.Now()
	}
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
	if err != nil{
		log.Println(err)
	}
}
