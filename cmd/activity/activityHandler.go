/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"encoding/json"
	"fmt"
	"github.com/estellegraef/Strava_Light/cmd/storageManagement"
	"github.com/estellegraef/Strava_Light/resources"
	"log"
	"mime/multipart"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var list []Activity

//TODO implement cache after tests

func GetActivities(user string) []Activity {
	userDir := resources.GetUserDir(user)
	files := filemanagement.GetAllFilesFromDir(userDir)
	var activities []Activity
	for _, file := range files {
		if filepath.Ext(file) == ".json" {
			activities = append(activities, GetActivity(user, GetIdByName(file)))
		}
	}
	return SortActivities(activities)
}

func SortActivities(activities []Activity) []Activity{
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].DateTime.Before(activities[j].DateTime)
	})
	return activities
}
//TODO request change to id string from uint32
func GetActivity(user string, id uint32) Activity {
	userDir := resources.GetUserDir(user)
	files := filemanagement.GetAllFilesFromDir(userDir)
	var activity Activity
	for _, file := range files {
		searchedFile := GetNameById(id)
		if filepath.Base(file) == searchedFile {
			content := filemanagement.ReadFile(file)
			activity = UnmarshalJSON(content)
		}
	}
	return activity
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

func AddActivity(username string, sportType string, file multipart.File, header *multipart.FileHeader, comment string) bool {
	//generate uuid
	//read + eval multipart file -> GPX info
		//read bytes -> save + read with gpxprocessing or readbytes and generate gpx without saving
	//new Activity object -> createActivity
	//save to usr if not already saved
	return true
}

func UpdateActivity(user string, id uint32, sportType string, comment string) bool {
	activity := GetActivity(user, id)
	activity.SportType = sportType
	activity.Comment = comment
	//save
	return true
}

func DeleteActivity(user string, id string) bool {
	//activity := GetActivity(user, id)
	//delete from cache
	//delete json + gpx/zip file from directory
	return true
}

func MarshalJSON(activity Activity) []byte {
	jsonData, err := json.MarshalIndent(activity, "", "		")
	if err != nil {
		log.Println(err)
	}
	return jsonData
}

func UnmarshalJSON(data []byte) Activity {
	var activity Activity
	err := json.Unmarshal(data, &activity)
	if err != nil {
		log.Println(err)
	}
	return activity
}

func GetNameById(id uint32) string {
	return fmt.Sprint(id) + ".json"
}

func GetIdByName(name string) uint32 {
	filename := strings.TrimSuffix(filepath.Base(name), ".json")
	val, _ := strconv.Atoi(filename)
	return uint32(val)
}