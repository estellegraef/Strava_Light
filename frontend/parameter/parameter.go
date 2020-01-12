/*
 * 2848869
 * 8089098
 * 3861852
 */

package parameter

import (
	"net/http"
)

func GetUserAndID(r *http.Request) (user string, id string) {
	user = GetUser(r)
	id = r.URL.Query().Get("id")

	return user, id
}

func GetUser(r *http.Request) string {
	user, ok := r.Context().Value("username").(string)
	if !ok {
		user = "unknown"
	}

	return user
}
