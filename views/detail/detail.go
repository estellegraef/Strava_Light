package detail

import (
	"../../cmd"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/templates/html/detail.html")

	if err != nil {
		//Add Logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := cmd.GetActivity()
	t.Execute(w, data)
}
