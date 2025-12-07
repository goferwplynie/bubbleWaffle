package main

import (
	"fmt"

	"github.com/goferwplynie/bubbleWaffle/internal/analyzer"
)

func main() {
	//cmd.Execute()
	comps, _ := analyzer.LoadComponents(".")
	fmt.Println(comps)
}
