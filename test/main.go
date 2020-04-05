package main

import "fmt"

func main() {
	a, b := 1, 1

	a, b = a+1, b+2
	append()

	fmt.Println(a, b)
}
