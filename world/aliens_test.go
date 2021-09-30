package world

import "testing"

func TestInvade(t *testing.T) {
	var err error

	file := openFile(t)
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	err = writeFile(file, "Foo north=Bar east=Hou south=Qu-ux west=Baz")
	if err != nil {
		t.Fatal(err)
	}

	_, _ = file.Seek(0, 0)
	data, err := GetCities(file)
	if err != nil {
		t.Fatal(err)
	}

	Invade(&data, len(data)*2)

	if data["Foo"].Aliens == 0 {
		t.Fatal("Aliens haven't invaded city")
	}
}
