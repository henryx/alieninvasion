package world

import "testing"

func TestCity_AddNear(t *testing.T) {
	cA := NewCity(1, "A")
	cB := NewCity(2, "B")

	if err := cA.AddNear(cB.Name, NORTH); err != nil {
		t.Fatal(err)
	}
}

func TestCity_AddNearSame(t *testing.T) {
	cA := NewCity(1, "A")

	err := cA.AddNear(cA.Name, NORTH)
	if err != nil {
		return
	}

	t.Fatal(err)
}
