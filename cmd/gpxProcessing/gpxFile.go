/*
 * 2848869
 * 8089098
 * 3861852
 */

package gpxProcessing

import (
	"time"
)

type GpxFile struct {
	Creator string   `xml:"creator,attr"`
	Meta    Metadata `xml:"metadata"`
	Tracks  []Track  `xml:"trk"`
}

type Metadata struct {
	Time time.Time `xml:"time"`
}

type Track struct {
	TrackSegments []TrackSegment `xml:"trkseg"`
}

type TrackSegment struct {
	TrackPoints []TrackPoint `xml:"trkpt"`
}

type TrackPoint struct {
	Latitude   float64   `xml:"lat,attr"`
	Longitude  float64   `xml:"lon,attr"`
	DateTime   time.Time `xml:"time"`
	Extensions Extension `xml:"extensions"`
}

type Extension struct {
	TrackPointExtensions TrackPointExtension `xml:"TrackPointExtension"`
}

type TrackPointExtension struct {
	Speed float64 `xml:"speed"`
}

func (g GpxFile) GetCreator() string {
	return g.Creator
}

func (g GpxFile) GetMeta() Metadata {
	return g.Meta
}

func (g GpxFile) GetTracks() []Track {
	return g.Tracks
}

func (m Metadata) GetTime() time.Time {
	return m.Time
}

func (t Track) GetTrackSegments() []TrackSegment {
	return t.TrackSegments
}

func (t TrackSegment) GetTrackPoints() []TrackPoint {
	return t.TrackPoints
}

func (t TrackPoint) GetLatitude() float64 {
	return t.Latitude
}

func (t TrackPoint) GetLongitude() float64 {
	return t.Longitude
}

func (t TrackPoint) GetDateTime() time.Time {
	return t.DateTime
}

func (t TrackPoint) GetExtension() Extension {
	return t.Extensions
}

func (e Extension) GetTrackPointExtension() TrackPointExtension {
	return e.TrackPointExtensions
}

func (t TrackPointExtension) GetSpeed() float64 {
	return t.Speed
}
