package views

import (
	"./detail"
	"./overview"
	"./upload"
	"net/http"
)

func CreateWebServer() {
	//Outsource to backend
	http.HandleFunc("/", overview.NewHandler)
	http.HandleFunc("/upload/", upload.NewHandler)
	http.HandleFunc("/detail/", detail.NewHandler)
	http.ListenAndServe(":8080", nil)
}
