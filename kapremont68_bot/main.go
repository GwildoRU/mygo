package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

//https://uk-167-116-1.friproxy0.eu:443 [UK]
//https://ua-139-170-1.fri-gate0.biz:443 [UA]
//https://fr-54-189-1.friproxy.eu:443 [FR]

func main() {
	//proxyUrl, err := url.Parse("https://uk-167-116-1.friproxy0.eu:443")
	//myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	//bot, err := tgbotapi.NewBotAPIWithClient("1173218587:AAFeNbm9qiJubZyMR79_VU4a3HFO9nkRrpw",tgbotapi.APIEndpoint,myClient)
	bot, err := tgbotapi.NewBotAPI("1173218587:AAFeNbm9qiJubZyMR79_VU4a3HFO9nkRrpw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		log.Println(update.Message)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text+" => qqqqqqqqqqqq")
		//msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
