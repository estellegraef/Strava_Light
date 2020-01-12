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
	return filepath.Join(GetBasePathStorage(), user)
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

func GetInvalidPath() string {
	return filepath.Join(GetBasePathStorage(), "gpx", "test.zip")
}

func GetUserCredsPath() string {
	return filepath.Join(GetResourcesPath(), "user_credentials", "users.txt")
}

func GetTestUserActivitiesPath() string {
	return filepath.Join(GetBasePathStorage(), "useractivities")
}

func GetCertPath() string {
	return filepath.Join(GetResourcesPath(), "cert.pem")
}

func GetKeyPath() string {
	return filepath.Join(GetResourcesPath(), "key.pem")
}
