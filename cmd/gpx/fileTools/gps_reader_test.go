package fileTools

import (
	"Strava_Light/cmd/gpx"
	"Strava_Light/cmd/gpx/gpx_info"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReadGpx(t *testing.T) {
	//generate GpxFile object  from path
	var actualFile = ReadGpx(gpx.GetTestGpxPath())

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
	files := ReadZip(gpx.GetTestZipPath())
	assert.True(t, len(files) == 1)
}

func TestReadFileWithGpx(t *testing.T) {
	files := ReadFile(gpx.GetTestGpxPath())
	assert.True(t, len(files) == 1)
}

func TestReadFileWithZip(t *testing.T) {
	files := ReadFile(gpx.GetTestZipPath())
	assert.True(t, len(files) == 1)
}

func TestReadFileInvalidPath(t *testing.T) {
	files := ReadFile(gpx.GetTestInvalidPath())
	assert.Equal(t, []gpx_info.GpxFile(nil), files)
}

func TestCheckFileNonExistentPositive(t *testing.T) {
	nonExistent := CheckFileNonExistent(gpx.GetTestInvalidPath())
	assert.True(t, nonExistent)
}

func TestCheckFileNonExistentNegative(t *testing.T) {
	nonExistent := CheckFileNonExistent(gpx.GetTestGpxPath())
	assert.False(t, nonExistent)
}


