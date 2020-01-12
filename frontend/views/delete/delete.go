/*
 * 2848869
 * 8089098
 * 3861852
 */

package delete

import (
	"github.com/estellegraef/Strava_Light/backend/activity"
	"github.com/estellegraef/Strava_Light/frontend/parameter"
	"github.com/estellegraef/Strava_Light/frontend/views/overview"
	"net/http"
)

//Delete Handler
//Calls Delete Function from backend with given id (from url)
func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, id := parameter.GetUserAndID(r)
	activity.DeleteActivity(username, id)

	//Redirect to overview
	overview.NewHandler(w, r)
}
