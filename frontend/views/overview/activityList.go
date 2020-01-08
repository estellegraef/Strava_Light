package overview

import (
	"../../../cmd"
	"../../templates/pages"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"frontend/templates/html/layout.html",
	"frontend/templates/html/index.html",
	"frontend/templates/html/items.html"))

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)

	if !ok {
		username = "unknown"
	}

	data := struct {
		Page    pages.Page
		Content []cmd.Activity
	}{
		Page:    pages.NewIndex(),
		Content: cmd.GetActivities(username),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		_ = fmt.Errorf("Template execution failed! \n %w", err)
	}
}
