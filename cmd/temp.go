package cmd

import (
	"mime/multipart"
	"time"
)

var list []Activity

//Creates mockup data
func GetActivities() []Activity {
	list = make([]Activity, 10)

	for i := 0; i < 5; i++ {
		list[i] = Activity{
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
			list[i] = Activity{
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

func SearchActivities(search string) []Activity {
	result := make([]Activity, 0)
	_ = GetActivities()

	for _, elem := range list {
		if elem.comment == search {
			result = append(result, elem)
		}
	}

	return result
}

func GetActivity() Activity {
	return New(1, "Radfahren", "I am a useful comment",
		12.3, 5.2, 13.4, 17.8, time.Now())
}

func CreateActivity(sportType string, file multipart.File, fileHeader *multipart.FileHeader, comment string) bool {
	return false
}
