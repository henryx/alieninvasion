package world

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// GetCities retrieve cities from map and sets their neighbors
func GetCities(file io.Reader) (map[string]*City, error) {
	var cities map[string]*City
	var err error

	cities = make(map[string]*City)
	id := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		split := strings.Split(line, " ")
		city := NewCity(id, split[0])
		for _, v := range split[1:] {
			pos := strings.Split(v, "=")
			if len(pos) != 2 {
				return nil, errors.New(fmt.Sprintf("position not correct %s", v))
			}
			direction, err := GetDirection(pos[0])
			if err != nil {
				return nil, err
			}
			if err := city.AddNear(pos[1], direction); err != nil {
				return nil, err
			}
		}
		cities[city.Name] = city
		id += 1
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}
