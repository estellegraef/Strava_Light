/*
 * 2848869
 * 8089098
 * 3861852
 */

package main

import (
	"github.com/estellegraef/Strava_Light/backend/activity"
	"github.com/estellegraef/Strava_Light/backend/webserver"
)

func main() {
	activity.Setup()
	webserver.CreateWebServer()
}
