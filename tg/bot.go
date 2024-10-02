package tg

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	chatID int64  // telegram chat channel ID here
	botKey string // telegram bot key here

)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botKey = os.Getenv("BOT_ID")
	chatIDInt, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	chatID = int64(chatIDInt)
}

func StartBot() *tg.BotAPI {
	fmt.Println(botKey)
	fmt.Println("Starting TG Bot...")

	bot, err := tg.NewBotAPI(botKey)
	if err != nil {
		fmt.Printf("Error Creating Telegram Bot %s", err)
	}

	return bot

}

func SendMessage(bot *tg.BotAPI, finalMessage string) int {

	msg := tg.NewMessage(chatID, finalMessage)
	msg.ParseMode = tg.ModeMarkdownV2

	// Send the message
	sentMessage, err := bot.Send(msg)
	if err != nil {
		fmt.Println(finalMessage)
		log.Panic(err)
	}

	fmt.Println("Message Sent!")
	return sentMessage.MessageID

}

func EditMessage(bot *tg.BotAPI, editMessage string, messageID int) {
	editMsg := tg.NewEditMessageText(chatID, messageID, editMessage)
	editMsg.ParseMode = tg.ModeMarkdownV2

	_, err := bot.Send(editMsg)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Message Edited!")
}
