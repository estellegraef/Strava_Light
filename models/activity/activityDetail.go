package activity

import (
	"strconv"
	"time"
)

type ActivityDetail struct {
	id          uint32
	sportType   string
	comment     string
	length      float64
	waitingTime float64
	avgSpeed    float64
	maxSpeed    float64
	dateTime    time.Time
}

func NewDetail(id uint32, sportType string, comment string, length float64, waitingTime float64, avgSpeed float64, maxSpeed float64, dateTime time.Time) ActivityDetail {
	return ActivityDetail{id: id, sportType: sportType, comment: comment, length: length, waitingTime: waitingTime, avgSpeed: avgSpeed, maxSpeed: maxSpeed, dateTime: dateTime}
}

func (a ActivityDetail) GetSportType() string {
	return a.sportType
}

func (a ActivityDetail) GetWeekDay() time.Weekday {
	return a.dateTime.Weekday()
}

func (a ActivityDetail) GetDate() string {
	return strconv.Itoa(a.dateTime.Day()) + "." + a.dateTime.Month().String() + " " + strconv.Itoa(a.dateTime.Year())
}

func (a ActivityDetail) GetTime() string {
	return strconv.Itoa(a.dateTime.Hour()) + ":" + strconv.Itoa(a.dateTime.Minute())
}

func (a ActivityDetail) GetLength() float64 {
	return a.length
}

func (a ActivityDetail) GetComment() string {
	return a.comment
}

func (a ActivityDetail) GetAvgSpeed() float64 {
	return a.avgSpeed
}

func (a ActivityDetail) GetMaxSpeed() float64 {
	return a.maxSpeed
}

func (a ActivityDetail) GetWaitingTime() float64 {
	return a.waitingTime
}

func (a ActivityDetail) GetID() uint32 {
	return a.id
}
