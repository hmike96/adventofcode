package adventofcode2023

import (
	_ "embed"
	"fmt"
	"strconv"

	"strings"
)

//go:embed  input.txt
var inputFile string

func getInput() []string{
	fmt.Println("Day One")
	var lines []string
	inputFile = strings.TrimRight(inputFile, "\n")
	lines = strings.Split(inputFile, "\n")
	fmt.Println(lines)
	return lines
}

func WeekOne() {
	lines := getInput()
	sum := 0
	for _, line := range lines {
		firstInteger := 0
		var lastInteger int
		chars := strings.Split(line, "")
		for _, char := range chars {
			if _, err := strconv.Atoi(char); err == nil {
				firstInteger, _ = strconv.Atoi(char)
				fmt.Println("First integer character:", char)
				break
			}
		}
		for i := len(chars) - 1; i >= 0; i-- {
			char := chars[i]
			if _, err := strconv.Atoi(char); err == nil {
				lastInteger, _ = strconv.Atoi(char)
				fmt.Println("First integer character in reverse order:", char)
				break
			}
		}
		concatenated := strconv.Itoa(firstInteger) + strconv.Itoa(lastInteger)
		fmt.Println("Concatenated integers:", concatenated)
		concatenatedInt, _ := strconv.Atoi(concatenated)
		sum += concatenatedInt
	}
	fmt.Println("Sum:", sum)
}


