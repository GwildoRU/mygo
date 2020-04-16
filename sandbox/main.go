package main

import (
	"fmt"
	os "os"
)

func main() {
	myfunc()
}

func myfunc() {
	fn1 := "q1.txt"
	fn2 := "q2.txt"
	f, _ := os.Create(fn1)
	defer func() {
		os.Rename(fn1, fn2)
	}()
	defer f.Close()
	fmt.Fprint(f, "1111111\n", "2222222\n", "333\n")
	fn2 = "q3.txt"
}
