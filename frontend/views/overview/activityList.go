package overview

import (
<<<<<<< HEAD:views/overview/activityList.go
	"../../cmd"
	"../pages"
	cmd2 "Strava_Light/cmd"
=======
	"../../../cmd"
	"../../templates/pages"
>>>>>>> master:frontend/views/overview/activityList.go
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/index.html",
	"frontend/templates/html/items.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		username = "unknown"
	}

	data := struct {
		Page    pages.Page
		Content []cmd2.Activity
	}{
		Page:    pages.NewIndex(),
		Content: cmd.GetActivities(username),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		_ = fmt.Errorf("Template execution failed! \n %w", err)
	}
}
