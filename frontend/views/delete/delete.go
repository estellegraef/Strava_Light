/*
 * 2848869
 * 8089098
 * 3861852
 */

package delete

import (
	"github.com/estellegraef/Strava_Light/backend/activity"
	"github.com/estellegraef/Strava_Light/frontend/views/overview"
	"net/http"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		username = "unknown"
	}
	id := r.URL.Query().Get("id")

	activity.DeleteActivity(username, id)
	overview.NewHandler(w, r)
}
