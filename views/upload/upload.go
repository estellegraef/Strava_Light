package upload

import (
	"../../cmd"
	"../pages"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"views/templates/html/layout.html",
	"views/templates/html/upload.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Page    pages.Page
		Content uint8
	}{
		Page: pages.NewUpload(),
	}

	if r.Method != http.MethodPost {
		data.Content = 0
	} else {
		sportType := r.FormValue("sportType")
		file, fileHeader, _ := r.FormFile("file")
		comment := r.FormValue("comment")

		//backend call
		success := cmd.CreateActivity(sportType, file, fileHeader, comment)

		if success {
			data.Content = 1
		} else {
			data.Content = 2
		}
	}

	_ = tmpl.Execute(w, data)
}
