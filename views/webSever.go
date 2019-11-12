package views

import (
	"./detail"
	"./login"
	"./overview"
	"./upload"
	"net/http"
)

func CreateWebServer() {
	//Outsource to backend
	http.HandleFunc("/", login.NewHandler)
	http.HandleFunc("/overview", basicAuth(overview.NewHandler))
	http.HandleFunc("/upload/", upload.NewHandler)
	http.HandleFunc("/detail/", detail.NewHandler)
	http.ListenAndServe(":8080", nil)
}

func basicAuth(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pwd, ok := r.BasicAuth()

		if !ok || user != "Rico" || pwd != "1234" {
			w.Header().Set("WWW-Authenticate", `Basic-Realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		hf(w, r)
	}
}
