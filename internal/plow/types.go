package plow

import "traffic.go/util"

type UsePlow struct {
	ID  string
	ID2 string

	Active bool
	State  string

	Position Point

	ClosestMile *util.MileMarker
	Route       string
}

type Point struct {
	Latitude  float64
	Longitude float64
}
