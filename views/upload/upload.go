package upload

import (
	"../../cmd"
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
	file, fileHeader, _ := r.FormFile("file")
	comment := r.FormValue("comment")

	//backend call
	success := cmd.CreateActivity(sportType, file, fileHeader, comment)

	tmpl.Execute(w, struct{ Success bool }{success})
}
