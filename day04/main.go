package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func solvePart1(inputFile string) int {
	data := loadDayFourData(inputFile)

	sum := 0

	re := regexp.MustCompile(`XMAS|SAMX`)
	for _, line := range data {
		matches := re.FindAllString(line, -1)
		if matches != nil {
			sum += len(matches)
		}
	}
	return sum
}

func solvePart2(inputFile string) int {
	return 0
}

func loadDayFourData(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []string
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	return data
}

func main() {
	fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("input1.txt"))
}
