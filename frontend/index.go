package frontend

import (
	"./detail"
	"./upload"
	"fmt"
	"net/http"
)

func CreateWebServer() {
	//Outsource to backend
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/upload/", upload.NewHandler)
	http.HandleFunc("/detail/", detail.NewHandler)
	http.ListenAndServe(":8080", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am an index page - func overview comming soon!")
}
