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
	id := r.URL.Query().Get("id")

	var data = struct {
		Page    pages.Page
		Content struct {
			IsDelete bool
			Activity activity.Activity
		}
	}{}

	data.Content.IsDelete = r.Method == http.MethodPost
	data.Content.Activity = activity.GetActivity(username, id)
	data.Page = pages.NewDetail(data.Content.Activity.GetSportType())

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("Template execution failed! \n", err)
	}
}
