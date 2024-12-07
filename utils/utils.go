package utils

import (
	"log"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func StringsToInts(arr []string) []int {
	var ints []int
	for _, str := range arr {
		n, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, n)
	}
	return ints
}

func SumArray(arr []int) int {
	result := 0
	for _, n := range arr {
		result += n
	}
	return result
}

func IntegerRange(end int) []int {
	if end < 0 {
		return []int{}
	}

	s := make([]int, 0, end+1)
	for i := 0; i <= end; i++ {
		s = append(s, i)
	}

	return s
}
