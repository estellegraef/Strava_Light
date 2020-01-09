package activityprocessing

import (
	"github.com/estellegraef/Strava_Light/cmd/gpx/fileTools"
	"github.com/estellegraef/Strava_Light/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllTrackPoints(t *testing.T) {
	var file = fileTools.ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	assert.Equal(t, 1755, len(actualTrackPoints))
}

func TestGetMaxSpeed(t *testing.T) {
	var file = fileTools.ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var maxSpeed = GetMaxSpeed(actualTrackPoints)
	assert.Equal(t, 14.65, maxSpeed)
}

func TestGetAvgSpeed(t *testing.T) {
	var file = fileTools.ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var avgSpeed = GetAvgSpeed(actualTrackPoints)
	assert.Equal(t, 6.189253561253568, avgSpeed)
}

//approximate result oriented on https://opensourceconnections.com/blog/uploads/2009/02/clientsidehaversinecalculation.html
func TestCalculateDistance2Points(t *testing.T) {
	var file = fileTools.ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var distance = CalculateDistance2Points(actualTrackPoints[0].GetLatitude(), actualTrackPoints[0].GetLongitude(),
		actualTrackPoints[1].GetLatitude(), actualTrackPoints[1].GetLongitude())
	//TODO check w/website: here 0.005045 there 0.005301 km
	assert.Equal(t, 0.005045829921162091, distance)
}

//approximate result oriented on https://www.sportdistancecalculator.com/import-gpx.php#map
func TestCalculateDistanceInKilometers(t *testing.T) {
	var file = fileTools.ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var distance = CalculateDistanceInKilometers(actualTrackPoints)
	//TODO check w/website: here 24.28 km, there 24.71 km
	assert.Equal(t, 24.28255382563406, distance)
}

func TestCorrectSpeed(t *testing.T) {
	var speed = 11.0
	var avgSpeed = 3.4
	var correctedSpeed = CorrectSpeed(speed, avgSpeed)
	assert.Equal(t, avgSpeed, correctedSpeed)
}

//TODO verify correctness
func TestCalculateStandbyTimeInSec(t *testing.T) {
	var file = fileTools.ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var totalStandbyTime = CalculateStandbyTimeInSec(actualTrackPoints)
	assert.Equal(t, 20.31, totalStandbyTime)
}
