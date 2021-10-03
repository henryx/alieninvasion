package world

import "errors"

// City node represents the city
type City struct {
	Id         int
	Name       string
	Directions map[Point]string
	Aliens     []int
	Trapped    bool
}

// NewCity is used as constructor for City
func NewCity(id int, city string) *City {
	c := City{
		Id:         id,
		Name:       city,
		Directions: make(map[Point]string),
		Aliens:     []int{},
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
