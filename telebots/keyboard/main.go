package main

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("01"),
		tgbotapi.NewKeyboardButton("02"),
		tgbotapi.NewKeyboardButton("03"),
		tgbotapi.NewKeyboardButton("04"),
		tgbotapi.NewKeyboardButton("05"),
		tgbotapi.NewKeyboardButton("06"),
		tgbotapi.NewKeyboardButton("07"),
		tgbotapi.NewKeyboardButton("08"),
		tgbotapi.NewKeyboardButton("09"),
		tgbotapi.NewKeyboardButton("10"),
		tgbotapi.NewKeyboardButton("11"),
		tgbotapi.NewKeyboardButton("12"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("13"),
		tgbotapi.NewKeyboardButton("14"),
		tgbotapi.NewKeyboardButton("15"),
		tgbotapi.NewKeyboardButton("16"),
		tgbotapi.NewKeyboardButton("17"),
		tgbotapi.NewKeyboardButton("18"),
		tgbotapi.NewKeyboardButton("19"),
		tgbotapi.NewKeyboardButton("20"),
		tgbotapi.NewKeyboardButton("20"),
		tgbotapi.NewKeyboardButton("22"),
		tgbotapi.NewKeyboardButton("23"),
		tgbotapi.NewKeyboardButton("24"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("25"),
		tgbotapi.NewKeyboardButton("26"),
		tgbotapi.NewKeyboardButton("27"),
		tgbotapi.NewKeyboardButton("28"),
		tgbotapi.NewKeyboardButton("29"),
		tgbotapi.NewKeyboardButton("30"),
		tgbotapi.NewKeyboardButton("31"),
		tgbotapi.NewKeyboardButton("32"),
		tgbotapi.NewKeyboardButton("33"),
		tgbotapi.NewKeyboardButton("34"),
		tgbotapi.NewKeyboardButton("35"),
		tgbotapi.NewKeyboardButton("36"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("37"),
		tgbotapi.NewKeyboardButton("38"),
		tgbotapi.NewKeyboardButton("39"),
		tgbotapi.NewKeyboardButton("40"),
		tgbotapi.NewKeyboardButton("41"),
		tgbotapi.NewKeyboardButton("42"),
		tgbotapi.NewKeyboardButton("43"),
		tgbotapi.NewKeyboardButton("44"),
		tgbotapi.NewKeyboardButton("45"),
		tgbotapi.NewKeyboardButton("46"),
		tgbotapi.NewKeyboardButton("47"),
		tgbotapi.NewKeyboardButton("48"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("49"),
		tgbotapi.NewKeyboardButton("50"),
		tgbotapi.NewKeyboardButton("51"),
		tgbotapi.NewKeyboardButton("52"),
		tgbotapi.NewKeyboardButton("53"),
		tgbotapi.NewKeyboardButton("54"),
		tgbotapi.NewKeyboardButton("55"),
		tgbotapi.NewKeyboardButton("56"),
		tgbotapi.NewKeyboardButton("57"),
		tgbotapi.NewKeyboardButton("58"),
		tgbotapi.NewKeyboardButton("59"),
		tgbotapi.NewKeyboardButton("60"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("61"),
		tgbotapi.NewKeyboardButton("62"),
		tgbotapi.NewKeyboardButton("63"),
		tgbotapi.NewKeyboardButton("64"),
		tgbotapi.NewKeyboardButton("65"),
		tgbotapi.NewKeyboardButton("66"),
		tgbotapi.NewKeyboardButton("67"),
		tgbotapi.NewKeyboardButton("68"),
		tgbotapi.NewKeyboardButton("69"),
		tgbotapi.NewKeyboardButton("70"),
		tgbotapi.NewKeyboardButton("71"),
		tgbotapi.NewKeyboardButton("72"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("73"),
		tgbotapi.NewKeyboardButton("74"),
		tgbotapi.NewKeyboardButton("75"),
		tgbotapi.NewKeyboardButton("76"),
		tgbotapi.NewKeyboardButton("77"),
		tgbotapi.NewKeyboardButton("78"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1173218587:AAFeNbm9qiJubZyMR79_VU4a3HFO9nkRrpw")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
