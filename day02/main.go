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
	matrix := loadDayTwoData(inputFile)

	safe := 0
	for _, arr := range matrix {
		if isSafe(arr) {
			safe += 1
		}
	}

	return safe
}

func isSafe(arr []int) bool {
	if len(arr) < 3 {
		return true
	}

	var prev int
	var asc bool
	var diff int
	unsafe := false
	for j, num := range arr {
		if j == 0 {
			prev = num
			continue
		}

		diff = num - prev
		prev = num
		if diff == 0 || diff > 3 || diff < -3 {
			unsafe = true
			break
		}

		if j == 1 {
			asc = diff > 0
			continue
		}

		if !((diff > 0 && asc) || (diff < 0 && !asc)) {
			unsafe = true
			break
		}
	}
	return !unsafe
}

func solvePart2(inputFile string) int {
	matrix := loadDayTwoData(inputFile)

	safe := 0
	for _, arr := range matrix {
		if isSafe(arr) {
			safe += 1
			continue
		}
		for i := 0; i < len(arr); i++ {
			if isSafeDampened(arr, i) {
				safe += 1
				break
			}
		}
	}

	return safe
}

func isSafeDampened(arr []int, dampIndex int) bool {
	if len(arr) < 3 {
		return true
	}

	var prev int
	var asc bool
	var diff int
	unsafe := false
	for j, num := range arr {
		if j == dampIndex {
			continue
		}

		if j == 0 || (dampIndex == 0 && j == 1) {
			prev = num
			continue
		}

		diff = num - prev
		prev = num
		if diff == 0 || diff > 3 || diff < -3 {
			unsafe = true
			break
		}

		if j == 1 || (j == 2 && (dampIndex == 1 || dampIndex == 0)) {
			asc = diff > 0
			continue
		}

		if (asc && diff < 0) || (!asc && diff > 0) {
			unsafe = true
			break
		}
	}

	return !unsafe
}

func loadDayTwoData(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		var arr []int
		for _, str := range line {
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, num)
		}

		matrix = append(matrix, arr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func main() {
	fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
