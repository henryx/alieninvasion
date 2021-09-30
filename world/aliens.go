package world

import (
	"fmt"
	"math/rand"
)

func move(cities *map[string]*City, city *City) {
	direction := RandDirection()
	if near, ok := city.Directions[direction]; ok {
		fmt.Println("An alien from", city.Name, "move to", near)
		(*cities)[near].Aliens++
		city.Aliens--
	}
}

// Invade randomly invade cities with aliens
func Invade(cities *map[string]*City, aliens int) {
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

func Attack(cities *map[string]*City) {
	for cityName, city := range *cities {
		if city.Aliens == 1 {
			if len(city.Directions) == 0 {
				fmt.Println("Aliens in city", cityName, "are trapped!")
				city.Trapped = true
			}

			if city.Trapped {
				continue
			}

			move(cities, city)
		}

		if city.Aliens == 2 {
			fmt.Println("City", cityName, "is destroyed")
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
