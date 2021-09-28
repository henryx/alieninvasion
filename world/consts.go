package world

import "errors"

type Point string

const NORTH Point = "north"
const SOUTH Point = "south"
const EAST Point = "east"
const WEST Point = "west"

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
