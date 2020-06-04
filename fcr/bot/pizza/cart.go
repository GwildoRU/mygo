package main

import (
	"strconv"
	"strings"
)

func cartList(cart []string) string {
	list := []string{
		"Ваш заказ",
		"_________",
	}
	n := 0
	sum := 0
	for i := range cart {
		p := menu[cart[i]]
		n++
		sum += p.price
		list = append(list, p.name+" - "+strconv.Itoa(p.price)+" руб.")
	}
	if len(cart) > 1 {
		list = append(list, "_________\nитого: "+strconv.Itoa(n)+" шт. на "+strconv.Itoa(sum)+" руб.")
	}

	return strings.Join(list, "\n")
}
