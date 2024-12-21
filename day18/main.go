package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const Exit = 'E'
const Start = 'S'
const Wall = '#'
const Empty = '.'
const Visited = '-'

type coor struct {
	x int
	y int
}

func (c1 coor) Add(c2 coor) coor {
	return coor{
		x: c1.x + c2.x,
		y: c1.y + c2.y,
	}
}

type Item struct {
	loc      coor
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func solvePart1(inputFile string, rows int, cols int, bytes int) int {
	grid := loadDayData(inputFile, rows, cols, bytes)

	start := coor{0, 0}
	end := coor{cols - 1, rows - 1}

	directions := []coor{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	distances := make(map[coor]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			var newInt int
			if grid[i][j] == '#' {
				newInt = -1
			} else {
				newInt = math.MaxInt
			}
			distances[coor{x: j, y: i}] = newInt
		}
	}
	distances[start] = 0

	predecessors := make(map[coor]coor)

	pq := &PriorityQueue{}
	heap.Init(pq)
	startItem := &Item{loc: start, priority: 0}
	heap.Push(pq, startItem)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentLoc := current.loc

		if currentLoc == end {
			break
		}

		newPriority := distances[currentLoc] + 1
		for _, direction := range directions {
			newLoc := currentLoc.Add(direction)
			if newLoc.x < 0 || newLoc.y < 0 || newLoc.x >= len(grid[0]) || newLoc.y >= len(grid) {
				continue
			}
			if predecessors[currentLoc] == newLoc {
				continue
			}

			if newPriority < distances[newLoc] {
				distances[newLoc] = newPriority
				predecessors[newLoc] = currentLoc
				heap.Push(pq, &Item{
					loc:      newLoc,
					priority: newPriority,
				})
			}
		}
	}

	grid[start.y][start.x] = 'o'
	grid[end.y][end.x] = 'o'

	stepsTaken := 1
	pastLoc := predecessors[end]
	for pastLoc != start {
		grid[pastLoc.y][pastLoc.x] = 'o'
		pastLoc = predecessors[pastLoc]
		stepsTaken++
	}

	printGrid(grid)

	return stepsTaken
}

// func solvePart2(inputFile string) int {
// 	return 0
// }

func printGrid(grid [][]rune) {
	// Bright ANSI color codes
	black := "\033[90m"
	red := "\033[91m"
	lightRed := "\033[31"
	green := "\033[92m"
	yellow := "\033[93m"
	reset := "\033[0m"

	for _, row := range grid {
		var sb strings.Builder
		for _, ch := range row {
			switch ch {
			case Empty:
				sb.WriteString(black)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			case Start:
				sb.WriteString(red)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			case Exit:
				sb.WriteString(yellow)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			case Wall:
				sb.WriteString(green)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			case Visited:
				sb.WriteString(lightRed)
				sb.WriteRune(ch)
				sb.WriteString(reset)
			default:
				sb.WriteRune(ch)
			}
		}

		fmt.Println(sb.String())
	}
}

func loadDayData(inputFile string, rows int, cols int, bytes int) [][]rune {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	grid := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]rune, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = '.'
		}
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() && bytes > 0 {
		nums := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		grid[y][x] = '#'
		bytes--
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt", 7, 7, 12))
	fmt.Println(solvePart1("input1.txt", 71, 71, 2987))
	// fmt.Println(solvePart2("input1.txt"))
}
