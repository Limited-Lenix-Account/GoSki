package merge

import (
	"traffic.go/internal/alerts"
	"traffic.go/internal/traffic"
)

type GrandObject struct {
	LovelandPass *PassStatus
	VailPass     *PassStatus
	BerthodPass  *PassStatus

	Traffic *[]traffic.UseableTraffic
}

type PassStatus struct {
	Name string
	Open int

	Alerts []alerts.UseableAlert
}
