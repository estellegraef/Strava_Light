package upload

import (
	"html/template"
	"net/http"
)

type Activity struct {
	File    int
	Type    string
	Comment string
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/upload/upload.html")

	if err != nil {
		//Add Logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}
