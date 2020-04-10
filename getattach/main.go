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
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Inbox struct {
	Inboxname    string `json:"inboxname"`
	Attachsuffix string `json:"attachsuffix"`
	Outfolder    string `json:"outfolder"`
}

func main() {
	const (
		YAHOST     = "imap.yandex.ru:993"
		YAUSER     = "reestr@kapremont68.ru"
		YAPASSWORD = "Qpwo1029"
	)

	start := time.Now()

	fmt.Printf("--------------------- %s -------------------\n", start)

	var inboxes []Inbox

	jsonBlob, err := ioutil.ReadFile("inbox.json")

	err = json.Unmarshal(jsonBlob, &inboxes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(YAHOST, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	// Don't forget to logout
	defer c.Logout() //nolint:errcheck

	// Login
	if err := c.Login(YAUSER, YAPASSWORD); err != nil {
		panic(err)
	}
	fmt.Println("Logged in")

	for i := range inboxes {
		parseInboxItem(c, inboxes[i])
	}

	secs := time.Since(start).Seconds()

	fmt.Println("Done!", secs)
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
	fmt.Println(inboxitem.Inboxname, "=> Unseen messages IDs:", ids)

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
								fmt.Println(inboxitem.Inboxname, "=> attach: ", fn)

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
