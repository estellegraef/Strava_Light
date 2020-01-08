package searchPart

import (
	"../../cmd"
	"../pages"
	cmd2 "Strava_Light/cmd"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/templates/html/layout.html", "views/templates/html/search.html", "views/templates/html/items.html"))

	var data = struct {
		Page    pages.Page
		Content []cmd2.Activity
	}{
		Page: pages.NewSearch(),
	}

	if r.Method != http.MethodPost {
		data.Content = nil
	} else {
		search := r.FormValue("search")
		data.Content = cmd.SearchActivities(search)
	}

	_ = temp.Execute(w, data)
}
