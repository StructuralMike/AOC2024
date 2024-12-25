package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solvePart1(inputFile string) int {
	locks, keys := loadDayData(inputFile)

	pairs := 0

	for _, lock := range locks {
		for _, key := range keys {
			isPair := true
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 5 {
					isPair = false
					break
				}
			}
			if isPair {
				pairs++
			}
		}
	}

	return pairs
}

// func solvePart2(inputFile string) int {
// 	return 0
// }

func loadDayData(inputFile string) ([][5]int, [][5]int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var locks [][5]int
	var keys [][5]int
	isNew := true
	var isKey bool
	var key [5]int
	var lock [5]int
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 2 {
			isNew = true
			if isKey {
				keys = append(keys, key)
			} else {
				locks = append(locks, lock)
			}
			continue
		}

		if isNew {
			isKey = line != "#####"
			for i := 0; i < 5; i++ {
				key[i] = -1
				lock[i] = -1
			}
			isNew = false
		}

		chars := []rune(line)
		for i, char := range chars {
			if char == '#' {
				if isKey {
					key[i]++
				} else {
					lock[i]++
				}
			}
		}
	}

	if isKey {
		keys = append(keys, key)
	} else {
		locks = append(locks, lock)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return locks, keys
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("input1.txt"))
}
