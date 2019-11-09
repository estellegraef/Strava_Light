package login

import (
	"html/template"
	"net/http"
)

type User struct {
	Name     string
	Password string
}

func Login_Handler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Rico", Password: "1234"}
	t, err := template.ParseFiles("frontend/login/login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, user)
}
