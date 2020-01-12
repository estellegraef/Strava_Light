/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var activities = []Activity{
	{
		Id:          "1",
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
	cache := NewCache()

	for _, word := range activities{
		cache.Check(word.Id, word)
	}
	isInCache, activity := cache.GetActivity("1")
	assert.Equal(t, true, isInCache)
	assert.Equal(t, activities[0], activity)
}

func TestCache_GetNodeTrue(t *testing.T) {
	cache := NewCache()

	cache.Check(activities[1].Id, activities[1])
	cache.Check(activities[2].Id, activities[2])

	isInCache, cacheActivity := cache.GetActivity("2")
	assert.True(t, isInCache)
	assert.Equal(t, activities[1], cacheActivity)
}

func TestCache_GetNodeFalse(t *testing.T) {
	cache := NewCache()

	cache.Check(activities[0].Id, activities[0])
	cache.Check(activities[1].Id, activities[1])

	isInCache, cacheActivity := cache.GetActivity("3")
	fmt.Println(isInCache)
	fmt.Println(cacheActivity)
	assert.False(t, isInCache)
}

