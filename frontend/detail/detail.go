package detail

import (
	"html/template"
	"net/http"
)

type ActivityDetail struct {
	Type    string
	Comment string
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	detail := ActivityDetail{Type: "Radfahren", Comment: "I am a useful comment"}
	t, err := template.ParseFiles("frontend/detail/detail.html")

	if err != nil {
		//Add Logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, detail)
}
