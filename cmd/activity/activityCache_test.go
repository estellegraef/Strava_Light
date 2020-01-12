/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"fmt"
	"testing"
	"time"
)

var activities = []Activity{
	{
		Id:          "2",
		SportType:   "Laufen",
		Comment:     "Let's go for a run!",
		Length:      24.6,
		WaitingTime: 120,
		AvgSpeed:    7.8,
		MaxSpeed:    12.6,
		DateTime:    time.Date(2018, 9, 22, 12, 42, 31, 0, time.UTC),
	},
	{
		Id:          "2",
		SportType:   "Laufen",
		Comment:     "Let's go for a run!",
		Length:      24.6,
		WaitingTime: 120,
		AvgSpeed:    7.8,
		MaxSpeed:    12.6,
		DateTime:     time.Date(2018, 9, 22, 12, 42, 31, 0, time.UTC),
	},
	{
		Id:          "3",
		SportType:   "Laufen",
		Comment:     "Let's go for a run!",
		Length:      24.6,
		WaitingTime: 120,
		AvgSpeed:    7.8,
		MaxSpeed:    12.6,
		DateTime:     time.Date(2018, 9, 22, 12, 42, 31, 0, time.UTC),
	},
}

//TODO test cache
func TestNewCache(t *testing.T) {
	fmt.Println("START CACHE")
	cache := NewCache()

	for _, word := range activities{
		cache.Check(word)
	}
}
