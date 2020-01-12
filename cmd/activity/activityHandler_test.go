/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"github.com/estellegraef/Strava_Light/cmd/storageManagement"
	"github.com/estellegraef/Strava_Light/resources"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var sortedActivities = []Activity{
	{
		Id:          "2",
		SportType:   "Radfahren",
		Comment:     "Let's go for a ride!",
		Length:      60.1,
		WaitingTime: 700,
		AvgSpeed:    24.3,
		MaxSpeed:    40.3,
		DateTime:     time.Date(2018, 9, 14, 12, 42, 31, 0000000, time.UTC),
	},
	{
		Id:          "1",
		SportType:   "Laufen",
		Comment:     "Let's go for a run!",
		Length:      24.6,
		WaitingTime: 120,
		AvgSpeed:    7.8,
		MaxSpeed:    12.6,
		DateTime:     time.Date(2018, 9, 22, 12, 42, 31, 0000000, time.UTC),
	},
}

var unsortedActivities = []Activity{
	{
		Id:          "1",
		SportType:   "Laufen",
		Comment:     "Let's go for a run!",
		Length:      24.6,
		WaitingTime: 120,
		AvgSpeed:    7.8,
		MaxSpeed:    12.6,
		DateTime:     time.Date(2018, 9, 22, 12, 42, 31, 0000000, time.UTC),
	},
	{
		Id:          "2",
		SportType:   "Radfahren",
		Comment:     "Let's go for a ride!",
		Length:      60.1,
		WaitingTime: 700,
		AvgSpeed:    24.3,
		MaxSpeed:    40.3,
		DateTime:     time.Date(2018, 9, 14, 12, 42, 31, 0000000, time.UTC),
	},
}

var firstActivityByte = []byte {123, 10, 9, 9, 34, 73, 100, 34, 58, 32, 34, 49, 34, 44, 10, 9, 9, 34, 83, 112, 111, 114,
	116, 84, 121, 112, 101, 34, 58, 32, 34, 76, 97, 117, 102, 101, 110, 34, 44, 10, 9, 9, 34, 67, 111, 109, 109, 101, 110,
	116, 34, 58, 32, 34, 76, 101, 116, 39, 115, 32, 103, 111, 32, 102, 111, 114, 32, 97, 32, 114, 117, 110, 33, 34, 44, 10,
	9, 9, 34, 76, 101, 110, 103, 116, 104, 34, 58, 32, 50, 52, 46, 54, 44, 10, 9, 9, 34, 87, 97, 105, 116, 105, 110, 103,
	84, 105, 109, 101, 34, 58, 32, 49, 50, 48, 44, 10, 9, 9, 34, 65, 118, 103, 83, 112, 101, 101, 100, 34, 58, 32, 55, 46,
	56, 44, 10, 9, 9, 34, 77, 97, 120, 83, 112, 101, 101, 100, 34, 58, 32, 49, 50, 46, 54, 44, 10, 9, 9, 34, 68, 97, 116,
	101, 84, 105, 109, 101, 34, 58, 32, 34, 50, 48, 49, 56, 45, 48, 57, 45, 50, 50, 84, 49, 50, 58, 52, 50, 58, 51, 49, 90,
	34, 10, 125}

func TestGetActivities(t *testing.T) {
	Setup()
	actualActivities := GetActivities("user1")
	assert.Equal(t, sortedActivities, actualActivities)
}

func TestSortActivities(t *testing.T) {
	actualActivities := SortActivities(unsortedActivities)
	assert.Equal(t, sortedActivities, actualActivities)
}

func TestGetActivity(t *testing.T) {
	Setup()
	actualActivity := GetActivity("user1", "1")
	assert.Equal(t, unsortedActivities[1], actualActivity)
}

func TestSearchActivitiesValidKeyword(t *testing.T) {
	Setup()
	matchingActivities := SearchActivities("user1", "go")
	assert.Equal(t, unsortedActivities, matchingActivities)
}

func TestSearchActivitiesInvalidKeyword(t *testing.T) {
	Setup()
	matchingActivities := SearchActivities("user1", "surfing")
	assert.Equal(t, []Activity(nil), matchingActivities)
}

func TestAddActivity(t *testing.T) {
	Setup()
	//TODO implement when multifile mocked
	/*
	isAdded := AddActivity("user1", "Laufen", file, fileHeader, "I love running")
	assert.True(t, isAdded)
	*/
}

func TestUpdateActivity(t *testing.T) {
	Setup()
	user := "user1"
	dir := resources.GetUserDir(user)
	id := "3"
	jsonFile := id + ".json"
	filemanagement.CreateFile(dir, jsonFile, firstActivityByte)
	isUpdated := UpdateActivity(user,  "3", "Radfahren", "They see me rollin")
	assert.True(t, isUpdated)
}

func TestDeleteActivity(t *testing.T) {
	Setup()
	user := "user1"
	dir := resources.GetUserDir(user)
	id := "3"
	jsonFile := id + ".json"
	zipFile := id + ".zip"
	filemanagement.CreateFile(dir, jsonFile, firstActivityByte)
	filemanagement.CreateFile(dir, zipFile, firstActivityByte)

	isDeleted := DeleteActivity(user, id)
	assert.True(t, isDeleted)
}

//TODO Marshal and Unmarshal individually work but not when file is executed in one go
func TestMarshalJSON(t *testing.T) {
	actual := MarshalJSON(unsortedActivities[0])
	assert.Equal(t, firstActivityByte, actual)
	time.Sleep(100)
}

func TestUnmarshalJSON(t *testing.T) {
	activity := UnmarshalJSON(firstActivityByte)
	assert.Equal(t, unsortedActivities[0], activity)
	time.Sleep(100)
}

func TestReturnFileForDownload(t *testing.T) {
	bytes, fileName := ReturnFileForDownload("user1", "1")
	expected := []byte{60, 63, 120, 109, 108, 32, 118, 101, 114, 115, 105, 111, 110, 61, 34, 49, 46, 48, 34, 32, 101, 110,
		99, 111, 100, 105, 110, 103, 61, 34, 85, 84, 70, 45, 56, 34, 63, 62, 13, 10, 60, 103, 112, 120, 32, 118, 101, 114,
		115, 105, 111, 110, 61, 34, 49, 46, 49, 34, 32, 99, 114, 101, 97, 116, 111, 114, 61, 34, 85, 114, 98, 97, 110, 32,
		66, 105, 107, 101, 114, 34, 32, 120, 115, 105, 58, 115, 99, 104, 101, 109, 97, 76, 111, 99, 97, 116, 105, 111, 110,
		61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 116, 111, 112, 111, 103, 114, 97, 102, 105, 120, 46, 99,
		111, 109, 47, 71, 80, 88, 47, 49, 47, 49, 32, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 116, 111, 112, 111,
		103, 114, 97, 102, 105, 120, 46, 99, 111, 109, 47, 71, 80, 88, 47, 49, 47, 49, 47, 103, 112, 120, 46, 120, 115, 100,
		32, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 103, 97, 114, 109, 105, 110, 46, 99, 111, 109, 47, 120, 109, 108,
		115, 99, 104, 101, 109, 97, 115, 47, 71, 112, 120, 69, 120, 116, 101, 110, 115, 105, 111, 110, 115, 47, 118, 51, 32, 104,
		116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 103, 97, 114, 109, 105, 110, 46, 99, 111, 109, 47, 120, 109, 108, 115, 99,
		104, 101, 109, 97, 115, 47, 71, 112, 120, 69, 120, 116, 101, 110, 115, 105, 111, 110, 115, 118, 51, 46, 120, 115, 100, 32,
		104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 103, 97, 114, 109, 105, 110, 46, 99, 111, 109, 47, 120, 109, 108, 115,
		99, 104, 101, 109, 97, 115, 47, 84, 114, 97, 99, 107, 80, 111, 105, 110, 116, 69, 120, 116, 101, 110, 115, 105, 111, 110,
		47, 118, 50, 32, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 103, 97, 114, 109, 105, 110, 46, 99, 111, 109, 47, 120,
		109, 108, 115, 99, 104, 101, 109, 97, 115, 47, 84, 114, 97, 99, 107, 80, 111, 105, 110, 116, 69, 120, 116, 101, 110, 115, 105,
		111, 110, 118, 50, 46, 120, 115, 100, 34, 32, 120, 109, 108, 110, 115, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119,
		46, 116, 111, 112, 111, 103, 114, 97, 102, 105, 120, 46, 99, 111, 109, 47, 71, 80, 88, 47, 49, 47, 49, 34, 32, 120, 109, 108,
		110, 115, 58, 103, 112, 120, 116, 112, 120, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 103, 97, 114, 109, 105,
		110, 46, 99, 111, 109, 47, 120, 109, 108, 115, 99, 104, 101, 109, 97, 115, 47, 84, 114, 97, 99, 107, 80, 111, 105, 110, 116, 69,
		120, 116, 101, 110, 115, 105, 111, 110, 47, 118, 50, 34, 32, 120, 109, 108, 110, 115, 58, 103, 112, 120, 112, 120, 61, 34, 104,
		116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 103, 97, 114, 109, 105, 110, 46, 99, 111, 109, 47, 120, 109, 108, 115, 99, 104, 101,
		109, 97, 115, 47, 80, 111, 119, 101, 114, 69, 120, 116, 101, 110, 115, 105, 111, 110, 118, 49, 46, 120, 115, 100, 34, 32, 120, 109,
		108, 110, 115, 58, 103, 112, 120, 120, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 103, 97, 114, 109, 105, 110, 46,
		99, 111, 109, 47, 120, 109, 108, 115, 99, 104, 101, 109, 97, 115, 47, 71, 112, 120, 69, 120, 116, 101, 110, 115, 105, 111, 110, 115,
		47, 118, 51, 34, 32, 120, 109, 108, 110, 115, 58, 120, 115, 105, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 119, 51,
		46, 111, 114, 103, 47, 50, 48, 48, 49, 47, 88, 77, 76, 83, 99, 104, 101, 109, 97, 45, 105, 110, 115, 116, 97, 110, 99, 101, 34, 62,
		13, 10, 13, 10, 60, 109, 101, 116, 97, 100, 97, 116, 97, 62, 13, 10, 60, 108, 105, 110, 107, 32, 104, 114, 101, 102, 61, 34, 104, 116,
		116, 112, 58, 47, 47, 119, 119, 119, 46, 117, 114, 98, 97, 110, 45, 98, 105, 107, 101, 45, 99, 111, 109, 112, 117, 116, 101, 114, 46,
		99, 111, 109, 47, 34, 62, 13, 10, 60, 116, 101, 120, 116, 62, 85, 114, 98, 97, 110, 32, 66, 105, 107, 101, 114, 60, 47, 116, 101, 120,
		116, 62, 13, 10, 60, 47, 108, 105, 110, 107, 62, 13, 10, 60, 116, 105, 109, 101, 62, 50, 48, 49, 57, 45, 48, 57, 45, 49, 52, 84, 49,
		51, 58, 49, 52, 58, 49, 55, 46, 48, 57, 52, 90, 60, 47, 116, 105, 109, 101, 62, 13, 10, 60, 47, 109, 101, 116, 97, 100, 97, 116, 97,
		62, 13, 10, 13, 10, 60, 116, 114, 107, 62, 13, 10, 60, 110, 97, 109, 101, 62, 67, 117, 98, 101, 32, 226, 128, 147, 32, 83, 97, 46, 44,
		32, 49, 52, 46, 32, 83, 101, 112, 46, 32, 50, 48, 49, 57, 44, 32, 49, 53, 58, 49, 52, 60, 47, 110, 97, 109, 101, 62, 13, 10, 60, 116,
		114, 107, 115, 101, 103, 62, 13, 10, 60, 116, 114, 107, 112, 116, 32, 108, 97, 116, 61, 34, 52, 57, 46, 51, 53, 52, 57, 56, 57, 48, 54,
		48, 48, 34, 32, 108, 111, 110, 61, 34, 57, 46, 49, 53, 49, 57, 54, 52, 57, 52, 48, 48, 34, 62, 60, 101, 108, 101, 62, 49, 55, 50, 46,
		53, 48, 60, 47, 101, 108, 101, 62, 60, 116, 105, 109, 101, 62, 50, 48, 49, 57, 45, 48, 57, 45, 49, 52, 84, 49, 51, 58, 49, 52, 58, 51,
		48, 46, 50, 55, 54, 90, 60, 47, 116, 105, 109, 101, 62, 60, 101, 120, 116, 101, 110, 115, 105, 111, 110, 115, 62, 60, 103, 112, 120,
		116, 112, 120, 58, 84, 114, 97, 99, 107, 80, 111, 105, 110, 116, 69, 120, 116, 101, 110, 115, 105, 111, 110, 62, 60, 103, 112, 120,
		116, 112, 120, 58, 97, 116, 101, 109, 112, 62, 50, 50, 46, 56, 60, 47, 103, 112, 120, 116, 112, 120, 58, 97, 116, 101, 109, 112, 62,
		60, 103, 112, 120, 116, 112, 120, 58, 115, 112, 101, 101, 100, 62, 53, 46, 53, 52, 60, 47, 103, 112, 120, 116, 112, 120, 58, 115, 112,
		101, 101, 100, 62, 60, 47, 103, 112, 120, 116, 112, 120, 58, 84, 114, 97, 99, 107, 80, 111, 105, 110, 116, 69, 120, 116, 101, 110, 115,
		105, 111, 110, 62, 60, 47, 101, 120, 116, 101, 110, 115, 105, 111, 110, 115, 62, 60, 47, 116, 114, 107, 112, 116, 62, 13, 10, 60, 47, 116,
		114, 107, 115, 101, 103, 62, 13, 10, 60, 47, 116, 114, 107, 62, 13, 10, 60, 47, 103, 112, 120, 62, 13, 10}
	assert.Equal(t,"1.gpx", fileName)
	assert.Equal(t, expected, bytes)
}