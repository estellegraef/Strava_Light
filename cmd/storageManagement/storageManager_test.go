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
	ReadFileContent(gpxFile)
	//TODO test byte result
}