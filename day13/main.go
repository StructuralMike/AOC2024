package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type claw struct {
	A     [2]int
	B     [2]int
	Prize [2]int
}

func solvePart1(inputFile string) int {
	data := loadDayData(inputFile)

	price := 0
	for _, claw := range data {
		A := claw.A[0]
		B := claw.A[1]
		C := claw.B[0]
		D := claw.B[1]
		K := claw.Prize[0]
		L := claw.Prize[1]

		bestPrice := 501
		for i := 0; i <= 100; i++ {
			posx := A * i
			posy := B * i
			for j := 0; j <= 100; j++ {
				if i*3+j > bestPrice {
					break
				}
				newposx := posx + C*j
				newposy := posy + D*j
				if newposx > K || newposy > L {
					break
				}
				if newposx == K && newposy == L {
					bestPrice = i*3 + j
					break
				}
			}
		}
		if bestPrice < 501 {
			price += bestPrice
		}
	}

	return price
}

func solvePart2(inputFile string) int {
	data := loadDayData(inputFile)

	price := 0
	for _, claw := range data {
		A := claw.A[0]
		B := claw.A[1]
		C := claw.B[0]
		D := claw.B[1]
		K := claw.Prize[0] + 10000000000000
		L := claw.Prize[1] + 10000000000000

	}

	return price
}

func loadDayData(inputFile string) []claw {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data []claw
	var newClaw claw
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "Button A"):
			newClaw = claw{
				A:     [2]int{0, 0},
				B:     [2]int{0, 0},
				Prize: [2]int{0, 0},
			}
			ax, err := strconv.Atoi(strings.Split(strings.Split(line, "+")[1], ",")[0])
			if err != nil {
				log.Fatal(err)
			}
			newClaw.A[0] = ax
			ay, err := strconv.Atoi(strings.Split(line, "+")[2])
			if err != nil {
				log.Fatal(err)
			}
			newClaw.A[1] = ay
		case strings.HasPrefix(line, "Button B"):
			bx, err := strconv.Atoi(strings.Split(strings.Split(line, "+")[1], ",")[0])
			if err != nil {
				log.Fatal(err)
			}
			newClaw.B[0] = bx
			by, err := strconv.Atoi(strings.Split(line, "+")[2])
			if err != nil {
				log.Fatal(err)
			}
			newClaw.B[1] = by

		case strings.HasPrefix(line, "Prize"):
			px, err := strconv.Atoi(strings.Split(strings.Split(line, "=")[1], ",")[0])
			if err != nil {
				log.Fatal(err)
			}
			newClaw.Prize[0] = px
			py, err := strconv.Atoi(strings.Split(line, "=")[2])
			if err != nil {
				log.Fatal(err)
			}
			newClaw.Prize[1] = py
		default:
			data = append(data, newClaw)
		}
	}
	data = append(data, newClaw)

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
