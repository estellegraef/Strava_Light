package activityList

import (
	"../../models/activity"
	"html/template"
	"net/http"
	"time"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	list := make([]activity.ActivityListItem, 10)

	for i := 0; i < 10; i++ {
		list[i] = activity.NewListItem(uint32(i), "Radfahren", time.Now())
	}

	t, err := template.ParseFiles("views/templates/html/index.html")

	if err != nil {
		//Add Logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, list)
}
