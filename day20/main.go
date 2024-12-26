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

func (c1 coor) Add(c2 coor) coor {
	return coor{
		x: c1.x + c2.x,
		y: c1.y + c2.y,
	}
}

func solvePart1(inputFile string) int {
	grid, start, end := loadDayData(inputFile)
	printGrid(grid)
	moves := []coor{
		{
			x: 0,
			y: 1,
		},
		{
			x: 0,
			y: -1,
		},
		{
			x: 1,
			y: 0,
		},
		{
			x: -1,
			y: 0,
		},
	}

	// Calculate times
	curr := start
	var newPos coor
	grid[curr.y][curr.x] = 0
	for curr != end {
		for _, move := range moves {
			newPos = curr.Add(move)
			if grid[newPos.y][newPos.x] == -2 {
				grid[newPos.y][newPos.x] = grid[curr.y][curr.x] + 1
				break
			}
		}
		curr = newPos
	}

	cheatTarget := 100
	cheatCount := 0
	var next coor

	for curr != start {
		currval := grid[curr.y][curr.x]
		if currval < cheatTarget {
			break
		}
		for _, move := range moves {
			cheat1 := curr
			cheat1 = cheat1.Add(move)
			if cheat1.y == 0 || cheat1.x == 0 || cheat1.y == len(grid)-1 || cheat1.x == len(grid[0])-1 {
				continue
			}

			if grid[cheat1.y][cheat1.x] == -1 {
				cheat2 := cheat1
				cheat2 = cheat2.Add(move)
				if grid[cheat2.y][cheat2.x] != -1 {
					cheatval := grid[cheat2.y][cheat2.x]
					if (currval - cheatval - 2) >= cheatTarget {
						cheatCount++
					}
				}
			}

			if grid[cheat1.y][cheat1.x] == currval-1 {
				next = cheat1
			}
		}
		curr = next
	}

	return cheatCount
}

// func solvePart2(inputFile string) int {
// 	return 0
// }

func printGrid(grid [][]int) {
	// ANSI color codes
	green := "\033[92m"
	reset := "\033[0m"

	for i := 0; i < len(grid); i++ {
		line := arrayToString(grid[i])
		// Replace -1 with '#'
		line = strings.Replace(line, "-1", "#", -1)
		line = strings.Replace(line, "-2", ".", -1)

		var sb strings.Builder
		for _, ch := range line {
			if ch == '#' {
				// Non-dot characters are non-zero, print them in green
				sb.WriteString(green)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			} else {
				// Dot remains uncolored
				sb.WriteRune(ch)
			}
		}

		fmt.Println(sb.String())
	}
}

func arrayToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]")
}

func loadDayData(inputFile string) ([][]int, coor, coor) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data [][]int
	var start coor
	var end coor
	for scanner.Scan() {
		line := []rune(scanner.Text())
		newLine := []int{}
		for i, char := range line {
			if char == '#' {
				newLine = append(newLine, -1)
			} else {
				newLine = append(newLine, -2)
			}
			if char == 'S' {
				start = coor{
					x: i,
					y: len(data),
				}
			}
			if char == 'E' {
				end = coor{
					x: i,
					y: len(data),
				}
			}
		}
		data = append(data, newLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data, start, end
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("input1.txt"))
}
