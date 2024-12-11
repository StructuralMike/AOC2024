package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solvePart1(data []string, blinks int) int {
	stones := data

	for i := 0; i < blinks; i++ {
		newStones := []string{}
		for _, stone := range stones {
			if stone == "0" {
				newStones = append(newStones, "1")
			} else if len(stone)%2 == 0 {
				left, err := strconv.Atoi(stone[0 : len(stone)/2])
				if err != nil {
					log.Fatal(err)
				}
				newStones = append(newStones, strconv.Itoa(left))
				right, err := strconv.Atoi(stone[len(stone)/2:])
				if err != nil {
					log.Fatal(err)
				}
				newStones = append(newStones, strconv.Itoa(right))
			} else {
				newNum, err := strconv.Atoi(stone)
				if err != nil {
					log.Fatal(err)
				}
				newStone := strconv.Itoa(newNum * 2024)
				newStones = append(newStones, newStone)
			}
		}
		stones = newStones[0:]
	}

	return len(stones)
}

func solvePart2(data []string, blinks int) int {
	stoneCount := 0
	for _, stone := range data[0] {
		stoneCount += solvePart1([]string{string(stone)}, blinks)
		fmt.Println(stoneCount)
	}

	return stoneCount
}

func loadDayElevenData(inputFile string) [][]string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]string
	for scanner.Scan() {
		line := scanner.Text()

		vars := strings.Split(line, " ")
		data = append(data, vars)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data

}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1(loadDayElevenData("input1.txt")[0], 25))
	fmt.Println(solvePart2(loadDayElevenData("input1.txt")[0], 75))
}
