/*
 * 2848869
 * 8089098
 * 3861852
 */

package upload

import (
	"github.com/estellegraef/Strava_Light/cmd/activity"
	"github.com/estellegraef/Strava_Light/frontend/templates/pages"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/upload.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		username = "unknown"
	}

	var data = struct {
		Page    pages.Page
		Content uint8
	}{
		Page: pages.NewUpload(),
	}

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

	_ = tmpl.Execute(w, data)
}
