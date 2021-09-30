package world

import "testing"

func TestGetDirection(t *testing.T) {
	_, err := GetDirection("west")
	if err != nil {
		t.Fatal()
	}
}

func TestGetDirection_Fail(t *testing.T) {
	_, err := GetDirection("foo")
	if err != nil {
		return
	}

	t.Fatal()
}

func TestRandDirection(t *testing.T) {
	direction := RandDirection()

	if direction == "" {
		t.Fatal()
	}
}
