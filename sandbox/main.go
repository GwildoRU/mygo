package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	s := "=?koi8-r?B?7yDQz9PU1dDMxc7JySDExc7F1s7ZyCDT0sXE09TXICjExdTBzNjO2Q==?= =?koi8-r?B?yikueGxzeA==?="
	s2 := "=?koi8-r?B?"
	s3 := "=?="

	fmt.Println(s)
	s = strings.ReplaceAll(s, s2, "\n")
	s = strings.ReplaceAll(s, s3, "\n")
	fmt.Println(s)
	ss := strings.Split(s, "\n")
	fmt.Println(ss)

	s4 := ""

	for n, s := range ss {
		fmt.Println(n)
		fmt.Println(s)



		data, _ := base64.StdEncoding.DecodeString(s+"=")
		ss[n] = string(data)
		//fmt.Println(data)
		fmt.Println(string(data))
		s4 += string(data)
	}
	fmt.Println(strings.ReplaceAll(strings.Join(ss,""),"\n",""))
	//fmt.Println("==>",strings.Join(ss,""))

}
