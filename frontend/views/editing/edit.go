package editing

import (
	"../../../cmd"
	"../../templates/pages"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/edit.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {

	var data = struct {
		Page    pages.Page
		Content struct {
			IsWalking bool
			IsBiking  bool
			Comment   string
		}
	}{
		Page: pages.NewEdit(),
	}

	var activity = cmd.GetActivity()
	data.Content.IsWalking = activity.GetSportType() == "Laufen"
	data.Content.IsBiking = !data.Content.IsWalking
	data.Content.Comment = activity.GetComment()

	_ = tmpl.Execute(w, data)
}
