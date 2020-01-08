package gpx

import (
	"os"
	"path/filepath"
)

func GetBasePath()  string {
	return filepath.Join(os.Getenv("GOPATH"), "\\src\\Strava_Light")
}

func GetTestGpxPath() string {
	return filepath.Join(GetBasePath(), "resources\\gpx\\2019-09-14_15-14.gpx")
}

func GetTestZipPath() string {
	return filepath.Join(GetBasePath(), "resources\\gpx\\2019-09-14_15-14.gpx.zip")
}

func GetTestInvalidPath() string {
	return filepath.Join(GetBasePath(), "resources\\gpx\\test.zip")
}