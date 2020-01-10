/*
 * 2848869
 * 8089098
 * 3861852
 */

package gpxProcessing

import (
	"github.com/estellegraef/Strava_Light/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllTrackPoints(t *testing.T) {
	var file = ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	assert.Equal(t, 1755, len(actualTrackPoints))
}

func TestGetMaxSpeed(t *testing.T) {
	var file = ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var maxSpeed = GetMaxSpeed(actualTrackPoints)
	assert.Equal(t, 14.65, maxSpeed)
}

func TestGetAvgSpeed(t *testing.T) {
	var file = ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var avgSpeed = GetAvgSpeed(actualTrackPoints)
	assert.Equal(t, 6.189253561253568, avgSpeed)
}

//approximate result oriented on https://opensourceconnections.com/blog/uploads/2009/02/clientsidehaversinecalculation.html
func TestCalculateDistance2Points(t *testing.T) {
	var file = ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var distance = CalculateDistance2Points(actualTrackPoints[0].GetLatitude(), actualTrackPoints[0].GetLongitude(),
		actualTrackPoints[1].GetLatitude(), actualTrackPoints[1].GetLongitude())
	assert.Equal(t, 0.005045829921162091, distance)
}

//approximate result oriented on https://www.sportdistancecalculator.com/import-gpx.php#map
func TestCalculateDistanceInKilometers(t *testing.T) {
	var file = ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var distance = CalculateDistanceInKilometers(actualTrackPoints)
	assert.Equal(t, 24.28255382563406, distance)
}

func TestCalculateStandbyTimeInSec(t *testing.T) {
	var file = ReadGpx(resources.GetTestGpxPath())
	var actualTrackPoints = GetAllTrackPoints(file)
	var totalStandbyTime = CalculateStandbyTimeInSec(actualTrackPoints)
	assert.Equal(t, 196.288, totalStandbyTime)
}

func TestVerifySportType(t *testing.T) {
	var verifiedSportType = VerifySportType("Laufen", 18.9)
	assert.Equal(t, "Radfahren", verifiedSportType)
}