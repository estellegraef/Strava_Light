/*
 * 2848869
 * 8089098
 * 3861852
 */

package download

import (
	"bytes"
	"github.com/estellegraef/Strava_Light/backend/activity"
	"github.com/estellegraef/Strava_Light/frontend/parameter"
	"net/http"
	"strconv"
	"time"
)

//Download Handler
//Provides the file bytes for downloading
func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, id := parameter.GetUserAndID(r)

	//Gets File bytes and name by id and detected file type
	downloadBytes, fileName := activity.ReturnFileForDownload(username, id)
	mime := http.DetectContentType(downloadBytes)
	fileSize := len(string(downloadBytes))

	//Sets Sever Header to trigger browser download dialog automatically
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	//Writes Bytes to Client
	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(downloadBytes))
}
