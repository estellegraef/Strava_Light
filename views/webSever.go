package views

import (
	"./detail"
	"./overview"
	"./upload"
	"fmt"
	"net/http"
	"strings"
)

func CreateWebServer() {
	//Outsource to backend
	http.HandleFunc("/", basicAuth(overview.NewHandler))
	http.HandleFunc("/upload/", basicAuth(upload.NewHandler))
	http.HandleFunc("/detail/", basicAuth(detail.NewHandler))
	http.Handle("/styles/", http.StripPrefix(strings.TrimRight("/styles/", "/"), http.FileServer(http.Dir("views/templates/styles"))))
	http.Handle("/images/", http.StripPrefix(strings.TrimRight("/images/", "/"), http.FileServer(http.Dir("resources/img"))))
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
