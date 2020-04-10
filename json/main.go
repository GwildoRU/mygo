package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	type Inbox struct {
		Inboxname    string `json:"inboxname"`
		Attachsuffix string `json:"attachsuffix"`
		Outfolder    string `json:"outfolder"`
	}

	inboxes2 := []Inbox{
		{
			Inboxname:    "Почта|Уварово",
			Attachsuffix: ".dbf",
			Outfolder:    "DBF",
		},
		{
			Inboxname:    "Почта|Мордовский",
			Attachsuffix: ".xls",
			Outfolder:    "XLS",
		},
	}

	b, err := json.Marshal(inboxes2)
	if err != nil {
		fmt.Println("error:", err)
	}

	err = ioutil.WriteFile("inbox2.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}

	jsonBlob, err := ioutil.ReadFile("inbox.json")

	var inboxes []Inbox

	err = json.Unmarshal(jsonBlob, &inboxes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inboxes)

}
