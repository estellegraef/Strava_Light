/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"time"
)

type Activity struct {
	id          string
	sportType   string
	comment     string
	length      float64
	waitingTime float64
	avgSpeed    float64
	maxSpeed    float64
	dateTime    time.Time
}

func New(id string, sportType string, comment string, length float64, waitingTime float64, avgSpeed float64, maxSpeed float64, dateTime time.Time) Activity {
	return Activity{id: id, sportType: sportType, comment: comment, length: length, waitingTime: waitingTime, avgSpeed: avgSpeed, maxSpeed: maxSpeed, dateTime: dateTime}
}

func (a Activity) GetSportType() string {
	return a.sportType
}

func (a Activity) GetWeekDay() time.Weekday {
	return a.dateTime.Weekday()
}

func (a Activity) GetLongDate() string {
	return a.dateTime.Format("02.January 2006")
}

func (a Activity) GetShortDate() string {
	return a.dateTime.Format("02.01.2006")
}

func (a Activity) GetTime() string {
	return a.dateTime.Format("15:04")
}

func (a Activity) GetLength() float64 {
	return a.length
}

func (a Activity) GetComment() string {
	return a.comment
}

func (a Activity) GetAvgSpeed() float64 {
	return a.avgSpeed
}

func (a Activity) GetMaxSpeed() float64 {
	return a.maxSpeed
}

func (a Activity) GetWaitingTime() float64 {
	return a.waitingTime
}

func (a Activity) GetID() string {
	return a.id
}
