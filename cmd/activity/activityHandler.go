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
	var success = true
	content := filemanagement.ReadReceiveFile(file)
	filename := header.Filename
	//TODO read bytes -> save + read with gpxprocessing or readbytes and generate gpx without saving
	gpxFiles := gpxProcessing.GenerateGpx(filename, content)
	for _, file := range gpxFiles{
		id := filemanagement.GenerateId()
		activity := New(string(id), id, sportType, comment, file.GetDistanceInKilometers(), file.GetWaitingTime(), file.GetAvgSpeed(), file.GetMaxSpeed(), file.GetMeta().GetTime())
		content := MarshalJSON(activity)
		success = filemanagement.CreateFile(resources.GetUserDir(username), filename, content)
	}
	return success
}

func UpdateActivity(user string, id uint32, sportType string, comment string) bool {
	activity := GetActivity(user, id)
	activity.SportType = sportType
	activity.Comment = comment
	//save
	return true
}

func DeleteActivity(user string, id string) bool {
	var success = true
	//TODO delete from cache
	//activity := GetActivity(user, id)
	success = DeleteActivity(user, id)
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