/*
 * 2848869
 */
package activityprocessing

import (
	"Strava_Light/cmd/gpx/gpx_info"
)

/*important:
- Strecke in km
- wenn Zeitstempel, dann Durchschnittsgeschwindigkeit (ohne Standzeit) und Maximalgeschwindigkeit
- Standzeit (Zeit in der sich nicht bewegt wurde)
- Eingabefehler anhand der Durchschnittsgeschwindigkeit korrigieren
- Datum
 */

func GetAllTrackPoints(file gpx_info.GpxFile) []gpx_info.TrackPoint {
	var allTrackPoints []gpx_info.TrackPoint
	for _, track := range file.GetTracks(){
		for _, segment := range track.GetTrackSegments(){
			allTrackPoints = append(allTrackPoints, segment.GetTrackPoints()...)
		}
	}
	return allTrackPoints
}

func GetMaxSpeed(points []gpx_info.TrackPoint) float64 {
	var maxSpeed float64 = 0
	for _, point := range points {
		var currentSpeed = point.GetExtension().GetTrackPointExtension().GetSpeed()
		if currentSpeed > maxSpeed {
			maxSpeed = currentSpeed
		}
	}
	return maxSpeed
}

func GetAvgSpeed(points []gpx_info.TrackPoint) float64 {
	var speedSum float64 = 0
	for _, point := range points{
		speedSum = speedSum + point.GetExtension().GetTrackPointExtension().GetSpeed()
	}
	return speedSum / float64(len(points))
}

func CalculateDistanceInKilometers(points []gpx_info.TrackPoint) float64 {
	//Harvesine Formula
	return 0
}