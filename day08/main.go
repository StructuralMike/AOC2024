package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solvePart1(inputFile string) int {
	grid := loadDaySevenData(inputFile)

	rowCount := len(grid)
	colCount := len(grid[0])

	antennas := make(map[rune][][2]int)

	for row, arr := range grid {
		for col, spot := range arr {
			if spot != '.' {
				location := [2]int{row, col}
				antennas[spot] = append(antennas[spot], location)
			}
		}
	}

	antinodes := make([][]rune, rowCount)
	for i := range antinodes {
		antinodes[i] = make([]rune, colCount)
	}

	for _, locations := range antennas {
		for a := 0; a < len(locations)-1; a++ {
			for b := a + 1; b < len(locations); b++ {
				arow := locations[a][0]
				acol := locations[a][1]
				brow := locations[b][0]
				bcol := locations[b][1]

				drow := brow - arow
				dcol := bcol - acol

				anti1row := arow - drow
				anti1col := acol - dcol
				anti2row := brow + drow
				anti2col := bcol + dcol

				if anti1row >= 0 && anti1row < rowCount && anti1col >= 0 && anti1col < colCount {
					antinodes[anti1row][anti1col] = '#'
				}

				if anti2row >= 0 && anti2row < rowCount && anti2col >= 0 && anti2col < colCount {
					antinodes[anti2row][anti2col] = '#'
				}
			}
		}
	}

	sum := 0
	for _, row := range antinodes {
		for _, col := range row {
			if col == '#' {
				sum++
			}
		}
	}

	return sum
}

func solvePart2(inputFile string) int {
	grid := loadDaySevenData(inputFile)

	rowCount := len(grid)
	colCount := len(grid[0])

	antennas := make(map[rune][][2]int)

	for row, arr := range grid {
		for col, spot := range arr {
			if spot != '.' {
				location := [2]int{row, col}
				antennas[spot] = append(antennas[spot], location)
			}
		}
	}

	antinodes := make([][]rune, rowCount)
	for i := range antinodes {
		antinodes[i] = make([]rune, colCount)
	}

	for _, locations := range antennas {
		for a := 0; a < len(locations)-1; a++ {
			for b := a + 1; b < len(locations); b++ {
				arow := locations[a][0]
				acol := locations[a][1]
				brow := locations[b][0]
				bcol := locations[b][1]

				antinodes[arow][acol] = '#'

				drow := brow - arow
				dcol := bcol - acol

				antirow := arow - drow
				anticol := acol - dcol

				for antirow >= 0 && antirow < rowCount && anticol >= 0 && anticol < colCount {
					antinodes[antirow][anticol] = '#'
					antirow = antirow - drow
					anticol = anticol - dcol
				}

				antirow = arow + drow
				anticol = acol + dcol

				for antirow >= 0 && antirow < rowCount && anticol >= 0 && anticol < colCount {
					antinodes[antirow][anticol] = '#'
					antirow = antirow + drow
					anticol = anticol + dcol
				}
			}
		}
	}

	sum := 0
	for _, row := range antinodes {
		for _, col := range row {
			if col == '#' {
				sum++
			}
		}
	}

	return sum
}

func loadDaySevenData(inputFile string) [][]rune {
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
