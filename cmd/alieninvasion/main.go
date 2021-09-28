package main

import (
	"alieninvasion/world"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %v <worldmap>\n", os.Args[0])
		os.Exit(1)
	}

	_, err := world.GetCities(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
