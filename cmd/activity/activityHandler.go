package activity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
)

func AddActivity(username string, sportType string, file multipart.File, fileHeader *multipart.FileHeader, comment string) bool {
	var success = true
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil{
		log.Print(err)
	}
	fmt.Println(fileBytes)
	fmt.Println(fileHeader.Filename)
	//calc bytes
	//obj
	//write obj
	//return success
	return success
}

/*func CreateActivity(){

}

//Get activity from json file
func GetActivity(user string, id string) Activity {
	userDir := resources.GetUserDir(user)
	files := filemanagement.GetAllFiles(userDir)
	var activity Activity
	for _, file := range files {
		if file == id {
			content := filemanagement.ReadFileContent(file)
			activity = UnmarshalJSON(content)
		}
	}
	return activity
}*/

func GetActivitiesFromUser(user string) []Activity {
	/*userDir := resources.GetUserDir(user)
	files := filemanagement.GetAllFiles(userDir)
	var activities []Activity
	for _, file := range files {
		activities = append(activities, GetActivity(user, file))
	}*/
	return []Activity{}
}

func GetActivitiesByKeyword(user string, keyword string) []Activity {
	return []Activity{}
}

func EditActivity(user string, id uint32, sportType string, comment string) bool {
	activity := GetActivity(user, id)
	activity.sportType = sportType
	activity.Comment = comment
	//save
	return true
}

func DeleteActivity(user string, id string) bool {
	//activity := GetActivity(user, id)
	//delete
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