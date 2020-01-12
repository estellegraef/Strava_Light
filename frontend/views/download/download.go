/*
 * 2848869
 * 8089098
 * 3861852
 */

package download

import (
	"github.com/estellegraef/Strava_Light/backend/activity"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		username = "unknown"
	}

	id := r.URL.Query().Get("id")
	filePath := activity.GetFile(username, id)

	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	w.Header().Set("Content-Type", "application/gpx")

}
