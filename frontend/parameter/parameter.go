/*
 * 2848869
 * 8089098
 * 3861852
 */

package parameter

import (
	"net/http"
)

//Reads username and ID from request (context and url)
func GetUserAndID(r *http.Request) (user string, id string) {
	user = GetUser(r)
	id = r.URL.Query().Get("id")

	return user, id
}

//Reads username from request context - default = "unknown"
func GetUser(r *http.Request) string {
	user, ok := r.Context().Value("username").(string)
	if !ok {
		user = "unknown"
	}

	return user
}
