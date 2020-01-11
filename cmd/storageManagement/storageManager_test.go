/*
 * 2848869
 * 8089098
 * 3861852
 */
package filemanagement

import (
	"github.com/estellegraef/Strava_Light/resources"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestGetAllFiles(t *testing.T) {
	actualDirs := GetAllFilesFromDir(resources.GetUserActivitiesPath())
	expectedDirs := []string{
		filepath.Join(resources.GetUserActivitiesPath(), "user1"),
		filepath.Join(resources.GetUserActivitiesPath(), "user2"),
	}
	assert.Equal(t, expectedDirs, actualDirs)
}

func TestReadFileContent(t *testing.T) {
	gpxFile := resources.GetShortTestGpx()
	ReadFile(gpxFile)
	//TODO test byte result
}

func TestCreateFile(t *testing.T) {
	dir := resources.GetUserDir("user1")
	fileName := "TestCreateFile.txt"
	content := []byte("Hello")
	isCreated := CreateFile(dir, fileName, content)
	assert.True(t, isCreated)
}

func TestDeleteFile(t *testing.T) {
	dir := resources.GetUserDir("user1")
	fileName := "TestCreateFile.txt"
	content := []byte("TestDelete")
	CreateFile(dir, fileName, content)
	isDeleted := DeleteFile(dir, fileName)
	assert.True(t, isDeleted)
}

func TestUpdateFile(t *testing.T) {
	dir := resources.GetUserDir("user1")
	fileName := "TestCreateFile.txt"
	content := []byte("Hello")
	newContent := []byte("Goodbye")
	isCreated := CreateFile(dir, fileName, content)
	assert.True(t, isCreated)
	isUpdated := UpdateFile(dir, fileName, newContent)
	assert.True(t, isUpdated)
}