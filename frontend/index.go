package frontend

import (
	"./login"
	"fmt"
	"net/http"
)

func CreateWebServer() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/login/", login.Login_Handler)
	http.ListenAndServe(":8080", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This ia a test!")
}
