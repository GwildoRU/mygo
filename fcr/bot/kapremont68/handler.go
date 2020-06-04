package main

import (
	"github.com/botogonia/flowbot"
	"log"
	"os/exec"
	"strings"
)

func botChatHandler(ch *flowbot.Chat) {

	defer ch.Close()

	_, _ = ch.WaitUpdate()

	ch.SendText(0, "Фонд капитального ремонта Тамбовской области")

	_, login := ch.Prompt(0, "Введите логин")
	_, password := ch.Prompt(0, "Введите пароль")

	log.Printf("%s - login: %s, password: %s \n", ch.GetUserNameOrId(), login, password)

	if logins[login] != password {
		ch.SendText(0, "Мы вас не знаем")
		log.Printf("%s - Error: Bad credentials \n", ch.GetUserNameOrId())
		return
	}

startChoice:
	startChoice, choice1 := ch.Choice(0,
		"Что сформировать?",
		&flowbot.Kbrd{
			{{Text: "Квитанцию", Data: "kvit"}},
			{{Text: "Акт сверки", Data: "act"}},
		},
	)

	log.Printf("%s - choice1: %s \n", ch.GetUserNameOrId(), choice1)

	_, account := ch.Prompt(startChoice, "Укажите номер счета")

	period := ""
	if choice1 == "kvit" {
		_, period = ch.Prompt(0, "Укажите период в формате ММ.ГГГГ")
		period = strings.Replace(period, ",", ".", -1)
	}

	log.Printf("%s - account: %s, period: %s \n", ch.GetUserNameOrId(), account, period)

	ch.SendText(0, "Подождите немного.")

	switch choice1 {
	case "kvit":
		exec.Command("kvit", login, account+"_"+period).Run()
		pdfName := ".\\OUT\\" + login + "\\kvit\\KVIT_" + account + "_" + period + ".pdf"
		ch.SendFile(pdfName)
		log.Printf("%s - %s\n", ch.GetUserNameOrId(), pdfName)
	case "act":
		exec.Command("act", login, account).Run()
		pdfName := ".\\OUT\\" + login + "\\act\\ACT_" + account + ".pdf"
		ch.SendFile(pdfName)
		log.Printf("%s - %s\n", ch.GetUserNameOrId(), pdfName)
	}

	endChoice, choice2 := ch.Choice(0,
		"Что дальше?",
		&flowbot.Kbrd{
			{{Text: "Продолжить работу", Data: "continue"}},
			{{Text: "Выход", Data: "exit"}},
		},
	)

	ch.DelMsg(endChoice)
	log.Printf("%s - endChoice: %s \n", ch.GetUserNameOrId(), endChoice)

	if choice2 == "continue" {
		goto startChoice
	}

}
