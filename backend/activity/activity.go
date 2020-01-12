/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"math"
	"time"
)

type Activity struct {
	Id          string
	SportType   string
	Comment     string
	Length      float64
	WaitingTime float64
	AvgSpeed    float64
	MaxSpeed    float64
	DateTime    time.Time
}

func New(id string, sportType string, comment string, length float64, waitingTime float64, avgSpeed float64, maxSpeed float64, dateTime time.Time) Activity {
	roundedLength := RoundToNearest(length)
	roundedAvg := RoundToNearest(avgSpeed)
	roundedMax := RoundToNearest(maxSpeed)
	roundedWaitingTime := SecondsToMinutes(waitingTime)
	return Activity{Id: id, SportType: sportType, Comment: comment, Length: roundedLength, WaitingTime: roundedWaitingTime, AvgSpeed: roundedAvg, MaxSpeed: roundedMax, DateTime: dateTime}
}

func (a Activity) GetSportType() string {
	return a.SportType
}

func (a Activity) GetWeekDay() time.Weekday {
	return a.DateTime.Weekday()
}

func (a Activity) GetLongDate() string {
	return a.DateTime.Format("02.January 2006")
}

func (a Activity) GetShortDate() string {
	return a.DateTime.Format("02.01.2006")
}

func (a Activity) GetTime() string {
	return a.DateTime.Format("15:04")
}

func (a Activity) GetLength() float64 {
	return a.Length
}

func (a Activity) GetComment() string {
	return a.Comment
}

func (a Activity) GetAvgSpeed() float64 {
	return a.AvgSpeed
}

func (a Activity) GetMaxSpeed() float64 {
	return a.MaxSpeed
}

func (a Activity) GetWaitingTime() float64 {
	return a.WaitingTime
}

func (a Activity) GetID() string {
	return a.Id
}

//from: https://yourbasic.org/golang/round-float-2-decimal-places/
func RoundToNearest(value float64) float64 {
	return math.Round(value*100) / 100
}

func SecondsToMinutes(seconds float64) float64 {
	mins := seconds / 60
	return RoundToNearest(mins)
}
