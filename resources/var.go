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
	//fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	return path.Dir(filename)
}

func GetTestGpxPath() string {
	return filepath.Join(GetBasePath(), "gpx", "2019-09-14_15-14.gpx")
}

func GetTestZipPath() string {
	return filepath.Join(GetBasePath(), "gpx", "2019-09-14_15-14.gpx.zip")
}

func GetTestInvalidPath() string {
	return filepath.Join(GetBasePath(), "gpx", "test.zip")
}

func GetUserCredsPath() string {
	return filepath.Join(GetBasePath(), "user_credentials", "users.txt")
}

func GetCertPath() string {
	return filepath.Join(GetBasePath(), "cert.pem")
}

func GetKeyPath() string {
	return filepath.Join(GetBasePath(), "key.pem")
}
