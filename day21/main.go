package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coor struct {
	x int
	y int
}

var numericgrid = map[rune]coor{
	'0': {1, 3},
	'A': {2, 3},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
}

var movegrid = map[rune]coor{
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

func solvePart1(inputFile string) int {
	data := loadDayData(inputFile)

	complexity := 0
	for _, code := range data {
		numeric, err := strconv.Atoi(strings.Replace(code, "A", "", -1))
		if err != nil {
			log.Fatal(err)
		}
		actions := pressAllTheButtons([]rune(code), 1)
		complexity += numeric * actions
		fmt.Println(actions)
	}
	return complexity
}

func pressAllTheButtons(code []rune, chain int) int {
	r1pos := coor{2, 3}
	r2pos := coor{2, 0}
	r3pos := coor{2, 0}

	r1moveseq := []rune{}
	for _, button := range code {
		r1target := numericgrid[button]
		if r1target != r1pos {
			if r1pos.x > 0 || r1target.y < 3 {
				for r1target.y > r1pos.y {
					r1moveseq = append(r1moveseq, 'v')
					r1pos.y++
				}
			}
			if r1pos.y < 3 || r1target.x > 0 {
				for r1target.x < r1pos.x {
					r1moveseq = append(r1moveseq, '<')
					r1pos.x--
				}
			}
			for r1target.x > r1pos.x {
				r1moveseq = append(r1moveseq, '>')
				r1pos.x++
			}
			for r1target.y < r1pos.y {
				r1moveseq = append(r1moveseq, '^')
				r1pos.y--
			}
			for r1target.y > r1pos.y {
				r1moveseq = append(r1moveseq, 'v')
				r1pos.y++
			}
			for r1target.x < r1pos.x {
				r1moveseq = append(r1moveseq, '<')
				r1pos.x--
			}
		}
		r1moveseq = append(r1moveseq, 'A')
	}

	moves := 0
	for _, move := range r1moveseq {
		targetseq := []rune{move}
		for r := 0; r < chain; r++ {
			r2pos.x = 2
			r2pos.y = 0
			nextseq := make([]rune, 0, len(targetseq)*3)
			for _, button := range targetseq {
				target := movegrid[button]
				if target != r2pos {
					if r2pos.y == 1 || target.x > 0 {
						for target.x < r2pos.x {
							nextseq = append(nextseq, '<')
							r2pos.x--
						}
					}
					if r2pos.x > 0 {
						for target.y < r2pos.y {
							nextseq = append(nextseq, '^')
							r2pos.y--
						}
					}
					for target.x > r2pos.x {
						nextseq = append(nextseq, '>')
						r2pos.x++
					}
					for target.y > r2pos.y {
						nextseq = append(nextseq, 'v')
						r2pos.y++
					}
					for target.x < r2pos.x {
						nextseq = append(nextseq, '<')
						r2pos.x--
					}
					for target.y < r2pos.y {
						nextseq = append(nextseq, '^')
						r2pos.y--
					}
				}
				nextseq = append(nextseq, 'A')
			}

			targetseq = nextseq
		}

		for _, button := range targetseq {
			r3target := movegrid[button]
			if r3target != r3pos {
				for r3target.x > r3pos.x {
					moves++
					r3pos.x++
				}
				for r3target.y < r3pos.y {
					moves++
					r3pos.y--
				}
				for r3target.y > r3pos.y {
					moves++
					r3pos.y++
				}
				for r3target.x < r3pos.x {
					moves++

					r3pos.x--
				}
			}
			moves++
		}
	}

	return moves

}

func solvePart2(inputFile string) int {
	data := loadDayData(inputFile)

	complexity := 0
	for _, code := range data {
		numeric, err := strconv.Atoi(strings.Replace(code, "A", "", -1))
		if err != nil {
			log.Fatal(err)
		}
		actions := pressAllTheButtons([]rune(code), 1)
		complexity += numeric * actions
		fmt.Println(actions)
	}
	return complexity
}

func loadDayData(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data []string
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
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
