package main

import (
	"alieninvasion/world"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	var err error

	if len(os.Args) < 2 {
		log.Printf("Usage: %v <worldmap>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(errors.New(fmt.Sprintf("Error opening file: %v\n", err)))
	}
	defer file.Close()

	_, err = world.GetCities(file)
	if err != nil {
		log.Fatal(err)
	}
}
