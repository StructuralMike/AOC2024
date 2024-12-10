package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solvePart1(inputFile string) int {
	GOAL := 9
	GRADIENTMIN := 1
	GRADIENTMAX := 1

	grid := loadDayTenData(inputFile)

	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 0 {
				found := [][]int{}
				sum += len(dfsPeaks(row, col, GOAL, GRADIENTMIN, GRADIENTMAX, grid, found))
			}
		}
	}
	return sum
}

func dfsPeaks(rowStart int, colStart int, goal int, gradientMin int, gradientMax int, grid [][]int, found [][]int) [][]int {
	paths := [][]int{
		{rowStart - 1, colStart},
		{rowStart + 1, colStart},
		{rowStart, colStart - 1},
		{rowStart, colStart + 1},
	}

	for _, path := range paths {
		row := path[0]
		col := path[1]
		if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
			continue
		}
		next := grid[row][col]
		gradient := next - grid[rowStart][colStart]
		if gradient < gradientMin || gradient > gradientMax {
			continue
		}
		if next == goal {
			newPeak := true
			for _, seen := range found {
				if seen[0] == row && seen[1] == col {
					newPeak = false
					break
				}
			}
			if newPeak {
				found = append(found, []int{row, col})
			}
		} else {
			found = dfsPeaks(row, col, goal, gradientMin, gradientMax, grid, found)
		}
	}

	return found

}

func solvePart2(inputFile string) int {
	GOAL := 9
	GRADIENTMIN := 1
	GRADIENTMAX := 1

	grid := loadDayTenData(inputFile)

	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 0 {
				found := [][]int{}
				sum += len(dfsPeaks2(row, col, GOAL, GRADIENTMIN, GRADIENTMAX, grid, found))
			}
		}
	}
	return sum
}

func dfsPeaks2(rowStart int, colStart int, goal int, gradientMin int, gradientMax int, grid [][]int, found [][]int) [][]int {
	paths := [][]int{
		{rowStart - 1, colStart},
		{rowStart + 1, colStart},
		{rowStart, colStart - 1},
		{rowStart, colStart + 1},
	}

	for _, path := range paths {
		row := path[0]
		col := path[1]
		if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
			continue
		}
		next := grid[row][col]
		gradient := next - grid[rowStart][colStart]
		if gradient < gradientMin || gradient > gradientMax {
			continue
		}
		if next == goal {
			newPeak := true
			if newPeak {
				found = append(found, []int{row, col})
			}
		} else {
			found = dfsPeaks2(row, col, goal, gradientMin, gradientMax, grid, found)
		}
	}

	return found

}

func loadDayTenData(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, val := range line {
			num, err := strconv.Atoi(string(val))
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data

}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
