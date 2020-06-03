package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("kvit", "1")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}