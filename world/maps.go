package world

import "errors"

// City node represents the city
type City struct {
	Name       string
	Directions map[Point]string
	Aliens     int
}

// NewCity is used as constructor for City
func NewCity(city string) *City {
	c := City{
		Name:       city,
		Directions: make(map[Point]string),
		Aliens:     0,
	}

	return &c
}

// AddNear method add a City defining the cardinal point where is located
func (c *City) AddNear(city string, direction Point) error {
	if c.Name == city {
		return errors.New("cannot add the same city")
	}

	c.Directions[direction] = city

	return nil
}
