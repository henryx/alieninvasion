package world

import (
	"errors"
	"math/rand"
)

type Point string

const NORTH Point = "north"
const SOUTH Point = "south"
const EAST Point = "east"
const WEST Point = "west"

// GetDirection return a Point constant based by cardinal point string passed. If string is not
// compliant, it returns an error
func GetDirection(direction string) (Point, error) {
	switch direction {
	case "north":
		return NORTH, nil
	case "south":
		return SOUTH, nil
	case "east":
		return EAST, nil
	case "west":
		return WEST, nil
	default:
		return "", errors.New("invalid cardinal point")
	}
}

// RandDirection returns a Point constant chosen randomically
func RandDirection() Point {
	direction := rand.Intn(4)

	switch direction {
	case 0:
		return NORTH
	case 1:
		return EAST
	case 2:
		return SOUTH
	case 3:
		return WEST
	}

	return ""
}
