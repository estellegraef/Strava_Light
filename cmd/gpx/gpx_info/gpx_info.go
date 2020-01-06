package gpx_info

import (
	"time"
)

type GpxFile struct{
	Creator string `xml:"creator,attr"`
	Meta Metadata `xml:"metadata"`
	Tracks []Track `xml:"trk"`
}

type Metadata struct {
	Time time.Time `xml:"time"`
}

type Track struct{
	TrackSegments []TrackSegment `xml:"trkseg"`
}

type TrackSegment struct{
	TrackPoints []TrackPoint `xml:"trkpt"`
}

type TrackPoint struct {
	Latitude float64 `xml:"lat,attr"`
	Longitude float64 `xml:"long,attr"`
	DateTime time.Time `xml:"time"`
	Speed float64 `xml:"speed"`
}

func NewGpx(creator string, metaData Metadata, tracks []Track) GpxFile {
	return GpxFile{Creator:creator, Meta:metaData, Tracks:tracks}
}

func NewMeta(time time.Time) Metadata {
	return Metadata{Time:time}
}

func NewTrack(segments []TrackSegment) Track {
	return Track{TrackSegments:segments}
}

func NewTrackSegment(points []TrackPoint) TrackSegment {
	return TrackSegment{TrackPoints:points}
}

func NewTrackPoint(lat float64, long float64, dateTime time.Time, speed float64) TrackPoint{
	return TrackPoint{Latitude:lat, Longitude:long, DateTime:dateTime, Speed:speed}
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

func (t TrackPoint) GetSpeed() float64 {
	return t.Speed
}