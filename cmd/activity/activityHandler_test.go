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

var firstActivityByte = []byte {123, 10, 9, 9, 34, 73, 100, 34, 58, 32, 34, 49, 34, 44, 10, 9, 9, 34, 83, 112, 111, 114, 116, 84, 121, 112, 101, 34, 58, 32, 34, 76, 97, 117, 102, 101, 110, 34, 44, 10, 9, 9, 34, 67, 111, 109, 109, 101, 110, 116, 34, 58, 32, 34, 76, 101, 116, 39, 115, 32, 103, 111, 32, 102, 111, 114, 32, 97, 32, 114, 117, 110, 33, 34, 44, 10, 9, 9, 34, 76, 101, 110, 103, 116, 104, 34, 58, 32, 50, 52, 46, 54, 44, 10, 9, 9, 34, 87, 97, 105, 116, 105, 110, 103, 84, 105, 109, 101, 34, 58, 32, 49, 50, 48, 44, 10, 9, 9, 34, 65, 118, 103, 83, 112, 101, 101, 100, 34, 58, 32, 55, 46, 56, 44, 10, 9, 9, 34, 77, 97, 120, 83, 112, 101, 101, 100, 34, 58, 32, 49, 50, 46, 54, 44, 10, 9, 9, 34, 68, 97, 116, 101, 84, 105, 109, 101, 34, 58, 32, 34, 50, 48, 49, 56, 45, 48, 57, 45, 50, 50, 84, 49, 50, 58, 52, 50, 58, 51, 49, 90, 34, 10, 125}

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