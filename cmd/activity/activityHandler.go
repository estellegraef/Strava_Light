/*
 * 2848869
 * 8089098
 * 3861852
 */

package activity

import (
	"encoding/json"
	"fmt"
	"github.com/estellegraef/Strava_Light/cmd/gpxProcessing"
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
var cache Cache
//TODO implement cache after tests

func Setup(){
	cache = NewCache()
}

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

func AddActivity(username string, sportType string, file multipart.File, header *multipart.FileHeader, comment string) bool {
	var success = false
	content := filemanagement.ReadReceiveFile(file)
	filename := header.Filename
	baseIsCreated, createdFile := filemanagement.CreateFile(resources.GetUserDir(username), filename, content)
	if baseIsCreated {
		gpxFiles := gpxProcessing.ReadFile(createdFile)
		for _, file := range gpxFiles {
			id := filemanagement.GenerateId()
			activity := New(string(id), id, sportType, comment, file.GetDistanceInKilometers(), file.GetWaitingTime(), file.GetAvgSpeed(), file.GetMaxSpeed(), file.GetMeta().GetTime())
			content := MarshalJSON(activity)
			//TODO create File with ending json
			jsonTitle := filename + ".json"
			activityIsCreated, _ := filemanagement.CreateFile(resources.GetUserDir(username), jsonTitle, content)
			if activityIsCreated {
				cache.Check(activity)
				success = true
			}
		}
	}
	return success
}

func UpdateActivity(user string, id uint32, sportType string, comment string) bool {
	activity := GetActivity(user, id)
	activity.SportType = sportType
	activity.Comment = comment
	content := MarshalJSON(activity)
	dir := resources.GetUserDir(user)
	isUpdated := filemanagement.UpdateFile(dir, GetNameById(id), content)
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

func GetNameById(id uint32) string {
	return fmt.Sprint(id) + ".json"
}

func GetIdByName(name string) uint32 {
	filename := strings.TrimSuffix(filepath.Base(name), ".json")
	val, _ := strconv.Atoi(filename)
	return uint32(val)
}