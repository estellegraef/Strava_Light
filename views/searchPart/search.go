package searchPart

import (
	"../../cmd"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/templates/html/search.html"))

	if r.Method != http.MethodPost {
		_ = temp.Execute(w, nil)
		return
	}

	search := r.FormValue("search")
	results := cmd.SearchActivities(search)

	data := struct {
		HasResults bool
		Results    []cmd.Activity
	}{
		HasResults: len(results) != 0,
		Results:    results,
	}

	_ = temp.Execute(w, data)
}
