package detail

import (
	"../../../cmd"
	"../../templates/pages"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/detail.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Page    pages.Page
		Content cmd.Activity
	}{}

	data.Content = cmd.GetActivity()
	data.Page = pages.NewDetail(data.Content.GetSportType())

	_ = tmpl.Execute(w, data)
}
