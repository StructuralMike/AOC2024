package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Registers struct {
	A int
	B int
	C int
}

func solvePart1(inputFile string) string {
	registers, program := loadDayData(inputFile)

	pointer := 0
	end := len(program) - 1
	var opcode int
	var literalOperand int
	var comboOperand int
	var output []string

	for 0 <= pointer && pointer < end {
		opcode = program[pointer]
		literalOperand = program[pointer+1]

		switch literalOperand {
		case 4:
			comboOperand = registers.A
		case 5:
			comboOperand = registers.B
		case 6:
			comboOperand = registers.C
		default:
			comboOperand = literalOperand
		}

		// if slices.Contains([]int{0, 2, 5, 6, 7}, opcode) && comboOperand == 7 {
		// 	log.Fatal("Reserved operand used.")
		// }

		switch opcode {
		case 0:
			registers.A = registers.A >> comboOperand
		case 1:
			registers.B = registers.B ^ literalOperand
		case 2:
			registers.B = comboOperand%2 ^ 3
		case 3:
			if registers.A != 0 {
				pointer = literalOperand
				continue
			}
		case 4:
			registers.B = registers.B ^ registers.C
		case 5:
			output = append(output, strconv.Itoa(comboOperand%8))
		case 6:
			registers.B = registers.A >> comboOperand
		case 7:
			registers.C = registers.A >> comboOperand
		}

		pointer += 2
	}

	return strings.Join(output, ",")
}

func solvePart2() int {
	program := []int{2, 4, 1, 3, 7, 5, 0, 3, 1, 4, 4, 7, 5, 5, 3, 0}

	end := len(program) - 1
	bestProgram := 0

	var opcode int
	var literalOperand int
	var comboOperand int
	var output []int
	var registers Registers
	var i int

	var mem = make(map[int][]int)
	for i = 3216114232161142; true; i++ {
		registers = Registers{
			A: i,
			B: 0,
			C: 0,
		}

		output = []int{}

		for pointer := 0; pointer < end; {
			if pointer == 0 && len(output) > 0 {
				if seen, ok := mem[registers.A]; ok {
					output = append(output, seen...)
					break
				}
			}
			opcode = program[pointer]
			literalOperand = program[pointer+1]

			switch literalOperand {
			case 4:
				comboOperand = registers.A
			case 5:
				comboOperand = registers.B
			case 6:
				comboOperand = registers.C
			default:
				comboOperand = literalOperand
			}

			switch opcode {
			case 0:
				registers.A = registers.A >> comboOperand
			case 1:
				registers.B = registers.B ^ literalOperand
			case 2:
				registers.B = comboOperand & 7
			case 3:
				if registers.A != 0 {
					pointer = literalOperand
					continue
				}
			case 4:
				registers.B = registers.B ^ registers.C
			case 5:
				newInt := comboOperand & 7
				output = append(output, newInt)
			case 6:
				registers.B = registers.A >> comboOperand
			case 7:
				registers.C = registers.A >> comboOperand
			}

			pointer += 2
		}

		mem[i] = output

		if len(output) > bestProgram {
			cnt := 0
			for j, num := range output {
				if program[j] != num {
					break
				}
				cnt++
			}
			if cnt >= bestProgram {
				bestProgram = cnt
				fmt.Println(bestProgram, i, output)
				if len(output) == len(program) {
					if intArraysEqual(output, program) {
						break
					}
				}
			}
		}

	}

	return i
}

func intArraysEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func loadDayData(inputFile string) (Registers, []int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	registers := Registers{
		A: 0,
		B: 0,
		C: 0,
	}

	var program []int

	for scanner.Scan() {
		line := strings.Split(strings.Replace(scanner.Text(), " ", "", -1), ":")
		switch line[0] {
		case "RegisterA":
			registers.A, err = strconv.Atoi(line[1])
		case "RegisterB":
			registers.B, err = strconv.Atoi(line[1])
		case "RegisterC":
			registers.C, err = strconv.Atoi(line[1])
		case "Program":
			vals := strings.Split(line[1], ",")
			for _, val := range vals {
				num, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal(err)
				}
				program = append(program, num)
			}
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return registers, program
}

func main() {
	// fmt.Println(solvePart1("sample_input.txt"))
	// fmt.Println(solvePart1("input1.txt"))
	fmt.Println(solvePart2())
}
