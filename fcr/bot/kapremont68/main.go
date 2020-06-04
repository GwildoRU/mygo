package main

import (
	"github.com/botogonia/flowbot"
	"log"
)

func main() {
	bot, err := flowbot.NewFlowBot(
		"1173218587:AAFeNbm9qiJubZyMR79_VU4a3HFO9nkRrpw",
		200,
		"Вас долго не было, начните сначала")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	bot.HandleUpdates(botChatHandler)
}
