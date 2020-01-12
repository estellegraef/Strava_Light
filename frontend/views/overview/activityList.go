/*
 * 2848869
 * 8089098
 * 3861852
 */

package overview

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
	html.GetLayoutPath(), html.GetIndexPath(), html.GetItemsPath(),
))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username := parameter.GetUser(r)

	data := struct {
		Page    pages.Page
		Content []activity.Activity
	}{
		Page:    pages.NewIndex(),
		Content: activity.GetActivities(username),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("Template execution failed! \n", err)
	}
}
