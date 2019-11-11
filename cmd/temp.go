package cmd

import (
	"mime/multipart"
	"time"
)

//Creates mockup data
func GetActivities() []ActivityDetail {
	list := make([]ActivityDetail, 10)

	for i := 0; i < 10; i++ {
		list[i] = NewDetail(
			uint32(i),
			"Radfahren",
			"test",
			12.3,
			5.2,
			12.6,
			15.9,
			time.Now())
	}

	return list
}

func GetActivity() ActivityDetail {
	return NewDetail(1, "Radfahren", "I am a useful comment",
		12.3, 5.2, 13.4, 17.8, time.Now())
}

func CreateActivity(sportType string, file multipart.File, fileHeader *multipart.FileHeader, comment string) bool {
	return true
}
