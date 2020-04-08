package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	type Inbox struct {
		Inboxname, Attachsuffix, Outfolder string
	}
	file, err := os.Open("inbox.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	dec := json.NewDecoder(file)

	inboxes := []Inbox{}

	for {
		var m Inbox
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		inboxes = append(inboxes, m)
	}
	fmt.Println(inboxes)
}
