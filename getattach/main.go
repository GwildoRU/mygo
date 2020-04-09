package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aglyzov/charmap"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Inbox struct {
	Inboxname, Attachsuffix, Outfolder string
}

func main() {
	const (
		YAHOST     = "imap.yandex.ru:993"
		YAUSER     = "reestr@kapremont68.ru"
		YAPASSWORD = "Qpwo1029"
	)

	var inboxes []Inbox

	file, err := os.Open("inbox.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	dec := json.NewDecoder(file)

	for {
		var m Inbox
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		inboxes = append(inboxes, m)
	}

	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(YAHOST, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout() //nolint:errcheck

	// Login
	if err := c.Login(YAUSER, YAPASSWORD); err != nil {
		panic(err)
	}
	log.Println("Logged in")

	start := time.Now()

	for i := range inboxes {
		parseInboxItem(c, inboxes[i])
	}

	secs := time.Since(start).Seconds()

	log.Println("Done!", secs)
}

func parseInboxItem(c *client.Client, inboxitem Inbox) {

	_, err := c.Select(inboxitem.Inboxname, false)
	if err != nil {
		panic(err)
	}
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
	//criteria.WithoutFlags = []string{}
	ids, err := c.Search(criteria)
	if err != nil {
		panic(err)
	}
	log.Println("Unseen messages IDs found:", ids)

	if len(ids) > 0 {
		seqset := new(imap.SeqSet)
		seqset.AddNum(ids...)

		messages := make(chan *imap.Message, 10)
		done := make(chan error, 1)
		go func() {
			done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchRFC822}, messages)
		}()

		for msg := range messages {
			for _, r := range msg.Body {
				entity, err := message.Read(r)
				if err != nil {
					panic(err)
				}

				if mr := entity.MultipartReader(); mr != nil {
					for {
						p, err := mr.NextPart()
						if err == io.EOF {
							break
						}
						_, params, _ := p.Header.ContentType()

						if params["name"] != "" {
							fn := getFileName(params["name"])
							if strings.EqualFold(path.Ext(fn), inboxitem.Attachsuffix) {
								log.Println(inboxitem.Inboxname, "=> attach: ", fn)

								c, err := ioutil.ReadAll(p.Body)
								if err != nil {
									panic(err)
								}
								if err = ioutil.WriteFile(getUniqFileName(filepath.Join(inboxitem.Outfolder, fn)), c, 0777); err != nil {
									panic(err)
								}
							}
						}
					}
				}
			}
		}

		if err := <-done; err != nil {
			panic(err)
		}
	}
}

func getUniqFileName(oldName string) (newName string) {
	newName = oldName
	i := 2
	for {
		if _, err := os.Stat(newName); err == nil {
			ext := path.Ext(oldName)
			noext := strings.TrimSuffix(oldName, ext)
			newName = fmt.Sprintf(noext+" (%v)"+ext, i)
			i++
		} else {
			break
		}
	}
	return newName
}

func getFileName(s string) string {
	s2 := "=?koi8-r?B?"
	if strings.HasPrefix(s, s2) {
		s = s[len(s2) : len(s)-2]
		data, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return s
		}
		data, _ = charmap.ANY_to_UTF8(data, "KOI8-R")
		return fmt.Sprintf("%s", data)
	}
	return s
}
