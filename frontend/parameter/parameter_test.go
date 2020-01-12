/*
 * 2848869
 * 8089098
 * 3861852
 */

package parameter

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, r, "Test Handler")
	}))
}

func TestGetUserWithUsername(t *testing.T) {
	expectedUser := "user"
	server := createServer()
	defer server.Close()

	request, err := http.NewRequest("GET", server.URL, nil)
	assert.NoError(t, err)

	ctx := context.WithValue(request.Context(), "username", expectedUser)
	request = request.WithContext(ctx)
	client := http.DefaultClient
	response, err := client.Do(request)
	assert.NoError(t, err)

	actualUser := GetUser(response.Request)
	assert.Equal(t, expectedUser, actualUser)
}
