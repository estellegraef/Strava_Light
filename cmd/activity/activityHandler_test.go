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
	{ Name:        "2",
		Id:          2,
		SportType:   "Radfahren",
		Comment:     "Let's go for a ride!",
		Length:      60.1,
		WaitingTime: 700,
		AvgSpeed:    24.3,
		MaxSpeed:    40.3,
		DateTime:     time.Date(2018, 9, 19, 12, 42, 31, 0000000, time.UTC),
	},
	{ Name:        "1",
		Id:          1,
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
	{ Name:        "2",
		Id:          2,
		SportType:   "Radfahren",
		Comment:     "Let's go for a ride!",
		Length:      60.1,
		WaitingTime: 700,
		AvgSpeed:    24.3,
		MaxSpeed:    40.3,
		DateTime:     time.Date(2018, 9, 19, 12, 42, 31, 0000000, time.UTC),
	},
	{ Name:        "1",
		Id:          1,
		SportType:   "Laufen",
		Comment:     "Let's go for a run!",
		Length:      24.6,
		WaitingTime: 120,
		AvgSpeed:    7.8,
		MaxSpeed:    12.6,
		DateTime:     time.Date(2018, 9, 22, 12, 42, 31, 0000000, time.UTC),
	},
}

var firstActivityByte = []byte{0x7b, 0xa, 0x9, 0x9, 0x22, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x22,
	0x2c, 0xa, 0x9, 0x9, 0x22, 0x49, 0x64, 0x22, 0x3a, 0x20, 0x32, 0x2c, 0xa, 0x9, 0x9, 0x22, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x52, 0x61, 0x64, 0x66, 0x61, 0x68, 0x72, 0x65, 0x6e, 0x22,
	0x2c, 0xa, 0x9, 0x9, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x4c, 0x65, 0x74, 0x27,
	0x73, 0x20, 0x67, 0x6f, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x72, 0x69, 0x64, 0x65, 0x21, 0x22, 0x2c, 0xa, 0x9,
	0x9, 0x22, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x3a, 0x20, 0x36, 0x30, 0x2e, 0x31, 0x2c, 0xa, 0x9, 0x9, 0x22,
	0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x37, 0x30, 0x30, 0x2c, 0xa, 0x9,
	0x9, 0x22, 0x41, 0x76, 0x67, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x32, 0x34, 0x2e, 0x33, 0x2c, 0xa, 0x9, 0x9,
	0x22, 0x4d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x34, 0x30, 0x2e, 0x33, 0x2c, 0xa, 0x9, 0x9, 0x22,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x31, 0x38, 0x2d, 0x30, 0x39, 0x2d, 0x31,
	0x39, 0x54, 0x31, 0x32, 0x3a, 0x34, 0x32, 0x3a, 0x33, 0x31, 0x5a, 0x22, 0xa, 0x7d}

func TestGetActivities(t *testing.T) {
	actualActivities := GetActivities("user1")
	assert.Equal(t, sortedActivities, actualActivities)
}

func TestSortActivities(t *testing.T) {
	actualActivities := SortActivities(unsortedActivities)
	assert.Equal(t, sortedActivities, actualActivities)
}

func TestGetActivity(t *testing.T) {
	actualActivity := GetActivity("user1", 1)
	assert.Equal(t, unsortedActivities[1], actualActivity)
}

func TestSearchActivitiesValidKeyword(t *testing.T) {
	matchingActivities := SearchActivities("user1", "go")
	assert.Equal(t, unsortedActivities, matchingActivities)
}

func TestSearchActivitiesInvalidKeyword(t *testing.T) {
	matchingActivities := SearchActivities("user1", "surfing")
	assert.Equal(t, []Activity(nil), matchingActivities)
}

func TestAddActivity(t *testing.T) {
	//TODO implement when multifile mocked
	/*
	isAdded := AddActivity("user1", "Laufen", file, fileHeader, "I love running")
	assert.True(t, isAdded)
	*/
}

func TestUpdateActivity(t *testing.T) {
	user := "user1"
	dir := resources.GetUserDir(user)
	id := "3"
	jsonFile := id + ".json"
	filemanagement.CreateFile(dir, jsonFile, firstActivityByte)
	isUpdated := UpdateActivity(user,  3, "Radfahren", "They see me rollin")
	assert.True(t, isUpdated)
}

func TestDeleteActivity(t *testing.T) {
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

func TestMarshalJSON(t *testing.T) {
	actual := MarshalJSON(unsortedActivities[0])
	assert.Equal(t, firstActivityByte, actual)
}

func TestUnmarshalJSON(t *testing.T) {
	activity := UnmarshalJSON(firstActivityByte)
	assert.Equal(t, unsortedActivities[0], activity)
}