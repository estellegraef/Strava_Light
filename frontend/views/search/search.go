/*
 * 2848869
 * 8089098
 * 3861852
 */

package search

import (
	"github.com/estellegraef/Strava_Light/backend/activity"
	"github.com/estellegraef/Strava_Light/frontend/parameter"
	"github.com/estellegraef/Strava_Light/frontend/templates/html"
	"github.com/estellegraef/Strava_Light/frontend/templates/pages"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	html.GetLayoutPath(),
	html.GetSearchPath(),
	html.GetItemsPath()))

//Search Handler
//User can enter a search term and gets the results displayed
func NewHandler(w http.ResponseWriter, r *http.Request) {
	username := parameter.GetUser(r)

	var data = struct {
		Page    pages.Page
		Content []activity.Activity
	}{
		Page: pages.NewSearch(),
	}

	//if request method is post, the activities with the given search term are queried and displayed afterwatds
	//else the user gets the search bar and no results
	if r.Method != http.MethodPost {
		data.Content = nil
	} else {
		search := r.FormValue("search")
		data.Content = activity.SearchActivities(username, search)
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println("Template execution failed! \n", err)
	}
}
