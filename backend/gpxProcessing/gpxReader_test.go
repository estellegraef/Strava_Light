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
	"time"
)

func TestReadGpx(t *testing.T) {
	//generate GpxFile object  from path
	var actualFile = ReadGpx(resources.GetTestShortGpx())

	//create expected file
	expectedFile := GpxFile{
		Creator: "Urban Biker",
		Meta: Metadata{
			Time: time.Date(2019, 9, 14, 13, 14, 17, 94000000, time.UTC),
		},
		Tracks: []Track{
			{TrackSegments: []TrackSegment{
				{TrackPoints: []TrackPoint{
					{Latitude: 49.35498906,
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
	assert.Equal(t, expectedFile, actualFile)
}

func TestReadZip(t *testing.T) {
	files := ReadZip(resources.GetTestZipPath())
	assert.True(t, len(files) == 1)
}

func TestReadFileWithGpx(t *testing.T) {
	files := ReadFile(resources.GetTestGpxPath())
	assert.True(t, len(files) == 1)
}

func TestReadFileWithZip(t *testing.T) {
	files := ReadFile(resources.GetTestZipPath())
	assert.True(t, len(files) == 1)
}

func TestReadFileInvalidPath(t *testing.T) {
	actualFiles := ReadFile(resources.GetTestInvalidPath())
	//since the path is invalid, an empty GpxFile array is expected
	expectedFiles := []GpxFile(nil)
	assert.Equal(t, expectedFiles, actualFiles)
}

func TestCheckFileNonExistentPositive(t *testing.T) {
	nonExistent := CheckFileNonExistent(resources.GetTestInvalidPath())
	assert.True(t, nonExistent)
}

func TestCheckFileNonExistentNegative(t *testing.T) {
	nonExistent := CheckFileNonExistent(resources.GetTestGpxPath())
	assert.False(t, nonExistent)
}
