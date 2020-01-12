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

//receive all activities from a user by its name
func GetActivities(userName string) []Activity {
	userDir := resources.GetUserDir(userName)
	files := filemanagement.GetAllFilesFromDir(userDir)
	var activities []Activity
	for _, fileName := range files {
		//only read .json files, since they contain the activity objects
		if filepath.Ext(fileName) == ".json" {
			title := strings.TrimSuffix(filepath.Base(fileName), ".json")
			//search activity and store in list
			var activity = GetActivity(userName, title)
			activities = append(activities, activity)
			//push activity to cache
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

//get a user activity by username and activity id
func GetActivity(user string, id string) Activity {
	var activity Activity
	//first, verify if activity is already in cache
	inCache, cachedActivity := cache.GetActivity(id)
	//if its in cache, retrieve object from cache
	if inCache {
		activity = cachedActivity
		//if activity is not in cache, get the file from the directory and read it
	} else {
		userDir := resources.GetUserDir(user)
		files := filemanagement.GetAllFilesFromDir(userDir)

		for _, file := range files {
			searchedFile := id + ".json"
			if filepath.Base(file) == searchedFile {
				content, _ := filemanagement.ReadFile(file)
				//create an activity object from the content bytes
				activity = UnmarshalJSON(content)
			}
		}
	}
	return activity
}

//get multiple activities based on a keyword, that may match a comment
func SearchActivities(username string, keyword string) []Activity {
	//first get all activities
	userActivities := GetActivities(username)
	var matchingActivities []Activity
	for _, activity := range userActivities{
		//look for each activity if any part of the commment matches the keyword
		if strings.Contains(activity.Comment, keyword) {
			matchingActivities = append(matchingActivities, activity)
		}
	}
	return matchingActivities
}

//add a new activity
func AddActivity(userName string, sportType string, file multipart.File, header *multipart.FileHeader, comment string) bool {
	var success = false
	content := filemanagement.ReadReceiveFile(file)
	fileName := header.Filename
	//create the new file inside the user directory
	baseIsCreated, createdFile := filemanagement.CreateFile(resources.GetUserDir(userName), fileName, content)
	if baseIsCreated {
		//read the contents of the received file and parse its structure to a GPX-Info object
		gpxFiles := gpxProcessing.ReadFile(createdFile)
		for _, file := range gpxFiles {
			//generate a unique id for each gpxFile
			id := filemanagement.GenerateId(fileName)
			//create an activity object
			activity := New(id, sportType, comment, file.GetDistanceInKilometers(), file.GetWaitingTime(), file.GetAvgSpeed(), file.GetMaxSpeed(), file.GetMeta().GetTime())
			//save the activity object
			content := MarshalJSON(activity)
			jsonTitle := id + ".json"
			activityIsCreated, _ := filemanagement.CreateFile(resources.GetUserDir(userName), jsonTitle, content)
			if activityIsCreated {
				//if the activity could be created successfully, put it into the cache
				cache.Check(activity.Id, activity)
				success = true
			}
		}
	}
	return success
}

//update the sporttype and comment of an activiry
func UpdateActivity(user string, id string, sportType string, comment string) bool {
	activity := GetActivity(user, id)
	//remove it from cache
	cache.RemoveById(id)
	//update the values
	activity.SportType = sportType
	activity.Comment = comment
	//put the new value into bytes and overwrite the old value
	content := MarshalJSON(activity)
	dir := resources.GetUserDir(user)
	isUpdated := filemanagement.UpdateFile(dir, id + ".json", content)
	//put the updated activity back into the cache
	cache.Check(id, activity)
	return isUpdated
}

//delete an activity based on its id and the username
func DeleteActivity(user string, id string) bool {
	var success = false
	//remove it from cache
	cache.RemoveById(id)
	dir := resources.GetUserDir(user)
	files := filemanagement.GetAllFilesFromDir(dir)
	originalName := filemanagement.GetOriginal(id)
	for _, file := range files {
		//search for its original file and json file and delete both
		if strings.Contains(file, originalName){
			success = filemanagement.DeleteFile(dir, filepath.Base(file))
		}
	}
	return success
}

//return a file content and its name
func ReturnFileForDownload(userName string, id string) (content []byte, name string){
	dir := resources.GetUserDir(userName)
	originalName := filemanagement.GetOriginal(id)
	var fileContent []byte
	var fileBase string

	files := filemanagement.GetAllFilesFromDir(dir)
	for _, file := range files {
		//only get the original file matching the id
		if strings.Contains(file, originalName) && filepath.Ext(file) != ".json" {
			fileContent, fileBase = filemanagement.ReadFile(file)
		}
	}
	return fileContent, fileBase
}

//marshal an activity to bytes for a json file
func MarshalJSON(activity Activity) []byte {
	jsonData, err := json.MarshalIndent(activity, "", "		")
	if err != nil {
		log.Println(err)
	}
	return jsonData
}

//unmarshal bytes from a json file to an activity
func UnmarshalJSON(data []byte) Activity {
	var activity Activity
	err := json.Unmarshal(data, &activity)
	if err != nil {
		log.Println(err)
	}
	return activity
}