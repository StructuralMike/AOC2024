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
	data := loadDayElevenData(inputFile)
	fmt.Println(data)
	return 0
}

// func solvePart2(inputFile string) int {
// 	return 0
// }

func loadDayElevenData(inputFile string) []int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []int
	for scanner.Scan() {
		line := scanner.Text()

		vars := strings.Split(line, " ")
		for _, strvar := range vars {
			num, err := strconv.Atoi(strvar)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, num)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data

}

func main() {
	fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("input1.txt"))
}
