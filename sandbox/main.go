package main

import "fmt"

type Exchanger interface {
	Exchange()
}

type StringPair struct{ first, second string }

func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

func (pair StringPair) String() string {
	return fmt.Sprintf("%q+%q", pair.first, pair.second)
}

func main() {
	sp := StringPair{"11111", "22222"}
	fmt.Println(sp)
}
