package overview

import (
	"../../cmd"
	"../pages"
	"Strava_Light/cmd/gpx/activity"
	"fmt"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/templates/html/layout.html", "views/templates/html/index.html", "views/templates/html/items.html"))

	data := struct {
		Page    pages.Page
		Content []activity.Activity
	}{
		Page:    pages.NewIndex(),
		Content: cmd.GetActivities(),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		_ = fmt.Errorf("Template execution failed! \n %w", err)
	}
}
