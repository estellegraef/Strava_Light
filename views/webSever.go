package views

import (
	"./activityList"
	"./detail"
	"./upload"
	"net/http"
)

func CreateWebServer() {
	//Outsource to backend
	http.HandleFunc("/", activityList.NewHandler)
	http.HandleFunc("/upload/", upload.NewHandler)
	http.HandleFunc("/detail/", detail.NewHandler)
	http.ListenAndServe(":8080", nil)
}
