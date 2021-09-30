package world

import "math/rand"

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
