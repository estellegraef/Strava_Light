package fileTools

import (
	"Strava_Light/cmd/gpx/gpx_info"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO convert to relative path
const testZip= "F:\\DHBW\\Semester 5\\Programmieren II\\Go Projects\\src\\Strava_Light\\resources\\gpx\\2019-09-14_15-14.gpx.zip"
const testPath = "F:\\DHBW\\Semester 5\\Programmieren II\\Go Projects\\src\\Strava_Light\\resources\\gpx\\2019-09-14_15-14.gpx"
const invalidPath = "F:\\DHBW\\Semester 5\\Programmieren II\\Go Projects\\src\\Strava_Light\\resources\\gpx\\test.zip"

func TestReadGpx(t *testing.T) {
	//generate GpxFile object  from path
	var actualFile = ReadGpx(testPath)

	//generate expected object
	var meta = gpx_info.NewMeta(time.Date(2019, 9, 14, 13, 14, 17,94000000, time.UTC))
	var ext = gpx_info.NewExtension(gpx_info.NewTrackPointExtension(5.54))
	var point = gpx_info.NewTrackPoint(49.3549890600, 9.1519649400, time.Date(2019, 9, 14, 13, 14, 30, 276000000, time.UTC), ext)
	var points []gpx_info.TrackPoint
	points = append(points, point)
	var segment = gpx_info.NewTrackSegment(points)
	var segments []gpx_info.TrackSegment
	segments = append(segments, segment)
	var track = gpx_info.NewTrack(segments)
	var tracks []gpx_info.Track
	tracks = append(tracks, track)
	var expectedFile = gpx_info.NewGpx("Urban Biker", meta, tracks)

	//extract same TrackPoint
	var actualTrackSegment = actualFile.GetTracks()[0].GetTrackSegments()[1].GetTrackPoints()[4]
	var expectedTrackSegment = expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0]

	//checkError similarity
	assert.Equal(t, actualFile.GetCreator(), expectedFile.GetCreator())
	assert.Equal(t, actualFile.GetMeta().GetTime(), expectedFile.GetMeta().GetTime())
	assert.Equal(t, actualTrackSegment.Longitude, expectedTrackSegment.Longitude)
	assert.Equal(t, actualTrackSegment.Latitude, expectedTrackSegment.Latitude)
	assert.Equal(t, actualTrackSegment.DateTime, expectedTrackSegment.DateTime)
	assert.Equal(t, actualTrackSegment.GetExtension().GetTrackPointExtension().GetSpeed(), expectedTrackSegment.GetExtension().GetTrackPointExtension().GetSpeed())
}

func TestReadZip(t *testing.T) {
	files := ReadZip(testZip)
	assert.True(t, len(files) == 1)
}

func TestReadFileWithGpx(t *testing.T) {
	files := ReadFile(testPath)
	assert.True(t, len(files) == 1)
}

func TestReadFileWithZip(t *testing.T) {
	files := ReadFile(testZip)
	assert.True(t, len(files) == 1)
}

func TestReadFileInvalidPath(t *testing.T) {
	files := ReadFile(invalidPath)
	assert.Equal(t, []gpx_info.GpxFile(nil), files)
}

func TestCheckFileNonExistentPositive(t *testing.T) {
	nonExistent := CheckFileNonExistent(invalidPath)
	assert.True(t, nonExistent)
}

func TestCheckFileNonExistentNegative(t *testing.T) {
	nonExistent := CheckFileNonExistent(testPath)
	assert.False(t, nonExistent)
}


