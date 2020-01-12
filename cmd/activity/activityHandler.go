/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"encoding/json"
	"github.com/estellegraef/Strava_Light/cmd/gpxProcessing"
	"github.com/estellegraef/Strava_Light/cmd/storageManagement"
	"github.com/estellegraef/Strava_Light/resources"
	"log"
	"mime/multipart"
	"path/filepath"
	"sort"
	"strings"
)

var cache Cache
//TODO implement cache after tests
//TODO CAUTION! GETIDBYNAME AND GETNAMEBYID HANDLED JSON EXT

func Setup(){
	cache = NewCache()
}

func GetActivities(userName string) []Activity {
	userDir := resources.GetUserDir(userName)
	files := filemanagement.GetAllFilesFromDir(userDir)
	var activities []Activity
	for _, fileName := range files {
		if filepath.Ext(fileName) == ".json" {
			title := strings.TrimSuffix(filepath.Base(fileName), ".json")
			var activity = GetActivity(userName, title)
			activities = append(activities, activity)
			cache.Check(activity)
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

func GetActivity(user string, id string) Activity {
	userDir := resources.GetUserDir(user)
	files := filemanagement.GetAllFilesFromDir(userDir)
	var activity Activity
	for _, file := range files {
		searchedFile := id + ".json"
		if filepath.Base(file) == searchedFile {
			content := filemanagement.ReadFile(file)
			activity = UnmarshalJSON(content)
		}
	}
	return activity
}

func SearchActivities(username string, keyword string) []Activity {
	userActivities := GetActivities(username)
	var matchingActivities []Activity
	for _, activity := range userActivities{
		if strings.Contains(activity.Comment, keyword) {
			matchingActivities = append(matchingActivities, activity)
		}
	}
	return matchingActivities
}

func AddActivity(userName string, sportType string, file multipart.File, header *multipart.FileHeader, comment string) bool {
	var success = false
	content := filemanagement.ReadReceiveFile(file)
	fileName := header.Filename
	baseIsCreated, createdFile := filemanagement.CreateFile(resources.GetUserDir(userName), fileName, content)
	if baseIsCreated {
		gpxFiles := gpxProcessing.ReadFile(createdFile)
		for _, file := range gpxFiles {
			id := filemanagement.GenerateId(fileName)
			activity := New(id, sportType, comment, file.GetDistanceInKilometers(), file.GetWaitingTime(), file.GetAvgSpeed(), file.GetMaxSpeed(), file.GetMeta().GetTime())
			content := MarshalJSON(activity)
			//TODO create File with ending json
			jsonTitle := fileName + ".json"
			activityIsCreated, _ := filemanagement.CreateFile(resources.GetUserDir(userName), jsonTitle, content)
			if activityIsCreated {
				//TODO push to cache
				success = true
			}
		}
	}
	return success
}

func UpdateActivity(user string, id string, sportType string, comment string) bool {
	activity := GetActivity(user, id)
	activity.SportType = sportType
	activity.Comment = comment
	content := MarshalJSON(activity)
	dir := resources.GetUserDir(user)
	isUpdated := filemanagement.UpdateFile(dir, id, content)
	return isUpdated
}

func DeleteActivity(user string, id string) bool {
	var success = true
	//TODO delete from cache
	//activity := GetActivity(user, id)
	dir := resources.GetUserDir(user)
	jsonFile := id + ".json"
	originalFile := id + ".zip"
	success = filemanagement.DeleteFile(dir, jsonFile)
	success = filemanagement.DeleteFile(dir, originalFile)
	return success
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