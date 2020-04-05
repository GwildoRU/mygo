package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	pdfFileName := os.Args[1]
	xmlFileName := strings.ReplaceAll(pdfFileName, ".pdf", ".xml")
	xml, err := ioutil.ReadFile(xmlFileName)
	if err != nil {
		fmt.Println(pdfFileName)
		os.Exit(1)
	}

	area := extractTagValue(string(xml), "area")
	adr := extractTagValue(string(xml), "readable_address")
	regnum := extractTagValue(string(xml), "registration_number")
	formdate := extractTagValue(string(xml), "date_formation")
	newPDFname := replaceInvalidChars(adr + " (" + area + ") " + formdate + " " + regnum + ".pdf")

	fmt.Println(newPDFname)
}

func extractTagValue(s, tag string) string {
	return s[strings.Index(s, "<"+tag+">")+len(tag)+2 : strings.Index(s, "</"+tag+">")]
}

func replaceInvalidChars(s string) string {
	re := regexp.MustCompile(`[^№()\-,.\s\wА-Яа-я]`)
	s = re.ReplaceAllString(s, "_")
	s = strings.ReplaceAll(s, "__", "_")
	s = strings.Join(strings.Fields(s), " ")
	return s
}
