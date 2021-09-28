package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %v <worldmap>\n", os.Args[0])
		os.Exit(1)
	}
}
