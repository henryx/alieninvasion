package world

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

func GetCities(file io.Reader) ([]*City, error) {
	var cities []*City
	var err error

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
		cities = append(cities, city)
		id += 1
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}

// Invade randomly invade cities with aliens
func Invade(cities *[]*City, aliens int) {
	for alien := 0; alien < aliens; alien++ {
		id := rand.Intn(len(*cities))
		for _, city := range *cities {
			if city.Aliens == 2 {
				continue
			}

			if city.Id == id {
				city.Aliens++
				break
			}
		}
	}
}
