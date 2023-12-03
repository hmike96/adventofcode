package adventofcode2023

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
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

func WeekThree() {
	lines := getInput()
	grid := make([][]rune, len(lines))
	addToSum := false
	numberToSkip := 0
	sum := 0
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	for i, row := range grid {
		for j, val := range row {
			if numberToSkip > 0 {
				numberToSkip--
				continue
			}
			if unicode.IsDigit(val) {
				fmt.Printf("Grid[%d][%d]: %c is an integer\n", i, j, val)
				addToSum = false
				if checkSurroundingSymbol(grid, i, j) {
					fmt.Printf("Grid[%d][%d]: %c is adjacent to a symbol\n", i, j, val)
					addToSum = true
				}
				numString := string(val)
				for k := j + 1; k < len(row); k++ {
					if unicode.IsDigit(row[k]) {
						numberToSkip++
						numString += string(row[k])
						if checkSurroundingSymbol(grid, i, k) {
							fmt.Printf("Grid[%d][%d]: %c is adjacent to a symbol\n", i, k, row[k])
							addToSum = true
						}
					} else {
						break
					}
				}
				fmt.Printf("Combined digits: %s\n", numString)
				if addToSum {
					num, err := strconv.Atoi(numString)
					if err != nil {
						fmt.Printf("Error converting %s to an integer: %v\n", numString, err)
					} else {
						sum += num
					}
				}
				fmt.Println("Number to skip:", numberToSkip)
			} else if unicode.IsSymbol(val) {
				fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i, j, val)
			} else {
				fmt.Printf("Grid[%d][%d]: %c is not an integer or symbol\n", i, j, val)
			}
		}
	}
	fmt.Println("Sum:", sum)
}

func checkSurroundingSymbol(grid [][]rune, i int, j int) bool {
	if boundsCheck(grid, i - 1, j - 1) && unicode.IsSymbol(grid[i - 1][j - 1]) {
		fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i - 1, j - 1, grid[i - 1][j - 1])
		return true
	} else if boundsCheck(grid, i - 1, j) && unicode.IsSymbol(grid[i - 1][j]) {
		fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i - 1, j, grid[i - 1][j])
		return true
	} else if boundsCheck(grid, i - 1, j + 1) && unicode.IsSymbol(grid[i - 1][j + 1]) {
		fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i - 1, j + 1, grid[i - 1][j + 1])
		return true
	} else if boundsCheck(grid, i, j - 1) && unicode.IsSymbol(grid[i][j - 1]) {
		fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i, j - 1, grid[i][j - 1])
		return true
	} else if boundsCheck(grid, i, j + 1) && unicode.IsSymbol(grid[i][j + 1]) {
		fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i, j + 1, grid[i][j + 1])
		return true
	} else if boundsCheck(grid, i + 1, j - 1) && unicode.IsSymbol(grid[i + 1][j - 1]) {
		fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i + 1, j - 1, grid[i + 1][j - 1])
		return true
	} else if boundsCheck(grid, i + 1, j) && unicode.IsSymbol(grid[i + 1][j]) {
		fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i + 1, j, grid[i + 1][j])
		return true
	} else if boundsCheck(grid, i + 1, j + 1) && unicode.IsSymbol(grid[i + 1][j + 1]) {
		fmt.Printf("Grid[%d][%d]: %c is a symbol\n", i + 1, j + 1, grid[i + 1][j + 1])
		return true
	}
	return false
}

func boundsCheck(grid [][]rune, i int, j int) bool {
	if i > -1 && j > -1 && i < len(grid) && j < len(grid[i]) {
		return true
	}
	return false
}