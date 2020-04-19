package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "abcйefg"
	fmt.Println(len(s), utf8.RuneCountInString(s))
	a := 1
loop:
	fmt.Println(a)
	a++
	if a < 10 {
		goto loop
	}

}
