package main

import (
	"flag"
	"fmt"
)

var inboxname = flag.String("inboxname", "INBOX", "Папка с письмами во входящих")
var attachext = flag.String("attachext", ".zip", "Расширение файлов вложений")
var outfolder = flag.String("outfolder", "TMP", "Папка для сохранения вложений на диске")

func main() {
	flag.Parse()
	fmt.Println(*inboxname, *attachext, *outfolder)
}
