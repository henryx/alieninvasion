package world

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func openFile(t *testing.T) *os.File {
	dir := t.TempDir()
	file, err := ioutil.TempFile(dir, "world.txt")
	if err != nil {
		t.Fatalf("unable to create temporary file for testing")
	}

	return file
}

func writeFile(file *os.File, str string) error {
	w := bufio.NewWriter(file)

	_, err := w.WriteString(fmt.Sprintf("%s\n", str))
	if err != nil {
		return err
	}
	_ = w.Flush()

	return nil
}

func TestGetCities(t *testing.T) {
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

	if len(data) == 0 {
		t.Fatal("Cannot read file")
	}
}

func TestGetCitiesFail(t *testing.T) {
	var err error

	file := openFile(t)
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	err = writeFile(file, "Foo err")
	if err != nil {
		t.Fatal(err)
	}

	_, _ = file.Seek(0, 0)
	data, err := GetCities(file)
	if err == nil {
		t.Fatal(err)
	}

	if len(data) > 0 {
		t.Fatal(err)
	}
}

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

	if data[0].Aliens == 0 {
		t.Fatal("Aliens haven't invaded city")
	}
}
