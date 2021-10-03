package world

import "testing"

func TestInvade(t *testing.T) {
	var err error

	file, err := createTempFile(t, "Foo north=Bar east=Hou south=Qu-ux west=Baz")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	data, err := GetCities(file)
	if err != nil {
		t.Fatal(err)
	}

	Invade(&data, len(data)*2)

	if len(data["Foo"].Aliens) == 0 {
		t.Fatal("Aliens haven't invaded city")
	}
}

func TestAttack(t *testing.T) {
	file, err := createTempFile(t, "Foo north=Bar east=Hou south=Qu-ux west=Baz")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	data, err := GetCities(file)
	if err != nil {
		t.Fatal(err)
	}

	Invade(&data, len(data)*2)
	Attack(&data)

	if _, ok := data["Foo"]; ok {
		t.Fatal("Aliens in Foo haven't destroyed the city:", data["Foo"].Aliens)
	}
}
