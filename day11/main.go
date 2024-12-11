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

func solvePart1(inputFile string) int {
	data := loadDayElevenData(inputFile)
	BLINKS := 25

	stoneCount := 0
	mem := make(map[int]map[string]int)
	var count int
	for _, stone := range data[0] {
		count, mem = blinkStones(stone, BLINKS, mem)
		stoneCount += count
	}

	return stoneCount
}

func solvePart2(inputfile string) int {
	data := loadDayElevenData(inputfile)
	BLINKS := 75

	stoneCount := 0
	mem := make(map[int]map[string]int)
	var count int
	for _, stone := range data[0] {
		count, mem = blinkStones(stone, BLINKS, mem)
		stoneCount += count
	}

	return stoneCount
}

func blinkStones(stone string, blinks int, mem map[int]map[string]int) (int, map[int]map[string]int) {
	if blinks == 0 {
		return 1, mem
	}
	final := 0
	newStones := blinkStone(string(stone))
	var count int
	for _, newStone := range newStones {
		if mem[blinks] == nil || mem[blinks][newStone] == 0 {
			count, mem = blinkStones(newStone, blinks-1, mem)
			final += count
			if mem[blinks] == nil {
				mem[blinks] = make(map[string]int)
			}
			mem[blinks][newStone] = count
		} else {
			final += mem[blinks][newStone]
		}
	}

	return final, mem
}

func blinkStone(stone string) []string {
	newStones := []string{}
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
	return newStones
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
	fmt.Println(solvePart1("input1.txt"))
	// Measure time
	start := time.Now()
	fmt.Println(solvePart2("input1.txt"))
	fmt.Println(time.Since(start))
}
