package main

import (
	"fmt"
	"os"

	"github.com/goferwplynie/bubbleWaffle/internal/analyzer"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Scanning:", cwd)
	components, err := analyzer.LoadComponents(cwd)
	if err != nil {
		panic(err)
	}
	for _, c := range components {
		fmt.Println("Found:", c.Name)
	}
	fmt.Printf("Total found: %d\n", len(components))
}
