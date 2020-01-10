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
	Name		string
	Id          uint32
	SportType   string
	Comment     string
	Length      float64
	WaitingTime float64
	AvgSpeed    float64
	MaxSpeed    float64
	DateTime    time.Time
}

func New(id uint32, sportType string, comment string, length float64, waitingTime float64, avgSpeed float64, maxSpeed float64, dateTime time.Time) Activity {
	return Activity{Id: id, SportType: sportType, Comment: comment, Length: length, WaitingTime: waitingTime, AvgSpeed: avgSpeed, MaxSpeed: maxSpeed, DateTime: dateTime}
}

func (a Activity) GetName() string {
	return a.GetName()
}

func (a Activity) GetSportType() string {
	return a.SportType
}

func (a Activity) GetComment() string {
	return a.Comment
}

func (a Activity) GetID() uint32 {
	return a.Id
}
