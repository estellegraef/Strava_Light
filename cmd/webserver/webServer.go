/*
 * 2848869
 * 8089098
 * 3861852
 */

package webserver

import (
	"context"
	"flag"
	"fmt"
	"github.com/estellegraef/Strava_Light/cmd/auth"
	"github.com/estellegraef/Strava_Light/frontend/views/detail"
	"github.com/estellegraef/Strava_Light/frontend/views/edit"
	"github.com/estellegraef/Strava_Light/frontend/views/overview"
	"github.com/estellegraef/Strava_Light/frontend/views/search"
	"github.com/estellegraef/Strava_Light/frontend/views/upload"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateWebServer() {
	http.HandleFunc("/", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), overview.NewHandler))
	http.HandleFunc("/upload/", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), upload.NewHandler))
	http.HandleFunc("/detail", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), detail.NewHandler))
	http.HandleFunc("/search/", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), search.NewHandler))
	http.HandleFunc("/edit", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), edit.NewHandler))
	http.Handle("/assets/", http.StripPrefix(strings.TrimRight("/assets/", "/"), http.FileServer(http.Dir("frontend/templates/assets"))))
	http.Handle("/images/", http.StripPrefix(strings.TrimRight("/images/", "/"), http.FileServer(http.Dir("resources/img"))))

	// Command-line-flag for port
	// the default value is 443
	portPtr := flag.Int("port", 443, "Webserver Port")
	flag.Parse()
	fmt.Println("Start Server on Port: ", *portPtr)

	//fmt.Println(http.ListenAndServe(":"+strconv.Itoa(*portPtr), nil))
	log.Fatalln(http.ListenAndServeTLS(":"+strconv.Itoa(*portPtr), "./resources/cert.pem", "./resources/key.pem", nil))
}

func basicAuth(authenticator auth.Authenticator, hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pwd, ok := r.BasicAuth()
		isValid := authenticator.Authenticate(user, pwd)

		if !ok || !isValid {
			w.Header().Add("WWW-Authenticate", "Basic Realm=\"Strava Login\"")
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			ctx := context.WithValue(r.Context(), "username", user)
			hf(w, r.WithContext(ctx))
		}
	}
}
