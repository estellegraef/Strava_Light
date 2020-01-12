/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var activity = Activity{
	id:          "2",
	SportType:   "Radfahren",
	Comment:     "Let's go for a ride!",
	Length:      60.1,
	WaitingTime: 700,
	AvgSpeed:    24.3,
	MaxSpeed:    40.3,
	DateTime:    time.Date(2018, 9, 19, 12, 42, 31, 0000000, time.UTC),
}

func TestNew(t *testing.T) {
	actualActivity := New("2", "Radfahren", "Let's go for a ride!",  60.1, 700, 24.3, 40.3, time.Date(2018, 9, 19, 12, 42, 31, 0000000, time.UTC))
	assert.Equal(t, activity, actualActivity)
}

func TestActivity_GetID(t *testing.T) {
	id := activity.GetID()
	assert.Equal(t, "2", id)
}

func TestActivity_GetSportType(t *testing.T) {
	sportType := activity.GetSportType()
	assert.Equal(t, "Radfahren", sportType)
}

func TestActivity_GetComment(t *testing.T) {
	comment := activity.GetComment()
	assert.Equal(t, "Let's go for a ride!", comment)
}
