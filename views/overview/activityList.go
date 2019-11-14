package overview

import (
	"../../cmd"
	"fmt"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/templates/html/index.html"))

	var data []cmd.Activity
	if r.Method == http.MethodPost {
		search := r.FormValue("search")
		data = cmd.SearchActivities(search)
	} else {
		data = cmd.GetActivities()
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		_ = fmt.Errorf("Template execution failed! \n %w", err)
	}
}
