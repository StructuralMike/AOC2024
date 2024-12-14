package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type coor struct {
	x int
	y int
}

type robot struct {
	pos coor
	vel coor
}

func solvePart1(inputFile string, gridx int, gridy int) int {
	robots := loadDayData(inputFile)

	SIM := 100

	quadrants := map[string]int{
		"tl": 0,
		"tr": 0,
		"bl": 0,
		"br": 0,
	}

	midx := gridx/2 + 1
	midy := gridy/2 + 1

	for _, robot := range robots {
		x := robot.pos.x + robot.vel.x*SIM
		y := robot.pos.y + robot.vel.y*SIM

		x = x % gridx
		y = y % gridy

		if x < 0 {
			x = x + gridx
		}

		if y < 0 {
			y = y + gridy
		}

		x++
		y++

		if x == midx || y == midy {
			continue
		}

		if x < midx {
			if y < midy {
				quadrants["bl"]++
			} else {
				quadrants["tl"]++
			}
		} else {
			if y < midy {
				quadrants["br"]++
			} else {
				quadrants["tr"]++
			}
		}
	}

	return quadrants["tl"] * quadrants["tr"] * quadrants["bl"] * quadrants["br"]
}

func solvePart2(inputFile string, gridx int, gridy int) int {
	robots := loadDayData(inputFile)

	SIM := 100000000

	minFactor := 15
	for i := 2000; i <= SIM; i++ {
		grid := make([][]int, gridx)
		for i := range grid {
			grid[i] = make([]int, gridy)
		}
		for _, robot := range robots {
			x := robot.pos.x + robot.vel.x*i
			y := robot.pos.y + robot.vel.y*i

			x = x % gridx
			y = y % gridy

			if x < 0 {
				x = x + gridx
			}

			if y < 0 {
				y = y + gridy
			}

			grid[x][y]++
		}

		factor := 0
		for k := 0; k < 25; k++ {
			for l := 0; l < 25; l++ {
				factor += grid[k][l]
			}
		}
		if factor < minFactor {
			//			fmt.Print("\033[H\033[2J")
			fmt.Println(i)
			printGrid(grid)
			time.Sleep(1000 * time.Millisecond)
		}

	}

	return 0
}

func printGrid(grid [][]int) {
	// ANSI color codes
	green := "\033[92m"
	reset := "\033[0m"

	for i := 0; i < len(grid); i++ {
		line := arrayToString(grid[i])
		// Replace zeros with '.'
		line = strings.Replace(line, "0", ".", -1)

		var sb strings.Builder
		for _, ch := range line {
			if ch != '.' {
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

func loadDayData(inputFile string) []robot {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data []robot
	for scanner.Scan() {
		line := scanner.Text()

		var newRobot robot
		p := strings.Split(strings.Split(strings.Split(line, " ")[0], "=")[1], ",")
		v := strings.Split(strings.Split(strings.Split(line, " ")[1], "=")[1], ",")

		errs := make([]error, 4)
		newRobot.pos.x, errs[0] = strconv.Atoi(p[0])
		newRobot.pos.y, errs[1] = strconv.Atoi(p[1])
		newRobot.vel.x, errs[2] = strconv.Atoi(v[0])
		newRobot.vel.y, errs[3] = strconv.Atoi(v[1])

		for _, err := range errs {
			if err != nil {
				log.Fatal(err)
			}
		}

		data = append(data, newRobot)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt", 11, 7))
	// fmt.Println(solvePart1("input1.txt", 101, 103))
	fmt.Println(solvePart2("input1.txt", 101, 103))
}
