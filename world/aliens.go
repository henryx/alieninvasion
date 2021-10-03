package world

import (
	"fmt"
	"math/rand"
)

// move move an alien from one city to its neighbors
func move(cities *map[string]*City, city *City) {
	direction := RandDirection()
	if near, ok := city.Directions[direction]; ok {
		fmt.Println("An alien from", city.Name, "move to", near)
		(*cities)[near].Aliens = append((*cities)[near].Aliens, city.Aliens[0])
		city.Aliens = city.Aliens[1:]
	}
}

// Invade randomly invade cities with aliens
func Invade(cities *map[string]*City, aliens int) {
	for alien := 0; alien < aliens; alien++ {
		id := rand.Intn(len(*cities))
		for _, city := range *cities {
			if len(city.Aliens) == 2 {
				continue
			}

			if city.Id == id {
				city.Aliens = append(city.Aliens, alien)
				break
			}
		}
	}
}

// Attack execute attack from aliens to near cities
func Attack(cities *map[string]*City) {
	for cityName, city := range *cities {
		switch len(city.Aliens) {
		case 1:
			{
				if len(city.Directions) == 0 {
					fmt.Println("Aliens in city", cityName, "are trapped!")
					city.Trapped = true
				}

				if city.Trapped {
					continue
				}

				move(cities, city)
			}

		case 2:
			{
				fmt.Println("City", cityName, "is destroyed by alien", city.Aliens[0], "and alien", city.Aliens[1])
				delete(*cities, cityName)
				for nearCity, near := range *cities {
					for direction, c := range near.Directions {
						if cityName == c {
							fmt.Println("Delete from", nearCity, "city", cityName, "at", direction)
							delete(near.Directions, direction)
						}
					}
				}
			}
		}
	}
}
