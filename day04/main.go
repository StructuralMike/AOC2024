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

	re1 := regexp.MustCompile(`XMAS`)
	for _, line := range data {
		matches := re1.FindAllString(line, -1)
		if matches != nil {
			sum += len(matches)
		}
	}

	re2 := regexp.MustCompile(`SAMX`)
	for _, line := range data {
		matches := re2.FindAllString(line, -1)
		if matches != nil {
			sum += len(matches)
		}
	}

	return sum
}

func solvePart2(inputFile string) int {
	data := loadDayFourDataTwo(inputFile)

	yMax := len(data)
	xMax := len(data[0])

	sum := 0

	for y := 1; y < (yMax - 1); y++ {
		for x := 1; x < (xMax - 1); x++ {
			if string(data[y][x]) != "A" {
				continue
			}
			dr := string(data[y-1][x-1]) + string(data[y+1][x+1])
			if dr != "MS" && dr != "SM" {
				continue
			}
			dl := string(data[y-1][x+1]) + string(data[y+1][x-1])
			if dl != "MS" && dl != "SM" {
				continue
			}
			sum++
		}
	}

	return sum
}

func loadDayFourData(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []string
	var horizontalChars [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
		runes := []rune(line)
		horizontalChars = append(horizontalChars, runes)
	}

	xMax := len(horizontalChars)
	yMax := len(horizontalChars[0])

	for x := 0; x < xMax; x++ {
		var line []rune
		for y := 0; y < yMax; y++ {
			line = append(line, horizontalChars[y][x])
		}
		data = append(data, string(line))
	}

	for i := 0; i < xMax; i++ {
		var line []rune
		x := i
		y := 0
		for x < xMax && y < yMax {
			line = append(line, horizontalChars[y][x])
			x++
			y++
		}
		data = append(data, string(line))
	}

	for i := 0; i < xMax; i++ {
		var line []rune
		x := i
		y := 0
		for x >= 0 && y < yMax {
			line = append(line, horizontalChars[y][x])
			x--
			y++
		}
		data = append(data, string(line))
	}

	for i := 1; i < yMax; i++ {
		var line []rune
		x := 0
		y := i
		for x < xMax && y < yMax {
			line = append(line, horizontalChars[y][x])
			x++
			y++
		}
		data = append(data, string(line))
	}

	for i := 1; i < yMax; i++ {
		var line []rune
		x := xMax - 1
		y := i
		for x >= 0 && y < yMax {
			line = append(line, horizontalChars[y][x])
			x--
			y++
		}
		data = append(data, string(line))
	}

	return data
}

func loadDayFourDataTwo(inputFile string) []string {
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
	fmt.Println(solvePart2("input1.txt"))
}
