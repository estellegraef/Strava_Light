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
	"github.com/estellegraef/Strava_Light/backend/auth"
	"github.com/estellegraef/Strava_Light/backend/user"
	"github.com/estellegraef/Strava_Light/frontend/views/delete"
	"github.com/estellegraef/Strava_Light/frontend/views/detail"
	"github.com/estellegraef/Strava_Light/frontend/views/download"
	"github.com/estellegraef/Strava_Light/frontend/views/edit"
	"github.com/estellegraef/Strava_Light/frontend/views/overview"
	"github.com/estellegraef/Strava_Light/frontend/views/search"
	"github.com/estellegraef/Strava_Light/frontend/views/upload"
	"github.com/estellegraef/Strava_Light/resources"
	"log"
	"net/http"
	"os"
	osUser "os/user"
	"strconv"
	"strings"
)

var (
	defaultPort = 443
	basePath    string
)

func CreateWebServer() {
	http.HandleFunc("/", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), overview.NewHandler))
	http.HandleFunc("/upload/", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), upload.NewHandler))
	http.HandleFunc("/detail", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), detail.NewHandler))
	http.HandleFunc("/search/", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), search.NewHandler))
	http.HandleFunc("/edit", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), edit.NewHandler))
	http.HandleFunc("/download", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), download.NewHandler))
	http.HandleFunc("/delete", basicAuth(auth.AuthenticatorFunc(auth.CheckUserIsValid), delete.NewHandler))
	http.Handle("/assets/", http.StripPrefix(strings.TrimRight("/assets/", "/"), http.FileServer(http.Dir("frontend/templates/assets"))))
	http.Handle("/images/", http.StripPrefix(strings.TrimRight("/images/", "/"), http.FileServer(http.Dir("resources/img"))))

	// Command-line-flag for port
	// the default value is 443
	portPtr := flag.Int("port", defaultPort, "Webserver Port")

	// Command-line-flag for storage location
	// the default value is the home directory
	u, err := osUser.Current()
	if err != nil {
		log.Println("Can't find current user directory", err)
	}
	defaultDir := u.HomeDir
	baseDir := flag.String("basedir", defaultDir, "base storage location for the web server")
	flag.Parse()

	checkAndHandleStoragePath(*baseDir, defaultDir)
	fmt.Println("Start Server on Port: ", *portPtr)

	// Start web server
	log.Fatalln(http.ListenAndServeTLS(":"+strconv.Itoa(*portPtr), resources.GetCertPath(), resources.GetKeyPath(), nil))
}

// Is responsible for querying the login information and checks the specified login data
func basicAuth(authenticator auth.Authenticator, hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pwd, ok := r.BasicAuth()
		isValid := authenticator.Authenticate(user, pwd)

		if !ok || !isValid {
			w.Header().Set("WWW-Authenticate", "Basic Realm=\"Strava Login\"")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			ctx := context.WithValue(r.Context(), "username", user)
			hf(w, r.WithContext(ctx))
		}
	}
}

// Checks the specified path and generates the user storage path
func checkAndHandleStoragePath(baseDir, defaultDir string) {
	if baseDir != defaultDir {
		if _, err := os.Stat(baseDir); os.IsNotExist(err) {
			log.Printf("The specified path does not exist. The storage location is moved to: %s", defaultDir)
			baseDir = defaultDir
		}
	}
	user.CreateStorageForUsers(baseDir)
}
