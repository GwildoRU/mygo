package main

import (
	"github.com/botogonia/flowbot"
	"strconv"
)

type Pizza struct {
	name  string
	price int
}

var menu = map[string]Pizza{
	"1": {"Ветчина грибы", 220},
	"2": {"Деревенская", 273},
	"3": {"Курица острая", 258},
	"4": {"Морской коктейль", 318},
	"5": {"С фаршем", 238},
	"6": {"Тамбовская", 316},
	"7": {"Четыре сыра", 306},
}

func makeMenuKbrd(rl int) *flowbot.Kbrd {

	var kbd flowbot.Kbrd
	var row []flowbot.KbrdBtn

	for k, v := range menu {
		if len(row) == rl {
			kbd = append(kbd, row)
			row = []flowbot.KbrdBtn{}
		}
		b := flowbot.KbrdBtn{v.name + " - " + strconv.Itoa(v.price), k}
		row = append(row, b)
	}
	if len(row) > 0 {
		kbd = append(kbd, row)
	}

	return &kbd
}
