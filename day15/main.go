package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type coor struct {
	x int
	y int
}

func solvePart1(inputFile string) int {
	grid, moves := loadDayData(inputFile)

	// Find robot
	var robot coor
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '@' {
				robot = coor{j, i}
				break
			}
		}
	}

	dirs := map[rune]coor{
		'^': {
			x: 0,
			y: -1,
		},
		'>': {
			x: 1,
			y: 0,
		},
		'v': {
			x: 0,
			y: 1,
		},
		'<': {
			x: -1,
			y: 0,
		},
	}

	var newPos coor
	var newBox coor
	var occupant rune
	for _, move := range moves {
		// time.Sleep(350 * time.Millisecond)
		// fmt.Print("\033[H\033[2J")
		for _, row := range grid {
			fmt.Println(string(row))
		}

		newPos.x = robot.x + dirs[move].x
		newPos.y = robot.y + dirs[move].y

		occupant = grid[newPos.y][newPos.x]

		// Wall
		if occupant == '#' {
			continue
		}

		// Unoccupied
		if occupant == '.' {
			grid[robot.y][robot.x] = '.'
			grid[newPos.y][newPos.x] = '@'
			robot.x = newPos.x
			robot.y = newPos.y
			continue
		}

		newBox = newPos

		// Box
		for occupant == 'O' {
			newBox.x = newBox.x + dirs[move].x
			newBox.y = newBox.y + dirs[move].y
			occupant = grid[newBox.y][newBox.x]
		}

		if occupant == '#' {
			continue
		}

		// Push
		for newBox != newPos {
			grid[newBox.y][newBox.x] = 'O'
			newBox.x = newBox.x + dirs[move].x*-1
			newBox.y = newBox.y + dirs[move].y*-1
		}

		grid[robot.y][robot.x] = '.'
		grid[newPos.y][newPos.x] = '@'
		robot.x = newPos.x
		robot.y = newPos.y

	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 'O' {
				sum += i*100 + j
			}
		}
	}

	return sum
}

func solvePart2(inputFile string) int {
	grid, moves := loadDayDataPart2(inputFile)

	var robot coor
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '@' {
				robot = coor{j, i}
				break
			}
		}
	}

	dirs := map[rune]coor{
		'^': {
			x: 0,
			y: -1,
		},
		'>': {
			x: 1,
			y: 0,
		},
		'v': {
			x: 0,
			y: 1,
		},
		'<': {
			x: -1,
			y: 0,
		},
	}

	var newPos coor
	var occupant rune
	for _, move := range moves {
		time.Sleep(750 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		for _, row := range grid {
			fmt.Println(string(row))
		}

		newPos.x = robot.x + dirs[move].x
		newPos.y = robot.y + dirs[move].y

		occupant = grid[newPos.y][newPos.x]

		// Immovable object
		if occupant == '#' {
			continue
		}

		// Unoccupied
		if occupant == '.' {
			grid[robot.y][robot.x] = '.'
			grid[newPos.y][newPos.x] = '@'
			robot.x = newPos.x
			robot.y = newPos.y
		}

		// Move boxes

	}

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '[' {
				sum += i*100 + j
			}
		}
	}

	return sum
}

func loadDayData(inputFile string) ([][]rune, []rune) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	var moves []rune
	isGrid := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			isGrid = false
			continue
		}
		if isGrid {
			grid = append(grid, []rune(line))
		} else {
			moves = append(moves, []rune(line)...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid, moves
}

func loadDayDataPart2(inputFile string) ([][]rune, []rune) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	var moves []rune
	isGrid := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			isGrid = false
			continue
		}
		if isGrid {
			var newRow []rune
			for _, pos := range line {
				switch pos {
				case '.':
					newRow = append(newRow, '.')
					newRow = append(newRow, '.')
				case 'O':
					newRow = append(newRow, '[')
					newRow = append(newRow, ']')
				case '#':
					newRow = append(newRow, '#')
					newRow = append(newRow, '#')
				case '@':
					newRow = append(newRow, '@')
					newRow = append(newRow, '.')
				}
			}
			grid = append(grid, newRow)
		} else {
			moves = append(moves, []rune(line)...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid, moves
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
