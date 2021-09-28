package main

import (
	"alieninvasion/world"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func parse(filename string) ([]*world.City, error) {
	var cities []*world.City
	var err error

	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error opening file: %v\n", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		city := world.NewCity(split[0])
		for _, v := range split[1:] {
			pos := strings.Split(v, "=")
			near := world.NewCity(pos[1])
			direction, err := world.GetDirection(pos[0])
			if err != nil {
				return nil, err
			}
			if err := city.AddNear(near, direction); err != nil {
				return nil, err
			}
		}
		cities = append(cities, city)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %v <worldmap>\n", os.Args[0])
		os.Exit(1)
	}

	_, err := parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
