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
	Inboxname    string `json:"inboxname"`
	Attachsuffix string `json:"attachsuffix"`
	Outfolder    string `json:"outfolder"`
}

func main() {
	start := time.Now()

	const (
		YAHOST     = "imap.yandex.ru:993"
		YAUSER     = "reestr@kapremont68.ru"
		YAPASSWORD = "Qpwo1029"
	)

	log.SetOutput(os.Stdout)

	fmt.Printf("--------------------- %s -------------------\n", start)

	var inboxes []Inbox

	jsonBlob, err := ioutil.ReadFile("inbox.json")

	err = json.Unmarshal(jsonBlob, &inboxes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(YAHOST, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")

	// Don't forget to logout
	defer c.Logout() //nolint:errcheck

	// Login
	if err := c.Login(YAUSER, YAPASSWORD); err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
	//criteria.WithoutFlags = []string{}
	ids, err := c.Search(criteria)
	if err != nil {
		log.Fatal(err)
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
					log.Fatal(err)
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
							//log.Println(fn)
							if strings.EqualFold(path.Ext(fn), inboxitem.Attachsuffix) {
								fmt.Println(inboxitem.Inboxname, "=> attach: ", fn)

								c, err := ioutil.ReadAll(p.Body)
								if err != nil {
									log.Fatal(err)
								}
								if err = ioutil.WriteFile(getUniqFileName(filepath.Join(inboxitem.Outfolder, fn)), c, 0777); err != nil {
									log.Fatal(err)
								}
							}
						}
					}
				}
			}
		}

		if err := <-done; err != nil {
			log.Fatal(err)
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
	//log.Println(s)
	s2 := "=?koi8-r?B?"
	s3 := "=?="
	if strings.HasPrefix(s, s2) {
		s = strings.ReplaceAll(s, s2, "\n")
		s = strings.ReplaceAll(s, s3, "\n")
		ss := strings.Split(s, "\n")

		for n, s := range ss {
			data, _ := base64.StdEncoding.DecodeString(s+"=")
			data, _ = charmap.ANY_to_UTF8(data, "KOI8-R")
			ss[n] = string(data)
		}
		//return fmt.Sprintf("%s", data)
		return strings.Join(ss,"")
	}
	return s
}
