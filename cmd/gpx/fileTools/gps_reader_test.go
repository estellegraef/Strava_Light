package fileTools

import (
	"Strava_Light/cmd/gpx"
	"Strava_Light/cmd/gpx/gpx_info"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO create own, smaller file with only one TP to get full coverage
func TestReadGpx(t *testing.T) {
	//generate GpxFile object  from path
	var actualFile = ReadGpx(gpx.GetTestGpxPath())

	//create expected file
	expectedFile := gpx_info.GpxFile{
		Creator: "Urban Biker",
		Meta: gpx_info.Metadata{
			Time: time.Date(2019, 9, 14, 13, 14, 17, 94000000, time.UTC),
		},
		Tracks: []gpx_info.Track{
			{	TrackSegments: []gpx_info.TrackSegment{
				{	TrackPoints: []gpx_info.TrackPoint{
					{	Latitude: 49.3549890600,
						Longitude: 9.1519649400,
						DateTime: time.Date(2019, 9, 14, 13, 14, 30, 276000000, time.UTC),
						Extensions: gpx_info.Extension{
						TrackPointExtensions: gpx_info.TrackPointExtension{
							Speed:5.54,
						},
						},
					},
				},
				},
			},
			},
		},
	}
	//extract same TrackPoint
	var actualTrackSegment = actualFile.GetTracks()[0].GetTrackSegments()[1].GetTrackPoints()[4]
	var expectedTrackSegment = expectedFile.GetTracks()[0].GetTrackSegments()[0].GetTrackPoints()[0]

	assert.Equal(t, expectedFile.GetCreator(), actualFile.GetCreator())
	assert.Equal(t, expectedFile.GetMeta(), actualFile.GetMeta())
	assert.Equal(t, expectedTrackSegment, actualTrackSegment)
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
	actualFiles := ReadFile(gpx.GetTestInvalidPath())
	//since the path is invalid, a empty object is expected
	expectedFiles := []gpx_info.GpxFile(nil)
	assert.Equal(t, expectedFiles, actualFiles)
}

func TestCheckFileNonExistentPositive(t *testing.T) {
	nonExistent := CheckFileNonExistent(gpx.GetTestInvalidPath())
	assert.True(t, nonExistent)
}

func TestCheckFileNonExistentNegative(t *testing.T) {
	nonExistent := CheckFileNonExistent(gpx.GetTestGpxPath())
	assert.False(t, nonExistent)
}


