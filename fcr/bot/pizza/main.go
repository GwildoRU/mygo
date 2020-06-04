package main

import (
	"github.com/botogonia/flowbot"
	"log"
)

func main() {
	bot, err := flowbot.NewFlowBot(
		"1256801395:AAENPIswnTiGpODZcucI6ycXcsfHumot8GY",
		60,
		"Вас долго не было, начните сначала")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	bot.HandleUpdates(pizzaChatHandler)
}
