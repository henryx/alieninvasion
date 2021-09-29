package world

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func GetCities(file io.Reader) ([]*City, error) {
	var cities []*City
	var err error

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		split := strings.Split(line, " ")
		city := NewCity(split[0])
		for _, v := range split[1:] {
			pos := strings.Split(v, "=")
			if len(pos) != 2 {
				return nil, errors.New(fmt.Sprintf("position not correct %s", v))
			}
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
