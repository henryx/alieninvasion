package world

import "testing"

func TestCity_AddNear(t *testing.T) {
	cA := NewCity("A")
	cB := NewCity("B")

	if err := cA.AddNear(cB, NORTH); err != nil {
		t.Fatal(err)
	}
}

func TestCity_AddNearSame(t *testing.T) {
	cA := NewCity("A")

	err := cA.AddNear(cA, NORTH)
	if err != nil {
		return
	}

	t.Fatal(err)
}
