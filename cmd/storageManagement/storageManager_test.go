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
	isCreated, createdFile := CreateFile(dir, fileName, content)
	assert.True(t, isCreated)
	assert.Equal(t, filepath.Join(dir, fileName), createdFile)
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
	isCreated, createdFile := CreateFile(dir, fileName, content)
	assert.True(t, isCreated)
	assert.Equal(t, filepath.Join(dir, fileName), createdFile)
	isUpdated := UpdateFile(dir, fileName, newContent)
	assert.True(t, isUpdated)
}

func TestGetAllFilesFromDir(t *testing.T) {
	dir := resources.GetUserDir("user2")
	actual := GetAllFilesFromDir(dir)
	expected := []string {filepath.Join(dir, "2019-09-21_15-54.gpx"), filepath.Join(dir, "3.json")}
	assert.Equal(t, expected, actual)
}

func TestReadReceiveFile(t *testing.T) {
	//TODO do when merged with frontend
}

func TestGenerateId(t *testing.T) {
	//TODO later
}