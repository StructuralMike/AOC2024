package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/structuralmike/adventofcode2024/utils"
)

func solvePart1(inputFile string) int {
	data1, data2 := loadDayOneData(inputFile)

	slices.Sort(data1)
	slices.Sort(data2)

	sumdiff := 0
	for i := range data1 {
		var diff = utils.Abs(data1[i] - data2[i])
		sumdiff += diff
	}

	return sumdiff
}

func solvePart2(inputFile string) int {
	data1, data2 := loadDayOneData(inputFile)

	counts := make(map[int]int)
	for _, n := range data2 {
		counts[n] += 1
	}

	sim := 0
	for _, n := range data1 {
		sim += n * counts[n]
	}

	return sim
}

func loadDayOneData(inputFile string) ([]int, []int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data1, data2 []int

	for scanner.Scan() {
		line := scanner.Text()

		var num1, num2 int
		n, err := fmt.Sscanf(line, "%d,%d", &num1, &num2)
		if err != nil || n != 2 {
			log.Fatalf("Failed to parse line '%s': %v", line, err)
		}

		data1 = append(data1, num1)
		data2 = append(data2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data1, data2
}

func main() {
	fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
