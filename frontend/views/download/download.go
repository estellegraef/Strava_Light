/*
 * 2848869
 * 8089098
 * 3861852
 */

package download

import (
	"bytes"
	"github.com/estellegraef/Strava_Light/backend/activity"
	"net/http"
	"strconv"
	"time"
)

func NewHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		username = "unknown"
	}

	id := r.URL.Query().Get("id")
	downloadBytes, fileName := activity.GetFile(username, id)

	// set the default MIME type to send
	mime := http.DetectContentType(downloadBytes)
	fileSize := len(string(downloadBytes))

	// Generate the server headers
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(downloadBytes))
}
