/*
 * 2848869
 * 8089098
 * 3861852
 */

package detail

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
	html.GetDetailPath()))

//Detail Handler
//Shows all details for an activity (get by id)
//If called with method post, the delete alert is triggered
func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, id := parameter.GetUserAndID(r)

	var data = struct {
		Page    pages.Page
		Content struct {
			IsDelete bool
			Activity activity.Activity
		}
	}{}

	//if IsDelete == true then trigger delete question alert
	//display activity in both cases
	data.Content.IsDelete = r.Method == http.MethodPost
	data.Content.Activity = activity.GetActivity(username, id)
	data.Page = pages.NewDetail(data.Content.Activity.GetSportType())

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("Template execution failed! \n", err)
	}
}
