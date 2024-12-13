package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solvePart1(inputFile string) int {
	grid := loadDayData(inputFile)

	total := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			plant := grid[i][j]
			if plant != '.' {
				mapped := make([][]rune, len(grid), len(grid[0]))
				for k := 0; k < len(grid); k++ {
					mapped[k] = make([]rune, len(grid[0]))
				}
				mapped[i][j] = plant
				mapped, grid = findRegion(i, j, plant, mapped, grid)
				total += evalRegion(mapped)
			}
		}
	}

	return total
}

func findRegion(i int, j int, plant rune, mapped [][]rune, grid [][]rune) ([][]rune, [][]rune) {
	directions := [][]int{
		{i, j - 1},
		{i, j + 1},
		{i - 1, j},
		{i + 1, j},
	}

	for _, coor := range directions {
		if coor[0] < 0 || coor[1] < 0 || coor[0] >= len(grid) || coor[1] >= len(grid[0]) {
			continue
		}
		if grid[coor[0]][coor[1]] != plant {
			continue
		}
		if mapped[coor[0]][coor[1]] == plant {
			continue
		}
		mapped[coor[0]][coor[1]] = plant
		grid[coor[0]][coor[1]] = '.'
		mapped, grid = findRegion(coor[0], coor[1], plant, mapped, grid)
	}

	return mapped, grid
}

func evalRegion(mapped [][]rune) int {
	plants := 0
	price := 0

	for i := 0; i < len(mapped); i++ {
		for j := 0; j < len(mapped[0]); j++ {
			plant := mapped[i][j]
			if plant == 0 {
				continue
			}

			plants++

			directions := [][]int{
				{i, j - 1},
				{i, j + 1},
				{i - 1, j},
				{i + 1, j},
			}

			for _, coor := range directions {
				if coor[0] < 0 || coor[1] < 0 || coor[0] >= len(mapped) || coor[1] >= len(mapped[0]) {
					price++
					continue
				}

				neighbourPlant := mapped[coor[0]][coor[1]]
				if neighbourPlant == 0 {
					price++
				}
			}
		}
	}

	return plants * price
}

func solvePart2(inputFile string) int {
	grid := loadDayData(inputFile)

	total := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			plant := grid[i][j]
			if plant != '.' {
				mapped := make([][]rune, len(grid), len(grid[0]))
				for k := 0; k < len(grid); k++ {
					mapped[k] = make([]rune, len(grid[0]))
				}
				mapped[i][j] = plant
				mapped, grid = findRegion(i, j, plant, mapped, grid)
				total += evalRegionDiscount(mapped)
			}
		}
	}

	return total
}

func evalRegionDiscount(mapped [][]rune) int {
	plants := 0
	sides := 0

	// Count plants
	for i := 0; i < len(mapped); i++ {
		for j := 0; j < len(mapped[0]); j++ {
			if mapped[i][j] != 0 {
				plants++
			}
		}
	}

	// Top
	for i := 0; i < len(mapped); i++ {
		section := false
		for j := 0; j < len(mapped[0]); j++ {
			curr := mapped[i][j]
			if !section {
				if curr != 0 {
					if i == 0 || mapped[i-1][j] == 0 {
						section = true
						sides++
					}
				}
			} else if section {
				if curr == 0 {
					section = false
				} else if i != 0 && mapped[i-1][j] != 0 {
					section = false
				}
			}
		}
	}

	// Bottom
	for i := 0; i < len(mapped); i++ {
		section := false
		for j := 0; j < len(mapped[0]); j++ {
			curr := mapped[i][j]
			if !section {
				if curr != 0 {
					if i == len(mapped)-1 || mapped[i+1][j] == 0 {
						section = true
						sides++
					}
				}
			} else if section {
				if curr == 0 {
					section = false
				} else if i != len(mapped)-1 && mapped[i+1][j] != 0 {
					section = false
				}
			}
		}
	}

	// Left
	for j := 0; j < len(mapped[0]); j++ {
		section := false
		for i := 0; i < len(mapped); i++ {
			curr := mapped[i][j]
			if !section {
				if curr != 0 {
					if j == 0 || mapped[i][j-1] == 0 {
						section = true
						sides++
					}
				}
			} else if section {
				if curr == 0 {
					section = false
				} else if j != 0 && mapped[i][j-1] != 0 {
					section = false
				}
			}
		}
	}

	// Right
	for j := 0; j < len(mapped[0]); j++ {
		section := false
		for i := 0; i < len(mapped); i++ {
			curr := mapped[i][j]
			if !section {
				if curr != 0 {
					if j == len(mapped[0])-1 || mapped[i][j+1] == 0 {
						section = true
						sides++
					}
				}
			} else if section {
				if curr == 0 {
					section = false
				} else if j != len(mapped[0])-1 && mapped[i][j+1] != 0 {
					section = false
				}
			}
		}
	}

	return plants * sides
}

func loadDayData(inputFile string) [][]rune {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	data := [][]rune{}
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
	// fmt.Println(solvePart2("sample_input.txt"))
	fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
