/*
 * 2848869
 * 8089098
 * 3861852
 */

package gpxProcessing

import (
	"math"
	"time"
)

func GetAllTrackPoints(file GpxFile) []TrackPoint {
	var allTrackPoints []TrackPoint
	for _, track := range file.GetTracks() {
		for _, segment := range track.GetTrackSegments() {
			allTrackPoints = append(allTrackPoints, segment.GetTrackPoints()...)
		}
	}
	return allTrackPoints
}

func GetMaxSpeed(points []TrackPoint) float64 {
	var maxSpeed float64 = 0
	for _, point := range points {
		var currentSpeed = point.GetExtension().GetTrackPointExtension().GetSpeed()
		if currentSpeed > maxSpeed {
			maxSpeed = currentSpeed
		}
	}
	return maxSpeed
}

func GetAvgSpeed(points []TrackPoint) float64 {
	var speedSum float64 = 0
	for _, point := range points {
		var currentSpeed = point.GetExtension().GetTrackPointExtension().GetSpeed()
		if currentSpeed != 0 {
			speedSum = speedSum + currentSpeed
		}
	}
	return speedSum / float64(len(points))
}

//calculate total distance in km by adding up distance between previous trackpoint and current
func CalculateDistanceInKilometers(points []TrackPoint) float64 {
	var previousTrkPt TrackPoint
	var totalDistance float64
	for index, point := range points {
		if index == 0 {
			previousTrkPt = point
		} else {
			var DistanceBetweenPoints = CalculateDistance2Points(previousTrkPt.GetLatitude(), previousTrkPt.GetLongitude(), point.GetLatitude(), point.GetLongitude())
			totalDistance = totalDistance + DistanceBetweenPoints
			previousTrkPt = point
		}
	}
	return totalDistance
}

//calculate the distance between 2 points using the haversine formula
func CalculateDistance2Points(lat1, lon1, lat2, lon2 float64) float64 {
	radiantLat1 := CalculateRadiant(lat1)
	radiantLat2 := CalculateRadiant(lat2)
	radiantLon1 := CalculateRadiant(lon1)
	radiantLon2 := CalculateRadiant(lon2)
	differenceLat := radiantLat1 - radiantLat2
	differenceLon := radiantLon1 - radiantLon2

	haversine := math.Pow(math.Sin(differenceLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(differenceLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(haversine), math.Sqrt(1-haversine))

	distanceInKm := c * 6371

	return distanceInKm
}

func CalculateRadiant(val float64) float64 {
	return val * math.Pi / 180
}

func CalculateStandbyTimeInMins(points []TrackPoint) float64 {
	var standbyTimeInSec float64
	var previousTrkPt TrackPoint
	for index, point := range points {
		var currentSpeed = point.GetExtension().GetTrackPointExtension().GetSpeed()
		if index == 0 {
			previousTrkPt = point
		} else {
			var previousSpeed = previousTrkPt.GetExtension().GetTrackPointExtension().GetSpeed()
			if previousSpeed <= 1 && currentSpeed <= 1 {
				var timeDifference = time.Time.Sub(point.GetDateTime(), previousTrkPt.GetDateTime()).Seconds()
				standbyTimeInSec = standbyTimeInSec + timeDifference
				previousTrkPt = point
			}
		}
	}
	return SecondsToMinutes(standbyTimeInSec)
}
func SecondsToMinutes(seconds float64) float64 {
	mins := seconds / 60
	return math.Round(mins*100) / 100
}

func VerifySportType(sportType string, avgSpeed float64) string {
	var matchingSportType = sportType
	switch sportType {
	case "Laufen":
		if avgSpeed > 16 {
			matchingSportType = "Radfahren"
		}
	case "Radfahren":
		if avgSpeed < 16 {
			matchingSportType = "Laufen"
		}
	}
	return matchingSportType
}