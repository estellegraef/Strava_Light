package detail

import (
	"html/template"
	"net/http"
	"time"
)

type ActivityDetail struct {
	Id          uint32
	Type        string
	Comment     string
	Length      float64
	WaitingTime float64
	AvgSpeed    float64
	MaxSpeed    float64
	DateTime    time.Time
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	detail := ActivityDetail{
		Id:          1,
		Type:        "Radfahren",
		Comment:     "I am a useful comment",
		Length:      12.3,
		WaitingTime: 5.2,
		AvgSpeed:    13.4,
		MaxSpeed:    17.8,
		DateTime:    time.Now(),
	}

	t, err := template.ParseFiles("frontend/detail/detail.html")

	if err != nil {
		//Add Logging
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, detail)
}
