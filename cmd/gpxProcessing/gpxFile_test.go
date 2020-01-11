/*
 * 2848869
 * 8089098
 * 3861852
 */

package gpxProcessing

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var expectedFile = GpxFile{
	Creator: "Urban Biker",
	Meta: Metadata {
		Time: time.Date(2019, 9, 14, 13, 14, 17, 94000000, time.UTC),
	},
	Tracks: []Track{
		{	TrackSegments: []TrackSegment {
			{	TrackPoints: []TrackPoint {
				{	Latitude: 49.35498906,
					Longitude: 9.15196494,
					DateTime:  time.Date(2019, 9, 14, 13, 14, 30, 276000000, time.UTC),
					Extensions: Extension{
						TrackPointExtensions: TrackPointExtension{
						Speed: 5.54,
						},
					},
				},
				},
			},
			},
		},
		},
	}

func TestGpxFile_GetMeta(t *testing.T) {
	expected := Metadata{Time:time.Date(2019, 9, 14, 13, 14, 17, 94000000, time.UTC)}
	actual := expectedFile.GetMeta()
	assert.Equal(t, expected, actual)
}

func TestTrackSegment_GetTrackPoints(t *testing.T) {
	expected := []Track{
		{	TrackSegments: []TrackSegment {
			{	TrackPoints: []TrackPoint {
				{	Latitude: 49.35498906,
					Longitude: 9.15196494,
					DateTime:  time.Date(2019, 9, 14, 13, 14, 30, 276000000, time.UTC),
					Extensions: Extension{
						TrackPointExtensions: TrackPointExtension{
							Speed: 5.54,
						},
					},
				},
			},
			},
		},
		},
	}
	actual := expectedFile.GetTracks()
	assert.Equal(t, expected, actual)
}

func TestGpxFile_GetTrackPoints(t *testing.T) {
	expected := []TrackPoint{
		{Latitude: 49.35498906,
			Longitude: 9.15196494,
			DateTime:  time.Date(2019, 9, 14, 13, 14, 30, 276000000, time.UTC),
			Extensions: Extension{
				TrackPointExtensions: TrackPointExtension{
					Speed: 5.54,
				},
			},
		},
	}
	actual := expectedFile.GetTrackPoints()
	assert.Equal(t, expected, actual)
}

func TestMetadata_GetTime(t *testing.T) {
	expected := time.Date(2019, 9, 14, 13, 14, 17, 94000000, time.UTC)
	actual := expectedFile.Meta.GetTime()
	assert.Equal(t, expected, actual)
}

func TestTrack_GetTrackSegments(t *testing.T) {
	expected := []TrackSegment {
		{	TrackPoints: []TrackPoint {
			{	Latitude: 49.35498906,
				Longitude: 9.15196494,
				DateTime:  time.Date(2019, 9, 14, 13, 14, 30, 276000000, time.UTC),
				Extensions: Extension{
					TrackPointExtensions: TrackPointExtension{
						Speed: 5.54,
					},
				},
			},
		},
		},
	}
	testTrack := expectedFile.GetTracks()[0]
	actual := testTrack.GetTrackSegments()
	assert.Equal(t, expected, actual)
}

func TestTrackPoint_GetLatitude(t *testing.T) {
	testTp := expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0]
	actualLatitude := testTp.GetLatitude()
	assert.Equal(t, 49.35498906, actualLatitude)
}

func TestTrackPoint_GetLongitude(t *testing.T) {
	testTp := expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0]
	actualLongitude := testTp.GetLongitude()
	assert.Equal(t, 9.15196494, actualLongitude)
}

func TestTrackPoint_GetDateTime(t *testing.T) {
	expectedTime := time.Date(2019, 9, 14, 13, 14, 30, 276000000, time.UTC)
	testTp := expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0]
	actualTime := testTp.GetDateTime()
	assert.Equal(t, expectedTime, actualTime)
}

func TestTrackPoint_GetExtension(t *testing.T) {
	expectedExt := Extension{TrackPointExtensions:TrackPointExtension{Speed:5.54}}
	testTp := expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0]
	actualExt := testTp.GetExtension()
	assert.Equal(t, expectedExt, actualExt)
}

func TestExtension_GetTrackPointExtension(t *testing.T) {
	expectedExt := TrackPointExtension{Speed:5.54}
	testExt := expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0].GetExtension()
	actualExt := testExt.GetTrackPointExtension()
	assert.Equal(t, expectedExt, actualExt)
}

func TestTrackPointExtension_GetSpeed(t *testing.T) {
	testExt := expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0].GetExtension().GetTrackPointExtension()
	actual := testExt.GetSpeed()
	assert.Equal(t, 5.54, actual)
}

func TestGpxFile_GetDistanceInKilometers(t *testing.T) {
	actual := expectedFile.GetDistanceInKilometers()
	assert.Equal(t, float64(0), actual)
}

func TestGpxFile_GetAvgSpeed(t *testing.T) {
	actual := expectedFile.GetAvgSpeed()
	assert.Equal(t, 5.54, actual)
}

func TestGpxFile_GetMaxSpeed(t *testing.T) {
	actual := expectedFile.GetMaxSpeed()
	assert.Equal(t, 5.54, actual)
}