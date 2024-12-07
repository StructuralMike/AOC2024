package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	combinations "github.com/mxschmitt/golang-combinations"
	"github.com/structuralmike/adventofcode2024/utils"
)

type calibration struct {
	sum  int
	nums []int
}

func solvePart1(inputFile string) int {
	data := loadDaySevenData(inputFile)

	finalSum := 0

	for _, entry := range data {
		// Try sum all
		tot := utils.SumArray(entry.nums)
		if tot == entry.sum {
			finalSum += entry.sum
			continue
		}

		indexCombinations := combinations.All(utils.IntegerRange(len(entry.nums) - 2))

		// Use multipliers
		for _, indices := range indexCombinations {
			tot := entry.nums[0]
			for i := 0; i < (len(entry.nums) - 1); i++ {
				num := entry.nums[i+1]
				if slices.Contains(indices, i) {
					tot *= num
				} else {
					tot += num
				}
			}
			if tot == entry.sum {
				finalSum += entry.sum
				break
			}
		}
	}

	return finalSum
}

// func solvePart2(inputFile string) int {
// 	return 0
// }

func loadDaySevenData(inputFile string) []calibration {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []calibration

	for scanner.Scan() {
		line := scanner.Text()

		var total int
		total, err = strconv.Atoi(strings.Split(line, ":")[0])
		if err != nil {
			log.Fatal(err)
		}
		vals := strings.Split(line, " ")
		var nums []int
		for i, val := range vals {
			if i == 0 {
				continue
			}
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}

		data = append(data, calibration{total, nums})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("input1.txt"))
}
