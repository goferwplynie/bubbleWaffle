package main

import (
	"github.com/goferwplynie/bubbleWaffle/internal/analyzer"
)

func main() {
	_, err := analyzer.LoadComponents("./")
	if err != nil {
		panic(err)
	}
	//cmd.Execute()
}
