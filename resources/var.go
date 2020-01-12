/*
 * 2848869
 * 8089098
 * 3861852
 */

package resources

import (
	"path"
	"path/filepath"
	"runtime"
)

func GetBasePath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
			panic("No caller information")
		}
	return path.Dir(filename)
}

func GetShortTestGpx() string {
	return filepath.Join(GetBasePath(), "gpx\\1.gpx")
}

func GetTestGpxPath() string {
	return filepath.Join(GetBasePath(), "gpx\\1.gpx")
}

func GetTestZipPath() string {
	return filepath.Join(GetBasePath(), "gpx\\1.gpx.zip")
}

func GetTestInvalidPath() string {
	return filepath.Join(GetBasePath(), "gpx\\test.zip")
}

func GetUserActivitiesPath() string {
	return filepath.Join(GetBasePath(), "useractivities")
}

func GetUserDir(user string) string {
	return filepath.Join(GetUserActivitiesPath(), user)
}