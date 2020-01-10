package activity

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	fmt.Println("START CACHE")
	cache := NewCache()
	activities := []Activity{
		{ Name:        "first",
			Id:          1,
			SportType:   "Laufen",
			Comment:     "Let's go for a run!",
			Length:      24.6,
			WaitingTime: 120,
			AvgSpeed:    7.8,
			MaxSpeed:    12.6,
			DateTime:     time.Date(2018, 9, 22, 12, 42, 31, 0, time.UTC),
		},
		{ Name:        "second",
			Id:          2,
			SportType:   "Laufen",
			Comment:     "Let's go for a run!",
			Length:      24.6,
			WaitingTime: 120,
			AvgSpeed:    7.8,
			MaxSpeed:    12.6,
			DateTime:     time.Date(2018, 9, 22, 12, 42, 31, 0, time.UTC),
		},
		{ Name:        "third",
			Id:          3,
			SportType:   "Laufen",
			Comment:     "Let's go for a run!",
			Length:      24.6,
			WaitingTime: 120,
			AvgSpeed:    7.8,
			MaxSpeed:    12.6,
			DateTime:     time.Date(2018, 9, 22, 12, 42, 31, 0, time.UTC),
		},

	}
	for _, word := range activities{
		cache.Check(word)
		cache.Display()
	}
}
