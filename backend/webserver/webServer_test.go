/*
 * 2848869
 * 8089098
 * 3861852
 */

package webserver

import (
	"crypto/tls"
	"fmt"
	"github.com/estellegraef/Strava_Light/backend/auth"
	"github.com/estellegraef/Strava_Light/backend/user"
	"github.com/estellegraef/Strava_Light/resources"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var (
	httpsPort = "8081"
)

func createServer(auth auth.Authenticator) *httptest.Server {
	return httptest.NewServer(basicAuth(auth, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello client")
	}))
}

func doRequestWithPassword(t *testing.T, url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	req.SetBasicAuth("<username>", "<password>")
	res, err := client.Do(req)
	assert.NoError(t, err)
	return res
}

// Test without Password
func TestWithoutPW(t *testing.T) {
	ts := createServer(auth.AuthenticatorFunc(func(name, pwd string) bool { return true }))
	defer ts.Close()
	res, err := http.Get(ts.URL)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "wrong status code")

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusText(http.StatusUnauthorized)+"\n", string(body), "wrong message")
}

// Test with wrong password
func TestWithWrongPW(t *testing.T) {
	var receivedName, receivedPwd string
	ts := createServer(auth.AuthenticatorFunc(func(name, pwd string) bool {
		receivedName = name
		receivedPwd = pwd
		return false // <--- deny every request
	}))
	defer ts.Close()

	res := doRequestWithPassword(t, ts.URL)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "wrong status")
	assert.Equal(t, "<username>", receivedName, "wrong username")
	assert.Equal(t, "<password>", receivedPwd, "wrong password")

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t,
		http.StatusText(http.StatusUnauthorized)+"\n",
		string(body), "wrong message")
}

// Test with correct password
func TestWithCorrectPW(t *testing.T) {
	var receivedName, receivedPwd string
	ts := createServer(auth.AuthenticatorFunc(func(name, pwd string) bool {
		receivedName = name
		receivedPwd = pwd
		return true // <--- accept every request
	}))
	defer ts.Close()

	res := doRequestWithPassword(t, ts.URL)

	assert.Equal(t, http.StatusOK, res.StatusCode, "wrong status code")
	assert.Equal(t, "<username>", receivedName, "wrong username")
	assert.Equal(t, "<password>", receivedPwd, "wrong password")

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello client\n", string(body), "wrong message")
}

// inspired by: https://blog.dnsimple.com/2017/08/how-to-test-golang-https-services/
func NewServer(port string) *http.Server {
	addr := fmt.Sprintf(":%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func buildSecureUrl(path string) string {
	return urlFor("https", httpsPort, path)
}

func urlFor(scheme string, serverPort string, path string) string {
	return scheme + "://localhost:" + serverPort + path
}

// test the https-Cert and private Key
func TestHTTPSServer(t *testing.T) {
	srv := NewServer(httpsPort)
	go srv.ListenAndServeTLS(resources.GetCertPath(), resources.GetKeyPath())
	defer srv.Close()
	time.Sleep(100 * time.Millisecond)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Get(buildSecureUrl("/"))

	assert.NoError(t, err)

	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode, "Response StatusCode is not 200")

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	expected := []byte("Hello World")

	assert.Equal(t, expected, body)
}

// test the function CheckAndHandleStoragePath with a path that does not exist
func TestCheckAndHandleStoragePathWithNonExistPath(t *testing.T) {
	defaultDir, err := os.Getwd()
	assert.NoError(t, err)

	path := filepath.Join(defaultDir, "storage")
	defer os.RemoveAll(path)

	baseDir := filepath.Join(defaultDir, "test1") // <- Path does not exist

	checkAndHandleStoragePath(baseDir, defaultDir)

	expectedPath := filepath.Join(defaultDir, "storage", "user1")

	users := user.GetUsers()

	assert.Equal(t, expectedPath, (*users)[0].GetStoragePath())
}

// test the function CheckAndHandleStoragePath with a path that exists
func TestCheckAndHandleStoragePathWithExistPath(t *testing.T) {
	defaultDir, err := os.Getwd()
	assert.NoError(t, err)
	path := filepath.Join(defaultDir, "storage")
	defer os.RemoveAll(path)

	checkAndHandleStoragePath(defaultDir, defaultDir)

	expectedPath := filepath.Join(defaultDir, "storage", "user1")
	users := user.GetUsers()

	assert.Equal(t, expectedPath, (*users)[0].GetStoragePath())
}
