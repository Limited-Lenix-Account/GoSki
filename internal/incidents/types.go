package incidents

import (
	"github.com/dhconnelly/rtreego"
	"traffic.go/util"
)

type UsableIncident struct {
	IncidentType string

	Route          string
	Direction      string
	StartMile      float64
	EndMile        float64
	PrimaryPoint   rtreego.Point
	SecondaryPoint rtreego.Point
	PrimaryMile    *util.MileMarker
	SecondaryMile  *util.MileMarker

	SingleMileStr string
	SinglePoint   rtreego.Point
	SingleMile    *util.MileMarker
}