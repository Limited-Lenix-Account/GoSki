package util

import "github.com/dhconnelly/rtreego"

type SpatialObject struct {
	Rect rtreego.Rect
	MileMarker
}

type MileMarker struct {
	RoadType string
	Marker   string
	Route    string

	Coordinates Coordinates

	Geom rtreego.Rect
}

func (m *MileMarker) Bounds() rtreego.Rect {
	return m.Geom
}

type Coordinates struct {
	Lat  float64
	Long float64
}
