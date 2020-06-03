package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("kvit", "444078560261836_05.2020")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}