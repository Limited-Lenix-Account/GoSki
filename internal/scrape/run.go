package scrape

import (
	"traffic.go/internal/merge"
	"traffic.go/tg"
)

func RunAndSend() {

	var Total merge.GrandObject

	Loveland, Vail, Berthoud := merge.GetValidAlerts()
	traffic := merge.GetRelivantTraffic()

	Total.LovelandPass = Loveland
	Total.VailPass = Vail
	Total.BerthodPass = Berthoud

	Total.Traffic = traffic

	tg.SendMessage(Total)

}
