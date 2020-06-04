package main

import (
	"github.com/botogonia/flowbot"
)

func pizzaChatHandler(ch *flowbot.Chat) {

	defer ch.Close()

	_, _ = ch.WaitUpdate()

	startChoice, data := ch.Choice(0,
		"Здесь можно заказать пиццу с доставкой",
		&flowbot.Kbrd{
			{{Text: "Заказать", Data: "order"}},
		},
		"Чтобы заказать пиццу нажмите кнопку [Заказать]",
	)

	if data != "order" {
		return
	}

	menuKbd := makeMenuKbrd(2)

	var cart []string

menuChoice:
	menuChoice, data := ch.Choice(startChoice,
		"Выберите пиццу",
		menuKbd,
		"Для выбора пиццы нажмите соответствующую кнопку",
	)

	cart = append(cart, data)

	addPizza, data := ch.Choice(menuChoice, cartList(cart),

		&flowbot.Kbrd{
			{{Text: "Добавить еще пиццу", Data: "add"}},
			{{Text: "Оформить заказ", Data: "order"}},
		},
		"",
	)

	if data == "add" {
		startChoice = addPizza
		goto menuChoice
	}

	ch.SendText(addPizza, cartList(cart))

	_, addr := ch.Prompt(0, "Укажите адрес доставки", "")

	_, phone := ch.Prompt(0, "Укажите телефон для связи", "")

	payChoice, payType := ch.Choice(0, "Оплата курьеру",
		&flowbot.Kbrd{
			{{Text: "Наличными", Data: "наличными"}, {Text: "Картой", Data: "картой"}},
		},
		"",
	)

	ch.SendText(payChoice, cartList(cart)+
		"\n\nадрес доставки: "+addr+
		"\nтелефон: "+phone+
		"\nоплата: "+payType)

	okChoice, data := ch.Choice(0, "Всё верно?", &flowbot.Kbrd{
		{{Text: "Да, всё верно. Отправить заказ.", Data: "ok"}},
	},
		"",
	)

	if data == "ok" {
		ch.SendText(okChoice, "Спасибо за заказ! \nОжидайте звонка оператора.")
	}

}
