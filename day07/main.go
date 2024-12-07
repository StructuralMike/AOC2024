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
	sum  uint64
	nums []uint64
}

func solvePart1(inputFile string) uint64 {
	data := loadDaySevenData(inputFile)

	var finalSum uint64

	for _, entry := range data {
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

func solvePart2(inputFile string) uint64 {
	data := loadDaySevenData(inputFile)

	var finalSum uint64
	operators := []rune{'+', '*', '|'}

	for _, entry := range data {
		for combination := range generateCombinations(operators, len(entry.nums)-1) {
			tot := entry.nums[0]
			for i, operation := range combination {
				num := entry.nums[i+1]
				switch operation {
				case '+':
					tot += num
				case '*':
					tot *= num
				case '|':
					left := strconv.FormatInt(int64(tot), 10)
					right := strconv.FormatInt(int64(num), 10)
					sum, err := strconv.ParseUint(left+right, 10, 64)
					if err != nil {
						log.Fatal(err)
					}
					tot = sum
				}
				if tot > entry.sum {
					break
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

func generateCombinations(operators []rune, length int) <-chan []rune {
	c := make(chan []rune)

	go func() {
		defer close(c)
		if length == 0 {
			return
		}

		combination := make([]rune, length)
		var generate func(pos int)
		generate = func(pos int) {
			if pos == length {
				result := append([]rune(nil), combination...)
				c <- result
				return
			}

			for _, op := range operators {
				combination[pos] = op
				generate(pos + 1)
			}
		}

		generate(0)

	}()

	return c
}

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

		var total uint64
		total, err = strconv.ParseUint(strings.Split(line, ":")[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		vals := strings.Split(line, " ")
		var nums []uint64
		for i, val := range vals {
			if i == 0 {
				continue
			}
			num, err := strconv.ParseUint(val, 10, 64)
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
	fmt.Println(solvePart2("input1.txt"))
}
