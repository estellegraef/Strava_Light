package activity

import (
	"time"
)

type Activity struct {
	Id          uint32
	sportType   string
	Comment     string
	length      float64
	waitingTime float64
	avgSpeed    float64
	maxSpeed    float64
	dateTime    time.Time
}

func New(id uint32, sportType string, comment string, length float64, waitingTime float64, avgSpeed float64, maxSpeed float64, dateTime time.Time) Activity {
	return Activity{Id: id, sportType: sportType, Comment: comment, length: length, waitingTime: waitingTime, avgSpeed: avgSpeed, maxSpeed: maxSpeed, dateTime: dateTime}
}

func (a Activity) GetSportType() string {
	return a.sportType
}

func (a Activity) GetComment() string {
	return a.Comment
}

func (a Activity) GetID() uint32 {
	return a.Id
}
