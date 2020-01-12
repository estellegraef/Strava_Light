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

var basePathStorage string

func SetBasePathStorage(path string) {
	basePathStorage = path
}

func GetBasePathStorage() string {
	return basePathStorage
}

func GetUserDir(user string) string {
	return filepath.Join(GetUserActivitiesPath(), user)
}

func GetResourcesPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename)
}

func GetTestShortGpx() string {
	return filepath.Join(GetBasePathStorage(), "gpx", "short.gpx")
}

func GetTestGpxPath() string {
	return filepath.Join(GetBasePathStorage(), "gpx", "2019-09-14_15-14.gpx")
}

func GetTestZipPath() string {
	return filepath.Join(GetBasePathStorage(), "gpx", "2019-09-14_15-14.gpx.zip")
}

func GetTestInvalidPath() string {
	return filepath.Join(GetBasePathStorage(), "gpx", "test.zip")
}

func GetTestUserCredsPath() string {
	return filepath.Join(GetBasePathStorage(), "user_credentials", "users.txt")
}

func GetUserActivitiesPath() string {
	return filepath.Join(GetBasePathStorage(), "useractivities")
}

func GetTestCertPath() string {
	return filepath.Join(GetBasePathStorage(), "cert.pem")
}

func GetTestKeyPath() string {
	return filepath.Join(GetBasePathStorage(), "key.pem")
}
