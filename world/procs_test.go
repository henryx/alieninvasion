package world

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func createTempFile(dir, data string) (*os.File, error) {
	var err error

	file, err := ioutil.TempFile(dir, "world.txt")
	if err != nil {
		return nil, errors.New("unable to create temporary file for testing")
	}

	w := bufio.NewWriter(file)
	_, err = w.WriteString(fmt.Sprintf("%s\n", data))
	if err != nil {
		return nil, err
	}

	err = w.Flush()
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func TestGetCities(t *testing.T) {
	var err error

	file, err := createTempFile(t.TempDir(), "Foo north=Bar east=Hou south=Qu-ux west=Baz")
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

	if len(data) == 0 {
		t.Fatal("Cannot read file")
	}
}

func TestGetCitiesFail(t *testing.T) {
	var err error

	file, err := createTempFile(t.TempDir(), "Foo err")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	data, err := GetCities(file)
	if err == nil {
		t.Fatal(err)
	}

	if len(data) > 0 {
		t.Fatal(err)
	}
}
