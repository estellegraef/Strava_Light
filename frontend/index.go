package frontend

import (
	"fmt"
	"net/http"
)

func CreateWebServer() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8080", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This ia a test!")
}
