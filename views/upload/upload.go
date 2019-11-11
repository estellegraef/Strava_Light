package upload

import (
	"../../cmd"
	"fmt"
	"html/template"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/templates/html/upload.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	sportType := r.FormValue("sportType")
	file, _, _ := r.FormFile("file")
	comment := r.FormValue("comment")

	fmt.Println(sportType)
	fmt.Println(file)
	fmt.Println(comment)
	//backend call
	success := cmd.CreateActivity(sportType, file, comment)

	tmpl.Execute(w, struct{ Success bool }{success})
}
