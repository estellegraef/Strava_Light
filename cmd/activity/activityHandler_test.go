package activity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetActivities(t *testing.T) {
	actualActivities := GetActivities("user1")
	expectedActivities := []Activity{
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
	assert.Equal(t, expectedActivities, actualActivities)
}

func TestGetActivity(t *testing.T) {
	actualActivity := GetActivity("user1", 1)
	expectedActivity := Activity{
		Name:        "Hello",
		Id:          1,
		SportType:   "Laufen",
		Comment:     "Let's go for a run!",
		Length:      24.6,
		WaitingTime: 120,
		AvgSpeed:    7.8,
		MaxSpeed:    12.6,
		DateTime:    time.Date(2018, 9, 22, 12, 42, 31, 0, time.UTC),
	}
	assert.Equal(t, expectedActivity, actualActivity)
}
