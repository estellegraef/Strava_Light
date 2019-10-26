/**
 * 2848869
 */

package test

import (
	"Strava_Light/cmd/gps"
	"fmt"
	//"github.com/stretchr/testify/assert"
	"testing"
)

//TODO hard coded files only for testing purposes, may be changed later when correctly implemented
const testGpxFile = "F:\\DHBW\\Semester 5\\Programmieren II\\Go Projects\\src\\Strava_Light\\resources\\2019-09-14_15-14.gpx"
const testZipFile = "F:\\DHBW\\Semester 5\\Programmieren II\\Go Projects\\src\\Strava_Light\\resources\\2019-09-21_15-54.gpx.zip"

//TODO implement ability to read zip files
func TestReadFile(t *testing.T) {
	fileData := gps.ReadFile(testGpxFile)
	fmt.Println(fileData)
}

func TestCalculateRouteInKilometers(t *testing.T) {
	fileData := gps.ReadFile(testGpxFile)
	routeInKilometers := gps.CalculateRouteInKilometers(fileData)
	fmt.Println(routeInKilometers)
}
