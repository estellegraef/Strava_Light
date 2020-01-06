package fileTools

import (
	"Strava_Light/cmd/gpx/gpx_info"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO convert to relative path
const testPath = "F:\\DHBW\\Semester 5\\Programmieren II\\Go Projects\\src\\Strava_Light\\resources\\gpx\\2019-09-14_15-14.gpx"

func TestParseXml(t *testing.T) {
	//generate GpxFile object  from path
	var actualFile = ParseXml(testPath)

	//generate expected object
	var meta = gpx_info.NewMeta(time.Date(2019, 9, 14, 13, 14, 17,94000000, time.UTC))
	var point = gpx_info.NewTrackPoint(49.3547198, 0, time.Date(2019, 9, 14, 13, 14, 0, 0, time.UTC), 0)
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
	var actualTrackSegment = actualFile.GetTracks()[0].GetTrackSegments()[1].GetTrackPoints()[0]
	var expectedTrackSegment = expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0]

	//check similarity
	assert.Equal(t, actualFile.GetCreator(), expectedFile.GetCreator())
	assert.Equal(t, actualFile.GetMeta().GetTime(), expectedFile.GetMeta().GetTime())
	assert.Equal(t, actualTrackSegment.Longitude, expectedTrackSegment.Longitude)
	assert.Equal(t, actualTrackSegment.Latitude, expectedTrackSegment.Latitude)
	assert.Equal(t, actualTrackSegment.DateTime, expectedTrackSegment.DateTime)
	assert.Equal(t, actualTrackSegment.Speed, expectedTrackSegment.Speed)
}
