/*
 * 2848869
 * 8089098
 * 3861852
 */

package detail

import (
	"github.com/estellegraef/Strava_Light/backend/activity"
	"github.com/estellegraef/Strava_Light/frontend/templates/pages"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/detail.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		username = "unknown"
	}

	var data = struct {
		Page    pages.Page
		Content activity.Activity
	}{}

	id := r.URL.Query().Get("id")
	data.Content = activity.GetActivity(username, id)
	data.Page = pages.NewDetail(data.Content.GetSportType())

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("Template execution failed! \n %w", err)
	}
}
