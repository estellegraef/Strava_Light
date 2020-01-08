package detail

import (
	"../../../cmd"
	"../../templates/pages"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/detail.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		username = "unknown"
	}

	var data = struct {
		Page    pages.Page
		Content cmd.Activity
	}{}

	urlValue := r.URL.Query().Get("id")
	id, _ := strconv.ParseUint(urlValue, 32, 32)
	data.Content = cmd.GetActivity(username, uint32(id))
	data.Page = pages.NewDetail(data.Content.GetSportType())

	_ = tmpl.Execute(w, data)
}
