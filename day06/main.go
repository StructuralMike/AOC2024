package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solvePart1(inputFile string) int {
	data := loadDaySixData(inputFile)

	OBSTACLE := '#'
	EMPTY := '.'
	VISITED := 'X'

	colMax := len(data[0]) - 1
	rowMax := len(data) - 1

	var currCol int
	var currRow int
	var nextCol int
	var nextRow int
	var direction rune
	var nextRune rune

	// Find starting position
	found := false
	for col := 0; col <= colMax; col++ {
		if found {
			break
		}
		for row := 0; row <= rowMax; row++ {
			r := data[row][col]
			if r != OBSTACLE && r != EMPTY {
				found = true
				currCol = col
				currRow = row
				direction = r
				data[row][col] = VISITED
				break
			}
		}
	}

	// Walk and Mark
	for {
		if direction == '^' {
			nextRow = currRow - 1
			nextCol = currCol
		}
		if direction == '>' {
			nextRow = currRow
			nextCol = currCol + 1
		}
		if direction == 'v' {
			nextRow = currRow + 1
			nextCol = currCol
		}
		if direction == '<' {
			nextRow = currRow
			nextCol = currCol - 1
		}

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
	for col := 0; col <= colMax; col++ {
		for row := 0; row <= rowMax; row++ {
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
		direction = '>'
	case '>':
		direction = 'v'
	case 'v':
		direction = '<'
	case '<':
		direction = '^'
	default:
		log.Fatal("unknown direction")
	}

	return direction
}

func solvePart2(inputFile string) int {
	data := loadDaySixData(inputFile)

	OBSTACLE := '#'
	EMPTY := '.'

	colMax := len(data[0]) - 1
	rowMax := len(data) - 1

	sum := 0

	var startCol int
	var startRow int
	var startDirection rune

	// Find starting position
	found := false
	for col := 0; col <= colMax; col++ {
		if found {
			break
		}
		for row := 0; row <= rowMax; row++ {
			r := data[row][col]
			if r != OBSTACLE && r != EMPTY {
				found = true
				startCol = col
				startRow = row
				startDirection = r
				break
			}
		}
	}

	for c := 0; c <= colMax; c++ {
		for r := 0; r <= rowMax; r++ {
			if data[r][c] != EMPTY {
				continue
			}

			data[r][c] = OBSTACLE
			track := make([][][]rune, rowMax+1)
			for i := 0; i <= colMax; i++ {
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
				if rotations >= 3 {
					sum++
					break
				}

				if direction == '^' {
					nextRow = currRow - 1
					nextCol = currCol
				}
				if direction == '>' {
					nextRow = currRow
					nextCol = currCol + 1
				}
				if direction == 'v' {
					nextRow = currRow + 1
					nextCol = currCol
				}
				if direction == '<' {
					nextRow = currRow
					nextCol = currCol - 1
				}

				if 0 > nextCol || nextCol > colMax || 0 > nextRow || nextRow > rowMax {
					break
				}

				nextRune = data[nextRow][nextCol]
				if nextRune == '#' {
					rotations++
					direction = rotateGuard(direction)
					continue
				}

				rotations = 0

				for _, r := range track[currRow][currCol] {
					if direction == r {
						loop = true
						break
					}
				}

				if loop {
					sum++
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
