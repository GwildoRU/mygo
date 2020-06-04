package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fh, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	log.SetOutput(fh)

	fn := os.Args[1]
	//fn := "d:\\PRG\\GO\\projects\\mygo\\fcr\\kp_xml\\OUT\\kp_06e61f6e-6914-43aa-b596-adac249332de.xml"
	//fn := "d:\\PRG\\GO\\projects\\mygo\\fcr\\kp_xml\\IN\\kp_5fd6bf9b-6f3a-4827-9d80-4a4cbde5abf8.xml"
	parseXml(fn)
}

func parseXml(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data)

	var s2 []string

	s2 = append(s2, filepath.Base(filename))

	a1 := getTag(string(data), "Area")
	a2 := getTagVal(a1[1:], "Area")

	if a2 != "-" {
		s2 = append(s2, a2)
	} else {
		s2 = append(s2, getTagVal(s, "Area"))
	}

	s = getTag(s, "Address")

	city := getParam(s, "adrs:City", "Type")
	if city == "-" {
		s2 = append(s2, getParam(s, "City", "Type"))
		s2 = append(s2, getParam(s, "City", "Name"))
		s2 = append(s2, getParam(s, "Street", "Type"))
		s2 = append(s2, getParam(s, "Street", "Name"))
		s2 = append(s2, getParam(s, "Level", "Type"))
		s2 = append(s2, getParam(s, "Level", "Name"))
		s2 = append(s2, getParam(s, "Apartment", "Type"))
		s2 = append(s2, getParam(s, "Apartment", "Name"))
	} else {
		s2 = append(s2, getParam(s, "adrs:City", "Type"))
		s2 = append(s2, getParam(s, "adrs:City", "Name"))
		s2 = append(s2, getParam(s, "adrs:Street", "Type"))
		s2 = append(s2, getParam(s, "adrs:Street", "Name"))
		s2 = append(s2, getParam(s, "adrs:Level", "Type"))
		s2 = append(s2, getParam(s, "adrs:Level", "Value"))
		s2 = append(s2, getParam(s, "adrs:Apartment", "Type"))
		s2 = append(s2, getParam(s, "adrs:Apartment", "Value"))
	}
	s = getTag(string(data), "Person")

	familyName := getTagVal(s, "FamilyName")
	if familyName != "-" {
		s2 = append(s2, familyName)
		s2 = append(s2, getTagVal(s, "FirstName"))
		s2 = append(s2, getTagVal(s, "Patronymic"))
	} else {
		s2 = append(s2, getTagVal(s, "Surname"))
		s2 = append(s2, getTagVal(s, "First"))
		s2 = append(s2, getTagVal(s, "Patronymic"))
	}

	fmt.Println(strings.Join(s2, "\t"))

}

func getTag(s, tag string) string {
	i1 := strings.Index(s, "<"+tag+">")
	i2 := strings.Index(s, "</"+tag+">")
	if i1 < 0 || i2 < 0 || i2 < i1 {
		return "-"
	}
	return s[i1 : i2+len(tag)+3]
}

func getTagVal(s, tag string) string {
	i1 := strings.Index(s, "<"+tag+">")
	i2 := strings.Index(s, "</"+tag+">")
	if i1 < 0 || i2 < 0 || i2 < i1 {
		return "-"
	}
	return s[i1+len(tag)+2 : i2]
}

func getParam(s, tag, param string) string {
	i1 := strings.Index(s, "<"+tag)
	if i1 < 0 {
		return "-"
	}
	s = s[i1+len(tag)+1:]

	i2 := strings.Index(s, "/>")
	if i2 < 0 {
		return "-"
	}
	s = s[0:i2]

	i3 := strings.Index(s, param+"=")
	if i3 < 0 {
		return "-"
	}

	s = s[i3+len(param)+2:]
	return s[0:strings.Index(s, "\" ")]
}
