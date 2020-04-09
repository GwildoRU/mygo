package main

import "fmt"

type T1 struct {
	F1 string
}
type T2 struct {
	F1 string
}
type T3 struct {
	T1
	T2
}

func main() {
	var V2 T3
	V2.T1.F1 = "123456"
	fmt.Println(V2)
}
