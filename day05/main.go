package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/structuralmike/adventofcode2024/utils"
)

func solvePart1(inputFile string) int {
	rules, updates := loadDayFiveData(inputFile)

	sum := 0

	for _, arr := range updates {
		var order []int
		valid := true
		for _, num := range arr {
			if !valid {
				break
			}

			blockers := rules[num]
			for _, blocker := range blockers {
				if slices.Contains(order, blocker) {
					valid = false
					break
				}
			}
			order = append(order, num)
		}

		if valid {
			sum += arr[int((len(arr)-1)/2)]
		}
	}

	return sum
}

func solvePart2(inputFile string) int {
	rules, updates := loadDayFiveData(inputFile)

	sum := 0

	for _, arr := range updates {
		var order []int
		valid := true
		for _, num := range arr {
			if !valid {
				break
			}

			blockers := rules[num]
			for _, blocker := range blockers {
				if slices.Contains(order, blocker) {
					valid = false
					break
				}
			}
			order = append(order, num)
		}

		if !valid {
			newOrder := make([]int, len(arr))

			for _, num := range arr {
				pos := len(arr) - 1
				for _, blocker := range rules[num] {
					if slices.Contains(arr, blocker) {
						pos -= 1
					}
				}
				newOrder[pos] = num
			}

			sum += newOrder[int((len(newOrder)-1)/2)]
		}
	}

	return sum
}

func loadDayFiveData(inputFile string) (map[int][]int, [][]int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[int][]int)
	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			nums := utils.StringsToInts(strings.Split(line, "|"))
			if len(nums) != 2 {
				log.Fatal("Rules format is wrong!")
			}
			rules[nums[0]] = append(rules[nums[0]], nums[1])

		} else if strings.Contains(line, ",") {
			nums := utils.StringsToInts(strings.Split(line, ","))
			updates = append(updates, nums)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules, updates
}

func main() {
	// fmt.Println(solvePart2("sample_input.txt"))
	fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
