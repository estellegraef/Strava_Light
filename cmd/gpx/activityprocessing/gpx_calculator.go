/*
 * 2848869
 */
package activityprocessing

import (
	"Strava_Light/cmd/gpx/gpx_info"
	"math"
)

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

//calculate total distance in km by adding up distance between previous trackpoint and current
func CalculateDistanceInKilometers(points []gpx_info.TrackPoint) float64 {
	var previousTrkPt gpx_info.TrackPoint
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
	return val*math.Pi/180
}

//return correct speed according to average speed
func CorrectSpeed(speed, avgspeed float64) float64{
	var correctSpeed = speed
	if speed > (avgspeed + avgspeed/2) {
		correctSpeed = avgspeed
	}
	return correctSpeed
}