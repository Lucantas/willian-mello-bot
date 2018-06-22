package main

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("WMTOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		r1, _ := regexp.Compile("\\A+(?:o\\s+)tel|\\s+(?:o\\s+)tel|\\A+(?:o\\s+)teu|\\s+(?:o\\s+)teu|hoste|mote|hote")

		if r1.MatchString(strings.ToLower(update.Message.Text)) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hotel? Trivago!")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
