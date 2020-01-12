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
	"strings"
	"testing"
)

var testDir = resources.GetResourcesPath()
var activitiesTest = filepath.Join(testDir, "useractivities")

func TestGetAllFiles(t *testing.T) {
	resources.SetBasePathStorage(testDir)
	actualDirs := GetAllFilesFromDir(resources.GetTestUserActivitiesPath())
	expectedDirs := []string{
		filepath.Join(resources.GetTestUserActivitiesPath(), "user1"),
		filepath.Join(resources.GetTestUserActivitiesPath(), "user2"),
	}
	assert.Equal(t, expectedDirs, actualDirs)
}

func TestGetSingleFileFromDir(t *testing.T) {
	resources.SetBasePathStorage(activitiesTest)
	dir := resources.GetUserDir("user1")
	actualFile := GetSingleFileFromDir(dir, "1", ".json")
	expectedFile := filepath.Join(resources.GetUserDir("user1"), "1.json")
	assert.Equal(t, expectedFile, actualFile)
}

func TestCreateFile(t *testing.T) {
	resources.SetBasePathStorage(activitiesTest)
	dir := resources.GetUserDir("user1")
	fileName := "TestCreateFile.txt"
	content := []byte("Hello")
	isCreated, createdFile := CreateFile(dir, fileName, content)
	assert.True(t, isCreated)
	assert.Equal(t, filepath.Join(dir, fileName), createdFile)
}

func TestDeleteFile(t *testing.T) {
	resources.SetBasePathStorage(activitiesTest)
	dir := resources.GetUserDir("user1")
	fileName := "TestCreateFile.txt"
	content := []byte("TestDelete")
	CreateFile(dir, fileName, content)
	isDeleted := DeleteFile(dir, fileName)
	assert.True(t, isDeleted)
}

func TestUpdateFile(t *testing.T) {
	resources.SetBasePathStorage(activitiesTest)
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
	resources.SetBasePathStorage(activitiesTest)
	dir := resources.GetUserDir("user2")
	actual := GetAllFilesFromDir(dir)
	expected := []string{filepath.Join(dir, "1.json"), filepath.Join(dir, "2.json")}
	assert.Equal(t, expected, actual)
}

func TestGenerateId(t *testing.T) {
	resources.SetBasePathStorage(testDir)
	name := "originalFileName"
	actual := GenerateId(name)
	assert.True(t, strings.Contains(actual, name))
}
