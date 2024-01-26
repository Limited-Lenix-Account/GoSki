package util

type MileMarker struct {
	RoadType string
	Marker   string
	Route    string

	Coordinates Coordinates
}

type Coordinates struct {
	Lat  float64
	Long float64
}
