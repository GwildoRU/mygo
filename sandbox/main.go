package main

import (
	"fmt"
	"runtime"
)

func main() {
	//ioutil.ReadDir
	fmt.Println(runtime.GOMAXPROCS(-1))
	//fmt.Println("11111111")
	//log.Println("22222222222")
}
