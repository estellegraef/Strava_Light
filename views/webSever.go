package views

import (
	"../cmd"
	"./detail"
	"./overview"
	"./searchPart"
	"./upload"
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func CreateWebServer() {
	//Outsource to backend
	http.HandleFunc("/", basicAuth(overview.NewHandler))
	http.HandleFunc("/upload/", basicAuth(upload.NewHandler))
	http.HandleFunc("/detail/", basicAuth(detail.NewHandler))
	http.HandleFunc("/search/", basicAuth(searchPart.NewHandler))
	http.Handle("/assets/", http.StripPrefix(strings.TrimRight("/assets/", "/"), http.FileServer(http.Dir("views/templates/assets"))))
	http.Handle("/images/", http.StripPrefix(strings.TrimRight("/images/", "/"), http.FileServer(http.Dir("resources/img"))))

	// Command-line-flag
	// the default value is 8081
	portPtr := flag.Int("port", 8080, "Webserver Port")
	flag.Parse()
	fmt.Println("Start Server on Port: ", *portPtr)

	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(*portPtr), nil))
}

func basicAuth(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pwd, ok := r.BasicAuth()
		isValid := authenticate(user, pwd)

		if !ok || !isValid {
			w.Header().Add("WWW-Authenticate", "Basic Realm=\"Strava Login\"")
			w.WriteHeader(401)
		} else {
			hf(w, r)
		}
	}
}

func authenticate(username, password string) bool {
	//return username == "Rico" && password == "1234"
	//user1: go!Project?2020
	//user2: user2Password
	var users []cmd.User
	//if users ==  {
	users = cmd.GetUserFromFile()
	//}
	for _, user := range users {
		if user.GetUserName() == username {
			passwordDecode, err1 := base64.StdEncoding.DecodeString(user.GetPassword())
			if err1 != nil {
				fmt.Println("Base64 Decoding error", err1)
				return false
			}
			saltDecode, err2 := base64.StdEncoding.DecodeString(user.GetSalt())
			if err2 != nil {
				fmt.Println("Base64 Decoding error", err2)
				return false
			}
			if cmd.Match([]byte(password), passwordDecode, saltDecode) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
