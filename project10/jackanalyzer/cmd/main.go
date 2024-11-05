package main

import (
	"errors"
	"jackanalyzer"
	"log"
)

func main() {
	tockenizer, err := jackanalyzer.NewTockenizer("cmd/Main.jack")
	if err != nil {
		log.Fatal("error starting tockenizer", err)
	}

	for {
		err := tockenizer.Advance()
		if errors.Is(err, jackanalyzer.ErrNoMoreCommands) {
			break
		}
		if err != nil {
			log.Fatal("error reading file", err)
		}
	}
}
