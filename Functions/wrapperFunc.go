package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	sliceBytes, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("error in main in readfile: %s", err)
	}

	fmt.Println(sliceBytes)
	fmt.Println(string(sliceBytes))
}

func readFile(fileName string) ([]byte, error) {
	sliceBytes, err := os.ReadFile("poem.txt")
	if err != nil {
		return nil, fmt.Errorf("there was an error in readfile: %s", err)
	}
	return sliceBytes, nil
}
