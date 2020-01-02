package searchPart

import (
	"../../cmd"
	"../pages"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"views/templates/html/layout.html",
	"views/templates/html/search.html",
	"views/templates/html/items.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Page    pages.Page
		Content []cmd.Activity
	}{
		Page: pages.NewSearch(),
	}

	if r.Method != http.MethodPost {
		data.Content = nil
	} else {
		search := r.FormValue("search")
		data.Content = cmd.SearchActivities(search)
	}

	_ = tmpl.Execute(w, data)
}
