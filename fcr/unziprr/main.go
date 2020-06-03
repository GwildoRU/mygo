package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fh, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	log.SetOutput(fh)

	filename := os.Args[1]
	if err := unpackZip(filename); err != nil {
		log.Fatal(fmt.Errorf("unpackZip (%s) \n=> %v", filename, err), "\n", strings.Repeat("-", 50), "\n")
	}
}

func unpackZip(filename string) error {
	fmt.Println(filename)
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, zipFile := range reader.Reader.File {
		if strings.HasSuffix(zipFile.Name, ".zip") {
			if err = unpackZippedFile("TMP\\"+zipFile.Name, zipFile); err != nil {
				return fmt.Errorf("unpackZippedFile (.zip) \n=> %v", err)
			}
			if err = unpackZip("TMP\\" + zipFile.Name); err != nil {
				return fmt.Errorf("unpackZip (%s) \n=> %v", zipFile.Name, err)
			}
		} else if re := regexp.MustCompile(`^(report|kp_).*\.(pdf|xml)$`); re.Match([]byte(zipFile.Name)) {
			if err = unpackZippedFile("OUT\\"+zipFile.Name, zipFile); err != nil {
				return fmt.Errorf("unpackZippedFile (pdf|xml) \n=> %v", err)
			}
		}
	}
	return nil
}

func unpackZippedFile(filename string, zipFile *zip.File) error {
	writer, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer writer.Close()
	reader, err := zipFile.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	if _, err = io.Copy(writer, reader); err != nil { //nolint:wsl
		return err
	}
	return nil //nolint:wsl
}
