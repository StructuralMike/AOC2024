package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func solvePart1(inputFile string) int {
	graph := loadDayData(inputFile)

	sets := make(map[string]struct{})
	chiefs := 0
	for c1, p1 := range graph {
		for c2 := range p1 {
			for c3 := range graph[c2] {
				if _, ok := graph[c1][c3]; ok {
					set := []string{c1, c2, c3}
					slices.Sort(set)
					norm := strings.Join(set, "-")
					if _, ok := sets[norm]; !ok {
						sets[norm] = struct{}{}
						if strings.HasPrefix(c1, "t") || strings.HasPrefix(c2, "t") || strings.HasPrefix(c3, "t") {
							chiefs++
							fmt.Println(norm)
						}
					}
				}
			}
		}
	}
	//	printGraph(graph)
	return chiefs
}

// func solvePart2(inputFile string) int {
// 	return 0
// }

func printGraph(graph map[string]map[string]struct{}) {
	for node, neighbours := range graph {
		fmt.Printf("%s: ", node)
		for neighbor := range neighbours {
			fmt.Printf("%s ", neighbor)
		}
		fmt.Println()
	}
}

func loadDayData(inputFile string) map[string]map[string]struct{} {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	data := make(map[string]map[string]struct{})
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), "-")
		A := nodes[0]
		B := nodes[1]
		if _, ok := data[A]; !ok {
			data[A] = make(map[string]struct{})
		}
		if _, ok := data[B]; !ok {
			data[B] = make(map[string]struct{})
		}
		data[A][B] = struct{}{}
		data[B][A] = struct{}{}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	fmt.Println(solvePart1("input1.txt"))
	// fmt.Println(solvePart2("input1.txt"))
}
