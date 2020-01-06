package cmd

import (
	"Strava_Light/cmd/gpx/activity"
	"mime/multipart"
	"time"
)

var list []activity.Activity

//Creates mockup data
func GetActivities() []activity.Activity {
	list = make([]activity.Activity, 10)

	for i := 0; i < 5; i++ {
		list[i] = activity.Activity{
			uint32(i),
			"Radfahren",
			"test",
			12.3,
			5.2,
			12.6,
			15.9,
			time.Now(),
		}
		for i := 5; i < 10; i++ {
			list[i] = activity.Activity{
				uint32(i),
				"Laufen",
				"turbo",
				12.3,
				5.2,
				12.6,
				15.9,
				time.Now().AddDate(0, 0, -2),
			}
		}
	}

	return list
}

func SearchActivities(search string) []activity.Activity {
	result := make([]activity.Activity, 0)
	_ = GetActivities()

	for _, elem := range list {
		if elem.comment == search {
			result = append(result, elem)
		}
	}

	return result
}

func GetActivity() activity.Activity {
	return activity.New(1, "Radfahren", "I am a useful comment",
		12.3, 5.2, 13.4, 17.8, time.Now())
}

func CreateActivity(sportType string, file multipart.File, fileHeader *multipart.FileHeader, comment string) bool {
	return false
}
