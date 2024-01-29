package scrape

import (
	"time"

	"traffic.go/internal/merge"
	"traffic.go/tg"
	"traffic.go/util"
)

const REFRESH_DELAY = 15

func RunAndSend() {
	tree := util.MakeTree()

	bot := tg.StartBot()
	var messageID int

	Total := merge.Merge(tree)
	initBody := tg.FormatMessage(Total)

	messageID = tg.SendMessage(bot, initBody)
	time.Sleep(REFRESH_DELAY * time.Second)
	for {
		if messageID != 0 {

			Total = merge.Merge(tree)
			body := tg.FormatMessage(Total)

			tg.EditMessage(bot, body, messageID)
			time.Sleep(REFRESH_DELAY * time.Second)
		}

	}

}
