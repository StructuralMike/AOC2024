package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solvePart1(inputFile string) int {
	data := loadDayNineData(inputFile)

	frag := make([]int, 0, len(data))

	for len(data) > 0 {
		next := data[0]
		if len(data) > 1 {
			data = data[1:]
		} else {
			data = nil
		}

		if next == -1 {
			for i := len(data) - 1; i >= 0; i-- {
				if data[i] != -1 {
					next = data[i]
					data = data[0:i]
					break
				}
			}
		}
		frag = append(frag, next)
	}

	checksum := 0
	for i, num := range frag {
		if num != -1 {
			checksum += i * num
		}
	}

	return checksum
}

func solvePart2(inputFile string) int {
	data := loadDayNineData(inputFile)

	var fileStart int
	var fileIndex int
	var fileSize int
	var emptySize int
	var emptyStart int
	newFile := true
	for i := len(data) - 1; i > 0; i-- {
		num := data[i]

		// If we are looking for a new file and this is just empty space
		if num == -1 && newFile {
			continue
		}

		// If we have found a new file, set parameters
		if newFile {
			fileIndex = num
			fileSize = 0
			newFile = false
		}

		// Increment filesize
		fileSize++

		// If the upcoming data belongs to this file
		if data[i-1] == fileIndex {
			continue
		}

		// If we are here we have found the file start and end
		fileStart = i

		// Find next empty space
		emptySize = 0
		for j := 0; j <= i; j++ {
			if data[j] == -1 {
				emptySize++
				if data[j+1] == -1 {
					continue
				}
			}

			//If we have found the entire empty slice but it is too small
			if emptySize < fileSize {
				emptySize = 0
				continue
			}

			//The empty slice is big enough
			emptyStart = j - emptySize + 1
			for k := range fileSize {
				data[emptyStart+k] = fileIndex
				data[fileStart+k] = -1
			}
			break
		}
		newFile = true
	}

	checksum := 0
	for i, num := range data {
		if num != -1 {
			checksum += i * num
		}
	}

	return checksum
}

func loadDayNineData(inputFile string) []int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []int
	for scanner.Scan() {
		line := scanner.Text()
		isFile := true
		fileIndex := 0
		var writeData int
		for _, digit := range line {
			num, err := strconv.Atoi(string(digit))
			if err != nil {
				log.Fatal(err)
			}

			if isFile {
				writeData = fileIndex
				fileIndex++
			} else {
				writeData = -1
			}
			for range num {
				data = append(data, writeData)
			}
			isFile = !isFile
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	// fmt.Println(solvePart2("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
