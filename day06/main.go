package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const OBSTACLE rune = '#'
const EMPTY rune = '.'
const VISITED rune = 'X'

func solvePart1(inputFile string) int {
	data := loadDaySixData(inputFile)

	colMax := len(data[0]) - 1
	rowMax := len(data) - 1

	var nextCol int
	var nextRow int
	var nextRune rune

	// Find starting position
	currRow, currCol, direction := startingPosition(data)
	data[currRow][currCol] = VISITED

	// Walk and Mark
	for {
		nextRow, nextCol = nextPosition(direction, currRow, currCol)

		if 0 > nextCol || nextCol > colMax || 0 > nextRow || nextRow > rowMax {
			break
		}

		nextRune = data[nextRow][nextCol]
		if nextRune == '#' {
			direction = rotateGuard(direction)
			continue
		}

		currCol = nextCol
		currRow = nextRow
		data[currRow][currCol] = VISITED
	}

	// count #
	sum := 0
	for row := 0; row <= rowMax; row++ {
		for col := 0; col <= colMax; col++ {
			if data[row][col] == VISITED {
				sum++
			}
		}
	}

	return sum
}

func rotateGuard(direction rune) rune {
	switch direction {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	default:
		log.Fatal("unknown direction")
		return '^' // Unreachable
	}
}

func nextPosition(direction rune, currRow, currCol int) (int, int) {
	switch direction {
	case '^':
		return currRow - 1, currCol
	case '>':
		return currRow, currCol + 1
	case 'v':
		return currRow + 1, currCol
	case '<':
		return currRow, currCol - 1
	}
	return currRow, currCol
}

func startingPosition(grid [][]rune) (int, int, rune) {
	colMax := len(grid[0]) - 1
	rowMax := len(grid) - 1

	for row := 0; row <= rowMax; row++ {
		for col := 0; col <= colMax; col++ {
			r := grid[row][col]
			if r != OBSTACLE && r != EMPTY {
				return row, col, r
			}
		}
	}

	log.Fatal("No starting position found")
	return 0, 0, EMPTY
}

func solvePart2(inputFile string) int {
	data := loadDaySixData(inputFile)

	sum := 0

	startRow, startCol, startDirection := startingPosition(data)

	colMax := len(data[0]) - 1
	rowMax := len(data) - 1

	for r := 0; r <= rowMax; r++ {
		for c := 0; c <= colMax; c++ {
			if data[r][c] != EMPTY {
				continue
			}

			data[r][c] = OBSTACLE
			track := make([][][]rune, rowMax+1)
			for i := 0; i <= rowMax; i++ {
				track[i] = make([][]rune, colMax+1)
			}

			direction := startDirection
			currCol := startCol
			currRow := startRow
			var nextRow int
			var nextCol int
			var nextRune rune
			rotations := 0
			loop := false
			// Walk
			for {
				nextRow, nextCol = nextPosition(direction, currRow, currCol)

				if 0 > nextCol || nextCol > colMax || 0 > nextRow || nextRow > rowMax {
					break
				}

				nextRune = data[nextRow][nextCol]
				if nextRune == '#' {
					rotations++
					if rotations >= 3 {
						sum++
						break
					}
					direction = rotateGuard(direction)
					continue
				}

				rotations = 0

				for _, r := range track[currRow][currCol] {
					if direction == r {
						sum++
						loop = true
						break
					}
				}
				if loop {
					break
				}

				track[currRow][currCol] = append(track[currRow][currCol], direction)

				currCol = nextCol
				currRow = nextRow
			}

			data[r][c] = EMPTY
		}
	}

	return sum
}

func loadDaySixData(inputFile string) [][]rune {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
