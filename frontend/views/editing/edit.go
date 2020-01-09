/*
 * 2848869
 * 8089098
 * 3861852
 */

package editing

import (
	"github.com/estellegraef/Strava_Light/cmd/activity"
	"github.com/estellegraef/Strava_Light/frontend/templates/pages"
	"html/template"
	"net/http"
	"strconv"
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
			ID        uint32
			IsWalking bool
			IsBiking  bool
			Comment   string
		}
	}{
		Page: pages.NewEdit(),
	}

	urlValue := r.URL.Query().Get("id")
	id, _ := strconv.ParseUint(urlValue, 32, 32)
	var activity = activity.GetActivity(username, uint32(id))
	data.Content.ID = activity.GetID()
	data.Content.IsWalking = activity.GetSportType() == "Laufen"
	data.Content.IsBiking = !data.Content.IsWalking
	data.Content.Comment = activity.GetComment()

	_ = tmpl.Execute(w, data)
}
