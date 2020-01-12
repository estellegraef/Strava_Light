/*
 * 2848869
 * 8089098
 * 3861852
 */

package edit

import (
	"github.com/estellegraef/Strava_Light/backend/activity"
	"github.com/estellegraef/Strava_Light/frontend/parameter"
	"github.com/estellegraef/Strava_Light/frontend/templates/pages"
	"github.com/estellegraef/Strava_Light/frontend/views/detail"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/edit.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, id := parameter.GetUserAndID(r)

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

	if r.Method == http.MethodPost {
		sportType := r.FormValue("sportType")
		comment := r.FormValue("comment")
		activity.UpdateActivity(username, id, sportType, comment)

		r.Method = http.MethodGet
		detail.NewHandler(w, r)
	} else {
		var act = activity.GetActivity(username, id)
		data.Content.ID = act.GetID()
		data.Content.IsWalking = act.GetSportType() == "Laufen"
		data.Content.IsBiking = !data.Content.IsWalking
		data.Content.Comment = act.GetComment()

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			log.Println("Template execution failed! \n", err)
		}
	}
}
