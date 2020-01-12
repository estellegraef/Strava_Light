/*
 * 2848869
 * 8089098
 * 3861852
 */

package edit

import (
	"github.com/estellegraef/Strava_Light/cmd/activity"
	"github.com/estellegraef/Strava_Light/frontend/templates/pages"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/edit.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		username = "unknown"
	}

	var data = struct {
		Page    pages.Page
		Content struct {
			ID        string
			IsWalking bool
			IsBiking  bool
			Comment   string
		}
	}{
		Page: pages.NewEdit(),
	}

	id := r.URL.Query().Get("id")
	var act = activity.GetActivity(username, id)
	data.Content.ID = act.GetID()
	data.Content.IsWalking = act.GetSportType() == "Laufen"
	data.Content.IsBiking = !data.Content.IsWalking
	data.Content.Comment = act.GetComment()

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("Template execution failed! \n %w", err)
	}
}
