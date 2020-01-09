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

func GetTestGpxPath() string {
	return filepath.Join(GetBasePath(), "gpx\\2019-09-14_15-14.gpx")
}

func GetTestZipPath() string {
	return filepath.Join(GetBasePath(), "gpx\\2019-09-14_15-14.gpx.zip")
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