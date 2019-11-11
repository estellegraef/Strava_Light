package overview

import (
	"../../cmd"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	data := cmd.GetActivities()

	t, err := template.ParseFiles("views/templates/html/index.html")

	if err != nil {
		//Add Logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}
