package world

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetCities(filename string) ([]*City, error) {
	var cities []*City
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
		city := NewCity(split[0])
		for _, v := range split[1:] {
			pos := strings.Split(v, "=")
			near := NewCity(pos[1])
			direction, err := GetDirection(pos[0])
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
