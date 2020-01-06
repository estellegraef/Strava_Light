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