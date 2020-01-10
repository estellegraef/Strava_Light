/*
 * 2848869
 * 8089098
 * 3861852
 */
package filemanagement

import (
	"github.com/estellegraef/Strava_Light/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllFiles(t *testing.T) {
	actualDirs := GetAllFiles(resources.GetUserActivitiesPath())
	expectedDirs := []string{"user1", "user2"}
	assert.Equal(t, expectedDirs, actualDirs)
}

func TestReadFileContent(t *testing.T) {
	gpxFile := resources.GetTestGpxPath()
	ReadFileContent(gpxFile)
}