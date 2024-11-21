package merge

import (
	"traffic.go/internal/alerts"
	"traffic.go/internal/incidents"
	"traffic.go/internal/plow"
	"traffic.go/internal/traffic"
)

type GrandObject struct {
	LovelandPass *PassStatus
	VailPass     *PassStatus
	BerthodPass  *PassStatus

	Traffic   *[]traffic.UseableTraffic
	Incidents *[]incidents.UsableIncident
}

type PassStatus struct {
	Name  string
	Open  bool
	Plows []plow.UsePlow

	Alerts []alerts.UseableAlert
}
