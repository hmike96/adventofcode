package adventofcode2023

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed  input.txt
var inputFile string

func getInput() []string {
	fmt.Println("Day Two")
	var lines []string
	inputFile = strings.TrimRight(inputFile, "\n")
	lines = strings.Split(inputFile, "\n")
	fmt.Println(lines)
	return lines
}

func WeekTwo() {
	lines := getInput()
	blue := 14
	red := 12
	green := 13
	sum := 0
	for _, line := range lines {
		chars := strings.Split(line, " ")
		if len(chars) > 2 {
			fmt.Println(line)
			game := strings.TrimRight(chars[1], ":")
			gameInt, err := strconv.Atoi(game)
			if err != nil {
				fmt.Println("Error converting game to integer:", err)
				continue
			}
			// Use gameInt as an integer
			badGame := false
			// Iterate over chars two indexes at a time
			for i := 2; i < len(chars)-1; i += 2 {
				fmt.Println("char at index", i, ":", chars[i], ", Color at index", i + 1, ":", chars[i + 1])
				// Rest of the code...
				trimmedColor := strings.TrimRight(strings.TrimRight(chars[i + 1], ";"), ",")
				if trimmedColor == "red" {
					nextIndex := i
					if nextIndex < len(chars) {
						nextInt, err := strconv.Atoi(chars[nextIndex])
						if err != nil {
							fmt.Println("Error converting next index to integer:", err)
							continue
						}
						if nextInt > red {
							badGame = true
							break // Skip to the next iteration of the loop
						}
					}
				} else if trimmedColor == "blue" {
					nextIndex := i
					if nextIndex < len(chars) {
						nextInt, err := strconv.Atoi(chars[nextIndex])
						if err != nil {
							fmt.Println("Error converting next index to integer:", err)
							continue
						}
						if nextInt > blue {
							badGame = true
							break // Skip to the next iteration of the loop
						}
					}
				} else if trimmedColor == "green" {
					nextIndex := i 
					if nextIndex < len(chars) {
						nextInt, err := strconv.Atoi(chars[nextIndex])
						if err != nil {
							fmt.Println("Error converting next index to integer:", err)
							continue
						}
						if nextInt > green {
							badGame = true
							break // Skip to the next iteration of the loop
						}
					}
				}
			}
			if !badGame {
				sum += gameInt
			}
		}
	}
	fmt.Println("Sum:", sum)
}

func WeekTwoParTwo() {
	lines := getInput()
	sum := 0
	for _, line := range lines {
		chars := strings.Split(line, " ")
		if len(chars) > 2 {
			// Use gameInt as an integer
			maxRed := 0
			maxBlue := 0
			maxGreen := 0
			// Iterate over chars two indexes at a time
			for i := 2; i < len(chars)-1; i += 2 {
				fmt.Println("char at index", i, ":", chars[i], ", Color at index", i + 1, ":", chars[i + 1])
				// Rest of the code...
				trimmedColor := strings.TrimRight(strings.TrimRight(chars[i + 1], ";"), ",")
				if trimmedColor == "red" {
					nextIndex := i
					if nextIndex < len(chars) {
						nextInt, err := strconv.Atoi(chars[nextIndex])
						if err != nil {
							fmt.Println("Error converting next index to integer:", err)
							continue
						}
						if maxRed < nextInt {
							maxRed = nextInt
						}
					}
				} else if trimmedColor == "blue" {
					nextIndex := i
					if nextIndex < len(chars) {
						nextInt, err := strconv.Atoi(chars[nextIndex])
						if err != nil {
							fmt.Println("Error converting next index to integer:", err)
							continue
						}
						if maxBlue < nextInt {
							maxBlue = nextInt
						}
					}
				} else if trimmedColor == "green" {
					nextIndex := i 
					if nextIndex < len(chars) {
						nextInt, err := strconv.Atoi(chars[nextIndex])
						if err != nil {
							fmt.Println("Error converting next index to integer:", err)
							continue
						}
						if maxGreen < nextInt {
							maxGreen = nextInt
						}
					}
				}
			}
			sum	+= maxRed * maxBlue * maxGreen
		}
	}
	fmt.Println("Sum:", sum)
}
