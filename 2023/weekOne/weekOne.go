package adventofcode2023

import (
	"fmt"
	"os"
)

func weekOne() {
	fmt.Println("Day One")
	fileContents, err := readFile("adventOfCode/2023/weekOne/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(fileContents))
}

func readFile(filePath string) ([]byte, error) {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return fileContents, nil
}