package main

import (
	"alieninvasion/world"
	"log"
	"os"
	"strconv"
)

func main() {
	var err error
	var aliens = 16

	if len(os.Args) < 2 {
		log.Printf("Usage: %v <worldmap> [aliens]\n", os.Args[0])
		os.Exit(1)
	}

	if len(os.Args) == 3 {
		aliens, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Aliens need to be a number: %city\n", err)
		}

		if aliens == 0 {
			log.Fatal("Aliens can't be 0")
		}
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %city\n", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error when closing file: %v\n", err)
		}
	}()

	cities, err := world.GetCities(file)
	if err != nil {
		log.Fatal(err)
	}

	if aliens > 2*len(cities) {
		log.Fatalf("Aliens are greater than the double of cities\n")
	}

	world.Invade(&cities, aliens)
}
