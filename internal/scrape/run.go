package scrape

import (
	"time"

	"traffic.go/internal/merge"
	"traffic.go/tg"
)

func RunAndSend() {

	bot := tg.StartBot()
	var messageID int

	var Total merge.GrandObject

	Loveland, Vail, Berthoud := merge.GetValidAlerts()
	traffic := merge.GetRelivantTraffic()

	Total.LovelandPass = Loveland
	Total.VailPass = Vail
	Total.BerthodPass = Berthoud

	Total.Traffic = traffic
	initBody := tg.FormatMessage(Total)
	messageID = tg.SendMessage(bot, initBody)
	time.Sleep(5 * time.Second)
	for {
		if messageID != 0 {

			Loveland, Vail, Berthoud := merge.GetValidAlerts()
			traffic := merge.GetRelivantTraffic()

			Total.LovelandPass = Loveland
			Total.VailPass = Vail
			Total.BerthodPass = Berthoud

			Total.Traffic = traffic
			body := tg.FormatMessage(Total)
			tg.EditMessage(bot, body, messageID)
			time.Sleep(5 * time.Second)
		}

	}

}
