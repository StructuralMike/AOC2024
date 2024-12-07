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
