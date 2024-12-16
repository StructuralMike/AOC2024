package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type coor struct {
	x int
	y int
}

const Exit = 'E'
const Start = 'S'
const Wall = '#'
const Empty = '.'
const Visited = '-'

var Dirs = map[rune]coor{
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

func (c1 coor) Add(c2 coor) coor {
	return coor{
		x: c1.x + c2.x,
		y: c1.y + c2.y,
	}
}

func solvePart1(inputFile string) int {
	grid := loadDayData(inputFile)
	printGrid(grid)

	var deer coor
	var goal coor
	// Find start and exit
	for y, row := range grid {
		for x, col := range row {
			if col == Start {
				deer = coor{x: x, y: y}
				grid[y][x] = '>'
			}
			if col == Exit {
				goal = coor{x: x, y: y}
			}
		}
	}

	grid = findShortestPath(grid, deer, goal)
	printGrid(grid)

	return 0
}

func findShortestPath(grid [][]rune, start coor, goal coor) [][]rune {
	position := grid[start.y][start.x]

	// Try the same direction
	newLoc := start.Add(Dirs[position])
	newPos := grid[newLoc.y][newLoc.x]
	if newPos == Wall {
		return grid
	}

	if newPos == Empty {
		grid[newLoc.y][newLoc.x] = position
		grid = findShortestPath(grid, newLoc, goal)
	}

	return grid
}

// func solvePart2(inputFile string) int {
// 	return 0
// }

func printGrid(grid [][]rune) {
	// Bright ANSI color codes
	black := "\033[90m"
	red := "\033[91m"
	lightRed := "\033[31"
	green := "\033[92m"
	yellow := "\033[93m"
	reset := "\033[0m"

	for _, row := range grid {
		var sb strings.Builder
		for _, ch := range row {
			switch ch {
			case Empty:
				sb.WriteString(black)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			case Start:
				sb.WriteString(red)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			case Exit:
				sb.WriteString(yellow)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			case Wall:
				sb.WriteString(green)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			case Visited:
				sb.WriteString(lightRed)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			default:
				sb.WriteRune(ch)
			}
		}

		fmt.Println(sb.String())
	}
}

func loadDayData(inputFile string) [][]rune {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

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
	fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("input1.txt"))
}
