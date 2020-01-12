/*
 * 2848869
 * 8089098
 * 3861852
 */

package pages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDetail(t *testing.T) {
	sportType := "Laufen"
	expectedPage := Page{
		title:            "Detail",
		headLine:         sportType,
		customStyleSheet: "../assets/css/detail.css",
		jsFile:           "../assets/js/custom.js"}
	actualPage := NewDetail(sportType)
	assert.Equal(t, expectedPage, actualPage)
}

func TestNewEdit(t *testing.T) {
	expectedPage := Page{
		title:    "Bearbeiten",
		headLine: "Bearbeiten einer Aktivität"}
	actualPage := NewEdit()
	assert.Equal(t, expectedPage, actualPage)
}

func TestNewIndex(t *testing.T) {
	expectedPage := Page{
		title:            "Strava",
		headLine:         "Deine Aktivitäten",
		customStyleSheet: "../assets/css/items.css",
		jsFile:           "../assets/js/custom.js"}
	actualPage := NewIndex()
	assert.Equal(t, expectedPage, actualPage)
}

func TestNewSearch(t *testing.T) {
	expectedPage := Page{
		title:            "Suche",
		headLine:         "Suche",
		customStyleSheet: "../assets/css/items.css",
		jsFile:           "../assets/js/custom.js"}
	actualPage := NewSearch()
	assert.Equal(t, expectedPage, actualPage)
}

func TestNewUpload(t *testing.T) {
	expectedPage := Page{
		title:    "Hinzufügen",
		headLine: "Anlegen einer Aktivität"}
	actualPage := NewUpload()
	assert.Equal(t, expectedPage, actualPage)
}

func TestPage_GetTitle(t *testing.T) {
	expectedTitle := "Strava"
	page := NewIndex()
	actualTitle := page.GetTitle()
	assert.Equal(t, expectedTitle, actualTitle)
}

func TestPage_GetCustomStyleSheet(t *testing.T) {
	expectedCSS := "../assets/css/items.css"
	page := NewIndex()
	actualCSS := page.GetCustomStyleSheet()
	assert.Equal(t, expectedCSS, actualCSS)
}

func TestPage_GetHeadline(t *testing.T) {
	expectedHeadline := "Deine Aktivitäten"
	page := NewIndex()
	actualHeadline := page.headLine
	assert.Equal(t, expectedHeadline, actualHeadline)
}

func TestPage_GetJsFile(t *testing.T) {
	expectedJSFile := "../assets/js/custom.js"
	page := NewIndex()
	actualJSFile := page.GetJsFile()
	assert.Equal(t, expectedJSFile, actualJSFile)
}

func TestPage_HasCustomStyleSheet(t *testing.T) {
	expectedHasCSS := true
	page := NewIndex()
	actualHasCSS := page.HasCustomStyleSheet()
	assert.Equal(t, expectedHasCSS, actualHasCSS)
}

func TestPage_HasJsFile(t *testing.T) {
	expectedHasJS := true
	page := NewIndex()
	actualHasJS := page.HasJsFile()
	assert.Equal(t, expectedHasJS, actualHasJS)
}
