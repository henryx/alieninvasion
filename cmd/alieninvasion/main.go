package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.New(fmt.Sprintf("Error opening file: %v\n", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %v <worldmap>\n", os.Args[0])
		os.Exit(1)
	}

	err := parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
