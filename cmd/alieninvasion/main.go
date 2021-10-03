package main

import (
	"alieninvasion/world"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var err error
	var aliens int
	var fileName string

	flag.IntVar(&aliens, "aliens", 16, "number of aliens invading (default 16)")
	flag.StringVar(&fileName, "world", "", "a file used as world map input")
	flag.Parse()

	if aliens == 0 {
		log.Fatal("Number of aliens can't be 0")
	}

	if fileName == "" {
		log.Fatal("World map isn't passed")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
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
	world.Attack(&cities)

	fmt.Println()
	trapped := 0
	for cityName, city := range cities {
		if city.Trapped {
			trapped++
		}

		if city.Aliens > 0 {
			fmt.Println("In city", cityName, "remains", city.Aliens, "aliens")
		}
	}

	fmt.Println("City trapped:", trapped)
}
