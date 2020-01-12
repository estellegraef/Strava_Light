/*
 * 2848869
 * 8089098
 * 3861852
 */

package download

import (
	"github.com/estellegraef/Strava_Light/cmd/activity"
	"github.com/estellegraef/Strava_Light/frontend/templates/pages"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/download.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		username = "unknown"
	}

	var data = struct {
		Page pages.Page
	}{
		Page: pages.NewDownload(),
	}

	filePath := activity.GetFile(username, r.URL.Query().Get("id"))

	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	w.Header().Set("Content-Type", "application/gpx")

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("Template execution failed! \n %w", err)
	}
}
