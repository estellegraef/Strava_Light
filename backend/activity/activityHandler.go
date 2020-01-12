/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"encoding/json"
	"github.com/estellegraef/Strava_Light/backend/gpxProcessing"
	"github.com/estellegraef/Strava_Light/backend/storageManagement"
	"github.com/estellegraef/Strava_Light/resources"
	"log"
	"mime/multipart"
	"path/filepath"
	"sort"
	"strings"
)

var cache Cache

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
			cache.Check(activity.Id, activity)
		}
	}
	return SortActivities(activities)
}

func SortActivities(activities []Activity) []Activity{
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].DateTime.After(activities[j].DateTime)
	})
	return activities
}

func GetActivity(user string, id string) Activity {
	var activity Activity
	inCache, cachedActivity := cache.GetActivity(id)
	if inCache {
		activity = cachedActivity
	} else {
		userDir := resources.GetUserDir(user)
		files := filemanagement.GetAllFilesFromDir(userDir)

		for _, file := range files {
			searchedFile := id + ".json"
			if filepath.Base(file) == searchedFile {
				content, _ := filemanagement.ReadFile(file)
				activity = UnmarshalJSON(content)
			}
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
			jsonTitle := id + ".json"
			activityIsCreated, _ := filemanagement.CreateFile(resources.GetUserDir(userName), jsonTitle, content)
			if activityIsCreated {
				cache.Check(activity.Id, activity)
				success = true
			}
		}
	}
	return success
}

func UpdateActivity(user string, id string, sportType string, comment string) bool {
	activity := GetActivity(user, id)
	cache.RemoveById(id)
	activity.SportType = sportType
	activity.Comment = comment
	content := MarshalJSON(activity)
	dir := resources.GetUserDir(user)
	isUpdated := filemanagement.UpdateFile(dir, id + ".json", content)
	cache.Check(id, activity)
	return isUpdated
}

func DeleteActivity(user string, id string) bool {
	var success = false
	cache.RemoveById(id)
	dir := resources.GetUserDir(user)
	files := filemanagement.GetAllFilesFromDir(dir)
	originalName := filemanagement.GetOriginal(id)
	for _, file := range files {
		if strings.Contains(file, originalName){
			success = filemanagement.DeleteFile(dir, filepath.Base(file))
		}
	}
	return success
}

func ReturnFileForDownload(userName string, id string) (content []byte, name string){
	dir := resources.GetUserDir(userName)
	originalName := filemanagement.GetOriginal(id)
	var fileContent []byte
	var fileBase string

	files := filemanagement.GetAllFilesFromDir(dir)
	for _, file := range files {
		if strings.Contains(file, originalName) && filepath.Ext(file) != ".json" {
			fileContent, fileBase = filemanagement.ReadFile(file)
		}
	}
	return fileContent, fileBase
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