/*
 * 2848869
 * 8089098
 * 3861852
 */

package searching

import (
	"../../../cmd/activity"
	"../../templates/pages"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/search.html",
	"frontend/templates/html/items.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Page    pages.Page
		Content []activity.Activity
	}{
		Page: pages.NewSearch(),
	}

	if r.Method != http.MethodPost {
		data.Content = nil
	} else {
		search := r.FormValue("search")
		data.Content = activity.SearchActivities(search)
	}

	_ = tmpl.Execute(w, data)
}
