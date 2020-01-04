/*
 * 2848869
 * 8089098
 * 3861852
 */

package overview

import (
	"../../../cmd/activity"
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
	data := struct {
		Page    pages.Page
		Content []activity.Activity
	}{
		Page:    pages.NewIndex(),
		Content: activity.GetActivities(),
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		_ = fmt.Errorf("Template execution failed! \n %w", err)
	}
}
