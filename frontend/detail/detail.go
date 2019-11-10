package detail

import (
	"../../models"
	"html/template"
	"net/http"
	"time"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	detail := activityDetail.New(1, "Radfahren", "I am a useful comment",
		12.3, 5.2, 13.4, 17.8, time.Now())

	t, err := template.ParseFiles("frontend/detail/detail.html")

	if err != nil {
		//Add Logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, detail)
}
