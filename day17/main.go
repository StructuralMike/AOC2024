package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

func solvePart2(inputFile string) int {
	var output int
	var i int
	start := (1 << 48)

	_, program := loadDayData(inputFile)

	registers := Registers{}
	for i = 100000000000; i < 1000000000000; i++ {
		registers.A = start + i
		registers.B = 0
		registers.C = 0

		output = 0

		for registers.A != 0 {
			registers.C = (registers.A & 7) ^ 3
			registers.B = ((registers.C ^ 4) ^ (registers.A >> registers.C)) & 7
			if program[output] == registers.B {
				output++
			} else {
				break
			}
			registers.A = registers.A >> 3
		}

		if output > 12 {
			fmt.Println(output, strconv.FormatInt(int64(i), 2))
		}

		if output == len(program) {
			break
		}
	}

	return output
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
	start := time.Now()
	fmt.Println(solvePart2("input2.txt"))
	fmt.Println(time.Since(start))
}
