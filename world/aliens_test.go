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

	if data["Foo"].Aliens == 0 {
		t.Fatal("Aliens haven't invaded city")
	}
}
