package activityprocessing

import (
	"Strava_Light/cmd/gpx/fileTools"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testPath = "F:\\DHBW\\Semester 5\\Programmieren II\\Go Projects\\src\\Strava_Light\\resources\\gpx\\2019-09-14_15-14.gpx"

func TestGetAllTrackPoints(t *testing.T) {
	var file = fileTools.ParseXml(testPath)
	var actualTrackPoints = GetAllTrackPoints(file)
	assert.Equal(t, len(actualTrackPoints), 1755)
}

func TestGetMaxSpeed(t *testing.T) {
	var file = fileTools.ParseXml(testPath)
	var actualTrackPoints = GetAllTrackPoints(file)
	var maxSpeed = GetMaxSpeed(actualTrackPoints)
	assert.Equal(t, maxSpeed, 14.65)
}

func TestGetAvgSpeed(t *testing.T) {
	var file = fileTools.ParseXml(testPath)
	var actualTrackPoints = GetAllTrackPoints(file)
	var avgSpeed = GetAvgSpeed(actualTrackPoints)
	assert.Equal(t, avgSpeed, 6.189253561253568)
}

//approximate result oriented on https://opensourceconnections.com/blog/uploads/2009/02/clientsidehaversinecalculation.html
func TestCalculateDistance2Points(t *testing.T) {
	var file = fileTools.ParseXml(testPath)
	var actualTrackPoints = GetAllTrackPoints(file)
	var distance = CalculateDistance2Points(actualTrackPoints[0].GetLatitude(), actualTrackPoints[0].GetLongitude(), actualTrackPoints[1].GetLatitude(), actualTrackPoints[1].GetLongitude())
	//TODO check w/website: here 0.005045 there 0.005301 km
	assert.Equal(t, 0.005045829921162091, distance)
}

//approximate result oriented on https://www.sportdistancecalculator.com/import-gpx.php#map
func TestCalculateDistanceInKilometers(t *testing.T) {
	var file = fileTools.ParseXml(testPath)
	var actualTrackPoints = GetAllTrackPoints(file)
	var distance = CalculateDistanceInKilometers(actualTrackPoints)
	//TODO check w/website: here 24.28 km, there 24.71 km
	assert.Equal(t, 24.28255382563406, distance)
}