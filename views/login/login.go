package login

import (
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/templates/html/login.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("user")
	password := r.FormValue("password")

	r.SetBasicAuth(username, password)
	user, pwd, _ := r.BasicAuth()

	tmpl.Execute(w, struct {
		Success bool
		User    string
		Pwd     string
	}{true, user, pwd})
}
