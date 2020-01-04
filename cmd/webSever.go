/*
 * 2848869
 * 8089098
 * 3861852
 */

package cmd

import (
	"../frontend/views/detail"
	"../frontend/views/editing"
	"../frontend/views/overview"
	"../frontend/views/searching"
	"../frontend/views/upload"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateWebServer() {
	http.HandleFunc("/", basicAuth(AuthenticatorFunc(CheckUserIsValid), overview.NewHandler))
	http.HandleFunc("/upload/", basicAuth(AuthenticatorFunc(CheckUserIsValid), upload.NewHandler))
	http.HandleFunc("/detail", basicAuth(AuthenticatorFunc(CheckUserIsValid), detail.NewHandler))
	http.HandleFunc("/search/", basicAuth(AuthenticatorFunc(CheckUserIsValid), searching.NewHandler))
	http.HandleFunc("/edit", basicAuth(AuthenticatorFunc(CheckUserIsValid), editing.NewHandler))
	http.Handle("/assets/", http.StripPrefix(strings.TrimRight("/assets/", "/"), http.FileServer(http.Dir("frontend/templates/assets"))))
	http.Handle("/images/", http.StripPrefix(strings.TrimRight("/images/", "/"), http.FileServer(http.Dir("resources/img"))))

	// Command-line-flag
	// the default value is 443
	portPtr := flag.Int("port", 443, "Webserver Port")
	flag.Parse()
	fmt.Println("Start Server on Port: ", *portPtr)

	//fmt.Println(http.ListenAndServe(":"+strconv.Itoa(*portPtr), nil))
	log.Fatalln(http.ListenAndServeTLS(":"+strconv.Itoa(*portPtr), "./resources/cert.pem", "./resources/key.pem", nil))
}

func basicAuth(authenticator Authenticator, hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pwd, ok := r.BasicAuth()
		//isValid := CheckUserIsValid(user, pwd)
		isValid := authenticator.Authenticate(user, pwd)

		if !ok || !isValid {
			w.Header().Add("WWW-Authenticate", "Basic Realm=\"Strava Login\"")
			w.WriteHeader(401)
		} else {
			ctx := context.WithValue(r.Context(), "username", user)
			hf(w, r.WithContext(ctx))
		}
	}
}
