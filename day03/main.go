package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solvePart1(inputFile string) int {
	return 0
}

func solvePart2(inputFile string) int {
	return 0
}

func loadDayTwoData(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 100 {
			continue
		}
	}
	return
}

func main() {
	fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
