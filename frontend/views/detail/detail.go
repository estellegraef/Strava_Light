/*
 * 2848869
 * 8089098
 * 3861852
 */

package detail

import (
	"../../../cmd/activity"
	"../../templates/pages"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/detail.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Page    pages.Page
		Content activity.Activity
	}{}

	urlValue := r.URL.Query().Get("id")
	id, _ := strconv.ParseUint(urlValue, 32, 32)
	data.Content = activity.GetActivity(uint32(id))
	data.Page = pages.NewDetail(data.Content.GetSportType())

	_ = tmpl.Execute(w, data)
}
