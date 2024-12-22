package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solvePart1(inputFile string) int {
	data := loadDayData(inputFile)

	N := 2000

	sum := 0
	for _, init := range data {
		secret := init
		for i := 0; i < N; i++ {
			secret = ((secret << 6) ^ secret) & (16777216 - 1)
			secret = ((secret >> 5) ^ secret) & (16777216 - 1)
			secret = ((secret << 11) ^ secret) & (16777216 - 1)
		}
		sum += secret
	}

	return sum
}

func solvePart2(inputFile string) int {
	buyers := loadDayData(inputFile)

	N := 2000

	var bigMap = make(map[string]int)
	for _, init := range buyers {
		var sequenceMap = make(map[string]bool)
		var deltas []int

		secret := init
		previousDigit := getLastDigit(secret)

		for i := 0; i < N; i++ {
			secret = getNextSecret(secret)
			currentDigit := getLastDigit(secret)
			delta := currentDigit - previousDigit
			deltas = append(deltas, delta)
			previousDigit = currentDigit
			if i < 3 {
				continue
			}
			sequence := intsToString(deltas[len(deltas)-4:])
			if !sequenceMap[sequence] {
				sequenceMap[sequence] = true
				bigMap[sequence] += currentDigit
			}
		}
	}

	bestPrice := -10
	for _, price := range bigMap {
		if price > bestPrice {
			bestPrice = price
		}
	}

	return bestPrice
}

func intsToString(arr []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(arr), " "), ""), "[]")
}

func getLastDigit(num int) int {
	digits := strconv.Itoa(num)
	digit, err := strconv.Atoi(digits[len(digits)-1:])
	if err != nil {
		log.Fatal(err)
	}
	return digit
}

func getNextSecret(num int) int {
	const mask = (1 << 24) - 1
	num = ((num << 6) ^ num) & mask
	num = ((num >> 5) ^ num) & mask
	num = ((num << 11) ^ num) & mask
	return num
}

func loadDayData(inputFile string) []int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data []int
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("sample_input.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
