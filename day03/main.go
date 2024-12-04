package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solvePart1(inputFile string) int {
	data := loadDayThreeData(inputFile)

	sum := 0
	for _, arr := range data {
		for _, str := range arr {
			sum += parseMultiplication(str)
		}
	}
	return sum
}

func parseMultiplication(str string) int {
	vals := strings.Split(strings.ReplaceAll(strings.ReplaceAll(str, "mul(", ""), ")", ""), ",")
	if len(vals) != 2 {
		log.Fatal("Parsing error")
	}
	val1, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Fatal(err)
	}

	val2, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Fatal(err)
	}

	return val1 * val2
}

func solvePart2(inputFile string) int {
	data := loadDayThreeDataTwo(inputFile)

	sum := 0
	do := true
	for _, arr := range data {
		for _, str := range arr {
			if str == "do()" {
				do = true
				continue
			}

			if str == "don't()" {
				do = false
				continue
			}

			if do {
				sum += parseMultiplication(str)
			}
		}
	}
	return sum
}

func loadDayThreeData(inputFile string) [][]string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`mul\(([\d]+),(\d+)\)`)
	var data [][]string

	for scanner.Scan() {
		line := scanner.Text()

		data = append(data, re.FindAllString(line, -1))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func loadDayThreeDataTwo(inputFile string) [][]string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(([\d]+),(\d+)\)`)
	var data [][]string

	for scanner.Scan() {
		line := scanner.Text()

		data = append(data, re.FindAllString(line, -1))
	}

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
