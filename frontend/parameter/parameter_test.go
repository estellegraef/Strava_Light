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

func createRequest(t *testing.T, username string) (req *http.Request, server *httptest.Server) {
	server = createServer()

	request, err := http.NewRequest("GET", server.URL, nil)
	assert.NoError(t, err)

	ctx := context.WithValue(request.Context(), "username", username)
	req = request.WithContext(ctx)
	return req, server
}

func executeRequest(t *testing.T, r *http.Request) *http.Response {
	client := http.DefaultClient
	response, err := client.Do(r)
	assert.NoError(t, err)

	return response
}

func TestGetUserWithUsername(t *testing.T) {
	expectedUser := "user"

	request, server := createRequest(t, expectedUser)
	response := executeRequest(t, request)
	defer server.Close()

	actualUser := GetUser(response.Request)
	assert.Equal(t, expectedUser, actualUser)
}

func TestGetUserWithEmptyUsername(t *testing.T) {
	expectedUser := "user"

	request, server := createRequest(t, "")
	response := executeRequest(t, request)
	defer server.Close()

	actualUser := GetUser(response.Request)
	assert.NotEqual(t, expectedUser, actualUser)
}

func TestGetUserAndIDWithID(t *testing.T) {
	expectedUser := "user"
	expectedID := "1"

	request, server := createRequest(t, expectedUser)
	url := request.URL.Query()
	url.Add("id", expectedID)
	request.URL.RawQuery = url.Encode()
	response := executeRequest(t, request)
	defer server.Close()

	actualUser, actualID := GetUserAndID(response.Request)
	assert.Equal(t, expectedUser, actualUser)
	assert.Equal(t, expectedID, actualID)
}

func TestGetUserAndIDWithEmptyID(t *testing.T) {
	expectedUser := "user"
	expectedID := "1"

	request, server := createRequest(t, expectedUser)
	url := request.URL.Query()
	url.Add("id", "")
	request.URL.RawQuery = url.Encode()
	defer server.Close()

	response := executeRequest(t, request)

	actualUser, actualID := GetUserAndID(response.Request)
	assert.Equal(t, expectedUser, actualUser)
	assert.NotEqual(t, expectedID, actualID)
}

func TestGetUserAndIDWithoutID(t *testing.T) {
	expectedUser := "user"
	expectedID := "1"

	request, server := createRequest(t, expectedUser)
	response := executeRequest(t, request)
	defer server.Close()

	actualUser, actualID := GetUserAndID(response.Request)
	assert.Equal(t, expectedUser, actualUser)
	assert.NotEqual(t, expectedID, actualID)
}
