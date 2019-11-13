package overview

import (
	"../../cmd"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/templates/html/index.html"))

	var data []cmd.Activity
	if r.Method != http.MethodPost {
		data = cmd.GetActivities()
	} else {
		search := r.FormValue("search")
		data = cmd.SearchActivities(search)
	}

	tmpl.Execute(w, data)
}
