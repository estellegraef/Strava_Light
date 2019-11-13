package views

import (
	"./detail"
	"./overview"
	"./upload"
	"fmt"
	"net/http"
)

func CreateWebServer() {
	//Outsource to backend
	http.HandleFunc("/", basicAuth(overview.NewHandler))
	http.HandleFunc("/upload/", basicAuth(upload.NewHandler))
	http.HandleFunc("/detail/", basicAuth(detail.NewHandler))
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func basicAuth(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pwd, ok := r.BasicAuth()
		isValid := authenticate(user, pwd)

		if !ok || !isValid {
			w.Header().Add("WWW-Authenticate", "Basic Realm=\"Strava\"")
			w.WriteHeader(401)
		} else {
			w.WriteHeader(200)
		}

		hf(w, r)
	}
}

func authenticate(username, password string) bool {
	return username == "Rico" && password == "1234"
}
