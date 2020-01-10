/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"mime/multipart"
	"time"
)

var list []Activity

//Creates mockup data
func GetActivities(username string) []Activity {
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

func SearchActivities(username string, search string) []Activity {
	result := make([]Activity, 0)

	for _, elem := range list {
		if elem.GetComment() == search {
			result = append(result, elem)
		}
	}

	return result
}

func GetActivity(username string, id uint32) Activity {
	for _, elem := range list {
		if elem.GetID() == id {
			return elem
		}
	}
	return Activity{}
}

func CreateActivity(username string, sportType string, file multipart.File, fileHeader *multipart.FileHeader, comment string) bool {
	return true
}
