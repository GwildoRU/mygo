package main

import (
	"encoding/base64"
	"fmt"
	"github.com/aglyzov/charmap"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	const (
		YA_HOST     = "imap.yandex.ru:993"
		YA_USER     = "reestr@kapremont68.ru"
		YA_PASSWORD = "Qpwo1029"
	)

	if len(os.Args) < 4 {
		err := fmt.Errorf("usage: %s inboxname attachsuffix outfolder", filepath.Base(os.Args[0]))
		fmt.Println(err)
		os.Exit(1)
	}

	inboxname := os.Args[1]
	attachsuffix := os.Args[2]
	outfolder := os.Args[3]

	fmt.Println(inboxname, attachsuffix, outfolder)

	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(YA_HOST, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(YA_USER, YA_PASSWORD); err != nil {
		panic(err)
	}
	log.Println("Logged in")

	_, err = c.Select(inboxname, false)
	if err != nil {
		panic(err)
	}

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
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
							if strings.HasSuffix(fn, attachsuffix) {
								log.Println("attach: ", fn)

								c, err := ioutil.ReadAll(p.Body)
								if err != nil {
									panic(err)
								}
								if err = ioutil.WriteFile(filepath.Join(outfolder, fn), c, 0777); err != nil {
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

	log.Println("Done!")
}

func getFileName(s string) string {
	s2 := "=?koi8-r?B?"
	if strings.Index(s, s2) == 0 {
		s = s[len(s2) : len(s)-2]
		data, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return s
		}
		data, err = charmap.ANY_to_UTF8(data, "KOI8-R")
		return fmt.Sprintf("%s", data)
	}
	return s
}
