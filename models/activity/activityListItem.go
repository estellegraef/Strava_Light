package activity

import (
	"time"
)

type ActivityListItem struct {
	id        uint32
	sportType string
	dateTime  time.Time
}

func NewListItem(id uint32, sportType string, dateTime time.Time) ActivityListItem {
	return ActivityListItem{id: id, sportType: sportType, dateTime: dateTime}
}

func (a ActivityListItem) GetID() uint32 {
	return a.id
}

func (a ActivityListItem) GetSportType() string {
	return a.sportType
}

func (a ActivityListItem) GetTimeStamp() time.Time {
	return a.dateTime
}
