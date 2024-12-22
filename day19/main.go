package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
)

func solvePart1(inputFile string) int {
	towels, patterns := loadDayData(inputFile)
	sort.Slice(towels, func(i, j int) bool {
		return len(towels[i]) < len(towels[j])
	})

	cnt := 0
	for _, pattern := range patterns {

		var subTowels []string
		var singles []string
		for _, towel := range towels {
			if len(towel) == 1 && strings.Contains(pattern, towel) {
				singles = append(singles, towel)
				subTowels = append(subTowels, towel)
				continue
			}

			if len(towel) > 1 {
				chars := strings.Split(towel, "")
				for _, char := range chars {
					if !slices.Contains(singles, char) {
						subTowels = append(subTowels, towel)
						break
					}
				}
			}

		}

		if canBuildPattern(pattern, subTowels) {
			cnt++
			fmt.Println("true:  " + pattern)
		} else {
			fmt.Println("false: " + pattern)
		}
	}
	return cnt
}

func canBuildPattern(pattern string, towels []string) bool {
	if pattern == "" || slices.Contains(towels, pattern) {
		return true
	}
	for _, towel := range towels {
		if !strings.Contains(pattern, towel) {
			continue
		}
		newPat := strings.Split(pattern, towel)
		canBuild := true
		for _, subPat := range newPat {
			if subPat == "" {
				continue
			}
			if !canBuildPattern(subPat, towels) {
				canBuild = false
				break
			}
		}
		if canBuild {
			return true
		}
	}
	return false
}

// func solvePart2(inputFile string) int {
// 	return 0
// }

func loadDayData(inputFile string) ([]string, []string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var patterns []string
	var towels []string
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineCount == 0 {
			towels = strings.Split(strings.Replace(line, " ", "", -1), ",")
		}
		lineCount++
		if lineCount < 3 {
			continue
		}
		patterns = append(patterns, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return towels, patterns
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("input1.txt"))
}
