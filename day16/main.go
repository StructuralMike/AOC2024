package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type coor struct {
	x int
	y int
}

type travel struct {
	loc  coor
	cost int
	dir  rune
}

type Item struct {
	point    coor
	dir      rune
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

const Exit = 'E'
const Start = 'S'
const Wall = '#'
const Empty = '.'
const Visited = '-'

var Dirs = map[rune]coor{
	'^': {
		x: 0,
		y: -1,
	},
	'>': {
		x: 1,
		y: 0,
	},
	'v': {
		x: 0,
		y: 1,
	},
	'<': {
		x: -1,
		y: 0,
	},
}

var nextDir = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

func (c1 coor) Add(c2 coor) coor {
	return coor{
		x: c1.x + c2.x,
		y: c1.y + c2.y,
	}
}

func solvePart1(inputFile string) int {
	grid := loadDayData(inputFile)

	var deer coor
	var goal coor

	for y, row := range grid {
		for x, col := range row {
			if col == Start {
				deer = coor{x: x, y: y}
				grid[y][x] = 'S'
			}
			if col == Exit {
				goal = coor{x: x, y: y}
			}
		}
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
	distances[deer] = 0

	predecessors := make(map[coor]coor)

	pq := &PriorityQueue{}
	heap.Init(pq)
	start := &Item{point: deer, dir: '>', priority: 0}
	heap.Push(pq, start)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentLoc := current.point
		currentDir := current.dir

		if currentLoc == goal {
			break
		}

		// Find possible nodes
		newCost := distances[currentLoc] + 1
		newLoc := currentLoc.Add(Dirs[currentDir])
		if newCost < distances[newLoc] {
			distances[newLoc] = newCost
			predecessors[newLoc] = currentLoc
			heap.Push(pq, &Item{
				point:    newLoc,
				priority: newCost,
				dir:      currentDir,
			})
		}

		newCost += 1000
		currentDir = nextDir[currentDir]
		newLoc = currentLoc.Add(Dirs[currentDir])
		if newCost < distances[newLoc] {
			distances[newLoc] = newCost
			predecessors[newLoc] = currentLoc
			heap.Push(pq, &Item{
				point:    newLoc,
				priority: newCost,
				dir:      currentDir,
			})
		}

		currentDir = nextDir[currentDir]
		currentDir = nextDir[currentDir]
		newLoc = currentLoc.Add(Dirs[currentDir])
		if newCost < distances[newLoc] {
			distances[newLoc] = newCost
			predecessors[newLoc] = currentLoc
			heap.Push(pq, &Item{
				point:    newLoc,
				priority: newCost,
				dir:      currentDir,
			})
		}
	}

	pastLoc := predecessors[goal]
	for pastLoc != deer {
		grid[pastLoc.y][pastLoc.x] = 'o'
		pastLoc = predecessors[pastLoc]
	}
	printGrid(grid)

	return distances[goal]
}

// solvePart2 loads the maze, finds the shortest distance using Dijkstra (via solvePart1),
// then finds all paths that match this shortest distance, and visualizes one set of solutions.
func solvePart2(inputFile string) int {
	grid := loadDayData(inputFile)

	var startCoor, goalCoor coor
	// Identify start and goal
	for y, row := range grid {
		for x, col := range row {
			if col == Start {
				startCoor = coor{x: x, y: y}
				grid[y][x] = 'S' // Mark the start on the grid
			}
			if col == Exit {
				goalCoor = coor{x: x, y: y}
			}
		}
	}

	// Use Dijkstra (assumed in solvePart1) to find the shortest distance
	targetDistance := solvePart1(inputFile)

	// Retrieve all shortest paths that match the targetDistance
	solutions := findAllShortestPaths(grid, startCoor, goalCoor, '>', 0, targetDistance, [][]bool{}, []coor{})

	// If solutions exist, apply one solution to the original grid for visualization
	if len(solutions) > 0 {
		for _, solution := range solutions {
			for i := 0; i < len(solution); i++ {
				for j := 0; j < len(solution[0]); j++ {
					if solution[i][j] == 'o' {
						grid[i][j] = 'o'
					}
				}
			}
		}
	}

	tiles := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'o' {
				tiles++
			}
		}
	}

	// Print the grid with the solution path marked
	printGrid(grid)

	return tiles
}
func findAllShortestPaths() {
	return solutions
}

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

func loadDayData(inputFile string) [][]rune {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	// fmt.Println(solvePart2("sample_input.txt"))
	// fmt.Println(solvePart2("sample_input2.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2("input1.txt"))
}
