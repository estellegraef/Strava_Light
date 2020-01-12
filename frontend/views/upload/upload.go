/*
 * 2848869
 * 8089098
 * 3861852
 */

package upload

import (
	"github.com/estellegraef/Strava_Light/backend/activity"
	"github.com/estellegraef/Strava_Light/frontend/parameter"
	"github.com/estellegraef/Strava_Light/frontend/templates/html"
	"github.com/estellegraef/Strava_Light/frontend/templates/pages"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	html.GetLayoutPath(),
	html.GetUploadPath()))

//Upload Handler
//Provides user with an form to upload files and add them to their activities
func NewHandler(w http.ResponseWriter, r *http.Request) {
	username := parameter.GetUser(r)

	var data = struct {
		Page    pages.Page
		Content uint8
	}{
		Page: pages.NewUpload(),
	}

	//if request method isn't post, then present a clear form
	//else read form parameter and call AddActivity to create the Activity
	//show message afterwards if the action was successfull or not
	if r.Method != http.MethodPost {
		data.Content = 0
	} else {
		sportType := r.FormValue("sportType")
		file, fileHeader, _ := r.FormFile("file")
		comment := r.FormValue("comment")
		//backend call
		success := activity.AddActivity(username, sportType, file, fileHeader, comment)
		if success {
			data.Content = 1
		} else {
			data.Content = 2
		}
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("Template execution failed! \n", err)
	}
}
